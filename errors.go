package herosignal

import "fmt"

// Error, API hata zarfından ({"error": {"code", "message", "request_id", "details"}}) türetilen hatadır.
//
// Ağ hatalarında (bağlantı, zaman aşımı) StatusCode 0 ve Code "NETWORK" olur; asıl hata Unwrap ile alınır.
//
//	var apiErr *herosignal.Error
//	if errors.As(err, &apiErr) && apiErr.StatusCode == 404 { … }
type Error struct {
	// StatusCode HTTP durum kodu (ağ hatasında 0).
	StatusCode int
	// Code makine tarafından okunur hata kodu (ör. "VALIDATION_FAILED", "NETWORK").
	Code string
	// Message insan tarafından okunur açıklama.
	Message string
	// RequestID destek taleplerinde kullanılacak istek kimliği (zarf veya X-Request-Id başlığı).
	RequestID string
	// Details alan bazlı ek bilgiler (varsa).
	Details map[string]any
	// Err ağ hatalarında alttaki hata.
	Err error
}

// Error, hata metnini döndürür.
func (e *Error) Error() string {
	msg := fmt.Sprintf("herosignal: [%d %s] %s", e.StatusCode, e.Code, e.Message)
	if e.RequestID != "" {
		msg += " (request_id=" + e.RequestID + ")"
	}
	return msg
}

// Unwrap, ağ hatalarında alttaki hatayı döndürür (errors.Is(err, context.DeadlineExceeded) vb. için).
func (e *Error) Unwrap() error { return e.Err }

// Retryable, hatanın yeniden denenebilir olup olmadığını bildirir (429, 5xx veya ağ hatası).
func (e *Error) Retryable() bool {
	return e.StatusCode == 0 || e.StatusCode == 429 || e.StatusCode >= 500
}
