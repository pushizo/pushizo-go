package herosignal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

type recorded struct {
	Method  string
	Path    string
	RawPath string
	Query   string
	Header  http.Header
	Body    []byte
}

type fakeResponse struct {
	status  int
	body    string
	headers map[string]string
	netErr  bool
}

// scriptTransport, sıradaki yanıt ağ hatasıysa sunucuya gitmeden hata döndürür (bağlantı kopması benzetimi).
type scriptTransport struct {
	mu        *sync.Mutex
	responses *[]fakeResponse
}

func (s scriptTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	s.mu.Lock()
	if len(*s.responses) > 0 && (*s.responses)[0].netErr {
		*s.responses = (*s.responses)[1:]
		s.mu.Unlock()
		return nil, errors.New("bağlantı koptu")
	}
	s.mu.Unlock()
	return http.DefaultTransport.RoundTrip(r)
}

var transports = map[*httptest.Server]scriptTransport{}

// fakeServer, sıralı yanıtlar döndüren ve istekleri kaydeden test sunucusudur.
func fakeServer(t *testing.T, responses ...fakeResponse) (*httptest.Server, *[]recorded) {
	t.Helper()
	var mu sync.Mutex
	var calls []recorded
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		body, _ := io.ReadAll(r.Body)
		calls = append(calls, recorded{r.Method, r.URL.Path, r.URL.EscapedPath(), r.URL.RawQuery, r.Header.Clone(), body})
		if len(responses) == 0 {
			t.Errorf("beklenmeyen istek: %s %s", r.Method, r.URL)
			w.WriteHeader(500)
			return
		}
		res := responses[0]
		responses = responses[1:]
		for k, v := range res.headers {
			w.Header().Set(k, v)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.status)
		io.WriteString(w, res.body)
	}))
	t.Cleanup(srv.Close)
	transports[srv] = scriptTransport{mu: &mu, responses: &responses}
	return srv, &calls
}

func newTestClient(t *testing.T, srv *httptest.Server, slept *[]time.Duration, opts ...ClientOption) *Client {
	t.Helper()
	all := append([]ClientOption{WithBaseURL(srv.URL + "/"), WithAppID("app_1"), WithHTTPClient(&http.Client{Transport: transports[srv]})}, opts...)
	c, err := NewClient("hs_live_test", all...)
	if err != nil {
		t.Fatal(err)
	}
	c.sleep = func(_ context.Context, d time.Duration) error {
		if slept != nil {
			*slept = append(*slept, d)
		}
		return nil
	}
	return c
}

var uuidRe = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

func TestCreateNotificationAuthAppAndIdempotency(t *testing.T) {
	srv, calls := fakeServer(t, fakeResponse{status: 202, body: `{"id":"msg_1","status":"queued"}`})
	hs := newTestClient(t, srv, nil)
	res, err := hs.Notifications.Create(context.Background(), &CreateNotificationDto{
		Target:   CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {"u1"}}},
		Contents: map[string]string{"tr": "Merhaba"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.ID != "msg_1" {
		t.Fatalf("id = %q", res.ID)
	}
	c := (*calls)[0]
	if c.Method != "POST" || c.Path != "/v1/notifications" {
		t.Fatalf("%s %s", c.Method, c.Path)
	}
	if got := c.Header.Get("Authorization"); got != "Bearer hs_live_test" {
		t.Fatalf("authorization = %q", got)
	}
	if !strings.HasPrefix(c.Header.Get("User-Agent"), "herosignal-go/") {
		t.Fatal("user-agent")
	}
	if !uuidRe.MatchString(c.Header.Get("Idempotency-Key")) {
		t.Fatalf("idempotency key = %q", c.Header.Get("Idempotency-Key"))
	}
	var sent map[string]any
	json.Unmarshal(c.Body, &sent)
	if sent["app_id"] != "app_1" || sent["contents"].(map[string]any)["tr"] != "Merhaba" {
		t.Fatalf("gövde = %s", c.Body)
	}
}

func TestExplicitBodyIdempotencyKeyAndForApp(t *testing.T) {
	srv, calls := fakeServer(t,
		fakeResponse{status: 202, body: `{"id":"m"}`},
		fakeResponse{status: 202, body: `{"id":"m"}`},
	)
	hs := newTestClient(t, srv, nil)
	ctx := context.Background()
	body := &CreateNotificationDto{Contents: map[string]string{"en": "x"}}
	if _, err := hs.Notifications.Create(ctx, body, WithIdempotencyKey("order-1")); err != nil {
		t.Fatal(err)
	}
	if body.AppID != nil {
		t.Fatal("kullanıcının gövdesi değiştirilmemeli")
	}
	if _, err := hs.Notifications.Create(ctx, &CreateNotificationDto{Contents: map[string]string{"en": "x"}, IdempotencyKey: String("order-2")}, ForApp("app_2")); err != nil {
		t.Fatal(err)
	}
	if (*calls)[0].Header.Get("Idempotency-Key") != "order-1" || (*calls)[1].Header.Get("Idempotency-Key") != "order-2" {
		t.Fatal("idempotency anahtarları")
	}
	var sent map[string]any
	json.Unmarshal((*calls)[1].Body, &sent)
	if sent["app_id"] != "app_2" {
		t.Fatalf("app_id = %v", sent["app_id"])
	}
}

func TestNoIdempotencyKeyWhenDisabledOrNotSupported(t *testing.T) {
	srv, calls := fakeServer(t,
		fakeResponse{status: 202, body: `{"id":"m"}`},
		fakeResponse{status: 200, body: `{"accepted":1}`},
	)
	hs := newTestClient(t, srv, nil, WithAutoIdempotencyKey(false))
	ctx := context.Background()
	hs.Notifications.Create(ctx, &CreateNotificationDto{Contents: map[string]string{"en": "x"}})
	hs.Events.Send(ctx, &APIEventsDto{Events: []APIEventsDtoEventsItem{{Name: "purchase"}}})
	for _, c := range *calls {
		if c.Header.Get("Idempotency-Key") != "" {
			t.Fatalf("%s beklenmeyen idempotency anahtarı", c.Path)
		}
	}
}

func TestQueryAndPathEncoding(t *testing.T) {
	srv, calls := fakeServer(t,
		fakeResponse{status: 200, body: `{"data":[],"has_more":false,"next_cursor":null}`},
		fakeResponse{status: 200, body: `{"id":"usr_1"}`},
	)
	hs := newTestClient(t, srv, nil)
	ctx := context.Background()
	if _, err := hs.Notifications.List(ctx, &NotificationsListParams{Limit: Int64(10), Status: String("completed")}); err != nil {
		t.Fatal(err)
	}
	if q := (*calls)[0].Query; q != "app_id=app_1&limit=10&status=completed" {
		t.Fatalf("sorgu = %q", q)
	}
	if _, err := hs.Users.GetByAlias(ctx, "crm id", "a/b", ForApp("app 2")); err != nil {
		t.Fatal(err)
	}
	if p := (*calls)[1].RawPath; p != "/v1/apps/app%202/users/by/crm%20id/a%2Fb" {
		t.Fatalf("yol = %q", p)
	}
}

func TestMissingAppID(t *testing.T) {
	hs, _ := NewClient("k")
	if _, err := hs.Apps.Get(context.Background()); err == nil || !strings.Contains(err.Error(), "app") {
		t.Fatalf("hata bekleniyordu: %v", err)
	}
}

func TestRetryAfterBackoffAndStableIdempotencyKey(t *testing.T) {
	srv, calls := fakeServer(t,
		fakeResponse{status: 429, body: `{"error":{"code":"HS_RATE_LIMITED","message":"yavaş"}}`, headers: map[string]string{"Retry-After": "2"}},
		fakeResponse{status: 503, body: `{"error":{"code":"HS_SERVICE_UNAVAILABLE","message":"x"}}`},
		fakeResponse{netErr: true},
		fakeResponse{status: 202, body: `{"id":"msg_1"}`},
	)
	var slept []time.Duration
	hs := newTestClient(t, srv, &slept)
	res, err := hs.Notifications.Create(context.Background(), &CreateNotificationDto{Contents: map[string]string{"en": "x"}})
	if err != nil {
		t.Fatal(err)
	}
	if res.ID != "msg_1" {
		t.Fatal(res.ID)
	}
	want := []time.Duration{2 * time.Second, time.Second, 3 * time.Second}
	if len(slept) != 3 || slept[0] != want[0] || slept[1] != want[1] || slept[2] != want[2] {
		t.Fatalf("beklemeler = %v", slept)
	}
	key := (*calls)[0].Header.Get("Idempotency-Key")
	for _, c := range *calls {
		if c.Header.Get("Idempotency-Key") != key {
			t.Fatal("yeniden denemelerde idempotency anahtarı değişti")
		}
	}
}

func TestTypedErrors(t *testing.T) {
	srv, calls := fakeServer(t,
		fakeResponse{status: 404, body: `{"error":{"code":"HS_NOT_FOUND","message":"yok","request_id":"req_1","details":{"id":"x"}}}`},
		fakeResponse{status: 422, body: `{"error":{"code":"HS_NO_RECIPIENTS","message":"boş"}}`},
		fakeResponse{status: 401, body: `{"error":{"code":"HS_UNAUTHENTICATED","message":"anahtar"}}`},
	)
	var slept []time.Duration
	hs := newTestClient(t, srv, &slept)
	ctx := context.Background()
	_, err := hs.Notifications.Get(ctx, "msg_x")
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("*Error bekleniyordu: %v", err)
	}
	if apiErr.StatusCode != 404 || apiErr.Code != CodeNotFound || apiErr.RequestID != "req_1" || apiErr.Details["id"] != "x" {
		t.Fatalf("%+v", apiErr)
	}
	if !errors.Is(err, ErrNotFound) || errors.Is(err, ErrServer) || apiErr.Retryable() {
		t.Fatal("kategori")
	}
	_, err = hs.Segments.Get(ctx, "seg_1")
	if !errors.Is(err, ErrInvalidRequest) {
		t.Fatal(err)
	}
	_, err = hs.Apps.List(ctx)
	if !errors.Is(err, ErrAuthentication) {
		t.Fatal(err)
	}
	if len(slept) != 0 || len(*calls) != 3 {
		t.Fatal("4xx yeniden denenmemeli")
	}
}

func TestRetriesExhausted(t *testing.T) {
	srv, calls := fakeServer(t,
		fakeResponse{status: 500, body: `{}`}, fakeResponse{status: 500, body: `{}`},
		fakeResponse{status: 500, body: `{}`}, fakeResponse{status: 500, body: `{}`},
	)
	var slept []time.Duration
	hs := newTestClient(t, srv, &slept)
	_, err := hs.Apps.List(context.Background())
	var apiErr *Error
	if !errors.As(err, &apiErr) || apiErr.Code != "HTTP_500" || !errors.Is(err, ErrServer) {
		t.Fatal(err)
	}
	if len(*calls) != 4 || len(slept) != 3 || slept[0] != 300*time.Millisecond {
		t.Fatalf("çağrı %d, bekleme %v", len(*calls), slept)
	}

	srv2, _ := fakeServer(t, fakeResponse{netErr: true}, fakeResponse{netErr: true})
	hs2 := newTestClient(t, srv2, nil, WithMaxRetries(1))
	_, err = hs2.Apps.List(context.Background())
	if !errors.Is(err, ErrNetwork) {
		t.Fatalf("ağ hatası bekleniyordu: %v", err)
	}
}

func TestPaginationIteratesAllPages(t *testing.T) {
	srv, calls := fakeServer(t,
		fakeResponse{status: 200, body: `{"data":[{"id":"s1"},{"id":"s2"}],"next_cursor":"c2","has_more":true}`},
		fakeResponse{status: 200, body: `{"data":[{"id":"s3"}],"next_cursor":null,"has_more":false}`},
	)
	hs := newTestClient(t, srv, nil)
	var ids []string
	for seg, err := range hs.Segments.ListAll(context.Background(), &SegmentsListParams{Limit: Int64(2)}) {
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, seg.ID)
	}
	if strings.Join(ids, ",") != "s1,s2,s3" {
		t.Fatalf("ids = %v", ids)
	}
	if (*calls)[0].Query != "limit=2" || (*calls)[1].Query != "cursor=c2&limit=2" {
		t.Fatalf("sorgular: %q, %q", (*calls)[0].Query, (*calls)[1].Query)
	}
}

func TestPaginationStopsOnError(t *testing.T) {
	srv, _ := fakeServer(t, fakeResponse{status: 403, body: `{"error":{"code":"HS_FORBIDDEN","message":"x"}}`})
	hs := newTestClient(t, srv, nil)
	n := 0
	for _, err := range hs.Segments.ListAll(context.Background(), nil) {
		n++
		if !errors.Is(err, ErrPermissionDenied) {
			t.Fatal(err)
		}
	}
	if n != 1 {
		t.Fatal(n)
	}
}

func TestMultipartImport(t *testing.T) {
	srv, calls := fakeServer(t, fakeResponse{status: 202, body: `{"id":"imp_1"}`})
	hs := newTestClient(t, srv, nil)
	_, err := hs.Imports.Create(context.Background(), &ImportsCreateForm{
		Type: "csv_users",
		File: &FileUpload{Filename: "users.csv", ContentType: "text/csv", Content: []byte("external_id\nu1\n")},
	})
	if err != nil {
		t.Fatal(err)
	}
	c := (*calls)[0]
	if !strings.HasPrefix(c.Header.Get("Content-Type"), "multipart/form-data; boundary=") {
		t.Fatal(c.Header.Get("Content-Type"))
	}
	body := string(c.Body)
	if !strings.Contains(body, `name="type"`) || !strings.Contains(body, "csv_users") || !strings.Contains(body, `filename="users.csv"`) || !strings.Contains(body, "external_id\nu1\n") {
		t.Fatal(body)
	}
}

func TestRawRequest(t *testing.T) {
	srv, calls := fakeServer(t, fakeResponse{status: 200, body: "day,sent\n"})
	hs := newTestClient(t, srv, nil)
	var csv string
	if err := hs.Request(context.Background(), "get", "/v1/apps/{appId}/analytics/messages", map[string][]string{"format": {"csv"}}, nil, &csv); err != nil {
		t.Fatal(err)
	}
	if csv != "day,sent\n" || (*calls)[0].Path != "/v1/apps/app_1/analytics/messages" {
		t.Fatal(csv)
	}
}

func TestClientStringHidesAPIKey(t *testing.T) {
	hs, _ := NewClient("hs_live_gizli", WithAppID("app_1"))
	for _, s := range []string{fmt.Sprintf("%v", hs), fmt.Sprintf("%+v", hs), fmt.Sprintf("%#v", hs)} {
		if strings.Contains(s, "hs_live_gizli") {
			t.Fatalf("API anahtarı sızdı: %s", s)
		}
	}
}
