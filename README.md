# herosignal-go (Go)

HeroSignal sunucu SDK'sı: bildirim gönderme, kullanıcı/abonelik yönetimi, segment/şablon/kampanya, analitik ve webhook imza doğrulaması. API anahtarı yalnızca sunucuda kullanılır.

## Kurulum
```
go get github.com/herosignal/herosignal-go
```

## Kullanım
```go
client, _ := herosignal.NewClient(os.Getenv("HS_API_KEY"), herosignal.WithAppID("app_…"))
res, err := client.Notifications.Create(ctx, &herosignal.CreateNotificationDto{
	Target:   herosignal.CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {"u1"}}},
	Contents: map[string]string{"en": "Shipped"},
}, herosignal.WithIdempotencyKey("order-42"))
```

- Kimlik doğrulama (Bearer API anahtarı), zaman aşımı, 429/5xx/ağ hatalarında üstel geri çekilmeli yeniden deneme (`Retry-After`'a uyar), otomatik `Idempotency-Key`, hata tipi (status, code, message, request_id), cursor sayfalama ve `HeroSignal-Signature` webhook doğrulaması (sabit zamanlı, tolerans, secret rotasyonu).
- Kaynaklar ve modeller `api/openapi.yaml`'dan `node sdk/scripts/generate-server-sdks.mjs` ile üretilir; üretilmiş dosyaları elle düzenlemeyin.

Ayrıntılı rehber: [docs/sdk/go.md](../../docs/sdk/go.md)

## Test
```
go vet ./... && go test ./...
```
