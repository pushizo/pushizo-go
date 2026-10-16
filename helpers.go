package pushizo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Ptr, herhangi bir değerin işaretçisini döndürür (isteğe bağlı alanlar için).
func Ptr[T any](v T) *T { return &v }

// String, isteğe bağlı string alanları için işaretçi döndürür.
func String(v string) *string { return &v }

// Int64, isteğe bağlı tamsayı alanları için işaretçi döndürür.
func Int64(v int64) *int64 { return &v }

// Float64, isteğe bağlı ondalık alanları için işaretçi döndürür.
func Float64(v float64) *float64 { return &v }

// Bool, isteğe bağlı mantıksal alanları için işaretçi döndürür.
func Bool(v bool) *bool { return &v }

// IdentityHash, SDK identity verification için identity_hash = HMAC-SHA256(sdkSecret, externalID) (hex) üretir.
// Sunucu tarafında hesaplanıp istemci SDK'sına iletilir; sdkSecret istemciye gönderilmez.
func IdentityHash(sdkSecret, externalID string) string {
	mac := hmac.New(sha256.New, []byte(sdkSecret))
	mac.Write([]byte(externalID))
	return hex.EncodeToString(mac.Sum(nil))
}
