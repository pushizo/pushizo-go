package herosignal

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

// recorded, sahte sunucunun aldığı isteğin kopyasıdır.
type recorded struct {
	Method string
	Path   string
	Raw    string
	Query  map[string][]string
	Header http.Header
	Body   []byte
}

type reply struct {
	status  int
	body    string
	headers map[string]string
	// hang, true ise yanıt verilmeden istek bekletilir (zaman aşımı testi).
	hang bool
}

type fakeServer struct {
	t        *testing.T
	mu       sync.Mutex
	replies  []reply
	requests []recorded
	srv      *httptest.Server
}

func newFakeServer(t *testing.T, replies ...reply) *fakeServer {
	f := &fakeServer{t: t, replies: replies}
	f.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		f.mu.Lock()
		f.requests = append(f.requests, recorded{
			Method: r.Method, Path: r.URL.Path, Raw: r.URL.EscapedPath(), Query: r.URL.Query(), Header: r.Header.Clone(), Body: body,
		})
		if len(f.replies) == 0 {
			f.mu.Unlock()
			t.Errorf("beklenmeyen istek: %s %s", r.Method, r.URL)
			w.WriteHeader(599)
			return
		}
		rep := f.replies[0]
		f.replies = f.replies[1:]
		f.mu.Unlock()
		if rep.hang {
			<-r.Context().Done()
			return
		}
		for k, v := range rep.headers {
			w.Header().Set(k, v)
		}
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "application/json")
		}
		w.WriteHeader(rep.status)
		_, _ = w.Write([]byte(rep.body))
	}))
	t.Cleanup(f.srv.Close)
	return f
}

func (f *fakeServer) client(t *testing.T, sleeps *[]time.Duration, opts ...ClientOption) *Client {
	all := append([]ClientOption{WithBaseURL(f.srv.URL), WithAppID("app_default")}, opts...)
	c, err := NewClient("hs_live_test", all...)
	if err != nil {
		t.Fatal(err)
	}
	c.sleep = func(_ context.Context, d time.Duration) error {
		if sleeps != nil {
			*sleeps = append(*sleeps, d)
		}
		return nil
	}
	return c
}

func notification() *CreateNotificationDto {
	return &CreateNotificationDto{
		Target:   CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {"u1"}}},
		Contents: map[string]string{"en": "Hi"},
	}
}

func decode(t *testing.T, raw []byte) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatalf("gövde JSON değil: %v (%s)", err, raw)
	}
	return m
}

func TestNewClientRequiresAPIKey(t *testing.T) {
	if _, err := NewClient(""); err == nil {
		t.Fatal("hata bekleniyordu")
	}
}

func TestCreateNotificationHeadersBodyAndAutoIdempotency(t *testing.T) {
	f := newFakeServer(t,
		reply{status: 503, body: `{"error":{"code":"UNAVAILABLE","message":"x"}}`},
		reply{status: 202, body: `{"id":"msg_1"}`},
	)
	var sleeps []time.Duration
	c := f.client(t, &sleeps)
	res, err := c.Notifications.Create(context.Background(), notification())
	if err != nil {
		t.Fatal(err)
	}
	if res.ID != "msg_1" {
		t.Fatalf("yanıt çözülmedi: %+v", res)
	}
	if len(f.requests) != 2 {
		t.Fatalf("2 istek bekleniyordu, %d", len(f.requests))
	}
	r := f.requests[0]
	if r.Method != "POST" || r.Path != "/v1/notifications" {
		t.Fatalf("yanlış istek: %s %s", r.Method, r.Path)
	}
	if got := r.Header.Get("Authorization"); got != "Bearer hs_live_test" {
		t.Fatalf("Authorization: %q", got)
	}
	if !strings.HasPrefix(r.Header.Get("User-Agent"), "herosignal-go/") || r.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("başlıklar: %v", r.Header)
	}
	body := decode(t, r.Body)
	if body["app_id"] != "app_default" {
		t.Fatalf("app_id eklenmedi: %v", body)
	}
	k1, k2 := r.Header.Get("Idempotency-Key"), f.requests[1].Header.Get("Idempotency-Key")
	if len(k1) != 36 || k1 != k2 {
		t.Fatalf("idempotency anahtarı denemelerde aynı olmalı: %q %q", k1, k2)
	}
	if len(sleeps) != 1 || sleeps[0] != 300*time.Millisecond {
		t.Fatalf("bekleme: %v", sleeps)
	}
}

func TestIdempotencyKeyExplicitBodyAndDisabled(t *testing.T) {
	f := newFakeServer(t,
		reply{status: 202, body: `{}`},
		reply{status: 202, body: `{}`},
		reply{status: 202, body: `{}`},
		reply{status: 201, body: `{}`},
	)
	c := f.client(t, nil)
	ctx := context.Background()
	if _, err := c.Notifications.Create(ctx, notification(), WithIdempotencyKey("order-1")); err != nil {
		t.Fatal(err)
	}
	n := notification()
	n.IdempotencyKey = Ptr("order-2")
	if _, err := c.Notifications.Create(ctx, n); err != nil {
		t.Fatal(err)
	}
	off := f.client(t, nil, WithAutoIdempotencyKey(false))
	if _, err := off.Notifications.Create(ctx, notification()); err != nil {
		t.Fatal(err)
	}
	if _, err := c.Suppressions.Create(ctx, &CreateSuppressionDto{Channel: "email", Value: "a@b.c"}); err != nil {
		t.Fatal(err)
	}
	if got := f.requests[0].Header.Get("Idempotency-Key"); got != "order-1" {
		t.Fatalf("açık anahtar: %q", got)
	}
	if got := f.requests[1].Header.Get("Idempotency-Key"); got != "order-2" {
		t.Fatalf("gövde anahtarı: %q", got)
	}
	if got := f.requests[2].Header.Get("Idempotency-Key"); got != "" {
		t.Fatalf("otomatik anahtar kapalıyken başlık olmamalı: %q", got)
	}
	if got := f.requests[3].Header.Get("Idempotency-Key"); got != "" {
		t.Fatalf("idempotent olmayan uçta başlık olmamalı: %q", got)
	}
}

func TestAppIDPathQueryAndOverride(t *testing.T) {
	f := newFakeServer(t,
		reply{status: 200, body: `{"id":"seg_1"}`},
		reply{status: 200, body: `{"data":[],"has_more":false}`},
	)
	c := f.client(t, nil)
	ctx := context.Background()
	if _, err := c.Segments.Get(ctx, "seg/1", WithRequestAppID("app_other")); err != nil {
		t.Fatal(err)
	}
	if f.requests[0].Raw != "/v1/apps/app_other/segments/seg%2F1" {
		t.Fatalf("yol: %s", f.requests[0].Raw)
	}
	if _, err := c.Notifications.List(ctx, &NotificationsListParams{Limit: Ptr[int64](10), Status: Ptr("completed")}); err != nil {
		t.Fatal(err)
	}
	q := f.requests[1].Query
	if q["app_id"][0] != "app_default" || q["limit"][0] != "10" || q["status"][0] != "completed" {
		t.Fatalf("sorgu: %v", q)
	}
	if _, ok := q["cursor"]; ok {
		t.Fatal("boş cursor gönderilmemeli")
	}
}

func TestMissingAppIDAndEmptyPathParam(t *testing.T) {
	c, _ := NewClient("k", WithBaseURL("http://127.0.0.1:1"))
	if _, err := c.Segments.List(context.Background(), nil); err == nil || !strings.Contains(err.Error(), "app ID") {
		t.Fatalf("app ID hatası bekleniyordu: %v", err)
	}
	c2, _ := NewClient("k", WithAppID("a"), WithBaseURL("http://127.0.0.1:1"))
	if _, err := c2.Segments.Get(context.Background(), ""); err == nil {
		t.Fatal("boş yol parametresi hatası bekleniyordu")
	}
}

func TestRetryBackoffAndRetryAfter(t *testing.T) {
	f := newFakeServer(t,
		reply{status: 500, body: `{}`},
		reply{status: 429, body: `{}`, headers: map[string]string{"Retry-After": "2"}},
		reply{status: 502, body: `{}`},
		reply{status: 200, body: `{"data":[]}`},
	)
	var sleeps []time.Duration
	c := f.client(t, &sleeps)
	if _, err := c.Apps.List(context.Background()); err != nil {
		t.Fatal(err)
	}
	want := []time.Duration{300 * time.Millisecond, 2 * time.Second, 3 * time.Second}
	if len(sleeps) != len(want) {
		t.Fatalf("bekleme: %v", sleeps)
	}
	for i := range want {
		if sleeps[i] != want[i] {
			t.Fatalf("bekleme %d: %v != %v", i, sleeps[i], want[i])
		}
	}
}

func TestRetryAfterHTTPDate(t *testing.T) {
	future := time.Date(2030, 1, 1, 0, 0, 10, 0, time.UTC)
	f := newFakeServer(t,
		reply{status: 429, body: `{}`, headers: map[string]string{"Retry-After": future.Format(http.TimeFormat)}},
		reply{status: 200, body: `{}`},
	)
	var sleeps []time.Duration
	c := f.client(t, &sleeps)
	c.now = func() time.Time { return time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC) }
	if _, err := c.Apps.List(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(sleeps) != 1 || sleeps[0] != 10*time.Second {
		t.Fatalf("bekleme: %v", sleeps)
	}
}

func TestRetriesExhaustedAndErrorMapping(t *testing.T) {
	f := newFakeServer(t,
		reply{status: 503, body: `{"error":{"code":"UNAVAILABLE","message":"down","request_id":"req_1"}}`},
		reply{status: 503, body: `{"error":{"code":"UNAVAILABLE","message":"down","request_id":"req_2"}}`},
	)
	var sleeps []time.Duration
	c := f.client(t, &sleeps, WithMaxRetries(1))
	_, err := c.Apps.List(context.Background())
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("*Error bekleniyordu: %v", err)
	}
	if apiErr.StatusCode != 503 || apiErr.Code != "UNAVAILABLE" || apiErr.RequestID != "req_2" || !apiErr.Retryable() {
		t.Fatalf("hata: %+v", apiErr)
	}
	if len(f.requests) != 2 || len(sleeps) != 1 {
		t.Fatalf("istek/bekleme sayısı: %d/%d", len(f.requests), len(sleeps))
	}
}

func TestClientErrorNotRetried(t *testing.T) {
	f := newFakeServer(t,
		reply{status: 422, body: `{"error":{"code":"VALIDATION_FAILED","message":"bad","request_id":"req_1","details":{"field":"x"}}}`},
		reply{status: 404, body: `not found`, headers: map[string]string{"Content-Type": "text/plain", "X-Request-Id": "req_h"}},
	)
	c := f.client(t, nil)
	_, err := c.Apps.List(context.Background())
	var apiErr *Error
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 422 || apiErr.Code != "VALIDATION_FAILED" || apiErr.Message != "bad" ||
		apiErr.RequestID != "req_1" || apiErr.Details["field"] != "x" || apiErr.Retryable() {
		t.Fatalf("hata eşlemesi: %+v", apiErr)
	}
	if !strings.Contains(err.Error(), "req_1") {
		t.Fatalf("hata metni: %s", err)
	}
	_, err = c.Apps.Get(context.Background())
	if !errors.As(err, &apiErr) || apiErr.Code != "HTTP_404" || apiErr.RequestID != "req_h" {
		t.Fatalf("zarfsız hata: %+v", apiErr)
	}
	if len(f.requests) != 2 {
		t.Fatalf("4xx yeniden denenmemeli: %d", len(f.requests))
	}
}

func TestNetworkErrorRetriedThenFails(t *testing.T) {
	var sleeps []time.Duration
	c, _ := NewClient("k", WithAppID("a"), WithBaseURL("http://127.0.0.1:1"), WithMaxRetries(2))
	c.sleep = func(_ context.Context, d time.Duration) error { sleeps = append(sleeps, d); return nil }
	_, err := c.Apps.List(context.Background())
	var apiErr *Error
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 0 || apiErr.Code != "NETWORK" || apiErr.Err == nil {
		t.Fatalf("ağ hatası: %+v", err)
	}
	if len(sleeps) != 2 {
		t.Fatalf("bekleme: %v", sleeps)
	}
}

func TestTimeoutIsRetried(t *testing.T) {
	f := newFakeServer(t, reply{hang: true}, reply{status: 200, body: `{}`})
	var sleeps []time.Duration
	c := f.client(t, &sleeps, WithTimeout(100*time.Millisecond))
	if _, err := c.Apps.List(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(sleeps) != 1 {
		t.Fatalf("zaman aşımı yeniden denenmeli: %v", sleeps)
	}
}

func TestContextCancelStops(t *testing.T) {
	f := newFakeServer(t, reply{status: 503, body: `{}`})
	c := f.client(t, nil)
	c.sleep = sleepContext
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := c.Apps.List(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("context.Canceled bekleniyordu: %v", err)
	}
}

func TestRequestOptionsHeaderAndRetries(t *testing.T) {
	f := newFakeServer(t, reply{status: 500, body: `{}`})
	c := f.client(t, nil)
	_, err := c.Apps.List(context.Background(), WithRequestMaxRetries(0), WithHeader("X-Trace", "t1"))
	if err == nil || len(f.requests) != 1 || f.requests[0].Header.Get("X-Trace") != "t1" {
		t.Fatalf("istek seçenekleri uygulanmadı: %v %d", err, len(f.requests))
	}
}

func TestAutoPaging(t *testing.T) {
	f := newFakeServer(t,
		reply{status: 200, body: `{"data":[{"id":"s1"},{"id":"s2"}],"next_cursor":"c2","has_more":true}`},
		reply{status: 200, body: `{"data":[{"id":"s3"}],"next_cursor":null,"has_more":false}`},
	)
	c := f.client(t, nil)
	items, err := c.Segments.ListAutoPaging(context.Background(), &SegmentsListParams{Limit: Ptr[int64](2)}).Collect()
	if err != nil {
		t.Fatal(err)
	}
	var ids []string
	for _, s := range items {
		ids = append(ids, s.ID)
	}
	if strings.Join(ids, ",") != "s1,s2,s3" {
		t.Fatalf("öğeler: %v", ids)
	}
	if _, ok := f.requests[0].Query["cursor"]; ok {
		t.Fatal("ilk sayfada cursor olmamalı")
	}
	if f.requests[1].Query["cursor"][0] != "c2" || f.requests[1].Query["limit"][0] != "2" {
		t.Fatalf("ikinci sayfa sorgusu: %v", f.requests[1].Query)
	}
}

func TestAutoPagingStopsOnError(t *testing.T) {
	f := newFakeServer(t, reply{status: 403, body: `{"error":{"code":"FORBIDDEN","message":"no"}}`})
	c := f.client(t, nil)
	it := c.Segments.ListAutoPaging(context.Background(), nil)
	if it.Next() {
		t.Fatal("öğe beklenmiyordu")
	}
	var apiErr *Error
	if !errors.As(it.Err(), &apiErr) || apiErr.StatusCode != 403 {
		t.Fatalf("hata: %v", it.Err())
	}
}

func TestMultipartUpload(t *testing.T) {
	f := newFakeServer(t, reply{status: 201, body: `{"id":"imp_1"}`})
	c := f.client(t, nil)
	_, err := c.Imports.Create(context.Background(), &ImportsCreateParams{
		Type: "csv_users",
		File: &FileUpload{Name: "users.csv", ContentType: "text/csv", Reader: strings.NewReader("external_id\nu1\n")},
	})
	if err != nil {
		t.Fatal(err)
	}
	r := f.requests[0]
	if r.Path != "/v1/apps/app_default/imports" || !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		t.Fatalf("multipart isteği: %s %s", r.Path, r.Header.Get("Content-Type"))
	}
	raw := string(r.Body)
	if !strings.Contains(raw, `name="type"`) || !strings.Contains(raw, "csv_users") || !strings.Contains(raw, `filename="users.csv"`) {
		t.Fatalf("multipart gövdesi: %s", raw)
	}
}

func TestDoRawAndQueryEncoding(t *testing.T) {
	f := newFakeServer(t,
		reply{status: 200, body: "a,b\n1,2\n", headers: map[string]string{"Content-Type": "text/csv"}},
		reply{status: 204},
	)
	c := f.client(t, nil)
	var csv string
	q := map[string][]string{"format": {"csv"}, "b": {"1", "2"}}
	if err := c.Do(context.Background(), "GET", "/v1/custom", q, nil, &csv); err != nil {
		t.Fatal(err)
	}
	if csv != "a,b\n1,2\n" || len(f.requests[0].Query["b"]) != 2 {
		t.Fatalf("ham yanıt/sorgu: %q %v", csv, f.requests[0].Query)
	}
	var out map[string]any
	if err := c.Do(context.Background(), "DELETE", "/v1/x", nil, nil, &out); err != nil || out != nil {
		t.Fatalf("boş yanıt: %v %v", err, out)
	}
}

func TestQueryValues(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{Ptr(true), "true"},
		{Ptr[int64](5), "5"},
		{Ptr(1.5), "1.5"},
		{[]string{"a", "b"}, "a|b"},
		{map[string]int{"x": 1}, `{"x":1}`},
		{(*string)(nil), ""},
	}
	for _, tc := range cases {
		if got := strings.Join(queryValues(tc.in), "|"); got != tc.want {
			t.Fatalf("%v: %q != %q", tc.in, got, tc.want)
		}
	}
}

func TestIdentityHash(t *testing.T) {
	// HMAC-SHA256("secret", "user-1") — Node/PHP/Python identityHash ile aynı (Python hmac ile hesaplandı).
	const want = "1d16fd7e96e8a9681f283b8a251822ffdc71e2a5bda0c5267a044874fa21b82b"
	if got := IdentityHash("secret", "user-1"); got != want {
		t.Fatalf("identity hash: %s", got)
	}
}

func TestNewUUIDFormat(t *testing.T) {
	u := newUUID()
	if len(u) != 36 || u[14] != '4' || strings.Count(u, "-") != 4 {
		t.Fatalf("uuid: %s", u)
	}
}
