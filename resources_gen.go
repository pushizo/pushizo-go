// Bu dosya api/openapi.yaml'dan sdk/scripts/generate-server-sdks.mjs ile üretilmiştir; elle düzenlemeyin.

package pushizo

import (
	"context"
	"iter"
	"net/url"
)

// APIVersion, SDK'nın üretildiği OpenAPI sözleşmesinin sürümüdür.
const APIVersion = "1.0.0"

// services, istemciye eklenen üretilmiş kaynak alanlarıdır.
type services struct {
	Notifications       *NotificationsService
	Users               *UsersService
	Aliases             *AliasesService
	Subscriptions       *SubscriptionsService
	Events              *EventsService
	Segments            *SegmentsService
	Templates           *TemplatesService
	Campaigns           *CampaignsService
	InAppMessages       *InAppMessagesService
	Inbox               *InboxService
	Suppressions        *SuppressionsService
	Exports             *ExportsService
	Imports             *ImportsService
	Webhooks            *WebhooksService
	Channels            *ChannelsService
	Analytics           *AnalyticsService
	Apps                *AppsService
	APIKeys             *APIKeysService
	Organization        *OrganizationService
	Ads                 *AdsService
	CannedResponses     *CannedResponsesService
	Categories          *CategoriesService
	ContentTables       *ContentTablesService
	Conversations       *ConversationsService
	DataFeeds           *DataFeedsService
	EmailDeliverability *EmailDeliverabilityService
	Forms               *FormsService
	IdentityKeys        *IdentityKeysService
	Integrations        *IntegrationsService
	Journeys            *JourneysService
	LandingPages        *LandingPagesService
	LiveActivities      *LiveActivitiesService
	Mcp                 *McpService
	Preferences         *PreferencesService
	SMSKeywords         *SMSKeywordsService
	Surveys             *SurveysService
	WarehouseExports    *WarehouseExportsService
	WebCampaigns        *WebCampaignsService
	Whatsapp            *WhatsappService
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
	c.Ads = &AdsService{client: c}
	c.CannedResponses = &CannedResponsesService{client: c}
	c.Categories = &CategoriesService{client: c}
	c.ContentTables = &ContentTablesService{client: c}
	c.Conversations = &ConversationsService{client: c}
	c.DataFeeds = &DataFeedsService{client: c}
	c.EmailDeliverability = &EmailDeliverabilityService{client: c}
	c.Forms = &FormsService{client: c}
	c.IdentityKeys = &IdentityKeysService{client: c}
	c.Integrations = &IntegrationsService{client: c}
	c.Journeys = &JourneysService{client: c}
	c.LandingPages = &LandingPagesService{client: c}
	c.LiveActivities = &LiveActivitiesService{client: c}
	c.Mcp = &McpService{client: c}
	c.Preferences = &PreferencesService{client: c}
	c.SMSKeywords = &SMSKeywordsService{client: c}
	c.Surveys = &SurveysService{client: c}
	c.WarehouseExports = &WarehouseExportsService{client: c}
	c.WebCampaigns = &WebCampaignsService{client: c}
	c.Whatsapp = &WhatsappService{client: c}
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

// UpdateByToken: Token ile abonelik güncelle (token_type: ios_push, android_push, hms_push, web_push, email, sms veya eski sağlayıcı adları iOSPush, AndroidPush, HuaweiPush, ChromePush, FirefoxPush, SafariPush, Email, SMS)
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

// EventsListDefinitionsParams: EventsService.ListDefinitions için isteğe bağlı sorgu parametreleri.
type EventsListDefinitionsParams struct {
	Limit  *int64
	Cursor *string
}

func (p *EventsListDefinitionsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
}

// ListDefinitions: Olay tanımları
//
// GET /v1/apps/{appId}/events/definitions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EventsService) ListDefinitions(ctx context.Context, params *EventsListDefinitionsParams, opts ...RequestOption) (*EventDefinitionList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/events/definitions"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out EventDefinitionList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListDefinitionsAll: ListDefinitions sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *EventsService) ListDefinitionsAll(ctx context.Context, params *EventsListDefinitionsParams, opts ...RequestOption) iter.Seq2[EventDefinition, error] {
	var p EventsListDefinitionsParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]EventDefinition, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.ListDefinitions(ctx, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
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
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
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
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
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

// SelectWinner: A/B kazananını elle seç
//
// POST /v1/apps/{appId}/in-app-messages/{messageId}/winner
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InAppMessagesService) SelectWinner(ctx context.Context, messageID string, body *InAppWinnerDto, opts ...RequestOption) (*InAppMessage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/in-app-messages/{messageId}/winner"}
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

// Winner: A/B kazananını elle seç
//
// POST /v1/apps/{appId}/in-app-messages/{messageId}/winner
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
//
// Deprecated: SelectWinner kullanın; bu ad geriye dönük uyumluluk için korunur.
func (s *InAppMessagesService) Winner(ctx context.Context, messageID string, body *InAppWinnerDto, opts ...RequestOption) (*InAppMessage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/in-app-messages/{messageId}/winner"}
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

// Teams: Takımlar
//
// GET /v1/apps/{appId}/inbox/teams
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) Teams(ctx context.Context, opts ...RequestOption) (*InboxTeamList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/inbox/teams"}
	var out InboxTeamList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Agents: Temsilciler ve varlık durumu
//
// GET /v1/apps/{appId}/inbox/agents
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) Agents(ctx context.Context, opts ...RequestOption) (*InboxAgentList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/inbox/agents"}
	var out InboxAgentList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Settings: Gelen kutusu ayarları
//
// GET /v1/apps/{appId}/inbox/settings
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) Settings(ctx context.Context, opts ...RequestOption) (*InboxSettings, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/inbox/settings"}
	var out InboxSettings
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateSettings: Gelen kutusu ayarlarını güncelle
//
// PUT /v1/apps/{appId}/inbox/settings
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) UpdateSettings(ctx context.Context, body *InboxSettingsUpdateDto, opts ...RequestOption) (*InboxSettings, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PUT", path: "/v1/apps/{appId}/inbox/settings"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out InboxSettings
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// InboxMetricsParams: InboxService.Metrics için isteğe bağlı sorgu parametreleri.
type InboxMetricsParams struct {
	From *string
	To   *string
	// Değerler: "sms", "whatsapp", "rcs", "email", "chat".
	Channel *string
}

func (p *InboxMetricsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
	addQuery(q, "channel", p.Channel)
}

// Metrics: Gelen kutusu metrikleri (ilk yanıt/çözüm ort. ve p90, SLA, temsilci tablosu)
//
// GET /v1/apps/{appId}/inbox/metrics
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *InboxService) Metrics(ctx context.Context, params *InboxMetricsParams, opts ...RequestOption) (*InboxMetrics, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/inbox/metrics"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out InboxMetrics
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
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
	// Değerler: "email", "sms", "whatsapp", "rcs".
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

// Preview: Event Streams önizlemesi: örnek olay + şablon + filtre → gövde/hata (gönderim yapmaz)
//
// POST /v1/apps/{appId}/webhooks/preview
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebhooksService) Preview(ctx context.Context, body *WebhookPreviewDto, opts ...RequestOption) (*WebhookPreviewResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/webhooks/preview"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out WebhookPreviewResult
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

// OrganizationUsageParams: OrganizationService.Usage için isteğe bağlı sorgu parametreleri.
type OrganizationUsageParams struct {
	Days *int64
}

func (p *OrganizationUsageParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "days", p.Days)
}

// Usage: Organizasyon kullanımı ve plan kotaları (org geneli anahtar)
//
// GET /v1/organization/usage
func (s *OrganizationService) Usage(ctx context.Context, params *OrganizationUsageParams, opts ...RequestOption) (*UsageReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/organization/usage"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out UsageReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsService: ads kaynağı.
type AdsService struct {
	client *Client
}

// ListAccounts: Bağlı reklam hesapları
//
// GET /v1/apps/{appId}/ads/accounts
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) ListAccounts(ctx context.Context, opts ...RequestOption) (*AdAccountList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/accounts"}
	var out AdAccountList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsListAudiencesParams: AdsService.ListAudiences için isteğe bağlı sorgu parametreleri.
type AdsListAudiencesParams struct {
	AdAccountID *string
}

func (p *AdsListAudiencesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "ad_account_id", p.AdAccountID)
}

// ListAudiences: Reklam kitleleri
//
// GET /v1/apps/{appId}/ads/audiences
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) ListAudiences(ctx context.Context, params *AdsListAudiencesParams, opts ...RequestOption) (*AdAudienceList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/audiences"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out AdAudienceList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateAudience: Reklam kitlesi oluştur
//
// POST /v1/apps/{appId}/ads/audiences
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) CreateAudience(ctx context.Context, body *CreateAdAudienceDto, opts ...RequestOption) (*AdAudience, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/ads/audiences"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out AdAudience
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAudience: Reklam kitlesi ayrıntısı
//
// GET /v1/apps/{appId}/ads/audiences/{audienceId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) GetAudience(ctx context.Context, audienceID string, opts ...RequestOption) (*AdAudience, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/audiences/{audienceId}"}
	op.pathParams = map[string]string{"audienceId": audienceID}
	var out AdAudience
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateAudience: Reklam kitlesini güncelle
//
// PATCH /v1/apps/{appId}/ads/audiences/{audienceId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) UpdateAudience(ctx context.Context, audienceID string, body *UpdateAdAudienceDto, opts ...RequestOption) (*AdAudience, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/ads/audiences/{audienceId}"}
	op.pathParams = map[string]string{"audienceId": audienceID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out AdAudience
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsRemoveAudienceParams: AdsService.RemoveAudience için isteğe bağlı sorgu parametreleri.
type AdsRemoveAudienceParams struct {
	// Değerler: "true", "false".
	DeleteOnPlatform *string
}

func (p *AdsRemoveAudienceParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "delete_on_platform", p.DeleteOnPlatform)
}

// RemoveAudience: Reklam kitlesini sil
//
// DELETE /v1/apps/{appId}/ads/audiences/{audienceId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) RemoveAudience(ctx context.Context, audienceID string, params *AdsRemoveAudienceParams, opts ...RequestOption) (*AdOk, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/ads/audiences/{audienceId}"}
	op.pathParams = map[string]string{"audienceId": audienceID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out AdOk
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Sync: Senkronu şimdi çalıştır
//
// POST /v1/apps/{appId}/ads/audiences/{audienceId}/sync
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Sync(ctx context.Context, audienceID string, body *SyncAdAudienceDto, opts ...RequestOption) (*AdAudienceSyncQueued, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/ads/audiences/{audienceId}/sync"}
	op.pathParams = map[string]string{"audienceId": audienceID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out AdAudienceSyncQueued
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsSyncsParams: AdsService.Syncs için isteğe bağlı sorgu parametreleri.
type AdsSyncsParams struct {
	Limit *int64
}

func (p *AdsSyncsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
}

// Syncs: Senkron geçmişi
//
// GET /v1/apps/{appId}/ads/audiences/{audienceId}/syncs
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Syncs(ctx context.Context, audienceID string, params *AdsSyncsParams, opts ...RequestOption) (*AdAudienceSyncList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/audiences/{audienceId}/syncs"}
	op.pathParams = map[string]string{"audienceId": audienceID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out AdAudienceSyncList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Lookalike: Benzer kitle (lookalike) oluştur
//
// POST /v1/apps/{appId}/ads/audiences/{audienceId}/lookalikes
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Lookalike(ctx context.Context, audienceID string, body *CreateLookalikeDto, opts ...RequestOption) (*AdAudience, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/ads/audiences/{audienceId}/lookalikes"}
	op.pathParams = map[string]string{"audienceId": audienceID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out AdAudience
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsListParams: AdsService.List için isteğe bağlı sorgu parametreleri.
type AdsListParams struct {
	AdAccountID *string
	// Değerler: "campaign", "ad_set", "ad".
	Level    *string
	ParentID *string
	// Değerler: "active", "paused", "pending", "rejected", "archived", "deleted", "unknown".
	Status *string
	Q      *string
	From   *string
	To     *string
	Limit  *int64
}

func (p *AdsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "ad_account_id", p.AdAccountID)
	addQuery(q, "level", p.Level)
	addQuery(q, "parent_id", p.ParentID)
	addQuery(q, "status", p.Status)
	addQuery(q, "q", p.Q)
	addQuery(q, "from", p.From)
	addQuery(q, "to", p.To)
	addQuery(q, "limit", p.Limit)
}

// List: Kampanya / reklam seti / reklam listesi ve metrikler
//
// GET /v1/apps/{appId}/ads/campaigns
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) List(ctx context.Context, params *AdsListParams, opts ...RequestOption) (*AdEntityList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/campaigns"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out AdEntityList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Kampanya / reklam seti / reklam ayrıntısı
//
// GET /v1/apps/{appId}/ads/campaigns/{entityId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Get(ctx context.Context, entityID string, opts ...RequestOption) (*AdEntity, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/campaigns/{entityId}"}
	op.pathParams = map[string]string{"entityId": entityID}
	var out AdEntity
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Durum ve günlük bütçe değişikliği
//
// PATCH /v1/apps/{appId}/ads/campaigns/{entityId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Update(ctx context.Context, entityID string, body *UpdateAdCampaignDto, opts ...RequestOption) (*AdEntity, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/ads/campaigns/{entityId}"}
	op.pathParams = map[string]string{"entityId": entityID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out AdEntity
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Refresh: Kampanyaları ve metrikleri şimdi yenile
//
// POST /v1/apps/{appId}/ads/accounts/{accountId}/refresh
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Refresh(ctx context.Context, accountID string, opts ...RequestOption) (*AdRefreshQueued, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/ads/accounts/{accountId}/refresh"}
	op.pathParams = map[string]string{"accountId": accountID}
	var out AdRefreshQueued
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsSpendParams: AdsService.Spend için isteğe bağlı sorgu parametreleri.
type AdsSpendParams struct {
	Currency *string
	// Değerler: "day", "platform", "account", "campaign", "ad_set".
	GroupBy     *string
	AdAccountID *string
	CampaignID  *string
}

func (p *AdsSpendParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "currency", p.Currency)
	addQuery(q, "group_by", p.GroupBy)
	addQuery(q, "ad_account_id", p.AdAccountID)
	addQuery(q, "campaign_id", p.CampaignID)
}

// Spend: Harcama ve performans raporu
//
// GET /v1/apps/{appId}/ads/reports/spend
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Spend(ctx context.Context, from string, to string, params *AdsSpendParams, opts ...RequestOption) (*AdSpendReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/reports/spend"}
	op.query = url.Values{}
	addQuery(op.query, "from", &from)
	addQuery(op.query, "to", &to)
	params.appendQuery(op.query)
	var out AdSpendReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsAttributionParams: AdsService.Attribution için isteğe bağlı sorgu parametreleri.
type AdsAttributionParams struct {
	// Değerler: "last_click", "first_click", "linear", "time_decay".
	Model        *string
	WindowDays   *int64
	HalfLifeDays *int64
	Currency     *string
	// Değerler: "channel", "campaign", "ad_set".
	GroupBy          *string
	ConversionEvents *string
}

func (p *AdsAttributionParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "model", p.Model)
	addQuery(q, "window_days", p.WindowDays)
	addQuery(q, "half_life_days", p.HalfLifeDays)
	addQuery(q, "currency", p.Currency)
	addQuery(q, "group_by", p.GroupBy)
	addQuery(q, "conversion_events", p.ConversionEvents)
}

// Attribution: Birleşik atribüsyon raporu (ROAS, CAC)
//
// GET /v1/apps/{appId}/ads/reports/attribution
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Attribution(ctx context.Context, from string, to string, params *AdsAttributionParams, opts ...RequestOption) (*AdAttributionReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/reports/attribution"}
	op.query = url.Values{}
	addQuery(op.query, "from", &from)
	addQuery(op.query, "to", &to)
	params.appendQuery(op.query)
	var out AdAttributionReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsLtvParams: AdsService.Ltv için isteğe bağlı sorgu parametreleri.
type AdsLtvParams struct {
	Months   *int64
	Cohorts  *int64
	Currency *string
}

func (p *AdsLtvParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "months", p.Months)
	addQuery(q, "cohorts", p.Cohorts)
	addQuery(q, "currency", p.Currency)
}

// Ltv: Kohort LTV raporu
//
// GET /v1/apps/{appId}/ads/reports/ltv
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Ltv(ctx context.Context, params *AdsLtvParams, opts ...RequestOption) (*AdLtvReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/reports/ltv"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out AdLtvReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CapiConfigs: Conversions API yapılandırmaları
//
// GET /v1/apps/{appId}/ads/conversions-api/configs
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) CapiConfigs(ctx context.Context, opts ...RequestOption) (*AdCapiConfigList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/conversions-api/configs"}
	var out AdCapiConfigList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsCapiDeliveriesParams: AdsService.CapiDeliveries için isteğe bağlı sorgu parametreleri.
type AdsCapiDeliveriesParams struct {
	ConfigID *string
	// Değerler: "pending", "sent", "failed", "skipped".
	Status *string
	Limit  *int64
}

func (p *AdsCapiDeliveriesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "config_id", p.ConfigID)
	addQuery(q, "status", p.Status)
	addQuery(q, "limit", p.Limit)
}

// CapiDeliveries: Conversions API teslim geçmişi
//
// GET /v1/apps/{appId}/ads/conversions-api/deliveries
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) CapiDeliveries(ctx context.Context, params *AdsCapiDeliveriesParams, opts ...RequestOption) (*AdCapiDeliveryList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/conversions-api/deliveries"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out AdCapiDeliveryList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Rules: Bütçe uyarı kuralları
//
// GET /v1/apps/{appId}/ads/alert-rules
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) Rules(ctx context.Context, opts ...RequestOption) (*AdBudgetAlertRuleList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/alert-rules"}
	var out AdBudgetAlertRuleList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateRule: Bütçe uyarı kuralı oluştur
//
// POST /v1/apps/{appId}/ads/alert-rules
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) CreateRule(ctx context.Context, body *CreateBudgetAlertRuleDto, opts ...RequestOption) (*AdBudgetAlertRule, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/ads/alert-rules"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out AdBudgetAlertRule
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateRule: Bütçe uyarı kuralını güncelle
//
// PATCH /v1/apps/{appId}/ads/alert-rules/{ruleId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) UpdateRule(ctx context.Context, ruleID string, body *UpdateBudgetAlertRuleDto, opts ...RequestOption) (*AdBudgetAlertRule, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/ads/alert-rules/{ruleId}"}
	op.pathParams = map[string]string{"ruleId": ruleID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out AdBudgetAlertRule
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RemoveRule: Bütçe uyarı kuralını sil
//
// DELETE /v1/apps/{appId}/ads/alert-rules/{ruleId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) RemoveRule(ctx context.Context, ruleID string, opts ...RequestOption) (*AdOk, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/ads/alert-rules/{ruleId}"}
	op.pathParams = map[string]string{"ruleId": ruleID}
	var out AdOk
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// AdsAlertHistoryParams: AdsService.AlertHistory için isteğe bağlı sorgu parametreleri.
type AdsAlertHistoryParams struct {
	RuleID *string
	Limit  *int64
}

func (p *AdsAlertHistoryParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "rule_id", p.RuleID)
	addQuery(q, "limit", p.Limit)
}

// AlertHistory: Tetiklenen bütçe uyarıları
//
// GET /v1/apps/{appId}/ads/alerts
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *AdsService) AlertHistory(ctx context.Context, params *AdsAlertHistoryParams, opts ...RequestOption) (*AdBudgetAlertList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/ads/alerts"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out AdBudgetAlertList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CannedResponsesService: cannedResponses kaynağı.
type CannedResponsesService struct {
	client *Client
}

// CannedResponsesCannedParams: CannedResponsesService.Canned için isteğe bağlı sorgu parametreleri.
type CannedResponsesCannedParams struct {
	Q *string
}

func (p *CannedResponsesCannedParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "q", p.Q)
}

// Canned: Hazır yanıtlar
//
// GET /v1/apps/{appId}/canned-responses
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CannedResponsesService) Canned(ctx context.Context, params *CannedResponsesCannedParams, opts ...RequestOption) (*CannedResponseList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/canned-responses"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out CannedResponseList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateCanned: Hazır yanıt oluştur
//
// POST /v1/apps/{appId}/canned-responses
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CannedResponsesService) CreateCanned(ctx context.Context, body *CannedInputDto, opts ...RequestOption) (*CannedResponse, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/canned-responses"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out CannedResponse
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateCanned: Hazır yanıtı güncelle
//
// PATCH /v1/apps/{appId}/canned-responses/{cannedId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CannedResponsesService) UpdateCanned(ctx context.Context, cannedID string, body *CannedUpdateDto, opts ...RequestOption) (*CannedResponse, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/canned-responses/{cannedId}"}
	op.pathParams = map[string]string{"cannedId": cannedID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out CannedResponse
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteCanned: Hazır yanıtı sil
//
// DELETE /v1/apps/{appId}/canned-responses/{cannedId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *CannedResponsesService) DeleteCanned(ctx context.Context, cannedID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/canned-responses/{cannedId}"}
	op.pathParams = map[string]string{"cannedId": cannedID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
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

// ContentTablesService: contentTables kaynağı.
type ContentTablesService struct {
	client *Client
}

// ListTables: İçerik tabloları
//
// GET /v1/apps/{appId}/content-tables
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ContentTablesService) ListTables(ctx context.Context, opts ...RequestOption) (*ContentTableList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/content-tables"}
	var out ContentTableList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ContentTablesCreateTableUploadForm: ContentTablesService.CreateTableUpload için multipart/form-data alanları.
type ContentTablesCreateTableUploadForm struct {
	// Zorunlu. UTF-8 CSV; ilk satır başlık. En fazla 100.000 satır, 20 MB.
	File *FileUpload
	// Zorunlu. Liquid tablo adı (küçük harf, rakam, alt çizgi).
	Name string
	// Zorunlu. Satır seçiminde kullanılan anahtar sütunu (değerleri benzersiz).
	KeyColumn   string
	Description *string
}

// CreateTableUpload: CSV ile içerik tablosu oluştur
//
// POST /v1/apps/{appId}/content-tables
//
// Şablonda `{% lookup '<ad>' key: user.tags.<etiket> as satir %}{{ satir.<sütun> }}`.
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ContentTablesService) CreateTableUpload(ctx context.Context, form *ContentTablesCreateTableUploadForm, opts ...RequestOption) (*ContentTable, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/content-tables"}
	if form == nil {
		return nil, errNilBody
	}
	mp := &multipartBody{}
	mp.file("file", form.File)
	mp.field("name", form.Name)
	mp.field("key_column", form.KeyColumn)
	if form.Description != nil {
		mp.field("description", *form.Description)
	}
	op.multipart = mp
	var out ContentTable
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTable: İçerik tablosu ayrıntısı
//
// GET /v1/apps/{appId}/content-tables/{tableId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ContentTablesService) GetTable(ctx context.Context, tableID string, opts ...RequestOption) (*ContentTable, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/content-tables/{tableId}"}
	op.pathParams = map[string]string{"tableId": tableID}
	var out ContentTable
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateTable: İçerik tablosu açıklamasını güncelle
//
// PATCH /v1/apps/{appId}/content-tables/{tableId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ContentTablesService) UpdateTable(ctx context.Context, tableID string, body *UpdateContentTableDto, opts ...RequestOption) (*ContentTable, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/content-tables/{tableId}"}
	op.pathParams = map[string]string{"tableId": tableID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out ContentTable
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteTable: İçerik tablosunu sil
//
// DELETE /v1/apps/{appId}/content-tables/{tableId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ContentTablesService) DeleteTable(ctx context.Context, tableID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/content-tables/{tableId}"}
	op.pathParams = map[string]string{"tableId": tableID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ContentTablesReplaceTableUploadForm: ContentTablesService.ReplaceTableUpload için multipart/form-data alanları.
type ContentTablesReplaceTableUploadForm struct {
	// Zorunlu. Yeni CSV (tablonun tüm satırları değiştirilir, sürüm artar).
	File *FileUpload
	// İsteğe bağlı yeni anahtar sütunu.
	KeyColumn *string
}

// ReplaceTableUpload: İçerik tablosunu yeniden yükle (sürüm artar)
//
// PUT /v1/apps/{appId}/content-tables/{tableId}/rows
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ContentTablesService) ReplaceTableUpload(ctx context.Context, tableID string, form *ContentTablesReplaceTableUploadForm, opts ...RequestOption) (*ContentTable, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PUT", path: "/v1/apps/{appId}/content-tables/{tableId}/rows"}
	op.pathParams = map[string]string{"tableId": tableID}
	if form == nil {
		return nil, errNilBody
	}
	mp := &multipartBody{}
	mp.file("file", form.File)
	if form.KeyColumn != nil {
		mp.field("key_column", *form.KeyColumn)
	}
	op.multipart = mp
	var out ContentTable
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ContentTablesTableRowsParams: ContentTablesService.TableRows için isteğe bağlı sorgu parametreleri.
type ContentTablesTableRowsParams struct {
	Limit  *int64
	Offset *int64
	Q      *string
}

func (p *ContentTablesTableRowsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "offset", p.Offset)
	addQuery(q, "q", p.Q)
}

// TableRows: İçerik tablosu satırları
//
// GET /v1/apps/{appId}/content-tables/{tableId}/rows
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ContentTablesService) TableRows(ctx context.Context, tableID string, params *ContentTablesTableRowsParams, opts ...RequestOption) (*ContentTableRows, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/content-tables/{tableId}/rows"}
	op.pathParams = map[string]string{"tableId": tableID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out ContentTableRows
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ConversationsService: conversations kaynağı.
type ConversationsService struct {
	client *Client
}

// ConversationsListParams: ConversationsService.List için isteğe bağlı sorgu parametreleri.
type ConversationsListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "sms", "whatsapp", "rcs", "email", "chat".
	Channel *string
	// Değerler: "open", "pending", "resolved", "closed".
	Status *string
	// Değerler: "true", "false".
	Unread   *string
	Q        *string
	Assignee *string
	TeamID   *string
}

func (p *ConversationsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "channel", p.Channel)
	addQuery(q, "status", p.Status)
	addQuery(q, "unread", p.Unread)
	addQuery(q, "q", p.Q)
	addQuery(q, "assignee", p.Assignee)
	addQuery(q, "team_id", p.TeamID)
}

// List: Konuşmalar (SMS, WhatsApp, RCS, e-posta, canlı sohbet; son mesaja göre; kanal/durum/okunmamış/atanan/takım/kişi filtresi)
//
// GET /v1/apps/{appId}/conversations
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) List(ctx context.Context, params *ConversationsListParams, opts ...RequestOption) (*ConversationPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/conversations"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out ConversationPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *ConversationsService) ListAll(ctx context.Context, params *ConversationsListParams, opts ...RequestOption) iter.Seq2[Conversation, error] {
	var p ConversationsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Conversation, *string, bool, error) {
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

// Get: Konuşma (atama, SLA, bot durumu, WhatsApp oturum penceresi ve açık rıza dahil)
//
// GET /v1/apps/{appId}/conversations/{conversationId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) Get(ctx context.Context, conversationID string, opts ...RequestOption) (*Conversation, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/conversations/{conversationId}"}
	op.pathParams = map[string]string{"conversationId": conversationID}
	var out Conversation
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SetStatus: Konuşma durumu: open / pending / resolved (closed = resolved)
//
// PATCH /v1/apps/{appId}/conversations/{conversationId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) SetStatus(ctx context.Context, conversationID string, body *ConversationUpdateDto, opts ...RequestOption) (*Conversation, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/conversations/{conversationId}"}
	op.pathParams = map[string]string{"conversationId": conversationID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Conversation
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ConversationsMessagesParams: ConversationsService.Messages için isteğe bağlı sorgu parametreleri.
type ConversationsMessagesParams struct {
	Limit    *int64
	Cursor   *string
	AfterSeq *int64
}

func (p *ConversationsMessagesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "after_seq", p.AfterSeq)
}

// Messages: Konuşma mesajları (yeniden eskiye; `after_seq` ile eskiden yeniye; giden mesaj durumu teslimattan veya okundu işaretçisinden)
//
// GET /v1/apps/{appId}/conversations/{conversationId}/messages
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) Messages(ctx context.Context, conversationID string, params *ConversationsMessagesParams, opts ...RequestOption) (*ConversationMessagePage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/conversations/{conversationId}/messages"}
	op.pathParams = map[string]string{"conversationId": conversationID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out ConversationMessagePage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// MessagesAll: Messages sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *ConversationsService) MessagesAll(ctx context.Context, conversationID string, params *ConversationsMessagesParams, opts ...RequestOption) iter.Seq2[ConversationMessagePageDataItem, error] {
	var p ConversationsMessagesParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]ConversationMessagePageDataItem, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.Messages(ctx, conversationID, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Reply: Yanıt gönder (client_message_id ile tekilleşir). Sohbet: doğrudan; e-posta: yanıt adresli işlemsel e-posta; WhatsApp: açık rıza zorunlu, son gelen mesajdan 24 saat içinde metin/medya, pencere dışında yalnızca onaylı şablon (aksi 422 WHATSAPP_WINDOW_CLOSED).
//
// POST /v1/apps/{appId}/conversations/{conversationId}/messages
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) Reply(ctx context.Context, conversationID string, body *ConversationReplyDto, opts ...RequestOption) (*ConversationReplyResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/conversations/{conversationId}/messages"}
	op.pathParams = map[string]string{"conversationId": conversationID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out ConversationReplyResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Read: Gelen mesajları okundu işaretle (sohbette ziyaretçiye okundu bilgisi gider)
//
// POST /v1/apps/{appId}/conversations/{conversationId}/read
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) Read(ctx context.Context, conversationID string, opts ...RequestOption) (*Conversation, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/conversations/{conversationId}/read"}
	op.pathParams = map[string]string{"conversationId": conversationID}
	var out Conversation
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Assign: Konuşmayı temsilciye ve/veya takıma ata (inbox.manage olmadan yalnızca kendine)
//
// POST /v1/apps/{appId}/conversations/{conversationId}/assign
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) Assign(ctx context.Context, conversationID string, body *ConversationAssignDto, opts ...RequestOption) (*Conversation, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/conversations/{conversationId}/assign"}
	op.pathParams = map[string]string{"conversationId": conversationID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Conversation
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Tags: Sohbetten kişiye etiket ekle/kaldır
//
// POST /v1/apps/{appId}/conversations/{conversationId}/tags
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) Tags(ctx context.Context, conversationID string, body *ConversationTagsDto, opts ...RequestOption) (*ConversationTagsResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/conversations/{conversationId}/tags"}
	op.pathParams = map[string]string{"conversationId": conversationID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out ConversationTagsResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Activity: Konuşma zaman çizelgesi (atama, durum, ilk yanıt, bot devri)
//
// GET /v1/apps/{appId}/conversations/{conversationId}/activity
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *ConversationsService) Activity(ctx context.Context, conversationID string, opts ...RequestOption) (*ConversationActivity, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/conversations/{conversationId}/activity"}
	op.pathParams = map[string]string{"conversationId": conversationID}
	var out ConversationActivity
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DataFeedsService: dataFeeds kaynağı.
type DataFeedsService struct {
	client *Client
}

// ListFeeds: Data feed listesi
//
// GET /v1/apps/{appId}/data-feeds
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *DataFeedsService) ListFeeds(ctx context.Context, opts ...RequestOption) (*DataFeedList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/data-feeds"}
	var out DataFeedList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateFeed: Data feed oluştur
//
// POST /v1/apps/{appId}/data-feeds
//
// URL şablonu Liquid ile kişiselleştirilir (ör. `https://api.ornek.com/u/{{ user.external_id }}`); şablonda `{{ feed.<ad>.alan }}`.
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *DataFeedsService) CreateFeed(ctx context.Context, body *CreateDataFeedDto, opts ...RequestOption) (*DataFeed, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/data-feeds"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out DataFeed
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetFeed: Data feed ayrıntısı
//
// GET /v1/apps/{appId}/data-feeds/{feedId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *DataFeedsService) GetFeed(ctx context.Context, feedID string, opts ...RequestOption) (*DataFeed, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/data-feeds/{feedId}"}
	op.pathParams = map[string]string{"feedId": feedID}
	var out DataFeed
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateFeed: Data feed güncelle
//
// PATCH /v1/apps/{appId}/data-feeds/{feedId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *DataFeedsService) UpdateFeed(ctx context.Context, feedID string, body *UpdateDataFeedDto, opts ...RequestOption) (*DataFeed, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/data-feeds/{feedId}"}
	op.pathParams = map[string]string{"feedId": feedID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out DataFeed
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteFeed: Data feed sil
//
// DELETE /v1/apps/{appId}/data-feeds/{feedId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *DataFeedsService) DeleteFeed(ctx context.Context, feedID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/data-feeds/{feedId}"}
	op.pathParams = map[string]string{"feedId": feedID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// TestFeed: Data feed'i örnek kullanıcıyla dene
//
// POST /v1/apps/{appId}/data-feeds/{feedId}/test
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *DataFeedsService) TestFeed(ctx context.Context, feedID string, body *TestDataFeedDto, opts ...RequestOption) (*DataFeedTestResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/data-feeds/{feedId}/test"}
	op.pathParams = map[string]string{"feedId": feedID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out DataFeedTestResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// EmailDeliverabilityService: emailDeliverability kaynağı.
type EmailDeliverabilityService struct {
	client *Client
}

// Settings: E-posta teslim ayarları (BCC/arşiv, gönderim öncesi adres doğrulama, üçüncü taraf doğrulayıcı)
//
// GET /v1/apps/{appId}/email-deliverability
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) Settings(ctx context.Context, opts ...RequestOption) (*EmailDeliverabilitySettings, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/email-deliverability"}
	var out EmailDeliverabilitySettings
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: E-posta teslim ayarlarını güncelle (doğrulayıcı anahtarı zarf şifreli saklanır, geri okunmaz)
//
// PUT /v1/apps/{appId}/email-deliverability
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) Update(ctx context.Context, body *DeliverabilitySettingsDto, opts ...RequestOption) (*EmailDeliverabilitySettings, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PUT", path: "/v1/apps/{appId}/email-deliverability"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out EmailDeliverabilitySettings
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Verify: Adres doğrulama testi (sözdizimi, MX, tek kullanımlık alan adı, rol hesabı, üçüncü taraf)
//
// POST /v1/apps/{appId}/email-deliverability/verify-address
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) Verify(ctx context.Context, body *VerifyAddressDto, opts ...RequestOption) (*EmailAddressCheck, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/email-deliverability/verify-address"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out EmailAddressCheck
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Plans: Isındırma planları (bugünkü limit/kullanım ve ileri günlere ayrılmış gönderimler)
//
// GET /v1/apps/{appId}/email-deliverability/warmup-plans
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) Plans(ctx context.Context, opts ...RequestOption) (*EmailWarmupPlanList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/email-deliverability/warmup-plans"}
	var out EmailWarmupPlanList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create: Isındırma planı oluştur (günlük limit eğrisi veya başlangıç/çarpan/hedef; sağlayıcı payları)
//
// POST /v1/apps/{appId}/email-deliverability/warmup-plans
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) Create(ctx context.Context, body *CreateWarmupPlanDto, opts ...RequestOption) (*EmailWarmupPlan, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/email-deliverability/warmup-plans"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out EmailWarmupPlan
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Isındırma planı
//
// GET /v1/apps/{appId}/email-deliverability/warmup-plans/{planId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) Get(ctx context.Context, planID string, opts ...RequestOption) (*EmailWarmupPlan, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/email-deliverability/warmup-plans/{planId}"}
	op.pathParams = map[string]string{"planId": planID}
	var out EmailWarmupPlan
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdatePlan: Isındırma planını güncelle (duraklat/sürdür, eğri, paylar)
//
// PATCH /v1/apps/{appId}/email-deliverability/warmup-plans/{planId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) UpdatePlan(ctx context.Context, planID string, body *UpdateWarmupPlanDto, opts ...RequestOption) (*EmailWarmupPlan, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/email-deliverability/warmup-plans/{planId}"}
	op.pathParams = map[string]string{"planId": planID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out EmailWarmupPlan
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Isındırma planını sil
//
// DELETE /v1/apps/{appId}/email-deliverability/warmup-plans/{planId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) Delete(ctx context.Context, planID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/email-deliverability/warmup-plans/{planId}"}
	op.pathParams = map[string]string{"planId": planID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// EmailDeliverabilityUsageParams: EmailDeliverabilityService.Usage için isteğe bağlı sorgu parametreleri.
type EmailDeliverabilityUsageParams struct {
	Days *int64
}

func (p *EmailDeliverabilityUsageParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "days", p.Days)
}

// Usage: Isındırma planı gün × sağlayıcı kullanımı
//
// GET /v1/apps/{appId}/email-deliverability/warmup-plans/{planId}/usage
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *EmailDeliverabilityService) Usage(ctx context.Context, planID string, params *EmailDeliverabilityUsageParams, opts ...RequestOption) (*EmailWarmupUsage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/email-deliverability/warmup-plans/{planId}/usage"}
	op.pathParams = map[string]string{"planId": planID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out EmailWarmupUsage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// FormsService: forms kaynağı.
type FormsService struct {
	client *Client
}

// FormsListParams: FormsService.List için isteğe bağlı sorgu parametreleri.
type FormsListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "active", "archived".
	Status *string
}

func (p *FormsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
}

// List: Formlar
//
// GET /v1/apps/{appId}/forms
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *FormsService) List(ctx context.Context, params *FormsListParams, opts ...RequestOption) (*FormPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/forms"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out FormPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *FormsService) ListAll(ctx context.Context, params *FormsListParams, opts ...RequestOption) iter.Seq2[Form, error] {
	var p FormsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Form, *string, bool, error) {
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

// Create: Form oluştur
//
// POST /v1/apps/{appId}/forms
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *FormsService) Create(ctx context.Context, body *CreateFormDto, opts ...RequestOption) (*Form, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/forms"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Form
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Form detayı
//
// GET /v1/apps/{appId}/forms/{formId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *FormsService) Get(ctx context.Context, formID string, opts ...RequestOption) (*Form, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/forms/{formId}"}
	op.pathParams = map[string]string{"formId": formID}
	var out Form
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// FormsSubmissionsParams: FormsService.Submissions için isteğe bağlı sorgu parametreleri.
type FormsSubmissionsParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "accepted", "pending_confirmation", "confirmed".
	Status *string
	Since  *string
}

func (p *FormsSubmissionsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
	addQuery(q, "since", p.Since)
}

// Submissions: Form gönderimleri
//
// GET /v1/apps/{appId}/forms/{formId}/submissions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *FormsService) Submissions(ctx context.Context, formID string, params *FormsSubmissionsParams, opts ...RequestOption) (*FormSubmissionPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/forms/{formId}/submissions"}
	op.pathParams = map[string]string{"formId": formID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out FormSubmissionPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SubmissionsAll: Submissions sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *FormsService) SubmissionsAll(ctx context.Context, formID string, params *FormsSubmissionsParams, opts ...RequestOption) iter.Seq2[FormSubmission, error] {
	var p FormsSubmissionsParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]FormSubmission, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.Submissions(ctx, formID, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Submit: Sunucu tarafı form gönderimi (`external_id` verilirse o kullanıcının profili güncellenir)
//
// POST /v1/apps/{appId}/forms/{formId}/submissions
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *FormsService) Submit(ctx context.Context, formID string, body *PublicFormSubmitDto, opts ...RequestOption) (*FormSubmitResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/forms/{formId}/submissions"}
	op.pathParams = map[string]string{"formId": formID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out FormSubmitResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// IdentityKeysService: identityKeys kaynağı.
type IdentityKeysService struct {
	client *Client
}

// List: SDK kimlik doğrulama (JWT) genel anahtarları
//
// GET /v1/apps/{appId}/identity-keys
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *IdentityKeysService) List(ctx context.Context, opts ...RequestOption) (*IdentityKeyList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/identity-keys"}
	var out IdentityKeyList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create: Genel anahtar ekle (ES256)
//
// POST /v1/apps/{appId}/identity-keys
//
// SDK JWT identity verification için ES256 (EC P-256) genel anahtarı ekler; özel anahtar gönderilmez. Belirteç sunucunuzda imzalanır (`sub` = external_id, `aud` = app_id, `exp` ≤ 24 saat) ve SDK isteklerinde `Authorization: Bearer <jwt>` başlığıyla taşınır. Kip app ayarında `identity_verification_mode` ile seçilir.
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *IdentityKeysService) Create(ctx context.Context, body *CreateIdentityKeyDto, opts ...RequestOption) (*IdentityKey, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/identity-keys"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out IdentityKey
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Revoke: Genel anahtarı iptal et
//
// DELETE /v1/apps/{appId}/identity-keys/{keyId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *IdentityKeysService) Revoke(ctx context.Context, keyID string, opts ...RequestOption) (*IdentityKey, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/identity-keys/{keyId}"}
	op.pathParams = map[string]string{"keyId": keyID}
	var out IdentityKey
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// IntegrationsService: integrations kaynağı.
type IntegrationsService struct {
	client *Client
}

// Me: Zapier bağlantı testi: anahtarın app’i ve kapsamları
//
// GET /v1/apps/{appId}/integrations/zapier/me
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *IntegrationsService) Me(ctx context.Context, opts ...RequestOption) (*ZapierMe, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/integrations/zapier/me"}
	var out ZapierMe
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Subscribe: REST Hook aboneliği (Zap açıldı)
//
// POST /v1/apps/{appId}/integrations/zapier/hooks
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *IntegrationsService) Subscribe(ctx context.Context, body *ZapierSubscribeDto, opts ...RequestOption) (*ZapierHook, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/integrations/zapier/hooks"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out ZapierHook
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Unsubscribe: REST Hook aboneliğini sil (Zap kapatıldı)
//
// DELETE /v1/apps/{appId}/integrations/zapier/hooks/{hookId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *IntegrationsService) Unsubscribe(ctx context.Context, hookID string, opts ...RequestOption) (*IntegrationOk, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/integrations/zapier/hooks/{hookId}"}
	op.pathParams = map[string]string{"hookId": hookID}
	var out IntegrationOk
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// IntegrationsSamplesParams: IntegrationsService.Samples için isteğe bağlı sorgu parametreleri.
type IntegrationsSamplesParams struct {
	EventName *string
}

func (p *IntegrationsSamplesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "event_name", p.EventName)
}

// Samples: Tetik örnekleri (performList): son kayıtlar, abonelik gövdesiyle aynı biçim
//
// GET /v1/apps/{appId}/integrations/zapier/samples
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *IntegrationsService) Samples(ctx context.Context, event string, params *IntegrationsSamplesParams, opts ...RequestOption) ([]ZapierSamplesItem, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/integrations/zapier/samples"}
	op.query = url.Values{}
	addQuery(op.query, "event", &event)
	params.appendQuery(op.query)
	var out []ZapierSamplesItem
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return out, nil
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
	// Değerler: "true", "false".
	AllVersions *string
}

func (p *JourneysStatsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "version", p.Version)
	addQuery(q, "all_versions", p.AllVersions)
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

// LandingPagesService: landingPages kaynağı.
type LandingPagesService struct {
	client *Client
}

// LandingPagesListParams: LandingPagesService.List için isteğe bağlı sorgu parametreleri.
type LandingPagesListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "draft", "published", "archived".
	Status *string
}

func (p *LandingPagesListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
}

// List: Landing sayfaları
//
// GET /v1/apps/{appId}/landing-pages
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LandingPagesService) List(ctx context.Context, params *LandingPagesListParams, opts ...RequestOption) (*LandingPageList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/landing-pages"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out LandingPageList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *LandingPagesService) ListAll(ctx context.Context, params *LandingPagesListParams, opts ...RequestOption) iter.Seq2[LandingPageListDataItem, error] {
	var p LandingPagesListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]LandingPageListDataItem, *string, bool, error) {
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

// Get: Landing sayfası detayı
//
// GET /v1/apps/{appId}/landing-pages/{pageId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LandingPagesService) Get(ctx context.Context, pageID string, opts ...RequestOption) (*LandingPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/landing-pages/{pageId}"}
	op.pathParams = map[string]string{"pageId": pageID}
	var out LandingPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Publish: Taslağı yayınla
//
// POST /v1/apps/{appId}/landing-pages/{pageId}/publish
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LandingPagesService) Publish(ctx context.Context, pageID string, opts ...RequestOption) (*LandingPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/landing-pages/{pageId}/publish"}
	op.pathParams = map[string]string{"pageId": pageID}
	var out LandingPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Unpublish: Yayından kaldır
//
// POST /v1/apps/{appId}/landing-pages/{pageId}/unpublish
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LandingPagesService) Unpublish(ctx context.Context, pageID string, opts ...RequestOption) (*LandingPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/landing-pages/{pageId}/unpublish"}
	op.pathParams = map[string]string{"pageId": pageID}
	var out LandingPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// LiveActivitiesService: liveActivities kaynağı.
type LiveActivitiesService struct {
	client *Client
}

// ListTypes: Live Activity türleri
//
// GET /v1/apps/{appId}/live-activities/types
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LiveActivitiesService) ListTypes(ctx context.Context, opts ...RequestOption) (*LiveActivityTypeList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/live-activities/types"}
	var out LiveActivityTypeList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateType: Live Activity türü oluştur
//
// POST /v1/apps/{appId}/live-activities/types
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LiveActivitiesService) CreateType(ctx context.Context, body *CreateLiveActivityTypeDto, opts ...RequestOption) (*LiveActivityType, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/live-activities/types"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out LiveActivityType
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateType: Live Activity türünü güncelle
//
// PATCH /v1/apps/{appId}/live-activities/types/{activityType}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LiveActivitiesService) UpdateType(ctx context.Context, activityType string, body *UpdateLiveActivityTypeDto, opts ...RequestOption) (*LiveActivityType, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/live-activities/types/{activityType}"}
	op.pathParams = map[string]string{"activityType": activityType}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out LiveActivityType
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteType: Live Activity türünü sil
//
// DELETE /v1/apps/{appId}/live-activities/types/{activityType}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LiveActivitiesService) DeleteType(ctx context.Context, activityType string, opts ...RequestOption) (*LiveActivityTokenResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/live-activities/types/{activityType}"}
	op.pathParams = map[string]string{"activityType": activityType}
	var out LiveActivityTokenResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Start: Live Activity başlat (push-to-start)
//
// POST /v1/apps/{appId}/live-activities/{activityType}/start
//
// Hedefteki iOS aboneliklerinin push-to-start belirteçleriyle etkinliği uzaktan başlatır (iOS 17.2+). `attributes` ve `content_state` türün şemasına uymalıdır. Dönen `activity_id` ile sonraki update/end yapılır.
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LiveActivitiesService) Start(ctx context.Context, activityType string, body *StartLiveActivityDto, opts ...RequestOption) (*LiveActivitySendResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/live-activities/{activityType}/start"}
	op.pathParams = map[string]string{"activityType": activityType}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out LiveActivitySendResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Live Activity güncelle
//
// POST /v1/apps/{appId}/live-activities/{activityType}/update
//
// Etkinliğin cihazlardan bildirilen güncelleme belirteçlerine yeni içerik durumu gönderir. Sessiz saat ve frekans sınırı uygulanmaz; priority 5 Apple bütçesini korur.
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LiveActivitiesService) Update(ctx context.Context, activityType string, body *UpdateLiveActivityDto, opts ...RequestOption) (*LiveActivitySendResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/live-activities/{activityType}/update"}
	op.pathParams = map[string]string{"activityType": activityType}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out LiveActivitySendResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// End: Live Activity bitir
//
// POST /v1/apps/{appId}/live-activities/{activityType}/end
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *LiveActivitiesService) End(ctx context.Context, activityType string, body *EndLiveActivityDto, opts ...RequestOption) (*LiveActivitySendResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/live-activities/{activityType}/end"}
	op.pathParams = map[string]string{"activityType": activityType}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out LiveActivitySendResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// McpService: mcp kaynağı.
type McpService struct {
	client *Client
}

// CreateCampaignDraft: Kampanya taslağı oluştur (MCP onaylı mod; gönderim yapmaz)
//
// POST /v1/apps/{appId}/mcp/campaign-drafts
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *McpService) CreateCampaignDraft(ctx context.Context, body *McpCreateCampaignDraftBody, opts ...RequestOption) (*Campaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/mcp/campaign-drafts"}
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

// CreateSegmentDraft: Segment taslağı oluştur (MCP onaylı mod; dinamik segment)
//
// POST /v1/apps/{appId}/mcp/segment-drafts
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *McpService) CreateSegmentDraft(ctx context.Context, body *McpCreateSegmentDraftBody, opts ...RequestOption) (*Segment, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/mcp/segment-drafts"}
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

// SMSKeywordsService: smsKeywords kaynağı.
type SMSKeywordsService struct {
	client *Client
}

// Keywords: İki yönlü mesaj anahtar sözcükleri (org geneli + bu app) ve yasal (ayrılmış) sözcükler
//
// GET /v1/apps/{appId}/sms-keywords
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SMSKeywordsService) Keywords(ctx context.Context, opts ...RequestOption) (*SMSKeywordList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/sms-keywords"}
	var out SMSKeywordList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateKeyword: Anahtar sözcük oluştur (tag ekle/kaldır, olay → journey tetiği, otomatik yanıt, abonelik)
//
// POST /v1/apps/{appId}/sms-keywords
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SMSKeywordsService) CreateKeyword(ctx context.Context, body *CreateKeywordDto, opts ...RequestOption) (*SMSKeyword, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/sms-keywords"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out SMSKeyword
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateKeyword: Anahtar sözcüğü güncelle
//
// PATCH /v1/apps/{appId}/sms-keywords/{keywordId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SMSKeywordsService) UpdateKeyword(ctx context.Context, keywordID string, body *UpdateKeywordDto, opts ...RequestOption) (*SMSKeyword, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/sms-keywords/{keywordId}"}
	op.pathParams = map[string]string{"keywordId": keywordID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out SMSKeyword
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteKeyword: Anahtar sözcüğü sil
//
// DELETE /v1/apps/{appId}/sms-keywords/{keywordId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SMSKeywordsService) DeleteKeyword(ctx context.Context, keywordID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/sms-keywords/{keywordId}"}
	op.pathParams = map[string]string{"keywordId": keywordID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SurveysService: surveys kaynağı.
type SurveysService struct {
	client *Client
}

// SurveysListParams: SurveysService.List için isteğe bağlı sorgu parametreleri.
type SurveysListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "draft", "active", "closed".
	Status *string
}

func (p *SurveysListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
}

// List: Anketler
//
// GET /v1/apps/{appId}/surveys
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SurveysService) List(ctx context.Context, params *SurveysListParams, opts ...RequestOption) (*SurveyPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/surveys"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out SurveyPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *SurveysService) ListAll(ctx context.Context, params *SurveysListParams, opts ...RequestOption) iter.Seq2[Survey, error] {
	var p SurveysListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]Survey, *string, bool, error) {
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

// Create: Anket oluştur (taslak)
//
// POST /v1/apps/{appId}/surveys
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SurveysService) Create(ctx context.Context, body *CreateSurveyDto, opts ...RequestOption) (*Survey, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/surveys"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Survey
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Anket detayı
//
// GET /v1/apps/{appId}/surveys/{surveyId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SurveysService) Get(ctx context.Context, surveyID string, opts ...RequestOption) (*Survey, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/surveys/{surveyId}"}
	op.pathParams = map[string]string{"surveyId": surveyID}
	var out Survey
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Status: Etkinleştir / kapat
//
// POST /v1/apps/{appId}/surveys/{surveyId}/status
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SurveysService) Status(ctx context.Context, surveyID string, body *SurveyStatusDto, opts ...RequestOption) (*Survey, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/surveys/{surveyId}/status"}
	op.pathParams = map[string]string{"surveyId": surveyID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out Survey
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SurveysReportParams: SurveysService.Report için isteğe bağlı sorgu parametreleri.
type SurveysReportParams struct {
	Days *int64
}

func (p *SurveysReportParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "days", p.Days)
}

// Report: NPS / CSAT raporu
//
// GET /v1/apps/{appId}/surveys/{surveyId}/report
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SurveysService) Report(ctx context.Context, surveyID string, params *SurveysReportParams, opts ...RequestOption) (*SurveyReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/surveys/{surveyId}/report"}
	op.pathParams = map[string]string{"surveyId": surveyID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out SurveyReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SurveysResponsesParams: SurveysService.Responses için isteğe bağlı sorgu parametreleri.
type SurveysResponsesParams struct {
	Limit        *int64
	Cursor       *string
	TextQuestion *string
	// Değerler: "promoter", "passive", "detractor".
	NpsCategory *string
}

func (p *SurveysResponsesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "text_question", p.TextQuestion)
	addQuery(q, "nps_category", p.NpsCategory)
}

// Responses: Anket yanıtları
//
// GET /v1/apps/{appId}/surveys/{surveyId}/responses
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SurveysService) Responses(ctx context.Context, surveyID string, params *SurveysResponsesParams, opts ...RequestOption) (*SurveyResponsePage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/surveys/{surveyId}/responses"}
	op.pathParams = map[string]string{"surveyId": surveyID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out SurveyResponsePage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ResponsesAll: Responses sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *SurveysService) ResponsesAll(ctx context.Context, surveyID string, params *SurveysResponsesParams, opts ...RequestOption) iter.Seq2[SurveyResponse, error] {
	var p SurveysResponsesParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]SurveyResponse, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.Responses(ctx, surveyID, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// Respond: Sunucu tarafı anket yanıtı (`external_id` ile kullanıcıya bağlanır)
//
// POST /v1/apps/{appId}/surveys/{surveyId}/responses
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *SurveysService) Respond(ctx context.Context, surveyID string, body *PublicSurveyResponseDto, opts ...RequestOption) (*SurveyResponseResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/surveys/{surveyId}/responses"}
	op.pathParams = map[string]string{"surveyId": surveyID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out SurveyResponseResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// WarehouseExportsService: warehouseExports kaynağı.
type WarehouseExportsService struct {
	client *Client
}

// Tables: Dışa aktarılabilir tablolar, kolonlar ve şema sürümleri
//
// GET /v1/apps/{appId}/warehouse-exports/tables
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) Tables(ctx context.Context, opts ...RequestOption) (*WarehouseTables, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/warehouse-exports/tables"}
	var out WarehouseTables
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Test: Kaydetmeden bağlantı testi
//
// POST /v1/apps/{appId}/warehouse-exports/test-connection
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) Test(ctx context.Context, body any, opts ...RequestOption) (*WarehouseConnectionResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/warehouse-exports/test-connection"}
	op.body = body
	var out WarehouseConnectionResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// List: Veri ambarı dışa aktarımları
//
// GET /v1/apps/{appId}/warehouse-exports
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) List(ctx context.Context, opts ...RequestOption) (*WarehouseExportList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/warehouse-exports"}
	var out WarehouseExportList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create: Dışa aktarım oluştur (kimlik bilgisi şifreli saklanır, yanıtta dönmez)
//
// POST /v1/apps/{appId}/warehouse-exports
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) Create(ctx context.Context, body any, opts ...RequestOption) (*WarehouseExport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/warehouse-exports"}
	op.body = body
	var out WarehouseExport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Dışa aktarım detayı
//
// GET /v1/apps/{appId}/warehouse-exports/{exportId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) Get(ctx context.Context, exportID string, opts ...RequestOption) (*WarehouseExport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/warehouse-exports/{exportId}"}
	op.pathParams = map[string]string{"exportId": exportID}
	var out WarehouseExport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Dışa aktarımı güncelle (ad, tablolar, zamanlama, etkin/duraklat, hedef ayarı, kimlik bilgisi)
//
// PATCH /v1/apps/{appId}/warehouse-exports/{exportId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) Update(ctx context.Context, exportID string, body *UpdateWarehouseExportDto, opts ...RequestOption) (*WarehouseExport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/warehouse-exports/{exportId}"}
	op.pathParams = map[string]string{"exportId": exportID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out WarehouseExport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete: Dışa aktarımı sil (hedefteki veriler silinmez)
//
// DELETE /v1/apps/{appId}/warehouse-exports/{exportId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) Delete(ctx context.Context, exportID string, opts ...RequestOption) (*Ok, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "DELETE", path: "/v1/apps/{appId}/warehouse-exports/{exportId}"}
	op.pathParams = map[string]string{"exportId": exportID}
	var out Ok
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// TestSaved: Kayıtlı dışa aktarımın bağlantı testi
//
// POST /v1/apps/{appId}/warehouse-exports/{exportId}/test
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) TestSaved(ctx context.Context, exportID string, opts ...RequestOption) (*WarehouseConnectionResult, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/warehouse-exports/{exportId}/test"}
	op.pathParams = map[string]string{"exportId": exportID}
	var out WarehouseConnectionResult
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Run: Şimdi çalıştır (son başarılı imleçten itibaren)
//
// POST /v1/apps/{appId}/warehouse-exports/{exportId}/run
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) Run(ctx context.Context, exportID string, opts ...RequestOption) (*WarehouseExportRun, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/warehouse-exports/{exportId}/run"}
	op.pathParams = map[string]string{"exportId": exportID}
	var out WarehouseExportRun
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// WarehouseExportsRunsParams: WarehouseExportsService.Runs için isteğe bağlı sorgu parametreleri.
type WarehouseExportsRunsParams struct {
	Limit  *int64
	Cursor *string
}

func (p *WarehouseExportsRunsParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
}

// Runs: Çalıştırma geçmişi (cursor)
//
// GET /v1/apps/{appId}/warehouse-exports/{exportId}/runs
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WarehouseExportsService) Runs(ctx context.Context, exportID string, params *WarehouseExportsRunsParams, opts ...RequestOption) (*WarehouseExportRunPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/warehouse-exports/{exportId}/runs"}
	op.pathParams = map[string]string{"exportId": exportID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out WarehouseExportRunPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RunsAll: Runs sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *WarehouseExportsService) RunsAll(ctx context.Context, exportID string, params *WarehouseExportsRunsParams, opts ...RequestOption) iter.Seq2[WarehouseExportRun, error] {
	var p WarehouseExportsRunsParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]WarehouseExportRun, *string, bool, error) {
		if cursor != nil {
			p.Cursor = cursor
		}
		page, err := s.Runs(ctx, exportID, &p, opts...)
		if err != nil {
			return nil, nil, false, err
		}
		return page.Data, page.NextCursor, page.HasMore, nil
	})
}

// WebCampaignsService: webCampaigns kaynağı.
type WebCampaignsService struct {
	client *Client
}

// WebCampaignsListParams: WebCampaignsService.List için isteğe bağlı sorgu parametreleri.
type WebCampaignsListParams struct {
	Limit  *int64
	Cursor *string
	// Değerler: "draft", "active", "paused", "archived".
	Status *string
}

func (p *WebCampaignsListParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "limit", p.Limit)
	addQuery(q, "cursor", p.Cursor)
	addQuery(q, "status", p.Status)
}

// List: Web kampanyaları
//
// GET /v1/apps/{appId}/web-campaigns
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebCampaignsService) List(ctx context.Context, params *WebCampaignsListParams, opts ...RequestOption) (*WebCampaignPage, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/web-campaigns"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out WebCampaignPage
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListAll: List sonuçlarını next_cursor ile tüm sayfalar boyunca tek tek döndürür.
// Hata olursa döngü (sıfır değer, hata) ile biter.
func (s *WebCampaignsService) ListAll(ctx context.Context, params *WebCampaignsListParams, opts ...RequestOption) iter.Seq2[WebCampaign, error] {
	var p WebCampaignsListParams
	if params != nil {
		p = *params
	}
	return paginate(func(cursor *string) ([]WebCampaign, *string, bool, error) {
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

// Create: Web kampanyası oluştur (taslak)
//
// POST /v1/apps/{appId}/web-campaigns
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebCampaignsService) Create(ctx context.Context, body *CreateWebCampaignDto, opts ...RequestOption) (*WebCampaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/web-campaigns"}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out WebCampaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get: Web kampanyası detayı ve istatistikleri
//
// GET /v1/apps/{appId}/web-campaigns/{campaignId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebCampaignsService) Get(ctx context.Context, campaignID string, opts ...RequestOption) (*WebCampaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/web-campaigns/{campaignId}"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	var out WebCampaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update: Web kampanyasını düzenle
//
// PATCH /v1/apps/{appId}/web-campaigns/{campaignId}
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebCampaignsService) Update(ctx context.Context, campaignID string, body *UpdateWebCampaignDto, opts ...RequestOption) (*WebCampaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "PATCH", path: "/v1/apps/{appId}/web-campaigns/{campaignId}"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out WebCampaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// WebCampaignsReportParams: WebCampaignsService.Report için isteğe bağlı sorgu parametreleri.
type WebCampaignsReportParams struct {
	Days *int64
}

func (p *WebCampaignsReportParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "days", p.Days)
}

// Report: Günlük seri
//
// GET /v1/apps/{appId}/web-campaigns/{campaignId}/report
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebCampaignsService) Report(ctx context.Context, campaignID string, params *WebCampaignsReportParams, opts ...RequestOption) (*WebCampaignReport, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/web-campaigns/{campaignId}/report"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out WebCampaignReport
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Status: Yayınla / duraklat / arşivle
//
// POST /v1/apps/{appId}/web-campaigns/{campaignId}/status
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WebCampaignsService) Status(ctx context.Context, campaignID string, body *WebCampaignStatusDto, opts ...RequestOption) (*WebCampaign, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/web-campaigns/{campaignId}/status"}
	op.pathParams = map[string]string{"campaignId": campaignID}
	if body == nil {
		return nil, errNilBody
	}
	op.body = body
	var out WebCampaign
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// WhatsappService: whatsapp kaynağı.
type WhatsappService struct {
	client *Client
}

// WhatsappTemplatesParams: WhatsappService.Templates için isteğe bağlı sorgu parametreleri.
type WhatsappTemplatesParams struct {
	Status *string
}

func (p *WhatsappTemplatesParams) appendQuery(q url.Values) {
	if p == nil {
		return
	}
	addQuery(q, "status", p.Status)
}

// Templates: WhatsApp şablonları (Meta onay durumu ile)
//
// GET /v1/apps/{appId}/whatsapp/templates
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WhatsappService) Templates(ctx context.Context, params *WhatsappTemplatesParams, opts ...RequestOption) (*WhatsAppTemplateList, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "GET", path: "/v1/apps/{appId}/whatsapp/templates"}
	op.query = url.Values{}
	params.appendQuery(op.query)
	var out WhatsAppTemplateList
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Sync: WhatsApp şablonlarını Meta Graph API’den senkronla
//
// POST /v1/apps/{appId}/whatsapp/templates/sync
//
// App kapsamı istemcinin varsayılan app’i veya ForApp seçeneğidir.
func (s *WhatsappService) Sync(ctx context.Context, opts ...RequestOption) (*WhatsAppTemplateSync, error) {
	ro := s.client.requestConfig(opts)
	op := &operation{method: "POST", path: "/v1/apps/{appId}/whatsapp/templates/sync"}
	var out WhatsAppTemplateSync
	if err := s.client.do(ctx, op, ro, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
