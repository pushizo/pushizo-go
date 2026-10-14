# HeroSignal Go SDK

HeroSignal public API'si için Go sunucu SDK'sı (Go 1.23+, yalnızca standart kütüphane: `net/http`).
Modeller (`models_gen.go`) ve servisler (`resources_gen.go`) `api/openapi.yaml`'dan üretilir; kimlik doğrulama,
yeniden deneme, idempotency, sayfalama ve webhook imza doğrulaması el yazımıdır.

## Kurulum

```bash
go get github.com/exec0de/herosignal/sdk/go
```

## Kullanım

```go
import (
	"context"
	"errors"
	"log"
	"os"

	herosignal "github.com/exec0de/herosignal/sdk/go"
)

func main() {
	ctx := context.Background()
	hs, err := herosignal.NewClient(os.Getenv("HS_API_KEY"),
		herosignal.WithAppID("app_…"),
		// herosignal.WithBaseURL("http://localhost:4000"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Bildirim gönder: Idempotency-Key verilmezse otomatik UUID (yeniden denemelerde aynı anahtar)
	res, err := hs.Notifications.Create(ctx, &herosignal.CreateNotificationDto{
		Target:   herosignal.CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {"user-42"}}},
		Contents: map[string]string{"en": "Your order shipped", "tr": "Siparişiniz kargoya verildi"},
		ThreadID: herosignal.String("orders"),
	}, herosignal.WithIdempotencyKey("order-1001"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.ID, res.Status)

	// Kullanıcı oluştur/güncelle
	_, err = hs.Users.Upsert(ctx, &herosignal.UpsertUserDto{
		ExternalID: herosignal.String("user-42"),
		Tags:       map[string]string{"plan": "pro"},
	})

	// Cursor sayfalama (Go 1.23 range-over-func)
	for segment, err := range hs.Segments.ListAll(ctx, &herosignal.SegmentsListParams{Limit: herosignal.Int64(100)}) {
		if err != nil {
			log.Fatal(err)
		}
		log.Println(segment.Name)
	}

	// Başka bir app için tek çağrı
	_, _ = hs.Apps.Get(ctx, herosignal.ForApp("app_diger"))

	// SDK identity verification
	_ = herosignal.IdentityHash(os.Getenv("HS_SDK_SECRET"), "user-42")
}
```

İsteğe bağlı alanlar işaretçidir: `herosignal.String`, `herosignal.Int64`, `herosignal.Bool`, `herosignal.Float64`
veya genel `herosignal.Ptr(v)` yardımcılarını kullanın. SDK'da metodu olmayan uçlar ve JSON dışı yanıtlar için:

```go
var csv string
err := hs.Request(ctx, "GET", "/v1/apps/{appId}/analytics/messages", url.Values{"format": {"csv"}}, nil, &csv)
```

### Ayarlar

| Seçenek | Varsayılan | Açıklama |
|---------|------------|----------|
| `WithAppID` | — | Çağrıda `ForApp` verilmezse kullanılan app |
| `WithBaseURL` | `https://api.herosignal.dev` | API adresi |
| `WithTimeout` | 30 sn | Deneme başına zaman aşımı |
| `WithMaxRetries` | 3 | 429/5xx ve ağ hatalarında yeniden deneme |
| `WithAutoIdempotencyKey` | `true` | `Idempotency-Key` destekleyen çağrılarda otomatik UUID |
| `WithHTTPClient` | `http.DefaultClient` | Özel `*http.Client` |

Çağrı başına: `ForApp`, `WithIdempotencyKey`, `WithHeader`, `WithRequestTimeout`, `WithRequestMaxRetries`.

### Hatalar

429, 5xx ve ağ hatalarında üstel geri çekilmeyle (0,3 / 1 / 3 sn) yeniden denenir; `Retry-After` başlığına uyulur.
API hataları `*herosignal.Error` tipindedir (`StatusCode`, `Code` (`HS_…`), `Message`, `RequestID`, `Details`):

```go
_, err := hs.Notifications.Get(ctx, "msg_yok")
var apiErr *herosignal.Error
switch {
case errors.Is(err, herosignal.ErrNotFound):
	// 404
case errors.Is(err, herosignal.ErrRateLimited):
	// 429 (yeniden denemeler tükendi)
case errors.As(err, &apiErr) && apiErr.Code == herosignal.CodeNoRecipients:
	// 422 HS_NO_RECIPIENTS
}
```

Kategoriler: `ErrNetwork`, `ErrInvalidRequest` (400/422), `ErrAuthentication`, `ErrPermissionDenied`, `ErrNotFound`,
`ErrConflict`, `ErrRateLimited`, `ErrServer`.

## Webhook imza doğrulaması

```go
http.HandleFunc("/hooks/herosignal", func(w http.ResponseWriter, r *http.Request) {
	payload, _ := io.ReadAll(r.Body) // ham gövde
	events, err := herosignal.ConstructWebhookEvents(payload,
		r.Header.Get(herosignal.WebhookSignatureHeader),
		[]string{os.Getenv("HS_WEBHOOK_SECRET")}, // rotasyonda eski ve yeni secret birlikte verilebilir
		nil)                                       // varsayılan tolerans 5 dk
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, ev := range events { // tekil veya batch gövde
		_ = ev.Type // HeroSignal-Event-Id ile tekrarları ayıklayın
	}
	w.WriteHeader(http.StatusNoContent)
})
```

Yalnızca doğrulama için `herosignal.VerifyWebhookSignature(...)` `WebhookVerifyResult{OK, Timestamp, Reason}` döner.

## Geliştirme

```bash
go vet ./... && go test ./...                          # sdk/go içinde
node sdk/scripts/generate-server-sdks.mjs --lang go    # depo kökünde; *_gen.go dosyaları yeniden üretilir
```

`*_gen.go` dosyaları elle düzenlenmez (gofmt uyumlu üretilir). Webhook testleri
`sdk/testdata/webhook-signature-vectors.json` ortak vektörlerini kullanır.
