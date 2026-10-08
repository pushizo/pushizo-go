// Package herosignal, HeroSignal sunucu SDK'sıdır: bildirim gönderme, kullanıcı/abonelik yönetimi,
// segment/şablon/kampanya, analitik ve webhook yönetimi.
//
//	client, err := herosignal.NewClient(os.Getenv("HS_API_KEY"), herosignal.WithAppID("app_…"))
//	res, err := client.Notifications.Create(ctx, &herosignal.CreateNotificationDto{
//		Target:   herosignal.CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {"u1"}}},
//		Contents: map[string]string{"en": "Shipped"},
//	}, herosignal.WithIdempotencyKey("order-42"))
//
// Kaynaklar ve modeller api/openapi.yaml'dan üretilir (*_gen.go); HTTP katmanı (kimlik doğrulama, yeniden deneme,
// idempotency, hata eşleme, sayfalama) el yazımıdır. Webhook imza doğrulaması için alt paket: herosignal/webhook.
package herosignal

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// Version, SDK sürümüdür.
	Version = "0.1.0"
	// DefaultBaseURL, varsayılan API adresidir.
	DefaultBaseURL = "https://api.herosignal.dev"
	// DefaultTimeout, deneme başına varsayılan zaman aşımıdır.
	DefaultTimeout = 30 * time.Second
	// DefaultMaxRetries, 429/5xx ve ağ hatalarında varsayılan en fazla yeniden deneme sayısıdır.
	DefaultMaxRetries = 3
)

// backoff, Retry-After yoksa denemeler arası bekleme süreleridir (Node/PHP/Python SDK'larıyla aynı).
var backoff = []time.Duration{300 * time.Millisecond, time.Second, 3 * time.Second}

// Client, HeroSignal API istemcisidir; eşzamanlı kullanım için güvenlidir.
type Client struct {
	resources

	apiKey             string
	appID              string
	baseURL            string
	userAgent          string
	httpClient         *http.Client
	timeout            time.Duration
	maxRetries         int
	autoIdempotencyKey bool
	sleep              func(ctx context.Context, d time.Duration) error
	now                func() time.Time
}

// NewClient, API anahtarı (hs_live_…) ile istemci oluşturur.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("herosignal: apiKey gerekli")
	}
	c := &Client{
		apiKey:             apiKey,
		baseURL:            DefaultBaseURL,
		userAgent:          "herosignal-go/" + Version,
		httpClient:         &http.Client{},
		timeout:            DefaultTimeout,
		maxRetries:         DefaultMaxRetries,
		autoIdempotencyKey: true,
		sleep:              sleepContext,
		now:                time.Now,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.resources.init(c)
	return c, nil
}

// AppID, istemcinin varsayılan app kimliğidir.
func (c *Client) AppID() string { return c.appID }

// IdentityHash, SDK identity verification için HMAC-SHA256(sdkSecret, externalID) değerini (hex) döndürür.
func IdentityHash(sdkSecret, externalID string) string {
	mac := hmac.New(sha256.New, []byte(sdkSecret))
	mac.Write([]byte(externalID))
	return hex.EncodeToString(mac.Sum(nil))
}

// Ptr, değerin işaretçisini döndürür (isteğe bağlı model alanları için: herosignal.Ptr("x")).
func Ptr[T any](v T) *T { return &v }

func deref[T any](p *T) T {
	var zero T
	if p == nil {
		return zero
	}
	return *p
}

func trimSlash(s string) string { return strings.TrimRight(s, "/") }

// FileUpload, multipart dosya alanıdır. İçerik ilk denemeden önce bir kez okunur; yeniden denemeler aynı baytları
// gönderir.
type FileUpload struct {
	// Name dosya adı (ör. "users.csv").
	Name string
	// ContentType içerik tipi; boşsa "application/octet-stream".
	ContentType string
	// Reader dosya içeriği.
	Reader io.Reader
}

// Do, düşük seviyeli istek gönderir (SDK'da henüz metodu olmayan uçlar için). query nil olabilir; body JSON'a
// çevrilir. out bir *string veya *[]byte ise ham gövde yazılır (ör. format=csv), aksi halde JSON çözülür; nil ise
// gövde yok sayılır.
func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body any, out any, opts ...RequestOption) error {
	call := newAPICall(method, path)
	for k, vs := range query {
		for _, v := range vs {
			call.query.Add(k, v)
		}
	}
	call.body = body
	return c.do(ctx, call, out, opts)
}

// apiCall, üretilmiş kaynak metotlarının doldurduğu istek tanımıdır.
type apiCall struct {
	method      string
	path        string
	pathParams  []string
	query       url.Values
	body        any
	multipart   bool
	formFields  [][2]string
	formFiles   []formFile
	appIDInPath bool
	queryAppID  bool
	bodyAppID   bool
	idempotent  bool
	err         error
}

type formFile struct {
	field string
	file  *FileUpload
}

func newAPICall(method, path string) *apiCall {
	return &apiCall{method: method, path: path, query: url.Values{}}
}

func (a *apiCall) pathParam(name, value string) {
	if value == "" && a.err == nil {
		a.err = fmt.Errorf("herosignal: %q yol parametresi boş olamaz", name)
	}
	a.pathParams = append(a.pathParams, name, value)
}

func (a *apiCall) setQuery(name string, value any) {
	for _, v := range queryValues(value) {
		a.query.Add(name, v)
	}
}

func (a *apiCall) formField(name string, value any) {
	for _, v := range queryValues(value) {
		a.formFields = append(a.formFields, [2]string{name, v})
	}
}

func (a *apiCall) formFile(name string, f *FileUpload) {
	if f != nil {
		a.formFiles = append(a.formFiles, formFile{field: name, file: f})
	}
}

// queryValues, sorgu/form değerini metne çevirir: nil işaretçiler atlanır, diziler tekrarlanan anahtar olur,
// bool "true"/"false", diğer karmaşık değerler JSON.
func queryValues(value any) []string {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		return []string{v}
	case *string:
		if v == nil {
			return nil
		}
		return []string{*v}
	case int64:
		return []string{strconv.FormatInt(v, 10)}
	case *int64:
		if v == nil {
			return nil
		}
		return []string{strconv.FormatInt(*v, 10)}
	case int:
		return []string{strconv.Itoa(v)}
	case float64:
		return []string{strconv.FormatFloat(v, 'f', -1, 64)}
	case *float64:
		if v == nil {
			return nil
		}
		return []string{strconv.FormatFloat(*v, 'f', -1, 64)}
	case bool:
		return []string{strconv.FormatBool(v)}
	case *bool:
		if v == nil {
			return nil
		}
		return []string{strconv.FormatBool(*v)}
	case []string:
		return v
	case []int64:
		out := make([]string, 0, len(v))
		for _, n := range v {
			out = append(out, strconv.FormatInt(n, 10))
		}
		return out
	case []any:
		out := []string{}
		for _, item := range v {
			out = append(out, queryValues(item)...)
		}
		return out
	default:
		raw, err := json.Marshal(v)
		if err != nil || string(raw) == "null" {
			return nil
		}
		return []string{string(raw)}
	}
}

func (c *Client) resolveAppID(cfg *requestConfig) (string, error) {
	if cfg.appID != "" {
		return cfg.appID, nil
	}
	if c.appID != "" {
		return c.appID, nil
	}
	return "", errors.New("herosignal: app ID verilmedi (WithAppID istemci seçeneği veya WithRequestAppID)")
}

// do, isteği hazırlar ve yeniden deneme politikasıyla gönderir.
func (c *Client) do(ctx context.Context, call *apiCall, out any, opts []RequestOption) error {
	if call.err != nil {
		return call.err
	}
	cfg := &requestConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	path := call.path
	if call.appIDInPath || call.queryAppID || call.bodyAppID {
		appID, err := c.resolveAppID(cfg)
		if err != nil {
			return err
		}
		if call.appIDInPath {
			path = strings.ReplaceAll(path, "{appId}", url.PathEscape(appID))
		}
		if call.queryAppID {
			call.query.Set("app_id", appID)
		}
		if call.bodyAppID {
			cfg.appID = appID
		}
	}
	for i := 0; i+1 < len(call.pathParams); i += 2 {
		path = strings.ReplaceAll(path, "{"+call.pathParams[i]+"}", url.PathEscape(call.pathParams[i+1]))
	}

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.apiKey)
	headers.Set("Accept", "application/json")
	headers.Set("User-Agent", c.userAgent)

	var payload []byte
	idempotencyKey := cfg.idempotencyKey
	if call.multipart {
		body, contentType, err := encodeMultipart(call.formFields, call.formFiles)
		if err != nil {
			return err
		}
		payload = body
		headers.Set("Content-Type", contentType)
	} else if call.body != nil {
		raw, err := json.Marshal(call.body)
		if err != nil {
			return fmt.Errorf("herosignal: gövde JSON'a çevrilemedi: %w", err)
		}
		if call.bodyAppID || (call.idempotent && idempotencyKey == "") {
			raw, idempotencyKey, err = adjustBody(raw, call.bodyAppID, cfg.appID, idempotencyKey)
			if err != nil {
				return err
			}
		}
		if string(raw) != "null" {
			payload = raw
			headers.Set("Content-Type", "application/json")
		}
	}
	if idempotencyKey == "" && call.idempotent && c.autoIdempotencyKey {
		idempotencyKey = newUUID()
	}
	if idempotencyKey != "" {
		headers.Set("Idempotency-Key", idempotencyKey)
	}
	for k, vs := range cfg.headers {
		headers[k] = vs
	}

	endpoint := c.baseURL + path
	if len(call.query) > 0 {
		endpoint += "?" + call.query.Encode()
	}
	maxRetries := c.maxRetries
	if cfg.maxRetries != nil {
		maxRetries = *cfg.maxRetries
	}
	timeout := c.timeout
	if cfg.timeout != nil {
		timeout = *cfg.timeout
	}

	for attempt := 0; ; attempt++ {
		resp, body, err := c.attempt(ctx, call.method, endpoint, headers, payload, timeout)
		var apiErr *Error
		if err != nil {
			if ctx.Err() != nil {
				return &Error{Code: "NETWORK", Message: ctx.Err().Error(), Err: ctx.Err()}
			}
			apiErr = &Error{Code: "NETWORK", Message: err.Error(), Err: err}
		} else if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return decodeBody(body, out)
		} else {
			apiErr = errorFromResponse(resp, body)
		}
		if !apiErr.Retryable() || attempt >= maxRetries {
			return apiErr
		}
		if err := c.sleep(ctx, retryDelay(resp, attempt, c.now())); err != nil {
			return &Error{Code: "NETWORK", Message: err.Error(), Err: err}
		}
	}
}

func (c *Client) attempt(ctx context.Context, method, endpoint string, headers http.Header, payload []byte, timeout time.Duration) (*http.Response, []byte, error) {
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	var reader io.Reader
	if payload != nil {
		reader = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}
	req.Header = headers.Clone()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return resp, body, nil
}

// adjustBody, gövdeye eksikse app_id ekler ve gövdedeki idempotency_key'i başlık için okur.
func adjustBody(raw []byte, addAppID bool, appID, key string) ([]byte, string, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil || fields == nil {
		// Nesne olmayan gövdeler olduğu gibi gönderilir.
		return raw, key, nil
	}
	if key == "" {
		if v, ok := fields["idempotency_key"]; ok {
			var s string
			if json.Unmarshal(v, &s) == nil {
				key = s
			}
		}
	}
	if addAppID {
		var existing string
		if v, ok := fields["app_id"]; ok {
			_ = json.Unmarshal(v, &existing)
		}
		if existing == "" {
			encoded, _ := json.Marshal(appID)
			fields["app_id"] = encoded
			out, err := json.Marshal(fields)
			if err != nil {
				return nil, "", err
			}
			return out, key, nil
		}
	}
	return raw, key, nil
}

func encodeMultipart(fields [][2]string, files []formFile) ([]byte, string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	for _, f := range fields {
		if err := w.WriteField(f[0], f[1]); err != nil {
			return nil, "", err
		}
	}
	for _, f := range files {
		name := f.file.Name
		if name == "" {
			name = "upload"
		}
		ct := f.file.ContentType
		if ct == "" {
			ct = "application/octet-stream"
		}
		h := textproto.MIMEHeader{}
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name=%q; filename=%q`, f.field, name))
		h.Set("Content-Type", ct)
		part, err := w.CreatePart(h)
		if err != nil {
			return nil, "", err
		}
		if f.file.Reader != nil {
			if _, err := io.Copy(part, f.file.Reader); err != nil {
				return nil, "", fmt.Errorf("herosignal: dosya okunamadı: %w", err)
			}
		}
	}
	if err := w.Close(); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), w.FormDataContentType(), nil
}

func decodeBody(body []byte, out any) error {
	if out == nil || len(body) == 0 {
		return nil
	}
	switch o := out.(type) {
	case *[]byte:
		*o = append([]byte(nil), body...)
		return nil
	case *string:
		*o = string(body)
		return nil
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("herosignal: yanıt çözülemedi: %w", err)
	}
	return nil
}

func errorFromResponse(resp *http.Response, body []byte) *Error {
	var envelope struct {
		Error *struct {
			Code      string         `json:"code"`
			Message   string         `json:"message"`
			RequestID string         `json:"request_id"`
			Details   map[string]any `json:"details"`
		} `json:"error"`
	}
	_ = json.Unmarshal(body, &envelope)
	e := &Error{StatusCode: resp.StatusCode}
	if envelope.Error != nil {
		e.Code = envelope.Error.Code
		e.Message = envelope.Error.Message
		e.RequestID = envelope.Error.RequestID
		e.Details = envelope.Error.Details
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

// retryDelay, Retry-After (saniye veya HTTP tarihi) varsa ona uyar; yoksa 300 ms / 1 sn / 3 sn.
func retryDelay(resp *http.Response, attempt int, now time.Time) time.Duration {
	if resp != nil {
		if h := strings.TrimSpace(resp.Header.Get("Retry-After")); h != "" {
			if secs, err := strconv.ParseFloat(h, 64); err == nil {
				if secs > 0 {
					return time.Duration(secs * float64(time.Second))
				}
			} else if t, err := http.ParseTime(h); err == nil {
				if d := t.Sub(now); d > 0 {
					return d
				}
			}
		}
	}
	if attempt >= len(backoff) {
		attempt = len(backoff) - 1
	}
	return backoff[attempt]
}

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

// newUUID, rastgele (v4) UUID üretir.
func newUUID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic("herosignal: rastgele sayı üretilemedi: " + err.Error())
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	h := hex.EncodeToString(b[:])
	return h[0:8] + "-" + h[8:12] + "-" + h[12:16] + "-" + h[16:20] + "-" + h[20:32]
}
