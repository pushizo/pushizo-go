// Package webhook, HeroSignal giden webhook imzalarını doğrular (alıcı tarafı).
//
// Başlık: "HeroSignal-Signature: t=<unix>,v1=<hex>"; v1 = HMAC-SHA256(secret, t + "." + ham gövde). Secret
// rotasyonunda başlıkta iki v1 bulunabilir; herhangi biri eşleşirse imza geçerlidir. Karşılaştırma sabit zamanlıdır.
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//		payload, _ := io.ReadAll(r.Body) // HAM gövde; JSON'u yeniden serileştirmeyin
//		events, err := webhook.ConstructEvents(payload, r.Header.Get(webhook.SignatureHeader), []string{secret}, nil)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		// HeroSignal-Event-Id ile tekrarları ayıklayın, 2xx ile hızlı yanıt verin.
//		w.WriteHeader(http.StatusNoContent)
//	}
package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

const (
	// SignatureHeader, imza başlığıdır.
	SignatureHeader = "HeroSignal-Signature"
	// EventIDHeader, olay kimliği başlığıdır; tekrar gönderimleri ayıklamak için kullanın.
	EventIDHeader = "HeroSignal-Event-Id"
	// DefaultTolerance, varsayılan zaman damgası toleransıdır; daha eski t tekrar (replay) sayılır.
	DefaultTolerance = 300 * time.Second
)

// Failure, doğrulama başarısızlık nedenidir (tüm SDK'larda aynı kodlar).
type Failure string

const (
	// MissingHeader: başlık yok veya boş.
	MissingHeader Failure = "missing_header"
	// MalformedHeader: t veya geçerli bir v1 yok.
	MalformedHeader Failure = "malformed_header"
	// TimestampOutOfTolerance: zaman damgası tolerans dışında.
	TimestampOutOfTolerance Failure = "timestamp_out_of_tolerance"
	// SignatureMismatch: hiçbir secret ile imza eşleşmedi.
	SignatureMismatch Failure = "signature_mismatch"
	// MalformedBody: imza geçerli ama gövde JSON nesnesi değil (yalnızca ConstructEvents).
	MalformedBody Failure = "malformed_body"
)

// VerifyOptions, doğrulama ayarlarıdır; nil veya sıfır değerler varsayılanı kullanır.
type VerifyOptions struct {
	// Tolerance zaman damgası toleransı (0 ise DefaultTolerance).
	Tolerance time.Duration
	// Now şimdiki zaman (test için; sıfırsa time.Now()).
	Now time.Time
}

// VerifyResult, doğrulama sonucudur. OK ise Timestamp, değilse Reason doludur.
type VerifyResult struct {
	OK        bool
	Timestamp int64
	Reason    Failure
}

// SignatureError, ConstructEvents'in imza veya gövde hatasıdır.
type SignatureError struct {
	Reason Failure
}

func (e *SignatureError) Error() string {
	return "herosignal: webhook imzası geçersiz: " + string(e.Reason)
}

// Event, webhook olayıdır. Batch gövdesi {"events": [...]} biçimindedir.
type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	CreatedAt string          `json:"created_at"`
	AppID     string          `json:"app_id"`
	Data      json.RawMessage `json:"data"`
}

// Verify, imzayı doğrular. payload HAM istek gövdesidir; secrets rotasyon sırasında eski ve yeni secret'ı içerebilir.
func Verify(payload []byte, header string, secrets []string, opts *VerifyOptions) VerifyResult {
	if header == "" {
		return VerifyResult{Reason: MissingHeader}
	}
	timestamp := int64(-1)
	var signatures [][]byte
	for _, part := range strings.Split(header, ",") {
		idx := strings.Index(part, "=")
		if idx <= 0 {
			continue
		}
		key := strings.TrimSpace(part[:idx])
		value := strings.TrimSpace(part[idx+1:])
		switch key {
		case "t":
			if isDigits(value) {
				if n, err := strconv.ParseInt(value, 10, 64); err == nil {
					timestamp = n
				}
			}
		case "v1":
			if isLowerHex64(value) {
				if b, err := hex.DecodeString(value); err == nil {
					signatures = append(signatures, b)
				}
			}
		}
	}
	if timestamp < 0 || len(signatures) == 0 {
		return VerifyResult{Reason: MalformedHeader}
	}
	tolerance := DefaultTolerance
	now := time.Now()
	if opts != nil {
		if opts.Tolerance > 0 {
			tolerance = opts.Tolerance
		}
		if !opts.Now.IsZero() {
			now = opts.Now
		}
	}
	diff := now.Unix() - timestamp
	if diff < 0 {
		diff = -diff
	}
	if diff > int64(tolerance/time.Second) {
		return VerifyResult{Reason: TimestampOutOfTolerance}
	}
	prefix := []byte(strconv.FormatInt(timestamp, 10) + ".")
	for _, secret := range secrets {
		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write(prefix)
		mac.Write(payload)
		expected := mac.Sum(nil)
		for _, sig := range signatures {
			if hmac.Equal(sig, expected) {
				return VerifyResult{OK: true, Timestamp: timestamp}
			}
		}
	}
	return VerifyResult{Reason: SignatureMismatch}
}

// ConstructEvents, imzayı doğrular ve gövdeyi olay listesine çevirir (tekil gövde veya {"events": [...]} batch).
// Geçersizse *SignatureError döner.
func ConstructEvents(payload []byte, header string, secrets []string, opts *VerifyOptions) ([]Event, error) {
	result := Verify(payload, header, secrets, opts)
	if !result.OK {
		return nil, &SignatureError{Reason: result.Reason}
	}
	var probe map[string]json.RawMessage
	if err := json.Unmarshal(payload, &probe); err != nil || probe == nil {
		return nil, &SignatureError{Reason: MalformedBody}
	}
	if raw, ok := probe["events"]; ok {
		var events []Event
		if err := json.Unmarshal(raw, &events); err != nil {
			return nil, &SignatureError{Reason: MalformedBody}
		}
		return events, nil
	}
	var event Event
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, &SignatureError{Reason: MalformedBody}
	}
	return []Event{event}, nil
}

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func isLowerHex64(s string) bool {
	if len(s) != 64 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') {
			return false
		}
	}
	return true
}
