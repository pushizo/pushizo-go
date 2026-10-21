<!-- pushizo:header -->
<p align="center">
  <a href="https://pushizo.com"><img src="https://raw.githubusercontent.com/pushizo/.github/main/profile/assets/mark.svg" alt="Pushizo" width="56" height="56"></a>
</p>
<p align="center">
  <a href="https://docs.pushizo.com/docs/sdk/go"><img src="https://img.shields.io/badge/docs-docs.pushizo.com-5b4bdb" alt="Docs"></a>
  <a href="https://docs.pushizo.com/docs/api"><img src="https://img.shields.io/badge/API-reference-b44fc4" alt="API reference"></a>
  <a href="https://github.com/pushizo/pushizo-go/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-16130f" alt="License: MIT"></a>
</p>
<p align="center"><a href="#english">English</a> · <a href="#türkçe">Türkçe</a></p>

## English

**Pushizo Go SDK** — Server-side SDK for Go 1.23+ using only the standard library (net/http). Models and services are generated from the Pushizo OpenAPI schema; authentication, retries, idempotency, pagination and webhook signature verification are hand-written.

```bash
go get github.com/pushizo/pushizo-go
```

The detailed guide below is in Turkish; code samples are language-neutral. Full documentation:
[docs.pushizo.com/docs/sdk/go](https://docs.pushizo.com/docs/sdk/go) · API reference: [docs.pushizo.com/docs/api](https://docs.pushizo.com/docs/api).

## Türkçe
<!-- /pushizo:header -->

# Pushizo Go SDK

Pushizo public API'si için Go sunucu SDK'sı (Go 1.23+, yalnızca standart kütüphane: `net/http`).
Modeller (`models_gen.go`) ve servisler (`resources_gen.go`) [Pushizo OpenAPI şemasından](https://github.com/pushizo/pushizo-openapi) üretilir; kimlik doğrulama,
yeniden deneme, idempotency, sayfalama ve webhook imza doğrulaması el yazımıdır.

## Kurulum

```bash
go get github.com/pushizo/pushizo-go
```

## Kullanım

```go
import (
	"context"
	"errors"
	"log"
	"os"

	pushizo "github.com/pushizo/pushizo-go"
)

func main() {
	ctx := context.Background()
	hs, err := pushizo.NewClient(os.Getenv("HS_API_KEY"),
		pushizo.WithAppID("app_…"),
		// pushizo.WithBaseURL("http://localhost:4000"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Bildirim gönder: Idempotency-Key verilmezse otomatik UUID (yeniden denemelerde aynı anahtar)
	res, err := hs.Notifications.Create(ctx, &pushizo.CreateNotificationDto{
		Target:   pushizo.CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {"user-42"}}},
		Contents: map[string]string{"en": "Your order shipped", "tr": "Siparişiniz kargoya verildi"},
		ThreadID: pushizo.String("orders"),
	}, pushizo.WithIdempotencyKey("order-1001"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.ID, res.Status)

	// Kullanıcı oluştur/güncelle
	_, err = hs.Users.Upsert(ctx, &pushizo.UpsertUserDto{
		ExternalID: pushizo.String("user-42"),
		Tags:       map[string]string{"plan": "pro"},
	})

	// Cursor sayfalama (Go 1.23 range-over-func)
	for segment, err := range hs.Segments.ListAll(ctx, &pushizo.SegmentsListParams{Limit: pushizo.Int64(100)}) {
		if err != nil {
			log.Fatal(err)
		}
		log.Println(segment.Name)
	}

	// Başka bir app için tek çağrı
	_, _ = hs.Apps.Get(ctx, pushizo.ForApp("app_diger"))

	// SDK identity verification
	_ = pushizo.IdentityHash(os.Getenv("HS_SDK_SECRET"), "user-42")
}
```

İsteğe bağlı alanlar işaretçidir: `pushizo.String`, `pushizo.Int64`, `pushizo.Bool`, `pushizo.Float64`
veya genel `pushizo.Ptr(v)` yardımcılarını kullanın. SDK'da metodu olmayan uçlar ve JSON dışı yanıtlar için:

```go
var csv string
err := hs.Request(ctx, "GET", "/v1/apps/{appId}/analytics/messages", url.Values{"format": {"csv"}}, nil, &csv)
```

### Ayarlar

| Seçenek | Varsayılan | Açıklama |
|---------|------------|----------|
| `WithAppID` | — | Çağrıda `ForApp` verilmezse kullanılan app |
| `WithBaseURL` | `https://api.pushizo.com` | API adresi |
| `WithTimeout` | 30 sn | Deneme başına zaman aşımı |
| `WithMaxRetries` | 3 | En fazla yeniden deneme (politika: aşağıda) |
| `WithAutoIdempotencyKey` | `true` | `Idempotency-Key` destekleyen çağrılarda otomatik UUID |
| `WithHTTPClient` | `http.DefaultClient` | Özel `*http.Client` |

Çağrı başına: `ForApp`, `WithIdempotencyKey`, `WithHeader`, `WithRequestTimeout`, `WithRequestMaxRetries`.

### Hatalar

429 her metotta; 5xx ve ağ hataları GET/HEAD/OPTIONS/PUT/DELETE'te, POST/PATCH'te yalnızca `Idempotency-Key` varsa (otomatik anahtar dahil) veya istek sunucuya hiç gönderilemediyse (bağlantı reddi, DNS hatası, bağlantı zaman aşımı) üstel geri çekilmeyle (0,3 / 1 / 3 sn) yeniden denenir; `Retry-After` başlığına uyulur.
API hataları `*pushizo.Error` tipindedir (`StatusCode`, `Code` (`HS_…`), `Message`, `RequestID`, `Details`):

```go
_, err := hs.Notifications.Get(ctx, "msg_yok")
var apiErr *pushizo.Error
switch {
case errors.Is(err, pushizo.ErrNotFound):
	// 404
case errors.Is(err, pushizo.ErrRateLimited):
	// 429 (yeniden denemeler tükendi)
case errors.As(err, &apiErr) && apiErr.Code == pushizo.CodeNoRecipients:
	// 422 HS_NO_RECIPIENTS
}
```

Kategoriler: `ErrNetwork`, `ErrInvalidRequest` (400/422), `ErrAuthentication`, `ErrPermissionDenied`, `ErrNotFound`,
`ErrConflict`, `ErrRateLimited`, `ErrServer`.

## Webhook imza doğrulaması

```go
http.HandleFunc("/hooks/pushizo", func(w http.ResponseWriter, r *http.Request) {
	payload, _ := io.ReadAll(r.Body) // ham gövde
	events, err := pushizo.ConstructWebhookEvents(payload,
		r.Header.Get(pushizo.WebhookSignatureHeader),
		[]string{os.Getenv("HS_WEBHOOK_SECRET")}, // rotasyonda eski ve yeni secret birlikte verilebilir
		nil)                                       // varsayılan tolerans 5 dk
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, ev := range events { // tekil veya batch gövde
		_ = ev.Type // Pushizo-Event-Id ile tekrarları ayıklayın
	}
	w.WriteHeader(http.StatusNoContent)
})
```

Yalnızca doğrulama için `pushizo.VerifyWebhookSignature(...)` `WebhookVerifyResult{OK, Timestamp, Reason}` döner.

## Geliştirme

```bash
go vet ./... && go test ./...                          # sdk/go içinde
node sdk/scripts/generate-server-sdks.mjs --lang go    # depo kökünde; *_gen.go dosyaları yeniden üretilir
```

`*_gen.go` dosyaları elle düzenlenmez (gofmt uyumlu üretilir). Webhook testleri
`sdk/testdata/webhook-signature-vectors.json` ortak vektörlerini kullanır.

<!-- pushizo:footer -->
---

## Support · Destek

- **Guide · Rehber:** [docs.pushizo.com/docs/sdk/go](https://docs.pushizo.com/docs/sdk/go)
- **API reference & OpenAPI · API referansı ve OpenAPI:** [docs.pushizo.com/docs/api](https://docs.pushizo.com/docs/api) ·
  [pushizo-openapi](https://github.com/pushizo/pushizo-openapi)
- **Bugs & ideas · Hata ve öneri:** [Issues](https://github.com/pushizo/pushizo-go/issues)
- **Security · Güvenlik:** please do not open public issues; email [security@pushizo.com](mailto:security@pushizo.com) ·
  herkese açık issue açmayın, bu adrese yazın.
- **Help · Destek:** [support@pushizo.com](mailto:support@pushizo.com)

This repository is mirrored automatically from the Pushizo main repository on every release. Please open an issue
instead of a pull request; fixes land in the main repository and ship with the next release.
Bu depo her sürümde Pushizo ana deposundan otomatik yansıtılır. Pull request yerine issue açın; düzeltmeler ana depoda
yapılır ve bir sonraki sürümle buraya gelir.

Released under the MIT License · MIT lisanslıdır — [LICENSE](https://github.com/pushizo/pushizo-go/blob/main/LICENSE).
<!-- /pushizo:footer -->
