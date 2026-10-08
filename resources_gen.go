// Bu dosya api/openapi.yaml'dan sdk/scripts/generate-server-sdks.mjs ile üretilmiştir; elle düzenlemeyin.

package herosignal

import (
	"context"
	"iter"
	"net/url"
)

// APIVersion, SDK'nın üretildiği OpenAPI sözleşmesinin sürümüdür.
const APIVersion = "0.1.0"

// services, istemciye eklenen üretilmiş kaynak alanlarıdır.
type services struct {
	Notifications *NotificationsService
	Users         *UsersService
	Aliases       *AliasesService
	Subscriptions *SubscriptionsService
	Events        *EventsService
	Segments      *SegmentsService
	Templates     *TemplatesService
	Campaigns     *CampaignsService
	InAppMessages *InAppMessagesService
	Inbox         *InboxService
	Suppressions  *SuppressionsService
	Exports       *ExportsService
	Imports       *ImportsService
	Webhooks      *WebhooksService
	Channels      *ChannelsService
	Analytics     *AnalyticsService
	Apps          *AppsService
	APIKeys       *APIKeysService
	Organization  *OrganizationService
	Categories    *CategoriesService
	Journeys      *JourneysService
	Preferences   *PreferencesService
}

func (c *Client) initServices() {
	c.Notifications = &NotificationsService{client: c}
	c.Users = &UsersService{client: c}
	c.Aliases = &AliasesService{client: c}
	c.Subscriptions = &SubscriptionsService{client: c}
	c.Events = &EventsService{client: c}
	c.Segments = &SegmentsService{client: c}
	c.Templates = &TemplatesService{client: c}
	c.Campaigns = &CampaignsService{client: c}
	c.InAppMessages = &InAppMessagesService{client: c}
	c.Inbox = &InboxService{client: c}
	c.Suppressions = &SuppressionsService{client: c}
	c.Exports = &ExportsService{client: c}
	c.Imports = &ImportsService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	c.Channels = &ChannelsService{client: c}
	c.Analytics = &AnalyticsService{client: c}
	c.Apps = &AppsService{client: c}
	c.APIKeys = &APIKeysService{client: c}
	c.Organization = &OrganizationService{client: c}
	c.Categories = &CategoriesService{client: c}
	c.Journeys = &JourneysService{client: c}
	c.Preferences = &PreferencesService{client: c}
}

// NotificationsService: Bildirim gönderme, durum, iptal, teslimatlar ve geçmiş dışa aktarımı.
type NotificationsService struct {
	client *Client
}

// Create: Bildirim oluştur (202, asenkron teslimat)
//
// POST /v1/notifications
//
// Idempotency-Key başlığı gönderilir: WithIdempotencyKey, gövdedeki idempotency_key veya otomatik UUID.
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *NotificationsService) Create(ctx context.Context, body *CreateNotificationDto, opts ...RequestOption) (*CreateNotificationResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/notifications", idempotent: true}
	if body == nil {
		return nil, errNilBody
	}
	b := *body
	if b.AppID == nil || *b.AppID == "" {
		appID, err := ro.resolveAppID()
		if err != nil {
			return nil, err
		}
		b.AppID = &appID
	}
	if b.IdempotencyKey != nil {
		op.idempotencyKey = *b.IdempotencyKey
	}
	op.body = &b
	var out CreateNotificationResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// NotificationsListParams: NotificationsService.List için isteğe bağlı sorgu parametreleri.
type NotificationsListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "campaign", "api", "journey", "test".
	Source *string
	// Değerler: "campaign", "api", "journey", "test".
	Kind *string
	// Değerler: "queued", "processing", "completed", "cancelled", "failed".
	Status *string
	// Değerler: "created_at", "-created_at".
	Sort  *string
	Label *string
}

func (p *NotificationsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "source", p.Source)
	addQuery(q, "kind", p.Kind)
	addQuery(q, "status", p.Status)
	addQuery(q, "sort", p.Sort)
	addQuery(q, "label", p.Label)
}

// List: Mesaj listesi (app_id zorunlu)
//
// GET /v1/notifications
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *NotificationsService) List(ctx context.Context, params *NotificationsListParams, opts ...RequestOption) (*MessageList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/notifications", queryAppID: true}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out MessageList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *NotificationsService) ListAll(ctx context.Context, params *NotificationsListParams, opts ...RequestOption) iter.Seq2[Message, error] {
	var p NotificationsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Message, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Estimate: Kaydedilmemiş hedef için alıcı sayısı (gönderim yapmaz)
//
// POST /v1/notifications/estimate
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *NotificationsService) Estimate(ctx context.Context, body *EstimateAudienceDto, opts ...RequestOption) (*EstimateAudienceResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/notifications/estimate"}
	if body == nil {
		return nil, errNilBody
	}
	b := *body
	if b.AppID == nil || *b.AppID == "" {
		appID, err := ro.resolveAppID()
		if err != nil {
			return nil, err
		}
		b.AppID = &appID
	}
	op.body = &b
	var out EstimateAudienceResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Mesaj durumu ve sayaçları
//
// GET /v1/notifications/{messageId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *NotificationsService) Get(ctx context.Context, messageID string, opts ...RequestOption) (*Message, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/notifications/{messageId}", queryAppID: true}
	op.pathParams = map[string]string{"messageId": messageID}
	var out Message
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Cancel: Zamanlanmış/kuyruktaki mesajı iptal et
//
// DELETE /v1/notifications/{messageId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *NotificationsService) Cancel(ctx context.Context, messageID string, opts ...RequestOption) (*Message, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/notifications/{messageId}", queryAppID: true}
	op.pathParams = map[string]string{"messageId": messageID}
	var out Message
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// NotificationsDeliveriesParams: NotificationsService.Deliveries için isteğe bağlı sorgu parametreleri.
type NotificationsDeliveriesParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "pending", "sent", "delivered", "failed", "bounced", "suppressed", "expired", "cancelled".
	Status *string
}

func (p *NotificationsDeliveriesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
}

// Deliveries: Abone bazlı teslimat sonuçları (cursor)
//
// GET /v1/notifications/{messageId}/deliveries
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *NotificationsService) Deliveries(ctx context.Context, messageID string, params *NotificationsDeliveriesParams, opts ...RequestOption) (*DeliveryList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/notifications/{messageId}/deliveries", queryAppID: true}
	op.pathParams = map[string]string{"messageId": messageID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out DeliveryList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeliveriesAll: Deliveries sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *NotificationsService) DeliveriesAll(ctx context.Context, messageID string, params *NotificationsDeliveriesParams, opts ...RequestOption) iter.Seq2[Delivery, error] {
	var p NotificationsDeliveriesParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Delivery, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.Deliveries(ctx, messageID, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// History: Bildirim geçmişi CSV işi (teslimatlar; tamamlanınca download_url)
//
// GET /v1/notifications/{messageId}/history
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *NotificationsService) History(ctx context.Context, messageID string, opts ...RequestOption) (*Export, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/notifications/{messageId}/history", queryAppID: true}
	op.pathParams = map[string]string{"messageId": messageID}
	var out Export
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ExportEvents: Mesaj olay CSV işi (gönderildi/teslim/açıldı/tıklandı/bounce…; tamamlanınca download_url)
//
// POST /v1/notifications/{messageId}/export_events
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *NotificationsService) ExportEvents(ctx context.Context, messageID string, opts ...RequestOption) (*Export, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/notifications/{messageId}/export_events", queryAppID: true}
	op.pathParams = map[string]string{"messageId": messageID}
	var out Export
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UsersService: Kullanıcı oluşturma/güncelleme (tag, özellik), arama ve abonelik ekleme.
type UsersService struct {
	client *Client
}

// UsersListParams: UsersService.List için isteğe bağlı sorgu parametreleri.
type UsersListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "last_active_at", "-last_active_at".
	Sort             *string
	Q                *string
	Tag              *string
	LastActiveAfter  *string
	LastActiveBefore *string
}

func (p *UsersListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "sort", p.Sort)
	addQuery(q, "q", p.Q)
	addQuery(q, "tag", p.Tag)
	addQuery(q, "last_active_after", p.LastActiveAfter)
	addQuery(q, "last_active_before", p.LastActiveBefore)
}

// List: Kullanıcı listesi
//
// GET /v1/apps/{appId}/users
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) List(ctx context.Context, params *UsersListParams, opts ...RequestOption) (*UserList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/users"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out UserList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *UsersService) ListAll(ctx context.Context, params *UsersListParams, opts ...RequestOption) iter.Seq2[User, error] {
	var p UsersListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]User, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Upsert: Kullanıcı oluştur/güncelle (external_id ile, isteğe bağlı aboneliklerle)
//
// POST /v1/apps/{appId}/users
//
// Idempotency-Key başlığı gönderilir: WithIdempotencyKey, gövdedeki idempotency_key veya otomatik UUID.
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) Upsert(ctx context.Context, body *UpsertUserDto, opts ...RequestOption) (*UserDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/users", idempotent: true}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out UserDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Search: Segment tanımı ile kullanıcı arama
//
// POST /v1/apps/{appId}/users/search
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) Search(ctx context.Context, body *UsersSearchBody, opts ...RequestOption) (*UserList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/users/search"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out UserList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetByAlias: Alias ile kullanıcı (hs_id, external_id veya özel etiket)
//
// GET /v1/apps/{appId}/users/by/{label}/{value}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) GetByAlias(ctx context.Context, label string, value string, opts ...RequestOption) (*UserDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/users/by/{label}/{value}"}
	op.pathParams = map[string]string{"label": label, "value": value}
	var out UserDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Tag/property merge-patch ve deltas
//
// PATCH /v1/apps/{appId}/users/by/{label}/{value}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) Update(ctx context.Context, label string, value string, body *PatchUserDto, opts ...RequestOption) (*UserDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/users/by/{label}/{value}"}
	op.pathParams = map[string]string{"label": label, "value": value}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out UserDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Kullanıcıyı sil (abonelikler silinir, olaylar anonimleşir)
//
// DELETE /v1/apps/{appId}/users/by/{label}/{value}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) Delete(ctx context.Context, label string, value string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/users/by/{label}/{value}"}
	op.pathParams = map[string]string{"label": label, "value": value}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetIdentity: Kullanıcının tüm kimlikleri (hs_id, external_id, özel alias'lar)
//
// GET /v1/apps/{appId}/users/by/{label}/{value}/identity
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) GetIdentity(ctx context.Context, label string, value string, opts ...RequestOption) (*UserIdentity, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/users/by/{label}/{value}/identity"}
	op.pathParams = map[string]string{"label": label, "value": value}
	var out UserIdentity
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Export: Kullanıcı verisini dışa aktar (KVKK/GDPR taşınabilirlik): profil, alias, abonelik, son olaylar
//
// GET /v1/apps/{appId}/users/by/{label}/{value}/export
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) Export(ctx context.Context, label string, value string, opts ...RequestOption) (*UserExport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/users/by/{label}/{value}/export"}
	op.pathParams = map[string]string{"label": label, "value": value}
	var out UserExport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AddSubscription: Kullanıcıya abonelik ekle
//
// POST /v1/apps/{appId}/users/by/{label}/{value}/subscriptions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *UsersService) AddSubscription(ctx context.Context, label string, value string, body *AddSubscriptionDto, opts ...RequestOption) (*Subscription, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/users/by/{label}/{value}/subscriptions"}
	op.pathParams = map[string]string{"label": label, "value": value}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Subscription
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AliasesService: Kullanıcı kimlikleri (alias): ekleme ve kaldırma.
type AliasesService struct {
	client *Client
}

// Add: Alias ekle
//
// PATCH /v1/apps/{appId}/users/by/{label}/{value}/identity
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AliasesService) Add(ctx context.Context, label string, value string, body *AddAliasDto, opts ...RequestOption) (*UserDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/users/by/{label}/{value}/identity"}
	op.pathParams = map[string]string{"label": label, "value": value}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out UserDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Alias kaldır
//
// DELETE /v1/apps/{appId}/users/by/{label}/{value}/identity/{aliasLabel}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AliasesService) Delete(ctx context.Context, label string, value string, aliasLabel string, opts ...RequestOption) (*UserDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/users/by/{label}/{value}/identity/{aliasLabel}"}
	op.pathParams = map[string]string{"label": label, "value": value, "aliasLabel": aliasLabel}
	var out UserDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SubscriptionsService: Abonelik (cihaz/e-posta/SMS) okuma, güncelleme, silme ve sahip atama.
type SubscriptionsService struct {
	client *Client
}

// Get: Abonelik detayı
//
// GET /v1/apps/{appId}/subscriptions/{subscriptionId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SubscriptionsService) Get(ctx context.Context, subscriptionID string, opts ...RequestOption) (*Subscription, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/subscriptions/{subscriptionId}"}
	op.pathParams = map[string]string{"subscriptionId": subscriptionID}
	var out Subscription
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Abonelik güncelle
//
// PATCH /v1/apps/{appId}/subscriptions/{subscriptionId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SubscriptionsService) Update(ctx context.Context, subscriptionID string, body *PatchSubscriptionDto, opts ...RequestOption) (*Subscription, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/subscriptions/{subscriptionId}"}
	op.pathParams = map[string]string{"subscriptionId": subscriptionID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Subscription
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Aboneliği sil
//
// DELETE /v1/apps/{appId}/subscriptions/{subscriptionId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SubscriptionsService) Delete(ctx context.Context, subscriptionID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/subscriptions/{subscriptionId}"}
	op.pathParams = map[string]string{"subscriptionId": subscriptionID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetUserIdentity: Aboneliğin sahibi kullanıcının kimlikleri
//
// GET /v1/apps/{appId}/subscriptions/{subscriptionId}/user/identity
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SubscriptionsService) GetUserIdentity(ctx context.Context, subscriptionID string, opts ...RequestOption) (*UserIdentity, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/subscriptions/{subscriptionId}/user/identity"}
	op.pathParams = map[string]string{"subscriptionId": subscriptionID}
	var out UserIdentity
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateByToken: Token ile abonelik güncelle (token_type: ios_push, android_push, web_push, email, sms veya eski sağlayıcı adları iOSPush, AndroidPush, ChromePush, FirefoxPush, SafariPush, Email, SMS)
//
// PATCH /v1/apps/{appId}/subscriptions_by_token/{tokenType}/{token}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SubscriptionsService) UpdateByToken(ctx context.Context, tokenType string, token string, body *PatchSubscriptionDto, opts ...RequestOption) (*Subscription, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/subscriptions_by_token/{tokenType}/{token}"}
	op.pathParams = map[string]string{"tokenType": tokenType, "token": token}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Subscription
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SetOwner: Aboneliği external_id ile kullanıcıya bağla
//
// POST /v1/apps/{appId}/subscriptions/{subscriptionId}/owner
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SubscriptionsService) SetOwner(ctx context.Context, subscriptionID string, body *SetOwnerDto, opts ...RequestOption) (*Subscription, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/subscriptions/{subscriptionId}/owner"}
	op.pathParams = map[string]string{"subscriptionId": subscriptionID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Subscription
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// EventsService: Özel olay gönderimi ve olay tanımları.
type EventsService struct {
	client *Client
}

// Send: Toplu olay gönder (en fazla 500)
//
// POST /v1/apps/{appId}/events
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EventsService) Send(ctx context.Context, body *APIEventsDto, opts ...RequestOption) (*EventsAccepted, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/events"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out EventsAccepted
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListDefinitions: Olay tanımları
//
// GET /v1/apps/{appId}/events/definitions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EventsService) ListDefinitions(ctx context.Context, opts ...RequestOption) (*EventDefinitionList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/events/definitions"}
	var out EventDefinitionList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetDefinition: Olay detayı
//
// GET /v1/apps/{appId}/events/definitions/{name}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EventsService) GetDefinition(ctx context.Context, name string, opts ...RequestOption) (*EventDefinitionDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/events/definitions/{name}"}
	op.pathParams = map[string]string{"name": name}
	var out EventDefinitionDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateDefinition: Görünen ad / dönüşüm işareti
//
// PATCH /v1/apps/{appId}/events/definitions/{name}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EventsService) UpdateDefinition(ctx context.Context, name string, body *PatchEventDefinitionDto, opts ...RequestOption) (*EventDefinition, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/events/definitions/{name}"}
	op.pathParams = map[string]string{"name": name}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out EventDefinition
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SegmentsService: Segment yönetimi, önizleme ve kitle tahmini.
type SegmentsService struct {
	client *Client
}

// SegmentsListParams: SegmentsService.List için isteğe bağlı sorgu parametreleri.
type SegmentsListParams struct {
	Limit  *int64
	Cursor *string
	Q      *string
}

func (p *SegmentsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "q", p.Q)
}

// List: Segment listesi
//
// GET /v1/apps/{appId}/segments
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SegmentsService) List(ctx context.Context, params *SegmentsListParams, opts ...RequestOption) (*SegmentPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/segments"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out SegmentPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *SegmentsService) ListAll(ctx context.Context, params *SegmentsListParams, opts ...RequestOption) iter.Seq2[Segment, error] {
	var p SegmentsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Segment, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Create: Segment oluştur
//
// POST /v1/apps/{appId}/segments
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SegmentsService) Create(ctx context.Context, body *SegmentsCreateBody, opts ...RequestOption) (*Segment, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/segments"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Segment
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Preview: Tanım ile sayım ve örnek kullanıcılar
//
// POST /v1/apps/{appId}/segments/preview
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SegmentsService) Preview(ctx context.Context, body *SegmentsPreviewBody, opts ...RequestOption) (*SegmentPreview, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/segments/preview"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out SegmentPreview
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Segment detayı
//
// GET /v1/apps/{appId}/segments/{segmentId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SegmentsService) Get(ctx context.Context, segmentID string, opts ...RequestOption) (*Segment, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/segments/{segmentId}"}
	op.pathParams = map[string]string{"segmentId": segmentID}
	var out Segment
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Segment güncelle
//
// PATCH /v1/apps/{appId}/segments/{segmentId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SegmentsService) Update(ctx context.Context, segmentID string, body *SegmentsUpdateBody, opts ...RequestOption) (*Segment, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/segments/{segmentId}"}
	op.pathParams = map[string]string{"segmentId": segmentID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Segment
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Segment sil
//
// DELETE /v1/apps/{appId}/segments/{segmentId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SegmentsService) Delete(ctx context.Context, segmentID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/segments/{segmentId}"}
	op.pathParams = map[string]string{"segmentId": segmentID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UsedBy: Segmenti kullanan kampanyalar
//
// GET /v1/apps/{appId}/segments/{segmentId}/campaigns
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SegmentsService) UsedBy(ctx context.Context, segmentID string, opts ...RequestOption) (*SegmentCampaigns, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/segments/{segmentId}/campaigns"}
	op.pathParams = map[string]string{"segmentId": segmentID}
	var out SegmentCampaigns
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SegmentsEstimateParams: SegmentsService.Estimate için isteğe bağlı sorgu parametreleri.
type SegmentsEstimateParams struct {
	// true: önbelleği atla, tam sayım Değerler: "true", "false".
	Force *string
}

func (p *SegmentsEstimateParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "force", p.Force)
}

// Estimate: Tahmini sayım
//
// POST /v1/apps/{appId}/segments/{segmentId}/estimate
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SegmentsService) Estimate(ctx context.Context, segmentID string, params *SegmentsEstimateParams, opts ...RequestOption) (*SegmentEstimate, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/segments/{segmentId}/estimate"}
	op.pathParams = map[string]string{"segmentId": segmentID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out SegmentEstimate
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// TemplatesService: Şablon yönetimi, sürümler ve önizleme.
type TemplatesService struct {
	client *Client
}

// TemplatesListParams: TemplatesService.List için isteğe bağlı sorgu parametreleri.
type TemplatesListParams struct {
	Limit  *int64
	Cursor *string
	Q      *string
	// Değerler: "ios_push", "android_push", "web_push", "email", "sms".
	Channel *string
}

func (p *TemplatesListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "q", p.Q)
	addQuery(q, "channel", p.Channel)
}

// List: Şablon listesi
//
// GET /v1/apps/{appId}/templates
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) List(ctx context.Context, params *TemplatesListParams, opts ...RequestOption) (*TemplatePage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/templates"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out TemplatePage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *TemplatesService) ListAll(ctx context.Context, params *TemplatesListParams, opts ...RequestOption) iter.Seq2[Template, error] {
	var p TemplatesListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Template, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Create: Şablon oluştur
//
// POST /v1/apps/{appId}/templates
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) Create(ctx context.Context, body *CreateTemplateDto, opts ...RequestOption) (*Template, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/templates"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Template
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Şablon detayı
//
// GET /v1/apps/{appId}/templates/{templateId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) Get(ctx context.Context, templateID string, opts ...RequestOption) (*Template, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/templates/{templateId}"}
	op.pathParams = map[string]string{"templateId": templateID}
	var out Template
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Şablon güncelle
//
// PATCH /v1/apps/{appId}/templates/{templateId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) Update(ctx context.Context, templateID string, body *UpdateTemplateDto, opts ...RequestOption) (*Template, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/templates/{templateId}"}
	op.pathParams = map[string]string{"templateId": templateID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Template
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Şablon sil
//
// DELETE /v1/apps/{appId}/templates/{templateId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) Delete(ctx context.Context, templateID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/templates/{templateId}"}
	op.pathParams = map[string]string{"templateId": templateID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Stats: Şablon analitiği
//
// GET /v1/apps/{appId}/templates/{templateId}/stats
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) Stats(ctx context.Context, templateID string, opts ...RequestOption) (*TemplateStats, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/templates/{templateId}/stats"}
	op.pathParams = map[string]string{"templateId": templateID}
	var out TemplateStats
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CopyToApp: Şablonu aynı organizasyondaki başka bir app'e kopyala
//
// POST /v1/apps/{appId}/templates/{templateId}/copy_to_app
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) CopyToApp(ctx context.Context, templateID string, body *CopyTemplateDto, opts ...RequestOption) (*Template, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/templates/{templateId}/copy_to_app"}
	op.pathParams = map[string]string{"templateId": templateID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Template
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Versions: Sürüm geçmişi
//
// GET /v1/apps/{appId}/templates/{templateId}/versions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) Versions(ctx context.Context, templateID string, opts ...RequestOption) (*TemplateVersionList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/templates/{templateId}/versions"}
	op.pathParams = map[string]string{"templateId": templateID}
	var out TemplateVersionList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Render: Şablonu örnek verilerle render et
//
// POST /v1/apps/{appId}/templates/{templateId}/render
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *TemplatesService) Render(ctx context.Context, templateID string, body *RenderTemplateDto, opts ...RequestOption) (*RenderPreview, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/templates/{templateId}/render"}
	op.pathParams = map[string]string{"templateId": templateID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out RenderPreview
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CampaignsService: Kampanya yaşam döngüsü (zamanlama, gönderim, duraklatma, onay) ve istatistikler.
type CampaignsService struct {
	client *Client
}

// CampaignsListParams: CampaignsService.List için isteğe bağlı sorgu parametreleri.
type CampaignsListParams struct {
	Limit  *int64
	Cursor *string
	Q      *string
	// Değerler: "draft", "scheduled", "running", "paused", "completed", "cancelled", "failed".
	Status *string
	// Değerler: "one_time", "scheduled", "recurring", "automated", "transactional".
	Type *string
	// Değerler: "ios_push", "android_push", "web_push", "email", "sms".
	Channel *string
	Label   *string
}

func (p *CampaignsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "q", p.Q)
	addQuery(q, "status", p.Status)
	addQuery(q, "type", p.Type)
	addQuery(q, "channel", p.Channel)
	addQuery(q, "label", p.Label)
}

// List: Kampanya listesi
//
// GET /v1/apps/{appId}/campaigns
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) List(ctx context.Context, params *CampaignsListParams, opts ...RequestOption) (*CampaignPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/campaigns"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out CampaignPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *CampaignsService) ListAll(ctx context.Context, params *CampaignsListParams, opts ...RequestOption) iter.Seq2[Campaign, error] {
	var p CampaignsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Campaign, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Create: Kampanya oluştur
//
// POST /v1/apps/{appId}/campaigns
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Create(ctx context.Context, body *CampaignsCreateBody, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Kampanya detayı
//
// GET /v1/apps/{appId}/campaigns/{campaignId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Get(ctx context.Context, campaignID string, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/campaigns/{campaignId}"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Kampanya güncelle
//
// PATCH /v1/apps/{appId}/campaigns/{campaignId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Update(ctx context.Context, campaignID string, body *CampaignsUpdateBody, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/campaigns/{campaignId}"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Kampanya sil
//
// DELETE /v1/apps/{appId}/campaigns/{campaignId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Delete(ctx context.Context, campaignID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/campaigns/{campaignId}"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Schedule: Zamanla
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/schedule
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Schedule(ctx context.Context, campaignID string, body *ScheduleCampaignDto, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/schedule"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Send: Hemen gönder
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/send
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Send(ctx context.Context, campaignID string, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/send"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Pause: Duraklat
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/pause
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Pause(ctx context.Context, campaignID string, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/pause"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Resume: Sürdür
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/resume
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Resume(ctx context.Context, campaignID string, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/resume"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Cancel: İptal
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/cancel
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Cancel(ctx context.Context, campaignID string, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/cancel"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Duplicate: Çoğalt
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/duplicate
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Duplicate(ctx context.Context, campaignID string, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/duplicate"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Approve: Onayla
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/approve
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Approve(ctx context.Context, campaignID string, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/approve"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Stats: İstatistikler
//
// GET /v1/apps/{appId}/campaigns/{campaignId}/stats
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Stats(ctx context.Context, campaignID string, opts ...RequestOption) (*CampaignStats, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/campaigns/{campaignId}/stats"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out CampaignStats
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// VariantStats: A/B varyant karşılaştırması
//
// GET /v1/apps/{appId}/campaigns/{campaignId}/variants/stats
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) VariantStats(ctx context.Context, campaignID string, opts ...RequestOption) (*CampaignVariantStats, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/campaigns/{campaignId}/variants/stats"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out CampaignVariantStats
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SelectWinner: A/B kazananını seç
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/winner
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) SelectWinner(ctx context.Context, campaignID string, body *SelectWinnerDto, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/winner"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CampaignsDeliveriesParams: CampaignsService.Deliveries için isteğe bağlı sorgu parametreleri.
type CampaignsDeliveriesParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "pending", "sent", "delivered", "failed", "bounced", "suppressed", "expired", "cancelled".
	Status *string
}

func (p *CampaignsDeliveriesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
}

// Deliveries: Abone bazlı teslimatlar
//
// GET /v1/apps/{appId}/campaigns/{campaignId}/deliveries
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Deliveries(ctx context.Context, campaignID string, params *CampaignsDeliveriesParams, opts ...RequestOption) (*DeliveryList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/campaigns/{campaignId}/deliveries"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out DeliveryList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeliveriesAll: Deliveries sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *CampaignsService) DeliveriesAll(ctx context.Context, campaignID string, params *CampaignsDeliveriesParams, opts ...RequestOption) iter.Seq2[Delivery, error] {
	var p CampaignsDeliveriesParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Delivery, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.Deliveries(ctx, campaignID, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// ResendFailed: Başarısız teslimatlara yeniden gönder
//
// POST /v1/apps/{appId}/campaigns/{campaignId}/resend-failed
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) ResendFailed(ctx context.Context, campaignID string, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/campaigns/{campaignId}/resend-failed"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out Campaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Messages: Kampanya çalıştırmaları
//
// GET /v1/apps/{appId}/campaigns/{campaignId}/messages
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CampaignsService) Messages(ctx context.Context, campaignID string, opts ...RequestOption) (*MessageList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/campaigns/{campaignId}/messages"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out MessageList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// InAppMessagesService: Uygulama içi (in-app) mesaj yönetimi ve durum değişiklikleri.
type InAppMessagesService struct {
	client *Client
}

// InAppMessagesListParams: InAppMessagesService.List için isteğe bağlı sorgu parametreleri.
type InAppMessagesListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "draft", "active", "paused", "archived".
	Status *string
}

func (p *InAppMessagesListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
}

// List: In-app mesajlar (durum filtresi; istatistik dahil)
//
// GET /v1/apps/{appId}/in-app-messages
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InAppMessagesService) List(ctx context.Context, params *InAppMessagesListParams, opts ...RequestOption) (*InAppMessagePage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/in-app-messages"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out InAppMessagePage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *InAppMessagesService) ListAll(ctx context.Context, params *InAppMessagesListParams, opts ...RequestOption) iter.Seq2[InAppMessage, error] {
	var p InAppMessagesListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]InAppMessage, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Create: In-app mesaj oluştur (taslak)
//
// POST /v1/apps/{appId}/in-app-messages
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InAppMessagesService) Create(ctx context.Context, body *CreateInAppDto, opts ...RequestOption) (*InAppMessage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/in-app-messages"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out InAppMessage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: In-app mesaj detayı ve istatistikleri
//
// GET /v1/apps/{appId}/in-app-messages/{messageId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InAppMessagesService) Get(ctx context.Context, messageID string, opts ...RequestOption) (*InAppMessage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/in-app-messages/{messageId}"}
	op.pathParams = map[string]string{"messageId": messageID}
	var out InAppMessage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: In-app mesajı düzenle
//
// PATCH /v1/apps/{appId}/in-app-messages/{messageId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InAppMessagesService) Update(ctx context.Context, messageID string, body *UpdateInAppDto, opts ...RequestOption) (*InAppMessage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/in-app-messages/{messageId}"}
	op.pathParams = map[string]string{"messageId": messageID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out InAppMessage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: In-app mesajı sil
//
// DELETE /v1/apps/{appId}/in-app-messages/{messageId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InAppMessagesService) Delete(ctx context.Context, messageID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/in-app-messages/{messageId}"}
	op.pathParams = map[string]string{"messageId": messageID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Status: Yayınla / duraklat / arşivle
//
// POST /v1/apps/{appId}/in-app-messages/{messageId}/status
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InAppMessagesService) Status(ctx context.Context, messageID string, body *InAppStatusDto, opts ...RequestOption) (*InAppMessage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/in-app-messages/{messageId}/status"}
	op.pathParams = map[string]string{"messageId": messageID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out InAppMessage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// InboxService: Kullanıcı bildirim kutusu (inbox): listeleme, okundu işaretleme ve arşivleme.
type InboxService struct {
	client *Client
}

// InboxListParams: InboxService.List için isteğe bağlı sorgu parametreleri.
type InboxListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "true", "false".
	IncludeArchived *string
}

func (p *InboxListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "include_archived", p.IncludeArchived)
}

// List: Kullanıcının bildirim kutusu (okunmamış sayısı dahil)
//
// GET /v1/apps/{appId}/users/by/{label}/{value}/inbox
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) List(ctx context.Context, label string, value string, params *InboxListParams, opts ...RequestOption) (*InboxPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/users/by/{label}/{value}/inbox"}
	op.pathParams = map[string]string{"label": label, "value": value}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out InboxPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *InboxService) ListAll(ctx context.Context, label string, value string, params *InboxListParams, opts ...RequestOption) iter.Seq2[InboxItem, error] {
	var p InboxListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]InboxItem, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, label, value, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// ReadAll: Tümünü okundu işaretle
//
// POST /v1/apps/{appId}/users/by/{label}/{value}/inbox/read-all
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) ReadAll(ctx context.Context, label string, value string, opts ...RequestOption) (*InboxReadAllResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/users/by/{label}/{value}/inbox/read-all"}
	op.pathParams = map[string]string{"label": label, "value": value}
	var out InboxReadAllResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Read: Okundu işaretle
//
// POST /v1/apps/{appId}/users/by/{label}/{value}/inbox/{itemId}/read
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) Read(ctx context.Context, label string, value string, itemID string, opts ...RequestOption) (*InboxItem, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/users/by/{label}/{value}/inbox/{itemId}/read"}
	op.pathParams = map[string]string{"label": label, "value": value, "itemId": itemID}
	var out InboxItem
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Archive: Kutudan kaldır (arşivle)
//
// DELETE /v1/apps/{appId}/users/by/{label}/{value}/inbox/{itemId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) Archive(ctx context.Context, label string, value string, itemID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/users/by/{label}/{value}/inbox/{itemId}"}
	op.pathParams = map[string]string{"label": label, "value": value, "itemId": itemID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SuppressionsService: E-posta/SMS gönderim engeli (suppression) listesi.
type SuppressionsService struct {
	client *Client
}

// SuppressionsListParams: SuppressionsService.List için isteğe bağlı sorgu parametreleri.
type SuppressionsListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "email", "sms".
	Channel *string
	// Değerler: "bounce", "complaint", "manual", "unsubscribe_link", "stop".
	Reason *string
	Value  *string
}

func (p *SuppressionsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "channel", p.Channel)
	addQuery(q, "reason", p.Reason)
	addQuery(q, "value", p.Value)
}

// List: Suppression list (e-posta/SMS gönderilmeyecek adresler; kanal/neden/değer filtresi)
//
// GET /v1/apps/{appId}/suppressions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SuppressionsService) List(ctx context.Context, params *SuppressionsListParams, opts ...RequestOption) (*SuppressionPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/suppressions"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out SuppressionPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *SuppressionsService) ListAll(ctx context.Context, params *SuppressionsListParams, opts ...RequestOption) iter.Seq2[Suppression, error] {
	var p SuppressionsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Suppression, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Create: Adres ekle (manuel engel; değer saklanmaz, özet + maskeli gösterim)
//
// POST /v1/apps/{appId}/suppressions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SuppressionsService) Create(ctx context.Context, body *CreateSuppressionDto, opts ...RequestOption) (*Suppression, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/suppressions"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Suppression
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Kaydı kaldır (adres yeniden gönderim alabilir)
//
// DELETE /v1/apps/{appId}/suppressions/{suppressionId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SuppressionsService) Delete(ctx context.Context, suppressionID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/suppressions/{suppressionId}"}
	op.pathParams = map[string]string{"suppressionId": suppressionID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ExportsService: Veri dışa aktarma işleri.
type ExportsService struct {
	client *Client
}

// Create: Export işi başlat
//
// POST /v1/apps/{appId}/exports
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ExportsService) Create(ctx context.Context, body *CreateExportDto, opts ...RequestOption) (*Export, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/exports"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Export
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ExportsListParams: ExportsService.List için isteğe bağlı sorgu parametreleri.
type ExportsListParams struct {
	Limit  *int64
	Cursor *string
}

func (p *ExportsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
}

// List: Export işleri
//
// GET /v1/apps/{appId}/exports
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ExportsService) List(ctx context.Context, params *ExportsListParams, opts ...RequestOption) (*ExportList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/exports"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out ExportList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *ExportsService) ListAll(ctx context.Context, params *ExportsListParams, opts ...RequestOption) iter.Seq2[Export, error] {
	var p ExportsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Export, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Get: Export durumu (download_url)
//
// GET /v1/apps/{appId}/exports/{exportId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ExportsService) Get(ctx context.Context, exportID string, opts ...RequestOption) (*Export, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/exports/{exportId}"}
	op.pathParams = map[string]string{"exportId": exportID}
	var out Export
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ImportsService: Veri içe aktarma işleri (CSV, multipart).
type ImportsService struct {
	client *Client
}

// ImportsCreateForm: ImportsService.Create için multipart/form-data alanları.
type ImportsCreateForm struct {
	// Zorunlu. Değerler: "eski sağlayıcı_players", "csv_users".
	Type string
	// Zorunlu. CSV dosyası (UTF-8, en fazla 50 MB)
	File *FileUpload
}

// Create: Import işi başlat (multipart: type, file)
//
// POST /v1/apps/{appId}/imports
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ImportsService) Create(ctx context.Context, form *ImportsCreateForm, opts ...RequestOption) (*Import, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/imports"}
	if form == nil {
		return nil, errNilBody
	}
	mp := &multipartBody{}
	mp.field("type", form.Type)
	mp.file("file", form.File)
	op.multipart = mp
	var out Import
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ImportsListParams: ImportsService.List için isteğe bağlı sorgu parametreleri.
type ImportsListParams struct {
	Limit  *int64
	Cursor *string
}

func (p *ImportsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
}

// List: Import işleri
//
// GET /v1/apps/{appId}/imports
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ImportsService) List(ctx context.Context, params *ImportsListParams, opts ...RequestOption) (*ImportList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/imports"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out ImportList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *ImportsService) ListAll(ctx context.Context, params *ImportsListParams, opts ...RequestOption) iter.Seq2[Import, error] {
	var p ImportsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Import, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Get: Import durumu
//
// GET /v1/apps/{appId}/imports/{importId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ImportsService) Get(ctx context.Context, importID string, opts ...RequestOption) (*Import, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/imports/{importId}"}
	op.pathParams = map[string]string{"importId": importID}
	var out Import
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// WebhooksService: Giden webhook uç noktaları, teslimatlar ve yeniden gönderim.
type WebhooksService struct {
	client *Client
}

// List: Webhook endpoint'leri
//
// GET /v1/apps/{appId}/webhooks
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) List(ctx context.Context, opts ...RequestOption) (*WebhookList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/webhooks"}
	var out WebhookList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create: Endpoint oluştur (secret yalnızca bu yanıtta döner)
//
// POST /v1/apps/{appId}/webhooks
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) Create(ctx context.Context, body *CreateWebhookDto, opts ...RequestOption) (*WebhookWithSecret, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/webhooks"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out WebhookWithSecret
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Endpoint detayı
//
// GET /v1/apps/{appId}/webhooks/{webhookId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) Get(ctx context.Context, webhookID string, opts ...RequestOption) (*Webhook, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/webhooks/{webhookId}"}
	op.pathParams = map[string]string{"webhookId": webhookID}
	var out Webhook
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Endpoint güncelle (URL, olaylar, batch, başlıklar, etkinleştir/durdur)
//
// PATCH /v1/apps/{appId}/webhooks/{webhookId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) Update(ctx context.Context, webhookID string, body *UpdateWebhookDto, opts ...RequestOption) (*Webhook, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/webhooks/{webhookId}"}
	op.pathParams = map[string]string{"webhookId": webhookID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Webhook
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Endpoint sil
//
// DELETE /v1/apps/{appId}/webhooks/{webhookId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) Delete(ctx context.Context, webhookID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/webhooks/{webhookId}"}
	op.pathParams = map[string]string{"webhookId": webhookID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RotateSecret: Secret rotasyonu (eski secret 24 saat ikinci imza olarak gönderilir)
//
// POST /v1/apps/{appId}/webhooks/{webhookId}/rotate-secret
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) RotateSecret(ctx context.Context, webhookID string, opts ...RequestOption) (*WebhookWithSecret, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/webhooks/{webhookId}/rotate-secret"}
	op.pathParams = map[string]string{"webhookId": webhookID}
	var out WebhookWithSecret
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Test: Test olayı gönder (anında, sonuç döner)
//
// POST /v1/apps/{appId}/webhooks/{webhookId}/test
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) Test(ctx context.Context, webhookID string, opts ...RequestOption) (*WebhookDeliveryDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/webhooks/{webhookId}/test"}
	op.pathParams = map[string]string{"webhookId": webhookID}
	var out WebhookDeliveryDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// WebhooksDeliveriesParams: WebhooksService.Deliveries için isteğe bağlı sorgu parametreleri.
type WebhooksDeliveriesParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "pending", "success", "failed", "dead".
	Status    *string
	EventType *string
}

func (p *WebhooksDeliveriesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
	addQuery(q, "event_type", p.EventType)
}

// Deliveries: Teslim geçmişi (cursor; durum/olay türü filtresi)
//
// GET /v1/apps/{appId}/webhooks/{webhookId}/deliveries
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) Deliveries(ctx context.Context, webhookID string, params *WebhooksDeliveriesParams, opts ...RequestOption) (*WebhookDeliveryPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/webhooks/{webhookId}/deliveries"}
	op.pathParams = map[string]string{"webhookId": webhookID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out WebhookDeliveryPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeliveriesAll: Deliveries sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *WebhooksService) DeliveriesAll(ctx context.Context, webhookID string, params *WebhooksDeliveriesParams, opts ...RequestOption) iter.Seq2[WebhookDeliveryPageDataItem, error] {
	var p WebhooksDeliveriesParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]WebhookDeliveryPageDataItem, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.Deliveries(ctx, webhookID, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Replay: Teslimi yeniden gönder (aynı olay kimliğiyle)
//
// POST /v1/apps/{appId}/webhooks/deliveries/{deliveryId}/replay
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) Replay(ctx context.Context, deliveryID string, opts ...RequestOption) (*WebhookDeliveryDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/webhooks/deliveries/{deliveryId}/replay"}
	op.pathParams = map[string]string{"deliveryId": deliveryID}
	var out WebhookDeliveryDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ChannelsService: Kanal (APNs, FCM, Web Push, e-posta, SMS) yapılandırması ve doğrulaması.
type ChannelsService struct {
	client *Client
}

// List: Kanal yapılandırmaları
//
// GET /v1/apps/{appId}/channels
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ChannelsService) List(ctx context.Context, opts ...RequestOption) (*ChannelList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/channels"}
	var out ChannelList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Upsert: Kimlik bilgisi yükle (APNs p8, FCM servis hesabı, VAPID, mock)
//
// PUT /v1/apps/{appId}/channels/{channel}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ChannelsService) Upsert(ctx context.Context, channel string, body any, opts ...RequestOption) (*ChannelConfig, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PUT", path: "/v1/apps/{appId}/channels/{channel}"}
	op.pathParams = map[string]string{"channel": channel}
	op.body = body
	var out ChannelConfig
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Validate: Sağlayıcıya test bağlantısı
//
// POST /v1/apps/{appId}/channels/{channel}/validate
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ChannelsService) Validate(ctx context.Context, channel string, opts ...RequestOption) (*ChannelValidation, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/channels/{channel}/validate"}
	op.pathParams = map[string]string{"channel": channel}
	var out ChannelValidation
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AnalyticsService: Analitik raporları.
type AnalyticsService struct {
	client *Client
}

// AnalyticsOverviewParams: AnalyticsService.Overview için isteğe bağlı sorgu parametreleri.
type AnalyticsOverviewParams struct {
	// Değerler: "7d", "30d".
	Range *string
	From  *string
	To    *string
}

func (p *AnalyticsOverviewParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "range", p.Range)
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
}

// Overview: Genel bakış KPI
//
// GET /v1/apps/{appId}/analytics/overview
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Overview(ctx context.Context, params *AnalyticsOverviewParams, opts ...RequestOption) (*OverviewStats, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/analytics/overview"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out OverviewStats
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AnalyticsMessagesParams: AnalyticsService.Messages için isteğe bağlı sorgu parametreleri.
type AnalyticsMessagesParams struct {
	From *string
	To   *string
	// Değerler: "true", "false".
	Compare *string
	// Değerler: "json", "csv".
	Format *string
	// Değerler: "day", "channel", "campaign".
	GroupBy *string
}

func (p *AnalyticsMessagesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
	addQuery(q, "compare", p.Compare)
	addQuery(q, "format", p.Format)
	addQuery(q, "group_by", p.GroupBy)
}

// Messages: Teslimat raporu
//
// GET /v1/apps/{appId}/analytics/messages
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Messages(ctx context.Context, params *AnalyticsMessagesParams, opts ...RequestOption) (*MessagesReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/analytics/messages"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out MessagesReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AnalyticsSubscriptionsParams: AnalyticsService.Subscriptions için isteğe bağlı sorgu parametreleri.
type AnalyticsSubscriptionsParams struct {
	From *string
	To   *string
	// Değerler: "true", "false".
	Compare *string
	// Değerler: "json", "csv".
	Format *string
}

func (p *AnalyticsSubscriptionsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
	addQuery(q, "compare", p.Compare)
	addQuery(q, "format", p.Format)
}

// Subscriptions: Abonelik büyümesi
//
// GET /v1/apps/{appId}/analytics/subscriptions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Subscriptions(ctx context.Context, params *AnalyticsSubscriptionsParams, opts ...RequestOption) (*SubscriptionsReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/analytics/subscriptions"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out SubscriptionsReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AnalyticsEventsParams: AnalyticsService.Events için isteğe bağlı sorgu parametreleri.
type AnalyticsEventsParams struct {
	From *string
	To   *string
	// Değerler: "true", "false".
	Compare *string
	// Değerler: "json", "csv".
	Format *string
	Name   *string
	// Değerler: "day", "name".
	GroupBy *string
}

func (p *AnalyticsEventsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
	addQuery(q, "compare", p.Compare)
	addQuery(q, "format", p.Format)
	addQuery(q, "name", p.Name)
	addQuery(q, "group_by", p.GroupBy)
}

// Events: Olay hacmi
//
// GET /v1/apps/{appId}/analytics/events
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Events(ctx context.Context, params *AnalyticsEventsParams, opts ...RequestOption) (*EventsReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/analytics/events"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out EventsReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AnalyticsConversionsParams: AnalyticsService.Conversions için isteğe bağlı sorgu parametreleri.
type AnalyticsConversionsParams struct {
	From *string
	To   *string
	// Değerler: "true", "false".
	Compare *string
	// Değerler: "json", "csv".
	Format     *string
	CampaignID *string
}

func (p *AnalyticsConversionsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
	addQuery(q, "compare", p.Compare)
	addQuery(q, "format", p.Format)
	addQuery(q, "campaign_id", p.CampaignID)
}

// Conversions: Dönüşüm/atribüsyon
//
// GET /v1/apps/{appId}/analytics/conversions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Conversions(ctx context.Context, params *AnalyticsConversionsParams, opts ...RequestOption) (*ConversionsReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/analytics/conversions"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out ConversionsReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Funnel: Funnel
//
// POST /v1/apps/{appId}/analytics/funnels
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Funnel(ctx context.Context, body *FunnelBodyDto, opts ...RequestOption) (*FunnelReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/analytics/funnels"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out FunnelReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AnalyticsProvidersParams: AnalyticsService.Providers için isteğe bağlı sorgu parametreleri.
type AnalyticsProvidersParams struct {
	From *string
	To   *string
	// Değerler: "true", "false".
	Compare *string
	// Değerler: "json", "csv".
	Format *string
}

func (p *AnalyticsProvidersParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
	addQuery(q, "compare", p.Compare)
	addQuery(q, "format", p.Format)
}

// Providers: Sağlayıcı sağlığı
//
// GET /v1/apps/{appId}/analytics/providers
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Providers(ctx context.Context, params *AnalyticsProvidersParams, opts ...RequestOption) (*ProvidersReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/analytics/providers"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out ProvidersReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AnalyticsRetentionParams: AnalyticsService.Retention için isteğe bağlı sorgu parametreleri.
type AnalyticsRetentionParams struct {
	Weeks *int64
	// Değerler: "json", "csv".
	Format *string
}

func (p *AnalyticsRetentionParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "weeks", p.Weeks)
	addQuery(q, "format", p.Format)
}

// Retention: Retention kohortu
//
// GET /v1/apps/{appId}/analytics/retention
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Retention(ctx context.Context, params *AnalyticsRetentionParams, opts ...RequestOption) (*RetentionReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/analytics/retention"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out RetentionReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AnalyticsEngagementParams: AnalyticsService.Engagement için isteğe bağlı sorgu parametreleri.
type AnalyticsEngagementParams struct {
	From *string
	To   *string
	// Değerler: "true", "false".
	Compare *string
	// Değerler: "json", "csv".
	Format *string
}

func (p *AnalyticsEngagementParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
	addQuery(q, "compare", p.Compare)
	addQuery(q, "format", p.Format)
}

// Engagement: Etkileşim ısı haritası
//
// GET /v1/apps/{appId}/analytics/engagement
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AnalyticsService) Engagement(ctx context.Context, params *AnalyticsEngagementParams, opts ...RequestOption) (*EngagementReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/analytics/engagement"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out EngagementReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AppsService: App listesi ve ayarları.
type AppsService struct {
	client *Client
}

// List: App listesi (org geneli admin anahtarı)
//
// GET /v1/apps
func (s *AppsService) List(ctx context.Context, opts ...RequestOption) (*AppList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps"}
	var out AppList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create: App oluştur (org geneli admin anahtarı)
//
// POST /v1/apps
func (s *AppsService) Create(ctx context.Context, body *CreateAppDto, opts ...RequestOption) (*App, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out App
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: App detayı
//
// GET /v1/apps/{appId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AppsService) Get(ctx context.Context, opts ...RequestOption) (*App, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}"}
	var out App
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: App ayarlarını güncelle
//
// PATCH /v1/apps/{appId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AppsService) Update(ctx context.Context, body *UpdateAppDto, opts ...RequestOption) (*App, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out App
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// APIKeysService: App API anahtarları: oluşturma, kapsam/IP izin listesi güncelleme, rotasyon ve iptal.
type APIKeysService struct {
	client *Client
}

// List: App'e bağlı API anahtarları
//
// GET /v1/apps/{appId}/api-keys
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *APIKeysService) List(ctx context.Context, opts ...RequestOption) (*APIKeyList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/api-keys"}
	var out APIKeyList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create: App'e bağlı API anahtarı oluştur (ham anahtar yalnızca bu yanıtta)
//
// POST /v1/apps/{appId}/api-keys
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *APIKeysService) Create(ctx context.Context, body *CreateAppAPIKeyDto, opts ...RequestOption) (*CreatedAPIKey, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/api-keys"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out CreatedAPIKey
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Anahtarı güncelle (ad, kapsam, IP izin listesi, bitiş)
//
// PATCH /v1/apps/{appId}/api-keys/{keyId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *APIKeysService) Update(ctx context.Context, keyID string, body *UpdateAPIKeyDto, opts ...RequestOption) (*APIKey, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/api-keys/{keyId}"}
	op.pathParams = map[string]string{"keyId": keyID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out APIKey
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Revoke: Anahtarı iptal et
//
// DELETE /v1/apps/{appId}/api-keys/{keyId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *APIKeysService) Revoke(ctx context.Context, keyID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/api-keys/{keyId}"}
	op.pathParams = map[string]string{"keyId": keyID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Rotate: Anahtarı döndür (eski değer anında geçersiz)
//
// POST /v1/apps/{appId}/api-keys/{keyId}/rotate
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *APIKeysService) Rotate(ctx context.Context, keyID string, opts ...RequestOption) (*CreatedAPIKey, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/api-keys/{keyId}/rotate"}
	op.pathParams = map[string]string{"keyId": keyID}
	var out CreatedAPIKey
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// OrganizationService: Organizasyon, üyeler ve denetim kayıtları.
type OrganizationService struct {
	client *Client
}

// Get: Anahtarın ait olduğu organizasyon (org geneli admin anahtarı)
//
// GET /v1/organization
func (s *OrganizationService) Get(ctx context.Context, opts ...RequestOption) (*Org, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/organization"}
	var out Org
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Members: Organizasyon üyeleri (org geneli admin anahtarı)
//
// GET /v1/organization/members
func (s *OrganizationService) Members(ctx context.Context, opts ...RequestOption) (*MemberList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/organization/members"}
	var out MemberList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// OrganizationAuditLogsParams: OrganizationService.AuditLogs için isteğe bağlı sorgu parametreleri.
type OrganizationAuditLogsParams struct {
	Limit   *int64
	Cursor  *string
	Action  *string
	ActorID *string
	AppID   *string
	From    *string
	To      *string
}

func (p *OrganizationAuditLogsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "action", p.Action)
	addQuery(q, "actor_id", p.ActorID)
	addQuery(q, "app_id", p.AppID)
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
}

// AuditLogs: Audit log (org geneli admin anahtarı; action/actor/app/tarih filtreleri, cursor)
//
// GET /v1/organization/audit-logs
func (s *OrganizationService) AuditLogs(ctx context.Context, params *OrganizationAuditLogsParams, opts ...RequestOption) (*AuditPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/organization/audit-logs"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out AuditPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AuditLogsAll: AuditLogs sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *OrganizationService) AuditLogsAll(ctx context.Context, params *OrganizationAuditLogsParams, opts ...RequestOption) iter.Seq2[AuditEntry, error] {
	var p OrganizationAuditLogsParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]AuditEntry, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.AuditLogs(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// CategoriesService: categories kaynağı.
type CategoriesService struct {
	client *Client
}

// List: Mesaj kategorileri
//
// GET /v1/apps/{appId}/categories
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CategoriesService) List(ctx context.Context, opts ...RequestOption) (*MessageCategoryList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/categories"}
	var out MessageCategoryList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create: Kategori oluştur
//
// POST /v1/apps/{appId}/categories
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CategoriesService) Create(ctx context.Context, body *CreateCategoryDto, opts ...RequestOption) (*MessageCategory, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/categories"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out MessageCategory
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Kategori güncelle
//
// PATCH /v1/apps/{appId}/categories/{categoryId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CategoriesService) Update(ctx context.Context, categoryID string, body *UpdateCategoryDto, opts ...RequestOption) (*MessageCategory, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/categories/{categoryId}"}
	op.pathParams = map[string]string{"categoryId": categoryID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out MessageCategory
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Kategori sil
//
// DELETE /v1/apps/{appId}/categories/{categoryId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CategoriesService) Delete(ctx context.Context, categoryID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/categories/{categoryId}"}
	op.pathParams = map[string]string{"categoryId": categoryID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// JourneysService: journeys kaynağı.
type JourneysService struct {
	client *Client
}

// JourneysListParams: JourneysService.List için isteğe bağlı sorgu parametreleri.
type JourneysListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "draft", "active", "paused", "archived".
	Status *string
	Q      *string
}

func (p *JourneysListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
	addQuery(q, "q", p.Q)
}

// List: Journey listesi
//
// GET /v1/apps/{appId}/journeys
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) List(ctx context.Context, params *JourneysListParams, opts ...RequestOption) (*JourneyPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/journeys"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out JourneyPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *JourneysService) ListAll(ctx context.Context, params *JourneysListParams, opts ...RequestOption) iter.Seq2[Journey, error] {
	var p JourneysListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Journey, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.List(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Create: Journey oluştur (taslak)
//
// POST /v1/apps/{appId}/journeys
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Create(ctx context.Context, body *JourneysCreateBody, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/journeys"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Journey detayı
//
// GET /v1/apps/{appId}/journeys/{journeyId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Get(ctx context.Context, journeyID string, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/journeys/{journeyId}"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Ad / açıklama
//
// PATCH /v1/apps/{appId}/journeys/{journeyId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Update(ctx context.Context, journeyID string, body *UpdateJourneyDto, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/journeys/{journeyId}"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Hiç yayınlanmamış journey'i sil
//
// DELETE /v1/apps/{appId}/journeys/{journeyId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Delete(ctx context.Context, journeyID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/journeys/{journeyId}"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SaveDraft: Taslağı kaydet
//
// PUT /v1/apps/{appId}/journeys/{journeyId}/draft
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) SaveDraft(ctx context.Context, journeyID string, body *JourneysSaveDraftBody, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PUT", path: "/v1/apps/{appId}/journeys/{journeyId}/draft"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateNode: Taslaktaki düğümü güncelle
//
// PATCH /v1/apps/{appId}/journeys/{journeyId}/nodes/{nodeId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) UpdateNode(ctx context.Context, journeyID string, nodeID string, body *UpdateNodeDto, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/journeys/{journeyId}/nodes/{nodeId}"}
	op.pathParams = map[string]string{"journeyId": journeyID, "nodeId": nodeID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Versions: Sürümler
//
// GET /v1/apps/{appId}/journeys/{journeyId}/versions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Versions(ctx context.Context, journeyID string, opts ...RequestOption) (*JourneyVersionList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/journeys/{journeyId}/versions"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	var out JourneyVersionList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Publish: Yayınla
//
// POST /v1/apps/{appId}/journeys/{journeyId}/publish
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Publish(ctx context.Context, journeyID string, opts ...RequestOption) (*JourneyPublishResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/journeys/{journeyId}/publish"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	var out JourneyPublishResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Pause: Duraklat
//
// POST /v1/apps/{appId}/journeys/{journeyId}/pause
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Pause(ctx context.Context, journeyID string, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/journeys/{journeyId}/pause"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Resume: Devam ettir
//
// POST /v1/apps/{appId}/journeys/{journeyId}/resume
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Resume(ctx context.Context, journeyID string, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/journeys/{journeyId}/resume"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Archive: Arşivle
//
// POST /v1/apps/{appId}/journeys/{journeyId}/archive
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Archive(ctx context.Context, journeyID string, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/journeys/{journeyId}/archive"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Duplicate: Kopyala
//
// POST /v1/apps/{appId}/journeys/{journeyId}/duplicate
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Duplicate(ctx context.Context, journeyID string, opts ...RequestOption) (*JourneyDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/journeys/{journeyId}/duplicate"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	var out JourneyDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// JourneysStatsParams: JourneysService.Stats için isteğe bağlı sorgu parametreleri.
type JourneysStatsParams struct {
	Version *int64
}

func (p *JourneysStatsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "version", p.Version)
}

// Stats: Düğüm bazlı istatistik
//
// GET /v1/apps/{appId}/journeys/{journeyId}/stats
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Stats(ctx context.Context, journeyID string, params *JourneysStatsParams, opts ...RequestOption) (*JourneyStats, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/journeys/{journeyId}/stats"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out JourneyStats
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// JourneysRunsParams: JourneysService.Runs için isteğe bağlı sorgu parametreleri.
type JourneysRunsParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "active", "waiting", "completed", "exited", "failed".
	Status  *string
	UserID  *string
	Version *int64
}

func (p *JourneysRunsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
	addQuery(q, "user_id", p.UserID)
	addQuery(q, "version", p.Version)
}

// Runs: Run listesi
//
// GET /v1/apps/{appId}/journeys/{journeyId}/runs
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Runs(ctx context.Context, journeyID string, params *JourneysRunsParams, opts ...RequestOption) (*JourneyRunPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/journeys/{journeyId}/runs"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out JourneyRunPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RunsAll: Runs sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *JourneysService) RunsAll(ctx context.Context, journeyID string, params *JourneysRunsParams, opts ...RequestOption) iter.Seq2[JourneyRun, error] {
	var p JourneysRunsParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]JourneyRun, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.Runs(ctx, journeyID, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Trigger: Kullanıcıyı journey'e sok (external_id, user_id veya alias; context Liquid custom_data)
//
// POST /v1/apps/{appId}/journeys/{journeyId}/runs
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Trigger(ctx context.Context, journeyID string, body *TriggerRunDto, opts ...RequestOption) (*JourneyTriggerResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/journeys/{journeyId}/runs"}
	op.pathParams = map[string]string{"journeyId": journeyID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out JourneyTriggerResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Run: Run detayı ve adımlar
//
// GET /v1/apps/{appId}/journeys/{journeyId}/runs/{runId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) Run(ctx context.Context, journeyID string, runID string, opts ...RequestOption) (*JourneyRunDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/journeys/{journeyId}/runs/{runId}"}
	op.pathParams = map[string]string{"journeyId": journeyID, "runId": runID}
	var out JourneyRunDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ExitRun: Run'ı sonlandır
//
// DELETE /v1/apps/{appId}/journeys/{journeyId}/runs/{runId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) ExitRun(ctx context.Context, journeyID string, runID string, opts ...RequestOption) (*JourneyRunDetail, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/journeys/{journeyId}/runs/{runId}"}
	op.pathParams = map[string]string{"journeyId": journeyID, "runId": runID}
	var out JourneyRunDetail
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListForUser: Kullanıcının journey run'ları
//
// GET /v1/apps/{appId}/users/by/{label}/{value}/journeys
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *JourneysService) ListForUser(ctx context.Context, label string, value string, opts ...RequestOption) (*UserJourneyRuns, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/users/by/{label}/{value}/journeys"}
	op.pathParams = map[string]string{"label": label, "value": value}
	var out UserJourneyRuns
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// PreferencesService: preferences kaynağı.
type PreferencesService struct {
	client *Client
}

// Get: Kullanıcının tercihleri (alias ile)
//
// GET /v1/apps/{appId}/users/by/{label}/{value}/preferences
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *PreferencesService) Get(ctx context.Context, label string, value string, opts ...RequestOption) (*UserPreferences, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/users/by/{label}/{value}/preferences"}
	op.pathParams = map[string]string{"label": label, "value": value}
	var out UserPreferences
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Set: Kullanıcı tercihlerini yaz (alias ile)
//
// PUT /v1/apps/{appId}/users/by/{label}/{value}/preferences
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *PreferencesService) Set(ctx context.Context, label string, value string, body *SetPreferencesDto, opts ...RequestOption) (*UserPreferences, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PUT", path: "/v1/apps/{appId}/users/by/{label}/{value}/preferences"}
	op.pathParams = map[string]string{"label": label, "value": value}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out UserPreferences
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
