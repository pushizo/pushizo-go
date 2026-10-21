package pushizo_test

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	pushizo "github.com/pushizo/pushizo-go"
)

// README'deki kullanım örneği (derlenir, çalıştırılmaz).
func ExampleNewClient() {
	ctx := context.Background()
	hs, err := pushizo.NewClient(os.Getenv("HS_API_KEY"), pushizo.WithAppID("app_…"))
	if err != nil {
		log.Fatal(err)
	}
	res, err := hs.Notifications.Create(ctx, &pushizo.CreateNotificationDto{
		Target:   pushizo.CreateNotificationDtoTarget{IncludeAliases: map[string][]string{"external_id": {"user-42"}}},
		Contents: map[string]string{"en": "Your order shipped", "tr": "Siparişiniz kargoya verildi"},
		ThreadID: pushizo.String("orders"),
	}, pushizo.WithIdempotencyKey("order-1001"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.ID, res.Status)

	_, _ = hs.Users.Upsert(ctx, &pushizo.UpsertUserDto{ExternalID: pushizo.String("user-42"), Tags: map[string]string{"plan": "pro"}})
	for segment, err := range hs.Segments.ListAll(ctx, &pushizo.SegmentsListParams{Limit: pushizo.Int64(100)}) {
		if err != nil {
			log.Fatal(err)
		}
		log.Println(segment.Name)
	}
	_, _ = hs.Apps.Get(ctx, pushizo.ForApp("app_diger"))

	var csv string
	_ = hs.Request(ctx, "GET", "/v1/apps/{appId}/analytics/messages", url.Values{"format": {"csv"}}, nil, &csv)

	_, err = hs.Notifications.Get(ctx, "msg_yok")
	var apiErr *pushizo.Error
	switch {
	case errors.Is(err, pushizo.ErrNotFound):
	case errors.As(err, &apiErr) && apiErr.Code == pushizo.CodeNoRecipients:
	}
	_ = pushizo.IdentityHash(os.Getenv("HS_SDK_SECRET"), "user-42")
}

// Webhook alıcısı örneği (derlenir, çalıştırılmaz).
func ExampleConstructWebhookEvents() {
	http.HandleFunc("/hooks/pushizo", func(w http.ResponseWriter, r *http.Request) {
		payload, _ := io.ReadAll(r.Body)
		events, err := pushizo.ConstructWebhookEvents(payload, r.Header.Get(pushizo.WebhookSignatureHeader),
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
