package pushizo

import (
	"errors"
	"fmt"
	"time"
)

// API hata zarfındaki error.code değerleri (HS_ önekli katalog).
const (
	CodeValidationError       = "HS_VALIDATION_ERROR"
	CodeBadRequest            = "HS_BAD_REQUEST"
	CodeUnauthenticated       = "HS_UNAUTHENTICATED"
	CodeForbidden             = "HS_FORBIDDEN"
	CodeTenantMismatch        = "HS_TENANT_MISMATCH"
	CodeNotFound              = "HS_NOT_FOUND"
	CodeConflict              = "HS_CONFLICT"
	CodeIdempotencyConflict   = "HS_IDEMPOTENCY_CONFLICT"
	CodeIdempotencyInProgress = "HS_IDEMPOTENCY_IN_PROGRESS"
	CodeNoRecipients          = "HS_NO_RECIPIENTS"
	CodeChannelNotConfigured  = "HS_CHANNEL_NOT_CONFIGURED"
	CodeContentInvalid        = "HS_CONTENT_INVALID"
	CodeSubscriptionInvalid   = "HS_SUBSCRIPTION_INVALID"
	CodeUnprocessable         = "HS_UNPROCESSABLE"
	CodeRateLimited           = "HS_RATE_LIMITED"
	CodeQuotaExceeded         = "HS_QUOTA_EXCEEDED"
	CodeInternal              = "HS_INTERNAL"
	CodeNotImplemented        = "HS_NOT_IMPLEMENTED"
	CodeProviderUnavailable   = "HS_PROVIDER_UNAVAILABLE"
	CodeServiceUnavailable    = "HS_SERVICE_UNAVAILABLE"
	// CodeNetwork istemci tarafıdır: bağlantı kurulamadı veya zaman aşımı (StatusCode 0).
	CodeNetwork = "NETWORK"
)

// HTTP durum sınıflarına karşılık gelen hata kategorileri; errors.Is ile kullanılır:
//
//	if errors.Is(err, pushizo.ErrNotFound) { ... }
var (
	ErrNetwork          = errors.New("pushizo: ağ hatası")
	ErrInvalidRequest   = errors.New("pushizo: geçersiz istek (400/422)")
	ErrAuthentication   = errors.New("pushizo: kimlik doğrulama (401)")
	ErrPermissionDenied = errors.New("pushizo: yetki yok (403)")
	ErrNotFound         = errors.New("pushizo: bulunamadı (404)")
	ErrConflict         = errors.New("pushizo: çakışma (409)")
	ErrRateLimited      = errors.New("pushizo: hız sınırı (429)")
	ErrServer           = errors.New("pushizo: sunucu hatası (5xx)")
)

// Error, API hata zarfından ({"error": {"code", "message", "request_id", "details"}}) türetilir.
// errors.As ile alınır; ağ hatalarında StatusCode 0 ve Code "NETWORK" olur.
type Error struct {
	StatusCode int
	Code       string
	Message    string
	RequestID  string
	Details    map[string]any
	// RetryAfter, 429/503 yanıtındaki Retry-After başlığıdır (yoksa 0).
	RetryAfter time.Duration
	cause      error
}

func (e *Error) Error() string {
	s := fmt.Sprintf("pushizo: [%d %s] %s", e.StatusCode, e.Code, e.Message)
	if e.RequestID != "" {
		s += " (request_id=" + e.RequestID + ")"
	}
	return s
}

// Unwrap, ağ hatasının altta yatan nedenini döndürür.
func (e *Error) Unwrap() error { return e.cause }

// Retryable: 429, 5xx veya ağ hatası.
func (e *Error) Retryable() bool {
	return e.StatusCode == 0 || e.StatusCode == 429 || e.StatusCode >= 500
}

// Is, hata kategorisi karşılaştırması sağlar (ErrNotFound, ErrRateLimited …).
func (e *Error) Is(target error) bool {
	switch target {
	case ErrNetwork:
		return e.StatusCode == 0
	case ErrInvalidRequest:
		return e.StatusCode == 400 || e.StatusCode == 422
	case ErrAuthentication:
		return e.StatusCode == 401
	case ErrPermissionDenied:
		return e.StatusCode == 403
	case ErrNotFound:
		return e.StatusCode == 404
	case ErrConflict:
		return e.StatusCode == 409
	case ErrRateLimited:
		return e.StatusCode == 429
	case ErrServer:
		return e.StatusCode >= 500
	}
	return false
}
