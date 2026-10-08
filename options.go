package herosignal

import (
	"net/http"
	"time"
)

// ClientOption, NewClient ile istemci ayarlarını değiştirir.
type ClientOption func(*Client)

// WithAppID, çağrılarda WithAppID verilmezse kullanılacak varsayılan app kimliğini ayarlar.
func WithAppID(appID string) ClientOption {
	return func(c *Client) { c.appID = appID }
}

// WithBaseURL, API adresini ayarlar (yerel geliştirmede "http://localhost:4000").
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) { c.baseURL = trimSlash(baseURL) }
}

// WithHTTPClient, özel *http.Client kullanır (proxy, TLS, izleme).
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// WithTimeout, her deneme için istek zaman aşımını ayarlar (varsayılan 30 sn; 0 = yalnızca context).
func WithTimeout(d time.Duration) ClientOption {
	return func(c *Client) { c.timeout = d }
}

// WithMaxRetries, 429/5xx ve ağ hatalarında en fazla yeniden deneme sayısını ayarlar (varsayılan 3).
func WithMaxRetries(n int) ClientOption {
	return func(c *Client) {
		if n >= 0 {
			c.maxRetries = n
		}
	}
}

// WithAutoIdempotencyKey, Idempotency-Key destekleyen çağrılarda anahtar verilmezse UUID üretilmesini açar/kapatır
// (varsayılan açık).
func WithAutoIdempotencyKey(enabled bool) ClientOption {
	return func(c *Client) { c.autoIdempotencyKey = enabled }
}

// WithUserAgent, User-Agent başlığını değiştirir.
func WithUserAgent(ua string) ClientOption {
	return func(c *Client) {
		if ua != "" {
			c.userAgent = ua
		}
	}
}

// RequestOption, tek bir çağrının ayarlarını değiştirir.
type RequestOption func(*requestConfig)

type requestConfig struct {
	appID          string
	idempotencyKey string
	headers        http.Header
	maxRetries     *int
	timeout        *time.Duration
}

// WithRequestAppID, bu çağrı için app kimliğini istemcinin varsayılanı yerine kullanır.
func WithRequestAppID(appID string) RequestOption {
	return func(r *requestConfig) { r.appID = appID }
}

// WithIdempotencyKey, Idempotency-Key başlığını ayarlar (ör. "order-42"); aynı anahtarla tekrar gönderim
// sunucuda tek işlem olarak sayılır. Yeniden denemelerde aynı anahtar kullanılır.
func WithIdempotencyKey(key string) RequestOption {
	return func(r *requestConfig) { r.idempotencyKey = key }
}

// WithHeader, bu çağrıya ek HTTP başlığı ekler.
func WithHeader(name, value string) RequestOption {
	return func(r *requestConfig) {
		if r.headers == nil {
			r.headers = http.Header{}
		}
		r.headers.Set(name, value)
	}
}

// WithRequestMaxRetries, bu çağrı için en fazla yeniden deneme sayısını ayarlar.
func WithRequestMaxRetries(n int) RequestOption {
	return func(r *requestConfig) { r.maxRetries = &n }
}

// WithRequestTimeout, bu çağrı için deneme başına zaman aşımını ayarlar.
func WithRequestTimeout(d time.Duration) RequestOption {
	return func(r *requestConfig) { r.timeout = &d }
}
