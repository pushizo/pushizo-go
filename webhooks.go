package pushizo

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
	// WebhookSignatureHeader, imza başlığıdır: "t=<unix>,v1=<hex>".
	WebhookSignatureHeader = "Pushizo-Signature"
	// WebhookEventIDHeader, olay kimliği başlığıdır; tekrar gönderimleri ayıklamak için kullanın.
	WebhookEventIDHeader = "Pushizo-Event-Id"
	// WebhookDefaultTolerance, varsayılan zaman damgası toleransıdır; daha eski t tekrar (replay) sayılır.
	WebhookDefaultTolerance = 300 * time.Second
)

// Webhook doğrulama hata nedenleri (tüm SDK'larda aynı).
const (
	WebhookReasonMissingHeader           = "missing_header"
	WebhookReasonMalformedHeader         = "malformed_header"
	WebhookReasonTimestampOutOfTolerance = "timestamp_out_of_tolerance"
	WebhookReasonSignatureMismatch       = "signature_mismatch"
	WebhookReasonMalformedBody           = "malformed_body"
)

// WebhookVerifyOptions, doğrulama ayarlarıdır; nil veya sıfır değerler varsayılanı kullanır.
type WebhookVerifyOptions struct {
	// Tolerance, zaman damgası toleransıdır (0 → 5 dk).
	Tolerance time.Duration
	// Now, test için şimdiki zamandır (sıfır → time.Now()).
	Now time.Time
}

// WebhookVerifyResult, doğrulama sonucudur. OK ise Timestamp, değilse Reason doludur.
type WebhookVerifyResult struct {
	OK        bool
	Timestamp int64
	Reason    string
}

// WebhookSignatureError, ConstructWebhookEvents'in geçersiz imzada döndürdüğü hatadır.
type WebhookSignatureError struct{ Reason string }

func (e *WebhookSignatureError) Error() string {
	return "pushizo: webhook imzası geçersiz: " + e.Reason
}

// WebhookEvent, webhook olayıdır (tekil gövde). Batch gövdesi {"events": [...]} biçimindedir.
type WebhookEvent struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	CreatedAt string          `json:"created_at"`
	AppID     string          `json:"app_id"`
	Data      json.RawMessage `json:"data"`
}

// VerifyWebhookSignature, Pushizo-Signature başlığını doğrular: v1 = HMAC-SHA256(secret, t + "." + ham gövde).
// payload HAM istek gövdesidir (JSON'u ayrıştırıp yeniden serileştirmeyin). Rotasyon sırasında eski ve yeni
// secret birlikte verilebilir; herhangi biri eşleşirse imza geçerlidir. Karşılaştırma sabit zamanlıdır.
//
//	body, _ := io.ReadAll(r.Body)
//	res := pushizo.VerifyWebhookSignature(body, r.Header.Get(pushizo.WebhookSignatureHeader), []string{secret}, nil)
//	if !res.OK { http.Error(w, res.Reason, http.StatusBadRequest); return }
func VerifyWebhookSignature(payload []byte, header string, secrets []string, opts *WebhookVerifyOptions) WebhookVerifyResult {
	if header == "" {
		return WebhookVerifyResult{Reason: WebhookReasonMissingHeader}
	}
	var timestamp int64
	hasTimestamp := false
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
				if ts, err := strconv.ParseInt(value, 10, 64); err == nil {
					timestamp, hasTimestamp = ts, true
				}
			}
		case "v1":
			if isLowerHex64(value) {
				if sig, err := hex.DecodeString(value); err == nil {
					signatures = append(signatures, sig)
				}
			}
		}
	}
	if !hasTimestamp || len(signatures) == 0 {
		return WebhookVerifyResult{Reason: WebhookReasonMalformedHeader}
	}
	tolerance := WebhookDefaultTolerance
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
		return WebhookVerifyResult{Reason: WebhookReasonTimestampOutOfTolerance}
	}
	prefix := []byte(strconv.FormatInt(timestamp, 10) + ".")
	for _, secret := range secrets {
		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write(prefix)
		mac.Write(payload)
		expected := mac.Sum(nil)
		for _, sig := range signatures {
			if hmac.Equal(sig, expected) {
				return WebhookVerifyResult{OK: true, Timestamp: timestamp}
			}
		}
	}
	return WebhookVerifyResult{Reason: WebhookReasonSignatureMismatch}
}

// ConstructWebhookEvents, imzayı doğrular ve gövdeyi olay listesine çevirir (tekil veya batch).
// Geçersiz imzada *WebhookSignatureError döner.
func ConstructWebhookEvents(payload []byte, header string, secrets []string, opts *WebhookVerifyOptions) ([]WebhookEvent, error) {
	res := VerifyWebhookSignature(payload, header, secrets, opts)
	if !res.OK {
		return nil, &WebhookSignatureError{Reason: res.Reason}
	}
	var probe map[string]json.RawMessage
	if err := json.Unmarshal(payload, &probe); err != nil {
		return nil, &WebhookSignatureError{Reason: WebhookReasonMalformedBody}
	}
	if raw, ok := probe["events"]; ok {
		var events []WebhookEvent
		if err := json.Unmarshal(raw, &events); err == nil {
			return events, nil
		}
	}
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, &WebhookSignatureError{Reason: WebhookReasonMalformedBody}
	}
	return []WebhookEvent{event}, nil
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
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f') {
			return false
		}
	}
	return true
}
