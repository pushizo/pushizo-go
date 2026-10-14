package herosignal_test

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	herosignal "github.com/exec0de/herosignal/sdk/go"
)

// README'deki kullanım örneği (derlenir, çalıştırılmaz).
func ExampleNewClient() {
	ctx := context.Background()
	hs, err := herosignal.NewClient(os.Getenv("HS_API_KEY"), herosignal.WithAppID("app_…"))
	if err != nil {
		log.Fatal(err)
	}
	res, err := hs.Notifications.Create(ctx, &herosignal.CreateNotificationDto{
		Target:   herosignal.CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {"user-42"}}},
		Contents: map[string]string{"en": "Your order shipped", "tr": "Siparişiniz kargoya verildi"},
		ThreadID: herosignal.String("orders"),
	}, herosignal.WithIdempotencyKey("order-1001"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.ID, res.Status)

	_, _ = hs.Users.Upsert(ctx, &herosignal.UpsertUserDto{ExternalID: herosignal.String("user-42"), Tags: map[string]string{"plan": "pro"}})
	for segment, err := range hs.Segments.ListAll(ctx, &herosignal.SegmentsListParams{Limit: herosignal.Int64(100)}) {
		if err != nil {
			log.Fatal(err)
		}
		log.Println(segment.Name)
	}
	_, _ = hs.Apps.Get(ctx, herosignal.ForApp("app_diger"))

	var csv string
	_ = hs.Request(ctx, "GET", "/v1/apps/{appId}/analytics/messages", url.Values{"format": {"csv"}}, nil, &csv)

	_, err = hs.Notifications.Get(ctx, "msg_yok")
	var apiErr *herosignal.Error
	switch {
	case errors.Is(err, herosignal.ErrNotFound):
	case errors.As(err, &apiErr) && apiErr.Code == herosignal.CodeNoRecipients:
	}
	_ = herosignal.IdentityHash(os.Getenv("HS_SDK_SECRET"), "user-42")
}

// Webhook alıcısı örneği (derlenir, çalıştırılmaz).
func ExampleConstructWebhookEvents() {
	http.HandleFunc("/hooks/herosignal", func(w http.ResponseWriter, r *http.Request) {
		payload, _ := io.ReadAll(r.Body)
		events, err := herosignal.ConstructWebhookEvents(payload, r.Header.Get(herosignal.WebhookSignatureHeader),
			[]string{os.Getenv("HS_WEBHOOK_SECRET")}, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, ev := range events {
			_ = ev.Type
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
