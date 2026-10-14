// Package herosignal, HeroSignal public API'si için Go sunucu SDK'sıdır.
//
//	hs, err := herosignal.NewClient(os.Getenv("HS_API_KEY"), herosignal.WithAppID("app_…"))
//	res, err := hs.Notifications.Create(ctx, &herosignal.CreateNotificationDto{
//		Target:   &herosignal.CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {userID}}},
//		Contents: map[string]string{"en": "Shipped"},
//	}, herosignal.WithIdempotencyKey("order-"+orderID))
//
// Modeller ve servisler api/openapi.yaml'dan üretilir (models_gen.go, resources_gen.go); HTTP katmanı,
// yeniden deneme, hatalar, sayfalama ve webhook doğrulaması el yazımıdır.
package herosignal

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Version, SDK sürümüdür (User-Agent başlığında gönderilir).
const Version = "0.1.0"

// DefaultBaseURL, varsayılan API adresidir; yerel geliştirmede WithBaseURL("http://localhost:4000").
const DefaultBaseURL = "https://api.herosignal.dev"

const (
	defaultTimeout    = 30 * time.Second
	defaultMaxRetries = 3
)

var backoff = []time.Duration{300 * time.Millisecond, time.Second, 3 * time.Second}

var errNilBody = errors.New("herosignal: istek gövdesi nil olamaz")

// Client, HeroSignal API istemcisidir; eşzamanlı kullanım için güvenlidir.
// Kaynak servisleri (Notifications, Users, Segments …) gömülü services yapısından gelir.
type Client struct {
	services
	apiKey             string
	baseURL            string
	appID              string
	httpClient         *http.Client
	timeout            time.Duration
	maxRetries         int
	userAgent          string
	autoIdempotencyKey bool
	sleep              func(ctx context.Context, d time.Duration) error
}

// ClientOption, NewClient ayarıdır.
type ClientOption func(*Client)

// WithAppID, çağrılarda ForApp verilmezse kullanılacak varsayılan app kimliğidir.
func WithAppID(appID string) ClientOption { return func(c *Client) { c.appID = appID } }

// WithBaseURL, API adresini değiştirir.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) { c.baseURL = strings.TrimRight(baseURL, "/") }
}

// WithHTTPClient, özel *http.Client (proxy, TLS, test transport'u) kullanır.
func WithHTTPClient(h *http.Client) ClientOption { return func(c *Client) { c.httpClient = h } }

// WithTimeout, deneme başına istek zaman aşımıdır (varsayılan 30 sn).
func WithTimeout(d time.Duration) ClientOption { return func(c *Client) { c.timeout = d } }

// WithMaxRetries, en fazla yeniden deneme sayısıdır (varsayılan 3). 429 her metotta; 5xx ve ağ hataları
// GET/HEAD/OPTIONS/PUT/DELETE'te, POST/PATCH'te yalnızca Idempotency-Key varsa veya istek sunucuya hiç
// gönderilmediyse (bağlantı reddi, DNS hatası, bağlantı zaman aşımı) yeniden denenir.
func WithMaxRetries(n int) ClientOption { return func(c *Client) { c.maxRetries = n } }

// WithUserAgent, User-Agent başlığını değiştirir.
func WithUserAgent(ua string) ClientOption { return func(c *Client) { c.userAgent = ua } }

// WithAutoIdempotencyKey, Idempotency-Key destekleyen çağrılarda anahtar verilmezse UUID üretilmesini
// açar/kapatır (varsayılan açık).
func WithAutoIdempotencyKey(enabled bool) ClientOption {
	return func(c *Client) { c.autoIdempotencyKey = enabled }
}

// NewClient, API anahtarıyla (hs_live_…) istemci oluşturur.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("herosignal: apiKey gerekli")
	}
	c := &Client{
		apiKey:             apiKey,
		baseURL:            DefaultBaseURL,
		httpClient:         http.DefaultClient,
		timeout:            defaultTimeout,
		maxRetries:         defaultMaxRetries,
		userAgent:          "herosignal-go/" + Version,
		autoIdempotencyKey: true,
		sleep:              sleepContext,
	}
	for _, o := range opts {
		o(c)
	}
	c.initServices()
	return c, nil
}

// String, istemciyi API anahtarını göstermeden biçimlendirir (log'larda gizli bilgi sızmasın).
func (c *Client) String() string {
	return fmt.Sprintf("herosignal.Client{baseURL: %q, appID: %q}", c.baseURL, c.appID)
}

// GoString, %#v biçiminde de API anahtarını gizler.
func (c *Client) GoString() string { return c.String() }

func sleepContext(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// RequestOption, tek bir çağrının ayarlarını değiştirir.
type RequestOption func(*requestConfig)

type requestConfig struct {
	client         *Client
	appID          string
	idempotencyKey string
	headers        http.Header
	timeout        time.Duration
	maxRetries     int
}

// ForApp, bu çağrı için app kimliğini belirler (istemcinin varsayılan app'i yerine).
func ForApp(appID string) RequestOption { return func(r *requestConfig) { r.appID = appID } }

// WithIdempotencyKey, Idempotency-Key başlığını belirler (aynı anahtarla 24 saat içindeki tekrar aynı yanıtı döner).
func WithIdempotencyKey(key string) RequestOption {
	return func(r *requestConfig) { r.idempotencyKey = key }
}

// WithHeader, bu çağrıya ek HTTP başlığı ekler.
func WithHeader(key, value string) RequestOption {
	return func(r *requestConfig) { r.headers.Add(key, value) }
}

// WithRequestTimeout, bu çağrı için deneme başına zaman aşımıdır.
func WithRequestTimeout(d time.Duration) RequestOption {
	return func(r *requestConfig) { r.timeout = d }
}

// WithRequestMaxRetries, bu çağrı için en fazla yeniden deneme sayısıdır.
func WithRequestMaxRetries(n int) RequestOption {
	return func(r *requestConfig) { r.maxRetries = n }
}

func (c *Client) requestConfig(opts []RequestOption) *requestConfig {
	r := &requestConfig{client: c, appID: c.appID, headers: http.Header{}, timeout: c.timeout, maxRetries: c.maxRetries}
	for _, o := range opts {
		o(r)
	}
	return r
}

func (r *requestConfig) resolveAppID() (string, error) {
	if r.appID == "" {
		return "", errors.New("herosignal: app kimliği verilmedi (WithAppID veya ForApp)")
	}
	return r.appID, nil
}

// operation, üretilmiş servislerin Client.do'ya verdiği istek tanımıdır.
type operation struct {
	method         string
	path           string
	pathParams     map[string]string
	query          url.Values
	queryAppID     bool
	body           any
	multipart      *multipartBody
	idempotent     bool
	idempotencyKey string
}

// Request, SDK'da metodu olmayan uçlar veya JSON dışı yanıtlar (ör. analitik format=csv) için düşük seviyeli
// istektir. path "/v1/…" biçimindedir; {appId} varsa app kimliğiyle doldurulur. out *[]byte veya *string ise
// ham gövde, diğer işaretçilerde JSON ayrıştırılır; nil ise yanıt gövdesi yok sayılır.
func (c *Client) Request(ctx context.Context, method, path string, query url.Values, body any, out any, opts ...RequestOption) error {
	return c.do(ctx, &operation{method: strings.ToUpper(method), path: path, query: query, body: body}, c.requestConfig(opts), out)
}

func (c *Client) do(ctx context.Context, op *operation, ro *requestConfig, out any) error {
	path := op.path
	if strings.Contains(path, "{appId}") {
		appID, err := ro.resolveAppID()
		if err != nil {
			return err
		}
		path = strings.ReplaceAll(path, "{appId}", url.PathEscape(appID))
	}
	for name, value := range op.pathParams {
		if value == "" {
			return fmt.Errorf("herosignal: %q yol parametresi boş olamaz", name)
		}
		path = strings.ReplaceAll(path, "{"+name+"}", url.PathEscape(value))
	}
	query := url.Values{}
	for k, v := range op.query {
		query[k] = v
	}
	if op.queryAppID {
		appID, err := ro.resolveAppID()
		if err != nil {
			return err
		}
		query.Set("app_id", appID)
	}
	fullURL := c.baseURL + path
	if len(query) > 0 {
		fullURL += "?" + query.Encode()
	}

	var payload []byte
	contentType := ""
	switch {
	case op.multipart != nil:
		data, ct, err := op.multipart.encode()
		if err != nil {
			return err
		}
		payload, contentType = data, ct
	case op.body != nil:
		data, err := json.Marshal(op.body)
		if err != nil {
			return fmt.Errorf("herosignal: gövde serileştirilemedi: %w", err)
		}
		payload, contentType = data, "application/json"
	}

	key := ro.idempotencyKey
	if key == "" {
		key = op.idempotencyKey
	}
	if key == "" && op.idempotent && c.autoIdempotencyKey {
		key = newUUID()
	}

	hasKey := key != "" || ro.headers.Get("Idempotency-Key") != ""
	for attempt := 0; ; attempt++ {
		resp, respBody, err := c.attempt(ctx, op.method, fullURL, payload, contentType, key, ro)
		var apiErr *Error
		notSent := false
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			apiErr = &Error{StatusCode: 0, Code: CodeNetwork, Message: err.Error(), cause: err}
			notSent = requestNotSent(err)
		} else if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return decodeResponse(resp, respBody, out)
		} else {
			apiErr = errorFromResponse(resp, respBody)
		}
		if attempt >= ro.maxRetries || !shouldRetry(op.method, apiErr, hasKey, notSent) {
			return apiErr
		}
		delay := apiErr.RetryAfter
		if delay <= 0 {
			delay = backoff[min(attempt, len(backoff)-1)]
		}
		if err := c.sleep(ctx, delay); err != nil {
			return err
		}
	}
}

// shouldRetry, yeniden deneme politikasıdır: 429 her metotta (istek işlenmeden reddedilmiştir); 5xx ve ağ
// hataları idempotent metotlarda, POST/PATCH'te yalnızca Idempotency-Key varsa ya da istek sunucuya hiç
// gönderilmediyse. Aksi hâlde sunucu isteği işlemiş olabilir ve tekrar çift kayıt üretebilir.
func shouldRetry(method string, e *Error, hasIdempotencyKey, notSent bool) bool {
	if e.StatusCode == 429 {
		return true
	}
	if !e.Retryable() {
		return false
	}
	switch strings.ToUpper(method) {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPut, http.MethodDelete:
		return true
	}
	return hasIdempotencyKey || notSent
}

// requestNotSent, hatanın bağlantı kurulmadan (DNS çözümlenemedi, bağlantı reddi, bağlantı zaman aşımı, proxy
// bağlantısı) oluştuğunu, yani isteğin sunucuya hiç gönderilmediğini bildirir.
func requestNotSent(err error) bool {
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return true
	}
	var opErr *net.OpError
	return errors.As(err, &opErr) && (opErr.Op == "dial" || opErr.Op == "proxyconnect")
}

func (c *Client) attempt(ctx context.Context, method, fullURL string, payload []byte, contentType, key string, ro *requestConfig) (*http.Response, []byte, error) {
	if ro.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, ro.timeout)
		defer cancel()
	}
	var reader io.Reader
	if payload != nil {
		reader = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(ctx, method, fullURL, reader)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if key != "" {
		req.Header.Set("Idempotency-Key", key)
	}
	for k, vs := range ro.headers {
		req.Header.Del(k)
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return resp, data, nil
}

func decodeResponse(resp *http.Response, data []byte, out any) error {
	switch o := out.(type) {
	case nil:
		return nil
	case *[]byte:
		*o = data
		return nil
	case *string:
		*o = string(data)
		return nil
	}
	if len(bytes.TrimSpace(data)) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, out); err != nil {
		return &Error{StatusCode: resp.StatusCode, Code: "DECODE", Message: "yanıt ayrıştırılamadı: " + err.Error(), RequestID: resp.Header.Get("X-Request-Id"), cause: err}
	}
	return nil
}

func errorFromResponse(resp *http.Response, data []byte) *Error {
	var envelope struct {
		Error struct {
			Code      string         `json:"code"`
			Message   string         `json:"message"`
			RequestID string         `json:"request_id"`
			Details   map[string]any `json:"details"`
		} `json:"error"`
	}
	_ = json.Unmarshal(data, &envelope)
	e := &Error{
		StatusCode: resp.StatusCode,
		Code:       envelope.Error.Code,
		Message:    envelope.Error.Message,
		RequestID:  envelope.Error.RequestID,
		Details:    envelope.Error.Details,
		RetryAfter: parseRetryAfter(resp.Header.Get("Retry-After")),
	}
	if e.Code == "" {
		e.Code = "HTTP_" + strconv.Itoa(resp.StatusCode)
	}
	if e.Message == "" {
		e.Message = "HTTP " + strconv.Itoa(resp.StatusCode)
	}
	if e.RequestID == "" {
		e.RequestID = resp.Header.Get("X-Request-Id")
	}
	return e
}

// parseRetryAfter, saniye veya HTTP tarihi biçimindeki Retry-After başlığını süreye çevirir.
func parseRetryAfter(v string) time.Duration {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0
	}
	if secs, err := strconv.ParseFloat(v, 64); err == nil {
		if secs > 0 {
			return time.Duration(secs * float64(time.Second))
		}
		return 0
	}
	if t, err := http.ParseTime(v); err == nil {
		if d := time.Until(t); d > 0 {
			return d
		}
	}
	return 0
}

// newUUID, rastgele (v4) UUID üretir.
func newUUID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic("herosignal: rastgele sayı üretilemedi: " + err.Error())
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// ---- sorgu yardımcıları ----------------------------------------------------------------------

func queryString(v any) string {
	switch x := v.(type) {
	case string:
		return x
	case bool:
		return strconv.FormatBool(x)
	case int64:
		return strconv.FormatInt(x, 10)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	default:
		data, err := json.Marshal(x)
		if err != nil {
			return fmt.Sprint(x)
		}
		return string(data)
	}
}

func addQuery[T any](q url.Values, key string, v *T) {
	if v == nil {
		return
	}
	q.Set(key, queryString(any(*v)))
}

func addQuerySlice[T any](q url.Values, key string, vs []T) {
	for _, v := range vs {
		q.Add(key, queryString(any(v)))
	}
}

func nonEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// ---- sayfalama ---------------------------------------------------------------------------------

// paginate, data / next_cursor / has_more zarfını sayfa sayfa dolaşan bir yineleyici döndürür.
func paginate[T any](fetch func(cursor *string) ([]T, *string, bool, error)) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var cursor *string
		for {
			items, next, more, err := fetch(cursor)
			if err != nil {
				var zero T
				yield(zero, err)
				return
			}
			for _, item := range items {
				if !yield(item, nil) {
					return
				}
			}
			if !more || next == nil || *next == "" {
				return
			}
			cursor = next
		}
	}
}

// ---- multipart ---------------------------------------------------------------------------------

// FileUpload, multipart dosya alanıdır. İçerik bellekte tutulur ki yeniden denemeler aynı gövdeyi göndersin.
type FileUpload struct {
	Filename    string
	ContentType string
	Content     []byte
}

type multipartPart struct {
	name  string
	value string
	file  *FileUpload
}

type multipartBody struct{ parts []multipartPart }

func (m *multipartBody) field(name string, v any) {
	m.parts = append(m.parts, multipartPart{name: name, value: queryString(v)})
}

func (m *multipartBody) file(name string, f *FileUpload) {
	if f != nil {
		m.parts = append(m.parts, multipartPart{name: name, file: f})
	}
}

func (m *multipartBody) encode() ([]byte, string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	for _, p := range m.parts {
		if p.file == nil {
			if err := w.WriteField(p.name, p.value); err != nil {
				return nil, "", err
			}
			continue
		}
		filename := p.file.Filename
		if filename == "" {
			filename = "upload"
		}
		ct := p.file.ContentType
		if ct == "" {
			ct = "application/octet-stream"
		}
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(p.name), escapeQuotes(filename)))
		h.Set("Content-Type", ct)
		part, err := w.CreatePart(h)
		if err != nil {
			return nil, "", err
		}
		if _, err := part.Write(p.file.Content); err != nil {
			return nil, "", err
		}
	}
	if err := w.Close(); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), w.FormDataContentType(), nil
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { return quoteEscaper.Replace(s) }
