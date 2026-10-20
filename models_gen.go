// Bu dosya api/openapi.yaml'dan sdk/scripts/generate-server-sdks.mjs ile üretilmiştir; elle düzenlemeyin.

package pushizo

// AdAccount: AdAccount şeması.
type AdAccount struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu. Değerler: "meta", "google", "tiktok", "linkedin".
	Platform string `json:"platform"`
	// Zorunlu.
	ExternalID string `json:"external_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Currency *string `json:"currency"`
	// Zorunlu.
	Timezone *string `json:"timezone"`
	// Zorunlu. Değerler: "active", "error", "revoked".
	Status string `json:"status"`
	// Zorunlu.
	Scopes []string `json:"scopes"`
	// Zorunlu.
	LoginCustomerID *string `json:"login_customer_id"`
	// Zorunlu.
	TokenExpiresAt *string `json:"token_expires_at"`
	// Zorunlu.
	LastRefreshedAt *string `json:"last_refreshed_at"`
	// Zorunlu.
	LastError *string `json:"last_error"`
	// Zorunlu.
	LastErrorAt *string `json:"last_error_at"`
	// Zorunlu.
	TermsAcceptedAt *string `json:"terms_accepted_at"`
	// Zorunlu.
	TermsVersion *string `json:"terms_version"`
	// Zorunlu.
	Audiences int64 `json:"audiences"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// AdAccountList: AdAccountList şeması.
type AdAccountList struct {
	// Zorunlu.
	Data []AdAccount `json:"data"`
}

// AdAttributionReport: AdAttributionReport şeması.
type AdAttributionReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu. Değerler: "last_click", "first_click", "linear", "time_decay".
	Model string `json:"model"`
	// Zorunlu.
	WindowDays int64 `json:"window_days"`
	// Zorunlu.
	HalfLifeDays int64 `json:"half_life_days"`
	// Zorunlu.
	Currency string `json:"currency"`
	// Zorunlu.
	GroupBy string `json:"group_by"`
	// Zorunlu.
	ConversionEvents []string `json:"conversion_events"`
	// Zorunlu.
	Rows []AdAttributionReportRowsItem `json:"rows"`
	// Zorunlu.
	Totals AdAttributionReportTotals `json:"totals"`
	// Zorunlu.
	Truncated bool `json:"truncated"`
	// Zorunlu.
	FxMissing []string `json:"fx_missing"`
}

// AdAttributionReportRowsItem: AdAttributionReportRowsItem şeması.
type AdAttributionReportRowsItem struct {
	// Zorunlu.
	Key string `json:"key"`
	// Zorunlu. Değerler: "meta", "google", "tiktok", "linkedin", "microsoft", "other_paid", "utm", "push", "email", "sms", "whatsapp", "rcs", "direct".
	Channel string `json:"channel"`
	// Zorunlu.
	Label string `json:"label"`
	// Zorunlu.
	Paid bool `json:"paid"`
	// Zorunlu.
	Conversions float64 `json:"conversions"`
	// Zorunlu.
	Revenue float64 `json:"revenue"`
	// Zorunlu.
	NewCustomers float64 `json:"new_customers"`
	// Zorunlu.
	Spend *float64 `json:"spend"`
	// Zorunlu.
	Roas *float64 `json:"roas"`
	// Zorunlu.
	Cac *float64 `json:"cac"`
}

// AdAttributionReportTotals: AdAttributionReportTotals şeması.
type AdAttributionReportTotals struct {
	// Zorunlu.
	Conversions float64 `json:"conversions"`
	// Zorunlu.
	Revenue float64 `json:"revenue"`
	// Zorunlu.
	Spend float64 `json:"spend"`
	// Zorunlu.
	NewCustomers float64 `json:"new_customers"`
	// Zorunlu.
	Roas *float64 `json:"roas"`
	// Zorunlu.
	PlatformRoas *float64 `json:"platform_roas"`
	// Zorunlu.
	Cac *float64 `json:"cac"`
	// Zorunlu.
	PaidCac *float64 `json:"paid_cac"`
}

// AdAudience: AdAudience şeması.
type AdAudience struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	AdAccountID string `json:"ad_account_id"`
	// Zorunlu. Değerler: "meta", "google", "tiktok", "linkedin".
	Platform string `json:"platform"`
	// Zorunlu. Değerler: "segment", "journey", "lookalike".
	Source string `json:"source"`
	// Zorunlu.
	SegmentID *string `json:"segment_id"`
	// Zorunlu.
	SegmentName *string `json:"segment_name"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu. Değerler: "pending", "active", "paused", "error".
	Status string `json:"status"`
	// Zorunlu. Değerler: "email", "phone", "maid".
	Identifiers []string `json:"identifiers"`
	// Zorunlu.
	SyncIntervalMinutes int64 `json:"sync_interval_minutes"`
	// Zorunlu.
	NextSyncAt *string `json:"next_sync_at"`
	// Zorunlu.
	LastSyncedAt *string `json:"last_synced_at"`
	// Zorunlu. Değerler: "succeeded", "partial", "failed".
	LastSyncStatus *string `json:"last_sync_status"`
	// Zorunlu.
	FullReloadRequested bool `json:"full_reload_requested"`
	// Zorunlu.
	MemberCount int64 `json:"member_count"`
	// Zorunlu.
	EstimatedSize *int64 `json:"estimated_size"`
	// Zorunlu.
	MatchRate *float64 `json:"match_rate"`
	// Zorunlu.
	SizeUpdatedAt *string `json:"size_updated_at"`
	// Zorunlu.
	LookalikeSourceID *string `json:"lookalike_source_id"`
	// Zorunlu.
	LookalikeSpec *AdAudienceLookalikeSpec `json:"lookalike_spec"`
	// Zorunlu.
	ConsecutiveFailures int64 `json:"consecutive_failures"`
	// Zorunlu.
	LastError *string `json:"last_error"`
	// Zorunlu.
	LastErrorAt *string `json:"last_error_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// AdAudienceList: AdAudienceList şeması.
type AdAudienceList struct {
	// Zorunlu.
	Data []AdAudience `json:"data"`
}

// AdAudienceLookalikeSpec: AdAudienceLookalikeSpec şeması.
type AdAudienceLookalikeSpec struct {
	// Zorunlu.
	Country string `json:"country"`
	// Zorunlu.
	Ratio float64 `json:"ratio"`
}

// AdAudienceSyncList: AdAudienceSyncList şeması.
type AdAudienceSyncList struct {
	// Zorunlu.
	Data []AdAudienceSyncListDataItem `json:"data"`
}

// AdAudienceSyncListDataItem: AdAudienceSyncListDataItem şeması.
type AdAudienceSyncListDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AudienceID string `json:"audience_id"`
	// Zorunlu. Değerler: "schedule", "manual", "journey".
	Trigger string `json:"trigger"`
	// Zorunlu. Değerler: "incremental", "full".
	Mode string `json:"mode"`
	// Zorunlu. Değerler: "running", "succeeded", "partial", "failed".
	Status string `json:"status"`
	// Zorunlu.
	Added int64 `json:"added"`
	// Zorunlu.
	Removed int64 `json:"removed"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	SkippedNoConsent int64 `json:"skipped_no_consent"`
	// Zorunlu.
	SkippedNoIdentifier int64 `json:"skipped_no_identifier"`
	// Zorunlu.
	Batches int64 `json:"batches"`
	// Zorunlu.
	Retries int64 `json:"retries"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	StartedAt string `json:"started_at"`
	// Zorunlu.
	FinishedAt *string `json:"finished_at"`
	// Zorunlu.
	DurationMs *int64 `json:"duration_ms"`
}

// AdAudienceSyncQueued: AdAudienceSyncQueued şeması.
type AdAudienceSyncQueued struct {
	// Zorunlu.
	Queued bool `json:"queued"`
	// Zorunlu.
	NextSyncAt string `json:"next_sync_at"`
	// Zorunlu. Değerler: "incremental", "full".
	Mode string `json:"mode"`
}

// AdBudgetAlertList: AdBudgetAlertList şeması.
type AdBudgetAlertList struct {
	// Zorunlu.
	Data []AdBudgetAlertListDataItem `json:"data"`
}

// AdBudgetAlertListDataItem: AdBudgetAlertListDataItem şeması.
type AdBudgetAlertListDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	RuleID string `json:"rule_id"`
	// Zorunlu.
	RuleName string `json:"rule_name"`
	// Zorunlu. Değerler: "daily_spend", "roas_drop", "budget_depletion".
	Kind string `json:"kind"`
	// Zorunlu.
	Day string `json:"day"`
	// Zorunlu.
	Value float64 `json:"value"`
	// Zorunlu.
	Threshold float64 `json:"threshold"`
	// Zorunlu.
	Details map[string]any `json:"details"`
	// Zorunlu.
	Emailed int64 `json:"emailed"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// AdBudgetAlertRule: AdBudgetAlertRule şeması.
type AdBudgetAlertRule struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "daily_spend", "roas_drop", "budget_depletion".
	Kind string `json:"kind"`
	// Zorunlu.
	AdAccountID *string `json:"ad_account_id"`
	// Zorunlu.
	CampaignExternalID *string `json:"campaign_external_id"`
	// Zorunlu.
	Threshold float64 `json:"threshold"`
	// Zorunlu.
	WindowDays int64 `json:"window_days"`
	// Zorunlu.
	Recipients []string `json:"recipients"`
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu.
	LastValue *float64 `json:"last_value"`
	// Zorunlu.
	LastEvaluatedAt *string `json:"last_evaluated_at"`
	// Zorunlu.
	LastTriggeredOn *string `json:"last_triggered_on"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// AdBudgetAlertRuleList: AdBudgetAlertRuleList şeması.
type AdBudgetAlertRuleList struct {
	// Zorunlu.
	Data []AdBudgetAlertRule `json:"data"`
}

// AdCapiConfig: AdCapiConfig şeması.
type AdCapiConfig struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu. Değerler: "meta", "google", "tiktok".
	Platform string `json:"platform"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	AdAccountID *string `json:"ad_account_id"`
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu.
	Settings AdCapiConfigSettings `json:"settings"`
	// Zorunlu.
	EventMappings []AdCapiConfigEventMappingsItem `json:"event_mappings"`
	// Zorunlu.
	TestEventCode *string `json:"test_event_code"`
	// Zorunlu.
	TestMode bool `json:"test_mode"`
	// Zorunlu. Değerler: "website", "app".
	ActionSource string `json:"action_source"`
	// Zorunlu.
	DefaultCurrency string `json:"default_currency"`
	// Zorunlu.
	HasAccessToken bool `json:"has_access_token"`
	// Zorunlu.
	LastSentAt *string `json:"last_sent_at"`
	// Zorunlu.
	LastError *string `json:"last_error"`
	// Zorunlu.
	LastErrorAt *string `json:"last_error_at"`
	// Zorunlu.
	Stats AdCapiConfigStats `json:"stats"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// AdCapiConfigEventMappingsItem: AdCapiConfigEventMappingsItem şeması.
type AdCapiConfigEventMappingsItem struct {
	// Zorunlu.
	EventName string `json:"event_name"`
	// Zorunlu.
	PlatformEvent string `json:"platform_event"`
	// Zorunlu.
	ValueProperty *string `json:"value_property"`
	// Zorunlu.
	CurrencyProperty *string `json:"currency_property"`
}

// AdCapiConfigList: AdCapiConfigList şeması.
type AdCapiConfigList struct {
	// Zorunlu.
	Data []AdCapiConfig `json:"data"`
}

// AdCapiConfigSettings: AdCapiConfigSettings şeması.
type AdCapiConfigSettings struct {
	// Zorunlu.
	PixelID *string `json:"pixel_id"`
	// Zorunlu.
	CustomerID *string `json:"customer_id"`
	// Zorunlu.
	LoginCustomerID *string `json:"login_customer_id"`
	// Zorunlu.
	PixelCode *string `json:"pixel_code"`
}

// AdCapiConfigStats: AdCapiConfigStats şeması.
type AdCapiConfigStats struct {
	// Zorunlu.
	Pending int64 `json:"pending"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Skipped int64 `json:"skipped"`
}

// AdCapiDelivery: AdCapiDelivery şeması.
type AdCapiDelivery struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	ConfigID string `json:"config_id"`
	// Zorunlu. Değerler: "meta", "google", "tiktok".
	Platform string `json:"platform"`
	// Zorunlu.
	EventID string `json:"event_id"`
	// Zorunlu.
	EventName string `json:"event_name"`
	// Zorunlu.
	PlatformEvent string `json:"platform_event"`
	// Zorunlu.
	DedupID string `json:"dedup_id"`
	// Zorunlu.
	UserID *string `json:"user_id"`
	// Zorunlu.
	Value *float64 `json:"value"`
	// Zorunlu.
	Currency *string `json:"currency"`
	// Zorunlu. Değerler: "pending", "sent", "failed", "skipped".
	Status string `json:"status"`
	// Zorunlu.
	SkipReason *string `json:"skip_reason"`
	// Zorunlu.
	Attempts int64 `json:"attempts"`
	// Zorunlu.
	NextAttemptAt *string `json:"next_attempt_at"`
	// Zorunlu.
	LastError *string `json:"last_error"`
	// Zorunlu.
	Test bool `json:"test"`
	// Zorunlu.
	EventOccurredAt string `json:"event_occurred_at"`
	// Zorunlu.
	SentAt *string `json:"sent_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// AdCapiDeliveryList: AdCapiDeliveryList şeması.
type AdCapiDeliveryList struct {
	// Zorunlu.
	Data []AdCapiDelivery `json:"data"`
}

// AddAliasDto: AddAliasDto şeması.
type AddAliasDto struct {
	// Zorunlu.
	Identity map[string]string `json:"identity"`
}

// AddSubscriptionDto: AddSubscriptionDto şeması.
type AddSubscriptionDto struct {
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu.
	Token  string                    `json:"token"`
	Device *AddSubscriptionDtoDevice `json:"device,omitempty"`
	// Değerler: "subscribed", "unsubscribed".
	Status *string                  `json:"status,omitempty"`
	OptIn  *AddSubscriptionDtoOptIn `json:"opt_in,omitempty"`
}

// AddSubscriptionDtoDevice: AddSubscriptionDtoDevice şeması.
type AddSubscriptionDtoDevice struct {
	Model       *string `json:"model,omitempty"`
	Os          *string `json:"os,omitempty"`
	OsVersion   *string `json:"os_version,omitempty"`
	AppVersion  *string `json:"app_version,omitempty"`
	SDKVersion  *string `json:"sdk_version,omitempty"`
	SDKPlatform *string `json:"sdk_platform,omitempty"`
	Carrier     *string `json:"carrier,omitempty"`
	Locale      *string `json:"locale,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
	PushCapable *bool   `json:"push_capable,omitempty"`
	// Değerler: "sandbox", "production".
	PushEnvironment *string `json:"push_environment,omitempty"`
}

// AddSubscriptionDtoOptIn: AddSubscriptionDtoOptIn şeması.
type AddSubscriptionDtoOptIn struct {
	// Zorunlu.
	Source string  `json:"source"`
	At     *string `json:"at,omitempty"`
}

// AdEntity: AdEntity şeması.
type AdEntity struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	AdAccountID string `json:"ad_account_id"`
	// Zorunlu. Değerler: "meta", "google", "tiktok".
	Platform string `json:"platform"`
	// Zorunlu. Değerler: "campaign", "ad_set", "ad".
	Level string `json:"level"`
	// Zorunlu.
	ExternalID string `json:"external_id"`
	// Zorunlu.
	CampaignExternalID string `json:"campaign_external_id"`
	// Zorunlu.
	AdSetExternalID *string `json:"ad_set_external_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "active", "paused", "pending", "rejected", "archived", "deleted", "unknown".
	Status string `json:"status"`
	// Zorunlu.
	PlatformStatus *string `json:"platform_status"`
	// Zorunlu.
	Objective *string `json:"objective"`
	// Zorunlu.
	OptimizationGoal *string `json:"optimization_goal"`
	// Zorunlu.
	Currency *string `json:"currency"`
	// Zorunlu.
	DailyBudget *string `json:"daily_budget"`
	// Zorunlu.
	LifetimeBudget *string `json:"lifetime_budget"`
	// Zorunlu.
	StartAt *string `json:"start_at"`
	// Zorunlu.
	EndAt *string `json:"end_at"`
	// Zorunlu.
	AudienceExternalIds []string `json:"audience_external_ids"`
	// Zorunlu.
	Creative *AdEntityCreative `json:"creative"`
	// Zorunlu.
	BudgetEditable bool `json:"budget_editable"`
	// Zorunlu.
	CreatedByPushizo bool `json:"created_by_pushizo"`
	// Zorunlu.
	LastSeenAt *string `json:"last_seen_at"`
	// Zorunlu.
	PlatformURL string `json:"platform_url"`
	// Zorunlu.
	Metrics *AdEntityMetrics `json:"metrics"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// AdEntityCreative: AdEntityCreative şeması.
type AdEntityCreative struct {
	// Zorunlu.
	Title *string `json:"title"`
	// Zorunlu.
	Body *string `json:"body"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	LinkURL *string `json:"link_url"`
	// Zorunlu.
	ImageURL *string `json:"image_url"`
	// Zorunlu.
	CallToAction *string `json:"call_to_action"`
}

// AdEntityList: AdEntityList şeması.
type AdEntityList struct {
	// Zorunlu.
	Data []AdEntity `json:"data"`
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	Accounts []AdEntityListAccountsItem `json:"accounts"`
}

// AdEntityListAccountsItem: AdEntityListAccountsItem şeması.
type AdEntityListAccountsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	CampaignsSyncedAt *string `json:"campaigns_synced_at"`
	// Zorunlu.
	InsightsSyncedAt *string `json:"insights_synced_at"`
	// Zorunlu.
	CampaignsSyncError *string `json:"campaigns_sync_error"`
	// Zorunlu.
	InsightsSyncError *string `json:"insights_sync_error"`
}

// AdEntityMetrics: AdEntityMetrics şeması.
type AdEntityMetrics struct {
	// Zorunlu.
	Impressions float64 `json:"impressions"`
	// Zorunlu.
	Clicks float64 `json:"clicks"`
	// Zorunlu.
	Spend float64 `json:"spend"`
	// Zorunlu.
	Conversions float64 `json:"conversions"`
	// Zorunlu.
	ConversionValue float64 `json:"conversion_value"`
	// Zorunlu.
	Ctr *float64 `json:"ctr"`
	// Zorunlu.
	Cpc *float64 `json:"cpc"`
	// Zorunlu.
	Roas *float64 `json:"roas"`
}

// AdLtvReport: AdLtvReport şeması.
type AdLtvReport struct {
	// Zorunlu.
	Currency string `json:"currency"`
	// Zorunlu.
	Months int64 `json:"months"`
	// Zorunlu.
	ConversionEvents []string `json:"conversion_events"`
	// Zorunlu.
	Cohorts []AdLtvReportCohortsItem `json:"cohorts"`
	// Zorunlu.
	Average []float64 `json:"average"`
	// Zorunlu.
	FxMissing []string `json:"fx_missing"`
}

// AdLtvReportCohortsItem: AdLtvReportCohortsItem şeması.
type AdLtvReportCohortsItem struct {
	// Zorunlu.
	Cohort string `json:"cohort"`
	// Zorunlu.
	Size int64 `json:"size"`
	// Zorunlu.
	Revenue []float64 `json:"revenue"`
	// Zorunlu.
	Ltv []float64 `json:"ltv"`
}

// AdOk: AdOk şeması.
type AdOk struct {
	// Zorunlu.
	Ok bool `json:"ok"`
}

// AdRefreshQueued: AdRefreshQueued şeması.
type AdRefreshQueued struct {
	// Zorunlu.
	Queued bool `json:"queued"`
	// Zorunlu.
	NextSyncAt string `json:"next_sync_at"`
}

// AdSpendReport: AdSpendReport şeması.
type AdSpendReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	Currency string `json:"currency"`
	// Zorunlu.
	GroupBy string `json:"group_by"`
	// Zorunlu.
	Rows []AdSpendReportRowsItem `json:"rows"`
	// Zorunlu.
	Totals AdSpendReportTotals `json:"totals"`
	// Zorunlu.
	FxMissing []string `json:"fx_missing"`
	// Zorunlu. Değerler: "postgres", "clickhouse".
	Store string `json:"store"`
}

// AdSpendReportRowsItem: AdSpendReportRowsItem şeması.
type AdSpendReportRowsItem struct {
	// Zorunlu.
	Key string `json:"key"`
	// Zorunlu.
	Label string `json:"label"`
	// Zorunlu.
	Platform *string `json:"platform"`
	// Zorunlu.
	Impressions float64 `json:"impressions"`
	// Zorunlu.
	Clicks float64 `json:"clicks"`
	// Zorunlu.
	Spend float64 `json:"spend"`
	// Zorunlu.
	Conversions float64 `json:"conversions"`
	// Zorunlu.
	ConversionValue float64 `json:"conversion_value"`
	// Zorunlu.
	Ctr *float64 `json:"ctr"`
	// Zorunlu.
	Cpc *float64 `json:"cpc"`
	// Zorunlu.
	Cpa *float64 `json:"cpa"`
	// Zorunlu.
	Roas *float64 `json:"roas"`
}

// AdSpendReportTotals: AdSpendReportTotals şeması.
type AdSpendReportTotals struct {
	// Zorunlu.
	Impressions float64 `json:"impressions"`
	// Zorunlu.
	Clicks float64 `json:"clicks"`
	// Zorunlu.
	Spend float64 `json:"spend"`
	// Zorunlu.
	Conversions float64 `json:"conversions"`
	// Zorunlu.
	ConversionValue float64 `json:"conversion_value"`
	// Zorunlu.
	Ctr *float64 `json:"ctr"`
	// Zorunlu.
	Cpc *float64 `json:"cpc"`
	// Zorunlu.
	Cpa *float64 `json:"cpa"`
	// Zorunlu.
	Roas *float64 `json:"roas"`
}

// APIEventsDto: ApiEventsDto şeması.
type APIEventsDto struct {
	// Zorunlu.
	Events []APIEventsDtoEventsItem `json:"events"`
}

// APIEventsDtoEventsItem: ApiEventsDtoEventsItem şeması.
type APIEventsDtoEventsItem struct {
	// Zorunlu.
	Name           string         `json:"name"`
	Properties     map[string]any `json:"properties,omitempty"`
	Value          *float64       `json:"value,omitempty"`
	OccurredAt     *string        `json:"occurred_at,omitempty"`
	EventID        *string        `json:"event_id,omitempty"`
	SubscriptionID *string        `json:"subscription_id,omitempty"`
	ExternalID     *string        `json:"external_id,omitempty"`
}

// APIKey: ApiKey şeması.
type APIKey struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	OrgID string `json:"org_id"`
	// Zorunlu.
	AppID *string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Prefix string `json:"prefix"`
	// Zorunlu.
	Scopes []string `json:"scopes"`
	// Zorunlu.
	LastUsedAt *string `json:"last_used_at"`
	// Zorunlu.
	ExpiresAt *string `json:"expires_at"`
	// Zorunlu.
	RevokedAt *string `json:"revoked_at"`
	// Zorunlu.
	IPAllowlist []string `json:"ip_allowlist"`
	// Zorunlu.
	RotatedAt *string `json:"rotated_at"`
	// Zorunlu.
	CreatedBy *string `json:"created_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// APIKeyList: ApiKeyList şeması.
type APIKeyList struct {
	// Zorunlu.
	Data []APIKey `json:"data"`
}

// App: App şeması.
type App struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	OrgID string `json:"org_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Slug string `json:"slug"`
	// Zorunlu.
	PlatformFlags AppPlatformFlags `json:"platform_flags"`
	// Zorunlu.
	IconURL *string `json:"icon_url"`
	// Zorunlu.
	DefaultLanguage string `json:"default_language"`
	// Zorunlu.
	Timezone string `json:"timezone"`
	// Zorunlu.
	Settings AppSettings `json:"settings"`
	// Zorunlu. Değerler: "hmac", "jwt", "both".
	IdentityVerificationMode string `json:"identity_verification_mode"`
	// Zorunlu. Değerler: "active", "suspended", "deleted".
	Status string `json:"status"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// AppList: AppList şeması.
type AppList struct {
	// Zorunlu.
	Data []App `json:"data"`
}

// AppPlatformFlags: AppPlatformFlags şeması.
type AppPlatformFlags struct {
	IOS     *bool `json:"ios,omitempty"`
	Android *bool `json:"android,omitempty"`
	Web     *bool `json:"web,omitempty"`
	Email   *bool `json:"email,omitempty"`
	SMS     *bool `json:"sms,omitempty"`
}

// AppSettings: AppSettings şeması.
type AppSettings struct {
	QuietHours                   *AppSettingsQuietHours           `json:"quietHours,omitempty"`
	FrequencyCap                 *AppSettingsFrequencyCap         `json:"frequencyCap,omitempty"`
	FrequencyCaps                []AppSettingsFrequencyCapsItem   `json:"frequencyCaps,omitempty"`
	InApp                        *AppSettingsInApp                `json:"inApp,omitempty"`
	AndroidChannels              []AppSettingsAndroidChannelsItem `json:"androidChannels,omitempty"`
	SMS                          *AppSettingsSMS                  `json:"sms,omitempty"`
	PreferenceCenter             *AppSettingsPreferenceCenter     `json:"preferenceCenter,omitempty"`
	DefaultTTLSeconds            *int64                           `json:"defaultTtlSeconds,omitempty"`
	IdentityVerificationRequired *bool                            `json:"identityVerificationRequired,omitempty"`
	AllowedOrigins               []string                         `json:"allowedOrigins,omitempty"`
	CampaignApprovalRequired     *bool                            `json:"campaignApprovalRequired,omitempty"`
	WebPush                      *AppSettingsWebPush              `json:"webPush,omitempty"`
	Onboarding                   *AppSettingsOnboarding           `json:"onboarding,omitempty"`
	SDKSetup                     *AppSettingsSDKSetup             `json:"sdkSetup,omitempty"`
}

// AppSettingsAndroidChannelsItem: AppSettingsAndroidChannelsItem şeması.
type AppSettingsAndroidChannelsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Zorunlu. Değerler: "low", "default", "high", "urgent".
	Importance string  `json:"importance"`
	Sound      *string `json:"sound,omitempty"`
	Vibration  *bool   `json:"vibration,omitempty"`
	Group      *string `json:"group,omitempty"`
}

// AppSettingsFrequencyCap: AppSettingsFrequencyCap şeması.
type AppSettingsFrequencyCap struct {
	// Zorunlu.
	Max int64 `json:"max"`
	// Zorunlu.
	WindowHours int64 `json:"windowHours"`
}

// AppSettingsFrequencyCapsItem: AppSettingsFrequencyCapsItem şeması.
type AppSettingsFrequencyCapsItem struct {
	// Zorunlu.
	Max int64 `json:"max"`
	// Zorunlu.
	Window AppSettingsFrequencyCapsItemWindow `json:"window"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels []string `json:"channels,omitempty"`
}

// AppSettingsFrequencyCapsItemWindow: AppSettingsFrequencyCapsItemWindow şeması.
type AppSettingsFrequencyCapsItemWindow struct {
	// Zorunlu.
	Value int64 `json:"value"`
	// Zorunlu. Değerler: "minutes", "hours", "days", "weeks".
	Unit string `json:"unit"`
}

// AppSettingsInApp: AppSettingsInApp şeması.
type AppSettingsInApp struct {
	MinIntervalSeconds *int64 `json:"minIntervalSeconds,omitempty"`
}

// AppSettingsOnboarding: AppSettingsOnboarding şeması.
type AppSettingsOnboarding struct {
	// Zorunlu.
	Step        int64   `json:"step"`
	Simulation  *bool   `json:"simulation,omitempty"`
	CompletedAt *string `json:"completed_at,omitempty"`
	Version     *int64  `json:"version,omitempty"`
	// Değerler: "ios", "android", "huawei", "web".
	Platforms []string `json:"platforms,omitempty"`
	// Değerler: "native", "flutter", "react-native", "unity", "capacitor", "cordova", "maui".
	Framework *string `json:"framework,omitempty"`
	// Değerler: "code", "wordpress".
	WebIntegration *string `json:"web_integration,omitempty"`
}

// AppSettingsPreferenceCenter: AppSettingsPreferenceCenter şeması.
type AppSettingsPreferenceCenter struct {
	Enabled              *bool                                            `json:"enabled,omitempty"`
	BrandName            *string                                          `json:"brandName,omitempty"`
	LogoURL              *string                                          `json:"logoUrl,omitempty"`
	AccentColor          *string                                          `json:"accentColor,omitempty"`
	CustomUnsubscribeURL *string                                          `json:"customUnsubscribeUrl,omitempty"`
	Texts                map[string]AppSettingsPreferenceCenterTextsValue `json:"texts,omitempty"`
}

// AppSettingsPreferenceCenterTextsValue: AppSettingsPreferenceCenterTextsValue şeması.
type AppSettingsPreferenceCenterTextsValue struct {
	Title          *string `json:"title,omitempty"`
	Description    *string `json:"description,omitempty"`
	UnsubscribeAll *string `json:"unsubscribeAll,omitempty"`
	Saved          *string `json:"saved,omitempty"`
}

// AppSettingsQuietHours: AppSettingsQuietHours şeması.
type AppSettingsQuietHours struct {
	// Zorunlu.
	Start string `json:"start"`
	// Zorunlu.
	End string `json:"end"`
	// Zorunlu. Değerler: "user", "app".
	TimezoneMode string `json:"timezoneMode"`
}

// AppSettingsSDKSetup: AppSettingsSdkSetup şeması.
type AppSettingsSDKSetup struct {
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	Platform *string `json:"platform"`
	// Zorunlu.
	Step string `json:"step"`
	// Zorunlu.
	Furthest               string  `json:"furthest"`
	Variant                *string `json:"variant,omitempty"`
	DetectedSubscriptionID *string `json:"detectedSubscriptionId,omitempty"`
	TestMessageID          *string `json:"testMessageId,omitempty"`
	CompletedAt            *string `json:"completedAt,omitempty"`
}

// AppSettingsSMS: AppSettingsSms şeması.
type AppSettingsSMS struct {
	HelpReply *string `json:"helpReply,omitempty"`
	StopReply *string `json:"stopReply,omitempty"`
	AutoReply *bool   `json:"autoReply,omitempty"`
}

// AppSettingsWebPush: AppSettingsWebPush şeması.
type AppSettingsWebPush struct {
	ServiceWorkerPath   *string                                `json:"serviceWorkerPath,omitempty"`
	ServiceWorkerScope  *string                                `json:"serviceWorkerScope,omitempty"`
	Prompt              *AppSettingsWebPushPrompt              `json:"prompt,omitempty"`
	WelcomeNotification *AppSettingsWebPushWelcomeNotification `json:"welcomeNotification,omitempty"`
	ClickBehavior       *AppSettingsWebPushClickBehavior       `json:"clickBehavior,omitempty"`
	CustomLink          *AppSettingsWebPushCustomLink          `json:"customLink,omitempty"`
	PromptBackoffDays   []int64                                `json:"promptBackoffDays,omitempty"`
}

// AppSettingsWebPushClickBehavior: AppSettingsWebPushClickBehavior şeması.
type AppSettingsWebPushClickBehavior struct {
	// Zorunlu. Değerler: "exact", "origin".
	Match string `json:"match"`
	// Zorunlu. Değerler: "navigate", "focus".
	Action string `json:"action"`
}

// AppSettingsWebPushCustomLink: AppSettingsWebPushCustomLink şeması.
type AppSettingsWebPushCustomLink struct {
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu. Değerler: "button", "link".
	Style string `json:"style"`
	// Zorunlu. Değerler: "small", "medium", "large".
	Size               string                                            `json:"size"`
	Color              *AppSettingsWebPushCustomLinkColor                `json:"color,omitempty"`
	Texts              map[string]AppSettingsWebPushCustomLinkTextsValue `json:"texts,omitempty"`
	UnsubscribeEnabled *bool                                             `json:"unsubscribeEnabled,omitempty"`
}

// AppSettingsWebPushCustomLinkColor: AppSettingsWebPushCustomLinkColor şeması.
type AppSettingsWebPushCustomLinkColor struct {
	Button *string `json:"button,omitempty"`
	Text   *string `json:"text,omitempty"`
}

// AppSettingsWebPushCustomLinkTextsValue: AppSettingsWebPushCustomLinkTextsValue şeması.
type AppSettingsWebPushCustomLinkTextsValue struct {
	Subscribe   *string `json:"subscribe,omitempty"`
	Unsubscribe *string `json:"unsubscribe,omitempty"`
	Explanation *string `json:"explanation,omitempty"`
}

// AppSettingsWebPushPrompt: AppSettingsWebPushPrompt şeması.
type AppSettingsWebPushPrompt struct {
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu. Değerler: "slidedown", "native".
	Type string `json:"type"`
	// Değerler: "push", "email", "sms", "email_sms".
	Collect       *string                                       `json:"collect,omitempty"`
	DelaySeconds  *int64                                        `json:"delaySeconds,omitempty"`
	PageViews     *int64                                        `json:"pageViews,omitempty"`
	ScrollPercent *int64                                        `json:"scrollPercent,omitempty"`
	Texts         map[string]AppSettingsWebPushPromptTextsValue `json:"texts,omitempty"`
	Categories    []AppSettingsWebPushPromptCategoriesItem      `json:"categories,omitempty"`
	AccentColor   *string                                       `json:"accentColor,omitempty"`
}

// AppSettingsWebPushPromptCategoriesItem: AppSettingsWebPushPromptCategoriesItem şeması.
type AppSettingsWebPushPromptCategoriesItem struct {
	// Zorunlu.
	Tag string `json:"tag"`
	// Zorunlu.
	Label map[string]string `json:"label"`
}

// AppSettingsWebPushPromptTextsValue: AppSettingsWebPushPromptTextsValue şeması.
type AppSettingsWebPushPromptTextsValue struct {
	Title   *string `json:"title,omitempty"`
	Message *string `json:"message,omitempty"`
	Accept  *string `json:"accept,omitempty"`
	Cancel  *string `json:"cancel,omitempty"`
}

// AppSettingsWebPushWelcomeNotification: AppSettingsWebPushWelcomeNotification şeması.
type AppSettingsWebPushWelcomeNotification struct {
	// Zorunlu.
	Enabled bool                                                       `json:"enabled"`
	URL     *string                                                    `json:"url,omitempty"`
	Texts   map[string]AppSettingsWebPushWelcomeNotificationTextsValue `json:"texts,omitempty"`
}

// AppSettingsWebPushWelcomeNotificationTextsValue: AppSettingsWebPushWelcomeNotificationTextsValue şeması.
type AppSettingsWebPushWelcomeNotificationTextsValue struct {
	Title   *string `json:"title,omitempty"`
	Message *string `json:"message,omitempty"`
}

// AuditEntry: AuditEntry şeması.
type AuditEntry struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	OrgID string `json:"org_id"`
	// Zorunlu.
	AppID *string `json:"app_id"`
	// Zorunlu. Değerler: "account", "api_key", "system".
	ActorType string `json:"actor_type"`
	// Zorunlu.
	ActorID *string `json:"actor_id"`
	// Zorunlu.
	Action string `json:"action"`
	// Zorunlu.
	TargetType *string `json:"target_type"`
	// Zorunlu.
	TargetID *string `json:"target_id"`
	// Zorunlu.
	Metadata map[string]any `json:"metadata"`
	// Zorunlu.
	IP *string `json:"ip"`
	// Zorunlu.
	RequestID *string `json:"request_id"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// AuditPage: AuditPage şeması.
type AuditPage struct {
	// Zorunlu.
	Data []AuditEntry `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// Campaign: Campaign şeması.
type Campaign struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Labels []string `json:"labels"`
	// Zorunlu.
	Category *string `json:"category"`
	// Zorunlu.
	AbTest CampaignAbTest `json:"ab_test"`
	// Zorunlu.
	Variants []CampaignVariantsItem `json:"variants"`
	// Zorunlu. Değerler: "one_time", "scheduled", "recurring", "automated", "transactional".
	Type string `json:"type"`
	// Zorunlu. Değerler: "draft", "scheduled", "running", "paused", "completed", "cancelled", "failed".
	Status string `json:"status"`
	// Zorunlu.
	Audience map[string]any `json:"audience"`
	// Zorunlu.
	Content map[string]any `json:"content"`
	// Zorunlu.
	TemplateID *string `json:"template_id"`
	// Zorunlu.
	TemplateVersion *int64 `json:"template_version"`
	// Zorunlu.
	Delivery CampaignDelivery `json:"delivery"`
	// Zorunlu.
	Recurrence *CampaignRecurrence `json:"recurrence"`
	// Zorunlu.
	Approval CampaignApproval `json:"approval"`
	// Zorunlu.
	Stats CampaignStats2 `json:"stats"`
	// Zorunlu.
	SendAt *string `json:"send_at"`
	// Zorunlu.
	NextRunAt *string `json:"next_run_at"`
	// Zorunlu.
	LastRunAt *string `json:"last_run_at"`
	// Zorunlu.
	RunCount int64 `json:"run_count"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	CreatedBy *string `json:"created_by"`
	// Zorunlu.
	StartedAt *string `json:"started_at"`
	// Zorunlu.
	CompletedAt *string `json:"completed_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
	// Zorunlu.
	EstimatedRecipients *int64 `json:"estimated_recipients"`
}

// CampaignAbTest: CampaignAbTest şeması.
type CampaignAbTest struct {
	Enabled        *bool                 `json:"enabled,omitempty"`
	TestPercentage *int64                `json:"test_percentage,omitempty"`
	Winner         *CampaignAbTestWinner `json:"winner,omitempty"`
	// Değerler: "testing", "winner_sent".
	Phase           *string                     `json:"phase,omitempty"`
	WinnerVariantID *string                     `json:"winner_variant_id,omitempty"`
	DecidedAt       *string                     `json:"decided_at,omitempty"`
	DecidedBy       *string                     `json:"decided_by,omitempty"`
	Significance    *CampaignAbTestSignificance `json:"significance,omitempty"`
}

// CampaignAbTestSignificance: CampaignAbTestSignificance şeması.
type CampaignAbTestSignificance struct {
	// Zorunlu.
	PValue float64 `json:"p_value"`
	// Zorunlu.
	Significant bool `json:"significant"`
	// Zorunlu.
	Metric string `json:"metric"`
}

// CampaignAbTestWinner: CampaignAbTestWinner şeması.
type CampaignAbTestWinner struct {
	// Zorunlu. Değerler: "manual", "auto".
	Mode string `json:"mode"`
	// Zorunlu. Değerler: "open_rate", "click_rate", "conversion_rate".
	Metric               string `json:"metric"`
	EvaluateAfterMinutes *int64 `json:"evaluate_after_minutes,omitempty"`
}

// CampaignApproval: CampaignApproval şeması.
type CampaignApproval struct {
	// Zorunlu.
	Required bool `json:"required"`
	// Değerler: "send", "schedule".
	PendingAction *string `json:"pending_action,omitempty"`
	RequestedBy   *string `json:"requested_by,omitempty"`
	RequestedAt   *string `json:"requested_at,omitempty"`
	ApprovedBy    *string `json:"approved_by,omitempty"`
	ApprovedAt    *string `json:"approved_at,omitempty"`
	RejectedBy    *string `json:"rejected_by,omitempty"`
	RejectedAt    *string `json:"rejected_at,omitempty"`
}

// CampaignDelivery: CampaignDelivery şeması.
type CampaignDelivery struct {
	// Değerler: "immediate", "per_user_timezone", "optimized".
	TimezoneMode      *string `json:"timezone_mode,omitempty"`
	DeliveryTimeOfDay *string `json:"delivery_time_of_day,omitempty"`
	ThrottlePerMinute *int64  `json:"throttle_per_minute,omitempty"`
	TTL               *int64  `json:"ttl,omitempty"`
	// Değerler: 5, 10.
	Priority           *int64  `json:"priority,omitempty"`
	CollapseID         *string `json:"collapse_id,omitempty"`
	FrequencyCapExempt *bool   `json:"frequency_cap_exempt,omitempty"`
	QuietHoursExempt   *bool   `json:"quiet_hours_exempt,omitempty"`
	OnePerUser         *bool   `json:"one_per_user,omitempty"`
}

// CampaignPage: CampaignPage şeması.
type CampaignPage struct {
	// Zorunlu.
	Data []Campaign `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// CampaignRecurrence: CampaignRecurrence şeması.
type CampaignRecurrence struct {
	// Zorunlu.
	Cron     string  `json:"cron"`
	Until    *string `json:"until,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// CampaignsCreateBody: CampaignsCreateBody şeması.
type CampaignsCreateBody struct {
	Name *string `json:"name,omitempty"`
	// Değerler: "one_time", "scheduled", "recurring".
	Type     *string                      `json:"type,omitempty"`
	Audience *CampaignsCreateBodyAudience `json:"audience,omitempty"`
	// messages.content ile aynı şema (headings, contents, url, buttons…)
	Content map[string]any `json:"content,omitempty"`
	Labels  []string       `json:"labels,omitempty"`
	// Tercih kategorisi anahtarı; kapalı kullanıcılar bastırılır
	Category *string                    `json:"category,omitempty"`
	AbTest   *CampaignsCreateBodyAbTest `json:"ab_test,omitempty"`
	// A/B varyantları (ağırlıklar toplamı 100; en fazla bir is_control)
	Variants        []CampaignsCreateBodyVariantsItem `json:"variants,omitempty"`
	TemplateID      *string                           `json:"template_id,omitempty"`
	TemplateVersion *int64                            `json:"template_version,omitempty"`
	Delivery        *CampaignsCreateBodyDelivery      `json:"delivery,omitempty"`
	Recurrence      *CampaignsCreateBodyRecurrence    `json:"recurrence,omitempty"`
	SendAt          *string                           `json:"send_at,omitempty"`
}

// CampaignsCreateBodyAbTest: CampaignsCreateBodyAbTest şeması.
type CampaignsCreateBodyAbTest struct {
	Enabled        *bool                            `json:"enabled,omitempty"`
	TestPercentage *int64                           `json:"test_percentage,omitempty"`
	Winner         *CampaignsCreateBodyAbTestWinner `json:"winner,omitempty"`
}

// CampaignsCreateBodyAbTestWinner: CampaignsCreateBodyAbTestWinner şeması.
type CampaignsCreateBodyAbTestWinner struct {
	// Değerler: "manual", "auto".
	Mode *string `json:"mode,omitempty"`
	// Değerler: "open_rate", "click_rate", "conversion_rate".
	Metric               *string `json:"metric,omitempty"`
	EvaluateAfterMinutes *int64  `json:"evaluate_after_minutes,omitempty"`
}

// CampaignsCreateBodyAudience: CampaignsCreateBodyAudience şeması.
type CampaignsCreateBodyAudience struct {
	All             *bool    `json:"all,omitempty"`
	Segments        []string `json:"segments,omitempty"`
	ExcludeSegments []string `json:"exclude_segments,omitempty"`
	// Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
	Filters *CampaignsCreateBodyAudienceFilters `json:"filters,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels        []string `json:"channels,omitempty"`
	TestDevicesOnly *bool    `json:"test_devices_only,omitempty"`
}

// CampaignsCreateBodyAudienceFilters: Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
type CampaignsCreateBodyAudienceFilters struct {
	All  []map[string]any `json:"all,omitempty"`
	Any  []map[string]any `json:"any,omitempty"`
	None []map[string]any `json:"none,omitempty"`
}

// CampaignsCreateBodyDelivery: CampaignsCreateBodyDelivery şeması.
type CampaignsCreateBodyDelivery struct {
	// Değerler: "immediate", "per_user_timezone", "optimized".
	TimezoneMode      *string `json:"timezone_mode,omitempty"`
	DeliveryTimeOfDay *string `json:"delivery_time_of_day,omitempty"`
	ThrottlePerMinute *int64  `json:"throttle_per_minute,omitempty"`
	TTL               *int64  `json:"ttl,omitempty"`
	// Değerler: 5, 10.
	Priority           *int64  `json:"priority,omitempty"`
	CollapseID         *string `json:"collapse_id,omitempty"`
	FrequencyCapExempt *bool   `json:"frequency_cap_exempt,omitempty"`
	QuietHoursExempt   *bool   `json:"quiet_hours_exempt,omitempty"`
	OnePerUser         *bool   `json:"one_per_user,omitempty"`
}

// CampaignsCreateBodyRecurrence: CampaignsCreateBodyRecurrence şeması.
type CampaignsCreateBodyRecurrence struct {
	Cron     *string `json:"cron,omitempty"`
	Until    *string `json:"until,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// CampaignsCreateBodyVariantsItem: CampaignsCreateBodyVariantsItem şeması.
type CampaignsCreateBodyVariantsItem struct {
	ID *string `json:"id,omitempty"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64          `json:"weight_percent"`
	Content       map[string]any `json:"content,omitempty"`
	IsControl     *bool          `json:"is_control,omitempty"`
}

// CampaignStats: CampaignStats şeması.
type CampaignStats struct {
	// Zorunlu.
	Campaign Campaign `json:"campaign"`
	// Zorunlu.
	Rates CampaignStatsRates `json:"rates"`
	// Zorunlu.
	Channels []CampaignStatsChannelsItem `json:"channels"`
	// Zorunlu.
	Hourly []CampaignStatsHourlyItem `json:"hourly"`
	// Zorunlu.
	Errors []CampaignStatsErrorsItem `json:"errors"`
	// Zorunlu.
	Messages []Message `json:"messages"`
}

// CampaignStats2: CampaignStats2 şeması.
type CampaignStats2 struct {
	Runs         *int64 `json:"runs,omitempty"`
	Targeted     *int64 `json:"targeted,omitempty"`
	Suppressed   *int64 `json:"suppressed,omitempty"`
	Sent         *int64 `json:"sent,omitempty"`
	Delivered    *int64 `json:"delivered,omitempty"`
	Failed       *int64 `json:"failed,omitempty"`
	Opened       *int64 `json:"opened,omitempty"`
	Clicked      *int64 `json:"clicked,omitempty"`
	Converted    *int64 `json:"converted,omitempty"`
	Bounced      *int64 `json:"bounced,omitempty"`
	Complained   *int64 `json:"complained,omitempty"`
	Unsubscribed *int64 `json:"unsubscribed,omitempty"`
	SMSSegments  *int64 `json:"sms_segments,omitempty"`
}

// CampaignStatsChannelsItem: CampaignStatsChannelsItem şeması.
type CampaignStatsChannelsItem struct {
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu.
	Targeted int64 `json:"targeted"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
}

// CampaignStatsErrorsItem: CampaignStatsErrorsItem şeması.
type CampaignStatsErrorsItem struct {
	// Zorunlu.
	Code string `json:"code"`
	// Zorunlu.
	Count int64 `json:"count"`
}

// CampaignStatsHourlyItem: CampaignStatsHourlyItem şeması.
type CampaignStatsHourlyItem struct {
	// Zorunlu.
	Hour string `json:"hour"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
}

// CampaignStatsRates: CampaignStatsRates şeması.
type CampaignStatsRates struct {
	// Zorunlu.
	DeliveryRate *float64 `json:"delivery_rate"`
	// Zorunlu.
	OpenRate *float64 `json:"open_rate"`
	// Zorunlu.
	ClickRate *float64 `json:"click_rate"`
	// Zorunlu.
	FailureRate *float64 `json:"failure_rate"`
	// Zorunlu.
	ConversionRate *float64 `json:"conversion_rate"`
}

// CampaignsUpdateBody: CampaignsUpdateBody şeması.
type CampaignsUpdateBody struct {
	Name *string `json:"name,omitempty"`
	// Değerler: "one_time", "scheduled", "recurring".
	Type     *string                      `json:"type,omitempty"`
	Audience *CampaignsUpdateBodyAudience `json:"audience,omitempty"`
	// messages.content ile aynı şema (headings, contents, url, buttons…)
	Content map[string]any `json:"content,omitempty"`
	Labels  []string       `json:"labels,omitempty"`
	// Tercih kategorisi anahtarı; kapalı kullanıcılar bastırılır
	Category *string                    `json:"category,omitempty"`
	AbTest   *CampaignsUpdateBodyAbTest `json:"ab_test,omitempty"`
	// A/B varyantları (ağırlıklar toplamı 100; en fazla bir is_control)
	Variants        []CampaignsUpdateBodyVariantsItem `json:"variants,omitempty"`
	TemplateID      *string                           `json:"template_id,omitempty"`
	TemplateVersion *int64                            `json:"template_version,omitempty"`
	Delivery        *CampaignsUpdateBodyDelivery      `json:"delivery,omitempty"`
	Recurrence      *CampaignsUpdateBodyRecurrence    `json:"recurrence,omitempty"`
	SendAt          *string                           `json:"send_at,omitempty"`
}

// CampaignsUpdateBodyAbTest: CampaignsUpdateBodyAbTest şeması.
type CampaignsUpdateBodyAbTest struct {
	Enabled        *bool                            `json:"enabled,omitempty"`
	TestPercentage *int64                           `json:"test_percentage,omitempty"`
	Winner         *CampaignsUpdateBodyAbTestWinner `json:"winner,omitempty"`
}

// CampaignsUpdateBodyAbTestWinner: CampaignsUpdateBodyAbTestWinner şeması.
type CampaignsUpdateBodyAbTestWinner struct {
	// Değerler: "manual", "auto".
	Mode *string `json:"mode,omitempty"`
	// Değerler: "open_rate", "click_rate", "conversion_rate".
	Metric               *string `json:"metric,omitempty"`
	EvaluateAfterMinutes *int64  `json:"evaluate_after_minutes,omitempty"`
}

// CampaignsUpdateBodyAudience: CampaignsUpdateBodyAudience şeması.
type CampaignsUpdateBodyAudience struct {
	All             *bool    `json:"all,omitempty"`
	Segments        []string `json:"segments,omitempty"`
	ExcludeSegments []string `json:"exclude_segments,omitempty"`
	// Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
	Filters *CampaignsUpdateBodyAudienceFilters `json:"filters,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels        []string `json:"channels,omitempty"`
	TestDevicesOnly *bool    `json:"test_devices_only,omitempty"`
}

// CampaignsUpdateBodyAudienceFilters: Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
type CampaignsUpdateBodyAudienceFilters struct {
	All  []map[string]any `json:"all,omitempty"`
	Any  []map[string]any `json:"any,omitempty"`
	None []map[string]any `json:"none,omitempty"`
}

// CampaignsUpdateBodyDelivery: CampaignsUpdateBodyDelivery şeması.
type CampaignsUpdateBodyDelivery struct {
	// Değerler: "immediate", "per_user_timezone", "optimized".
	TimezoneMode      *string `json:"timezone_mode,omitempty"`
	DeliveryTimeOfDay *string `json:"delivery_time_of_day,omitempty"`
	ThrottlePerMinute *int64  `json:"throttle_per_minute,omitempty"`
	TTL               *int64  `json:"ttl,omitempty"`
	// Değerler: 5, 10.
	Priority           *int64  `json:"priority,omitempty"`
	CollapseID         *string `json:"collapse_id,omitempty"`
	FrequencyCapExempt *bool   `json:"frequency_cap_exempt,omitempty"`
	QuietHoursExempt   *bool   `json:"quiet_hours_exempt,omitempty"`
	OnePerUser         *bool   `json:"one_per_user,omitempty"`
}

// CampaignsUpdateBodyRecurrence: CampaignsUpdateBodyRecurrence şeması.
type CampaignsUpdateBodyRecurrence struct {
	Cron     *string `json:"cron,omitempty"`
	Until    *string `json:"until,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// CampaignsUpdateBodyVariantsItem: CampaignsUpdateBodyVariantsItem şeması.
type CampaignsUpdateBodyVariantsItem struct {
	ID *string `json:"id,omitempty"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64          `json:"weight_percent"`
	Content       map[string]any `json:"content,omitempty"`
	IsControl     *bool          `json:"is_control,omitempty"`
}

// CampaignVariantsItem: CampaignVariantsItem şeması.
type CampaignVariantsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	Content map[string]any `json:"content"`
	// Zorunlu.
	IsControl bool `json:"is_control"`
	// Zorunlu.
	Position int64 `json:"position"`
}

// CampaignVariantStats: CampaignVariantStats şeması.
type CampaignVariantStats struct {
	// Zorunlu.
	AbTest CampaignVariantStatsAbTest `json:"ab_test"`
	// Zorunlu.
	Variants []CampaignVariantStatsVariantsItem `json:"variants"`
	// Zorunlu.
	Leader *string `json:"leader"`
	// Zorunlu.
	Significance CampaignVariantStatsSignificance `json:"significance"`
}

// CampaignVariantStatsAbTest: CampaignVariantStatsAbTest şeması.
type CampaignVariantStatsAbTest struct {
	Enabled        *bool                             `json:"enabled,omitempty"`
	TestPercentage *int64                            `json:"test_percentage,omitempty"`
	Winner         *CampaignVariantStatsAbTestWinner `json:"winner,omitempty"`
	// Değerler: "testing", "winner_sent".
	Phase           *string                                 `json:"phase,omitempty"`
	WinnerVariantID *string                                 `json:"winner_variant_id,omitempty"`
	DecidedAt       *string                                 `json:"decided_at,omitempty"`
	DecidedBy       *string                                 `json:"decided_by,omitempty"`
	Significance    *CampaignVariantStatsAbTestSignificance `json:"significance,omitempty"`
}

// CampaignVariantStatsAbTestSignificance: CampaignVariantStatsAbTestSignificance şeması.
type CampaignVariantStatsAbTestSignificance struct {
	// Zorunlu.
	PValue float64 `json:"p_value"`
	// Zorunlu.
	Significant bool `json:"significant"`
	// Zorunlu.
	Metric string `json:"metric"`
}

// CampaignVariantStatsAbTestWinner: CampaignVariantStatsAbTestWinner şeması.
type CampaignVariantStatsAbTestWinner struct {
	// Zorunlu. Değerler: "manual", "auto".
	Mode string `json:"mode"`
	// Zorunlu. Değerler: "open_rate", "click_rate", "conversion_rate".
	Metric               string `json:"metric"`
	EvaluateAfterMinutes *int64 `json:"evaluate_after_minutes,omitempty"`
}

// CampaignVariantStatsSignificance: CampaignVariantStatsSignificance şeması.
type CampaignVariantStatsSignificance struct {
	// Zorunlu.
	PValue float64 `json:"p_value"`
	// Zorunlu.
	Significant bool `json:"significant"`
	// Zorunlu.
	Metric string `json:"metric"`
}

// CampaignVariantStatsVariantsItem: CampaignVariantStatsVariantsItem şeması.
type CampaignVariantStatsVariantsItem struct {
	// Zorunlu.
	VariantID string `json:"variant_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	IsControl bool `json:"is_control"`
	// Zorunlu.
	Targeted int64 `json:"targeted"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Delivered int64 `json:"delivered"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
	// Zorunlu.
	Converted int64 `json:"converted"`
	// Zorunlu.
	Holdout int64 `json:"holdout"`
	// Zorunlu.
	OpenRate float64 `json:"open_rate"`
	// Zorunlu.
	ClickRate float64 `json:"click_rate"`
	// Zorunlu.
	ConversionRate float64 `json:"conversion_rate"`
}

// CannedInputDto: CannedInputDto şeması.
type CannedInputDto struct {
	// Zorunlu.
	Shortcut string `json:"shortcut"`
	// Zorunlu.
	Title string `json:"title"`
	// Zorunlu.
	Body     string  `json:"body"`
	Language *string `json:"language,omitempty"`
}

// CannedResponse: CannedResponse şeması.
type CannedResponse struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Shortcut string `json:"shortcut"`
	// Zorunlu.
	Title string `json:"title"`
	// Zorunlu.
	Body string `json:"body"`
	// Zorunlu.
	Language *string `json:"language"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// CannedResponseList: CannedResponseList şeması.
type CannedResponseList struct {
	// Zorunlu.
	Data []CannedResponse `json:"data"`
}

// CannedUpdateDto: CannedUpdateDto şeması.
type CannedUpdateDto struct {
	Shortcut *string `json:"shortcut,omitempty"`
	Title    *string `json:"title,omitempty"`
	Body     *string `json:"body,omitempty"`
	Language *string `json:"language,omitempty"`
}

// ChannelConfig: ChannelConfig şeması.
type ChannelConfig struct {
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu. Değerler: "apns", "fcm", "hms", "webpush", "smtp", "ses", "resend", "twilio", "netgsm", "meta_whatsapp", "google_rbm", "mock".
	Provider *string `json:"provider"`
	// Zorunlu. Değerler: "unconfigured", "active", "invalid", "suspended".
	Status string `json:"status"`
	// Zorunlu.
	Config ChannelConfigConfig `json:"config"`
	// Zorunlu.
	HasSecret bool `json:"has_secret"`
	// Zorunlu.
	HasHook bool `json:"has_hook"`
	// Zorunlu.
	LastValidatedAt *string `json:"last_validated_at"`
	// Zorunlu.
	LastError *string `json:"last_error"`
	// Zorunlu.
	UpdatedAt *string `json:"updated_at"`
	// Zorunlu.
	SuspendedReason *string `json:"suspended_reason"`
	// Zorunlu.
	DnsCheck *ChannelConfigDnsCheck `json:"dns_check"`
}

// ChannelConfigConfig: ChannelConfigConfig şeması.
type ChannelConfigConfig struct {
	// Değerler: "token", "certificate".
	AuthType *string `json:"auth_type,omitempty"`
	// Değerler: "sandbox", "production", "universal".
	CertificateEnvironment *string `json:"certificate_environment,omitempty"`
	CertificateSubject     *string `json:"certificate_subject,omitempty"`
	CertificateExpiresAt   *string `json:"certificate_expires_at,omitempty"`
	KeyID                  *string `json:"key_id,omitempty"`
	TeamID                 *string `json:"team_id,omitempty"`
	BundleID               *string `json:"bundle_id,omitempty"`
	// Değerler: "sandbox", "production".
	Environment         *string `json:"environment,omitempty"`
	ProjectID           *string `json:"project_id,omitempty"`
	ClientEmail         *string `json:"client_email,omitempty"`
	HmsAppID            *string `json:"hms_app_id,omitempty"`
	HighPriority        *bool   `json:"high_priority,omitempty"`
	Subject             *string `json:"subject,omitempty"`
	VAPIDPublicKey      *string `json:"vapid_public_key,omitempty"`
	SiteURL             *string `json:"site_url,omitempty"`
	FromAddress         *string `json:"from_address,omitempty"`
	FromName            *string `json:"from_name,omitempty"`
	ReplyTo             *string `json:"reply_to,omitempty"`
	DkimSelector        *string `json:"dkim_selector,omitempty"`
	PostalAddress       *string `json:"postal_address,omitempty"`
	Host                *string `json:"host,omitempty"`
	Port                *int64  `json:"port,omitempty"`
	Secure              *bool   `json:"secure,omitempty"`
	Username            *string `json:"username,omitempty"`
	Region              *string `json:"region,omitempty"`
	AccessKeyID         *string `json:"access_key_id,omitempty"`
	ConfigurationSet    *string `json:"configuration_set,omitempty"`
	AccountSid          *string `json:"account_sid,omitempty"`
	FromNumber          *string `json:"from_number,omitempty"`
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	Usercode            *string `json:"usercode,omitempty"`
	Msgheader           *string `json:"msgheader,omitempty"`
	// Değerler: "0", "11", "12".
	IYSFilter          *string `json:"iys_filter,omitempty"`
	WabaID             *string `json:"waba_id,omitempty"`
	PhoneNumberID      *string `json:"phone_number_id,omitempty"`
	GraphAPIVersion    *string `json:"graph_api_version,omitempty"`
	DisplayPhoneNumber *string `json:"display_phone_number,omitempty"`
	VerifiedName       *string `json:"verified_name,omitempty"`
	RbmAgentID         *string `json:"rbm_agent_id,omitempty"`
	ContentSid         *string `json:"content_sid,omitempty"`
	SMSFallback        *bool   `json:"sms_fallback,omitempty"`
}

// ChannelConfigDnsCheck: ChannelConfigDnsCheck şeması.
type ChannelConfigDnsCheck struct {
	// Zorunlu.
	Domain string `json:"domain"`
	// Zorunlu.
	CheckedAt string `json:"checked_at"`
	// Zorunlu.
	Spf ChannelConfigDnsCheckSpf `json:"spf"`
	// Zorunlu.
	Dkim ChannelConfigDnsCheckDkim `json:"dkim"`
	// Zorunlu.
	Dmarc ChannelConfigDnsCheckDmarc `json:"dmarc"`
	// Zorunlu.
	Issues []string `json:"issues"`
}

// ChannelConfigDnsCheckDkim: ChannelConfigDnsCheckDkim şeması.
type ChannelConfigDnsCheckDkim struct {
	// Zorunlu.
	Ok bool `json:"ok"`
	// Zorunlu.
	Selector string `json:"selector"`
	// Zorunlu.
	Record *string `json:"record"`
	// Zorunlu.
	Expected string `json:"expected"`
}

// ChannelConfigDnsCheckDmarc: ChannelConfigDnsCheckDmarc şeması.
type ChannelConfigDnsCheckDmarc struct {
	// Zorunlu.
	Ok bool `json:"ok"`
	// Zorunlu.
	Record *string `json:"record"`
	// Zorunlu.
	Expected string `json:"expected"`
}

// ChannelConfigDnsCheckSpf: ChannelConfigDnsCheckSpf şeması.
type ChannelConfigDnsCheckSpf struct {
	// Zorunlu.
	Ok bool `json:"ok"`
	// Zorunlu.
	Record *string `json:"record"`
	// Zorunlu.
	Expected string `json:"expected"`
}

// ChannelList: ChannelList şeması.
type ChannelList struct {
	// Zorunlu.
	Data []ChannelConfig `json:"data"`
}

// ChannelValidation: ChannelValidation şeması.
type ChannelValidation struct {
	// Zorunlu.
	Ok bool `json:"ok"`
	// Zorunlu. Değerler: "auth", "config", "transient".
	Kind *string `json:"kind"`
	// Zorunlu.
	Code *string `json:"code"`
	// Zorunlu.
	Message *string `json:"message"`
	// Zorunlu.
	Details map[string]any `json:"details"`
	// Zorunlu.
	Channel ChannelConfig `json:"channel"`
}

// ContentTable: ContentTable şeması.
type ContentTable struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	KeyColumn string `json:"key_column"`
	// Zorunlu.
	Columns []string `json:"columns"`
	// Zorunlu.
	RowCount int64 `json:"row_count"`
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	SourceFileName *string `json:"source_file_name"`
	// Zorunlu.
	Liquid string `json:"liquid"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// ContentTableList: ContentTableList şeması.
type ContentTableList struct {
	// Zorunlu.
	Data []ContentTable `json:"data"`
}

// ContentTableRows: ContentTableRows şeması.
type ContentTableRows struct {
	// Zorunlu.
	Columns []string `json:"columns"`
	// Zorunlu.
	KeyColumn string `json:"key_column"`
	// Zorunlu.
	Total int64 `json:"total"`
	// Zorunlu.
	Data []map[string]string `json:"data"`
}

// Conversation: Conversation şeması.
type Conversation struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "sms", "whatsapp", "rcs", "email", "chat".
	Channel string `json:"channel"`
	// Zorunlu. Değerler: "open", "pending", "resolved", "closed".
	Status string `json:"status"`
	// Zorunlu.
	SubscriptionID *string `json:"subscription_id"`
	// Zorunlu.
	SubscriptionStatus *string `json:"subscription_status"`
	// Zorunlu.
	UserID *string `json:"user_id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu.
	Contact string `json:"contact"`
	// Zorunlu.
	OptedIn bool `json:"opted_in"`
	// Zorunlu.
	LastMessageAt string `json:"last_message_at"`
	// Zorunlu.
	LastInboundAt *string `json:"last_inbound_at"`
	// Zorunlu. Değerler: "inbound", "outbound".
	LastDirection *string `json:"last_direction"`
	// Zorunlu.
	LastMessagePreview *string `json:"last_message_preview"`
	// Zorunlu.
	UnreadCount int64 `json:"unread_count"`
	// Zorunlu.
	Window *ConversationWindow `json:"window"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	AssigneeID *string `json:"assignee_id"`
	// Zorunlu.
	TeamID *string `json:"team_id"`
	// Zorunlu.
	AssignedAt *string `json:"assigned_at"`
	// Zorunlu.
	Language *string `json:"language"`
	// Zorunlu.
	BotActive bool `json:"bot_active"`
	// Zorunlu.
	LastSeq int64 `json:"last_seq"`
	// Zorunlu.
	VisitorLastReadSeq int64 `json:"visitor_last_read_seq"`
	// Zorunlu.
	VisitorLastDeliveredSeq int64 `json:"visitor_last_delivered_seq"`
	// Zorunlu.
	AgentLastReadSeq int64 `json:"agent_last_read_seq"`
	// Zorunlu.
	WaitingSince *string `json:"waiting_since"`
	// Zorunlu.
	FirstResponseAt *string `json:"first_response_at"`
	// Zorunlu.
	SlaDueAt *string `json:"sla_due_at"`
	// Zorunlu.
	ResolvedAt *string `json:"resolved_at"`
	// Zorunlu.
	Visitor *ConversationVisitor `json:"visitor"`
}

// ConversationActivity: ConversationActivity şeması.
type ConversationActivity struct {
	// Zorunlu.
	Data []ConversationActivityDataItem `json:"data"`
}

// ConversationActivityDataItem: ConversationActivityDataItem şeması.
type ConversationActivityDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Type string `json:"type"`
	// Zorunlu.
	ActorType string `json:"actor_type"`
	// Zorunlu.
	ActorID *string `json:"actor_id"`
	// Zorunlu.
	Data map[string]any `json:"data"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// ConversationAssignDto: ConversationAssignDto şeması.
type ConversationAssignDto struct {
	AssigneeID *string `json:"assignee_id,omitempty"`
	TeamID     *string `json:"team_id,omitempty"`
}

// ConversationMessagePage: ConversationMessagePage şeması.
type ConversationMessagePage struct {
	// Zorunlu.
	Data []ConversationMessagePageDataItem `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// ConversationMessagePageDataItem: ConversationMessagePageDataItem şeması.
type ConversationMessagePageDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "inbound", "outbound".
	Direction string `json:"direction"`
	// Zorunlu.
	Status string `json:"status"`
	// Zorunlu.
	Body *string `json:"body"`
	// Zorunlu.
	Content map[string]any `json:"content"`
	// Zorunlu.
	MessageID *string `json:"message_id"`
	// Zorunlu.
	DeliveryID *string `json:"delivery_id"`
	// Zorunlu.
	SentBy *string `json:"sent_by"`
	// Zorunlu.
	ErrorCode *string `json:"error_code"`
	// Zorunlu.
	ReadAt *string `json:"read_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	Seq int64 `json:"seq"`
	// Zorunlu.
	SenderType string `json:"sender_type"`
	// Zorunlu.
	ClientMessageID *string `json:"client_message_id"`
	// Zorunlu.
	Attachments []ConversationMessagePageDataItemAttachmentsItem `json:"attachments"`
	// Zorunlu.
	QuickReplies []ConversationMessagePageDataItemQuickRepliesItem `json:"quick_replies"`
}

// ConversationMessagePageDataItemAttachmentsItem: ConversationMessagePageDataItemAttachmentsItem şeması.
type ConversationMessagePageDataItemAttachmentsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Type string `json:"type"`
	// Zorunlu.
	Size int64 `json:"size"`
}

// ConversationMessagePageDataItemQuickRepliesItem: ConversationMessagePageDataItemQuickRepliesItem şeması.
type ConversationMessagePageDataItemQuickRepliesItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Label string `json:"label"`
}

// ConversationPage: ConversationPage şeması.
type ConversationPage struct {
	// Zorunlu.
	Data []Conversation `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// ConversationReplyDto: ConversationReplyDto şeması.
type ConversationReplyDto struct {
	Text            *string                          `json:"text,omitempty"`
	Template        *ConversationReplyDtoTemplate    `json:"template,omitempty"`
	Media           *ConversationReplyDtoMedia       `json:"media,omitempty"`
	Interactive     *ConversationReplyDtoInteractive `json:"interactive,omitempty"`
	ClientMessageID *string                          `json:"client_message_id,omitempty"`
	// Değerler: "manual", "canned", "ai_suggestion".
	Source *string `json:"source,omitempty"`
}

// ConversationReplyDtoInteractive: ConversationReplyDtoInteractive şeması.
type ConversationReplyDtoInteractive struct {
	// Zorunlu. Değerler: "button", "list".
	Type    string                                       `json:"type"`
	Header  *string                                      `json:"header,omitempty"`
	Footer  *string                                      `json:"footer,omitempty"`
	Buttons []ConversationReplyDtoInteractiveButtonsItem `json:"buttons,omitempty"`
	List    *ConversationReplyDtoInteractiveList         `json:"list,omitempty"`
}

// ConversationReplyDtoInteractiveButtonsItem: ConversationReplyDtoInteractiveButtonsItem şeması.
type ConversationReplyDtoInteractiveButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title string `json:"title"`
}

// ConversationReplyDtoInteractiveList: ConversationReplyDtoInteractiveList şeması.
type ConversationReplyDtoInteractiveList struct {
	// Zorunlu.
	Button string `json:"button"`
	// Zorunlu.
	Sections []ConversationReplyDtoInteractiveListSectionsItem `json:"sections"`
}

// ConversationReplyDtoInteractiveListSectionsItem: ConversationReplyDtoInteractiveListSectionsItem şeması.
type ConversationReplyDtoInteractiveListSectionsItem struct {
	Title *string `json:"title,omitempty"`
	// Zorunlu.
	Rows []ConversationReplyDtoInteractiveListSectionsItemRowsItem `json:"rows"`
}

// ConversationReplyDtoInteractiveListSectionsItemRowsItem: ConversationReplyDtoInteractiveListSectionsItemRowsItem şeması.
type ConversationReplyDtoInteractiveListSectionsItemRowsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

// ConversationReplyDtoMedia: ConversationReplyDtoMedia şeması.
type ConversationReplyDtoMedia struct {
	// Zorunlu. Değerler: "image", "video", "document", "audio".
	Type string `json:"type"`
	// Zorunlu.
	URL      string  `json:"url"`
	Filename *string `json:"filename,omitempty"`
}

// ConversationReplyDtoTemplate: ConversationReplyDtoTemplate şeması.
type ConversationReplyDtoTemplate struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Language     string                                         `json:"language"`
	Header       any                                            `json:"header,omitempty"`
	BodyParams   []string                                       `json:"body_params,omitempty"`
	ButtonParams []ConversationReplyDtoTemplateButtonParamsItem `json:"button_params,omitempty"`
}

// ConversationReplyDtoTemplateButtonParamsItem: ConversationReplyDtoTemplateButtonParamsItem şeması.
type ConversationReplyDtoTemplateButtonParamsItem struct {
	// Zorunlu.
	Index int64 `json:"index"`
	// Zorunlu. Değerler: "url", "quick_reply", "copy_code".
	SubType string `json:"sub_type"`
	// Zorunlu.
	Value string `json:"value"`
}

// ConversationReplyResult: ConversationReplyResult şeması.
type ConversationReplyResult struct {
	// Zorunlu.
	MessageID *string `json:"message_id"`
	// Zorunlu.
	ConversationMessage ConversationReplyResultConversationMessage `json:"conversation_message"`
}

// ConversationReplyResultConversationMessage: ConversationReplyResultConversationMessage şeması.
type ConversationReplyResultConversationMessage struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "inbound", "outbound".
	Direction string `json:"direction"`
	// Zorunlu.
	Status string `json:"status"`
	// Zorunlu.
	Body *string `json:"body"`
	// Zorunlu.
	Content map[string]any `json:"content"`
	// Zorunlu.
	MessageID *string `json:"message_id"`
	// Zorunlu.
	DeliveryID *string `json:"delivery_id"`
	// Zorunlu.
	SentBy *string `json:"sent_by"`
	// Zorunlu.
	ErrorCode *string `json:"error_code"`
	// Zorunlu.
	ReadAt *string `json:"read_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	Seq int64 `json:"seq"`
	// Zorunlu.
	SenderType string `json:"sender_type"`
	// Zorunlu.
	ClientMessageID *string `json:"client_message_id"`
	// Zorunlu.
	Attachments []ConversationReplyResultConversationMessageAttachmentsItem `json:"attachments"`
	// Zorunlu.
	QuickReplies []ConversationReplyResultConversationMessageQuickRepliesItem `json:"quick_replies"`
}

// ConversationReplyResultConversationMessageAttachmentsItem: ConversationReplyResultConversationMessageAttachmentsItem şeması.
type ConversationReplyResultConversationMessageAttachmentsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Type string `json:"type"`
	// Zorunlu.
	Size int64 `json:"size"`
}

// ConversationReplyResultConversationMessageQuickRepliesItem: ConversationReplyResultConversationMessageQuickRepliesItem şeması.
type ConversationReplyResultConversationMessageQuickRepliesItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Label string `json:"label"`
}

// ConversationTagsDto: ConversationTagsDto şeması.
type ConversationTagsDto struct {
	Add    map[string]string `json:"add,omitempty"`
	Remove []string          `json:"remove,omitempty"`
}

// ConversationTagsResult: ConversationTagsResult şeması.
type ConversationTagsResult struct {
	// Zorunlu.
	UserID string `json:"user_id"`
	// Zorunlu.
	Tags map[string]any `json:"tags"`
}

// ConversationUpdateDto: ConversationUpdateDto şeması.
type ConversationUpdateDto struct {
	// Zorunlu. Değerler: "open", "pending", "resolved", "closed".
	Status string `json:"status"`
}

// ConversationVisitor: ConversationVisitor şeması.
type ConversationVisitor struct {
	// Zorunlu.
	Name *string `json:"name"`
	// Zorunlu.
	Email *string `json:"email"`
}

// ConversationWindow: ConversationWindow şeması.
type ConversationWindow struct {
	// Zorunlu.
	Open bool `json:"open"`
	// Zorunlu.
	ExpiresAt *string `json:"expires_at"`
}

// ConversionsReport: ConversionsReport şeması.
type ConversionsReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	Totals ConversionsReportTotals `json:"totals"`
	// Zorunlu.
	Previous *ConversionsReportPrevious `json:"previous"`
	// Zorunlu.
	ByCampaign []ConversionsReportByCampaignItem `json:"by_campaign"`
	// Zorunlu.
	Rows []ConversionsReportRowsItem `json:"rows"`
}

// ConversionsReportByCampaignItem: ConversionsReportByCampaignItem şeması.
type ConversionsReportByCampaignItem struct {
	// Zorunlu.
	CampaignID *string `json:"campaign_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Conversions int64 `json:"conversions"`
	// Zorunlu.
	Direct int64 `json:"direct"`
	// Zorunlu.
	Influenced int64 `json:"influenced"`
	// Zorunlu.
	Unattributed int64 `json:"unattributed"`
	// Zorunlu.
	Revenue float64 `json:"revenue"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	ConversionRate *float64 `json:"conversion_rate"`
}

// ConversionsReportPrevious: ConversionsReportPrevious şeması.
type ConversionsReportPrevious struct {
	// Zorunlu.
	Conversions int64 `json:"conversions"`
	// Zorunlu.
	Direct int64 `json:"direct"`
	// Zorunlu.
	Influenced int64 `json:"influenced"`
	// Zorunlu.
	Unattributed int64 `json:"unattributed"`
	// Zorunlu.
	Revenue float64 `json:"revenue"`
}

// ConversionsReportRowsItem: ConversionsReportRowsItem şeması.
type ConversionsReportRowsItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	Conversions int64 `json:"conversions"`
	// Zorunlu.
	Revenue float64 `json:"revenue"`
}

// ConversionsReportTotals: ConversionsReportTotals şeması.
type ConversionsReportTotals struct {
	// Zorunlu.
	Conversions int64 `json:"conversions"`
	// Zorunlu.
	Direct int64 `json:"direct"`
	// Zorunlu.
	Influenced int64 `json:"influenced"`
	// Zorunlu.
	Unattributed int64 `json:"unattributed"`
	// Zorunlu.
	Revenue float64 `json:"revenue"`
}

// CopyTemplateDto: CopyTemplateDto şeması.
type CopyTemplateDto struct {
	// Zorunlu.
	TargetAppID string  `json:"target_app_id"`
	Name        *string `json:"name,omitempty"`
}

// CreateAdAudienceDto: CreateAdAudienceDto şeması.
type CreateAdAudienceDto struct {
	// Zorunlu.
	AdAccountID string `json:"ad_account_id"`
	// Zorunlu. Değerler: "segment", "journey".
	Source    string  `json:"source"`
	SegmentID *string `json:"segment_id,omitempty"`
	// Zorunlu.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Değerler: "email", "phone", "maid".
	Identifiers         []string `json:"identifiers,omitempty"`
	SyncIntervalMinutes *int64   `json:"sync_interval_minutes,omitempty"`
}

// CreateAppAPIKeyDto: CreateAppApiKeyDto şeması.
type CreateAppAPIKeyDto struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "admin", "notifications:write", "users:read", "users:write", "events:write", "segments:read", "analytics:read", "mcp:write".
	Scopes []string `json:"scopes"`
	// Değerler: "live", "test".
	Environment *string  `json:"environment,omitempty"`
	ExpiresAt   *string  `json:"expires_at,omitempty"`
	IPAllowlist []string `json:"ip_allowlist,omitempty"`
}

// CreateAppDto: CreateAppDto şeması.
type CreateAppDto struct {
	// Zorunlu.
	Name            string                     `json:"name"`
	Slug            *string                    `json:"slug,omitempty"`
	PlatformFlags   *CreateAppDtoPlatformFlags `json:"platform_flags,omitempty"`
	DefaultLanguage *string                    `json:"default_language,omitempty"`
	Timezone        *string                    `json:"timezone,omitempty"`
}

// CreateAppDtoPlatformFlags: CreateAppDtoPlatformFlags şeması.
type CreateAppDtoPlatformFlags struct {
	IOS     *bool `json:"ios,omitempty"`
	Android *bool `json:"android,omitempty"`
	Web     *bool `json:"web,omitempty"`
	Email   *bool `json:"email,omitempty"`
	SMS     *bool `json:"sms,omitempty"`
}

// CreateBudgetAlertRuleDto: CreateBudgetAlertRuleDto şeması.
type CreateBudgetAlertRuleDto struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "daily_spend", "roas_drop", "budget_depletion".
	Kind               string  `json:"kind"`
	AdAccountID        *string `json:"ad_account_id,omitempty"`
	CampaignExternalID *string `json:"campaign_external_id,omitempty"`
	// Zorunlu.
	Threshold  float64  `json:"threshold"`
	WindowDays *int64   `json:"window_days,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
	Enabled    *bool    `json:"enabled,omitempty"`
}

// CreateCategoryDto: CreateCategoryDto şeması.
type CreateCategoryDto struct {
	// Zorunlu.
	Key string `json:"key"`
	// Zorunlu.
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels       []string `json:"channels,omitempty"`
	DefaultOptedIn *bool    `json:"default_opted_in,omitempty"`
	Position       *int64   `json:"position,omitempty"`
}

// CreatedAPIKey: CreatedApiKey şeması.
type CreatedAPIKey struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	OrgID string `json:"org_id"`
	// Zorunlu.
	AppID *string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Prefix string `json:"prefix"`
	// Zorunlu.
	Scopes []string `json:"scopes"`
	// Zorunlu.
	LastUsedAt *string `json:"last_used_at"`
	// Zorunlu.
	ExpiresAt *string `json:"expires_at"`
	// Zorunlu.
	RevokedAt *string `json:"revoked_at"`
	// Zorunlu.
	IPAllowlist []string `json:"ip_allowlist"`
	// Zorunlu.
	RotatedAt *string `json:"rotated_at"`
	// Zorunlu.
	CreatedBy *string `json:"created_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	Key string `json:"key"`
}

// CreateDataFeedDto: CreateDataFeedDto şeması.
type CreateDataFeedDto struct {
	// Zorunlu.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Zorunlu.
	URLTemplate string `json:"url_template"`
	// Değerler: "GET", "POST".
	Method          *string           `json:"method,omitempty"`
	BodyTemplate    *string           `json:"body_template,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	SecretHeaders   map[string]string `json:"secret_headers,omitempty"`
	CacheTTLSeconds *int64            `json:"cache_ttl_seconds,omitempty"`
	TimeoutMs       *int64            `json:"timeout_ms,omitempty"`
	Enabled         *bool             `json:"enabled,omitempty"`
}

// CreateExportDto: CreateExportDto şeması.
type CreateExportDto struct {
	// Zorunlu. Değerler: "users", "subscriptions", "events", "deliveries", "message_events".
	Type    string                  `json:"type"`
	Filters *CreateExportDtoFilters `json:"filters,omitempty"`
}

// CreateExportDtoFilters: CreateExportDtoFilters şeması.
type CreateExportDtoFilters struct {
	From       *string `json:"from,omitempty"`
	To         *string `json:"to,omitempty"`
	Name       *string `json:"name,omitempty"`
	MessageID  *string `json:"message_id,omitempty"`
	CampaignID *string `json:"campaign_id,omitempty"`
	Status     *string `json:"status,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel   *string `json:"channel,omitempty"`
	SegmentID *string `json:"segment_id,omitempty"`
}

// CreateFormDto: CreateFormDto şeması.
type CreateFormDto struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Fields   []CreateFormDtoFieldsItem `json:"fields"`
	Settings *CreateFormDtoSettings    `json:"settings,omitempty"`
	// Değerler: "active", "archived".
	Status *string `json:"status,omitempty"`
}

// CreateFormDtoFieldsItem: CreateFormDtoFieldsItem şeması.
type CreateFormDtoFieldsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "email", "phone", "first_name", "last_name", "full_name", "text", "textarea", "select", "radio", "checkbox", "consent".
	Type string `json:"type"`
	// Zorunlu.
	Label       map[string]string                    `json:"label"`
	Placeholder map[string]string                    `json:"placeholder,omitempty"`
	Required    *bool                                `json:"required,omitempty"`
	Options     []CreateFormDtoFieldsItemOptionsItem `json:"options,omitempty"`
	MapsTo      *string                              `json:"maps_to,omitempty"`
	MaxLength   *int64                               `json:"max_length,omitempty"`
	PolicyURL   *string                              `json:"policy_url,omitempty"`
}

// CreateFormDtoFieldsItemOptionsItem: CreateFormDtoFieldsItemOptionsItem şeması.
type CreateFormDtoFieldsItemOptionsItem struct {
	// Zorunlu.
	Value string `json:"value"`
	// Zorunlu.
	Label map[string]string `json:"label"`
}

// CreateFormDtoSettings: CreateFormDtoSettings şeması.
type CreateFormDtoSettings struct {
	SubmitLabel     map[string]string                 `json:"submit_label,omitempty"`
	SuccessMessage  map[string]string                 `json:"success_message,omitempty"`
	RedirectURL     *string                           `json:"redirect_url,omitempty"`
	Tags            map[string]string                 `json:"tags,omitempty"`
	Subscribe       *CreateFormDtoSettingsSubscribe   `json:"subscribe,omitempty"`
	DoubleOptIn     *CreateFormDtoSettingsDoubleOptIn `json:"double_opt_in,omitempty"`
	EventName       *string                           `json:"event_name,omitempty"`
	Captcha         *bool                             `json:"captcha,omitempty"`
	DefaultLanguage *string                           `json:"default_language,omitempty"`
}

// CreateFormDtoSettingsDoubleOptIn: CreateFormDtoSettingsDoubleOptIn şeması.
type CreateFormDtoSettingsDoubleOptIn struct {
	// Zorunlu.
	Enabled bool              `json:"enabled"`
	Subject map[string]string `json:"subject,omitempty"`
	Body    map[string]string `json:"body,omitempty"`
}

// CreateFormDtoSettingsSubscribe: CreateFormDtoSettingsSubscribe şeması.
type CreateFormDtoSettingsSubscribe struct {
	Email *bool `json:"email,omitempty"`
	SMS   *bool `json:"sms,omitempty"`
}

// CreateIdentityKeyDto: CreateIdentityKeyDto şeması.
type CreateIdentityKeyDto struct {
	Kid          *string        `json:"kid,omitempty"`
	Label        *string        `json:"label,omitempty"`
	PublicJwk    map[string]any `json:"public_jwk,omitempty"`
	PublicKeyPem *string        `json:"public_key_pem,omitempty"`
}

// CreateInAppDto: CreateInAppDto şeması.
type CreateInAppDto struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Triggers      []any                        `json:"triggers"`
	TriggerGroups [][]any                      `json:"trigger_groups,omitempty"`
	Variants      []CreateInAppDtoVariantsItem `json:"variants,omitempty"`
	AbTest        *CreateInAppDtoAbTest        `json:"ab_test,omitempty"`
	// Zorunlu.
	Display     CreateInAppDtoDisplay    `json:"display"`
	Audience    *CreateInAppDtoAudience  `json:"audience,omitempty"`
	Frequency   *CreateInAppDtoFrequency `json:"frequency,omitempty"`
	JourneyOnly *bool                    `json:"journey_only,omitempty"`
	Priority    *int64                   `json:"priority,omitempty"`
	StartAt     *string                  `json:"start_at,omitempty"`
	EndAt       *string                  `json:"end_at,omitempty"`
}

// CreateInAppDtoAbTest: CreateInAppDtoAbTest şeması.
type CreateInAppDtoAbTest struct {
	Winner *CreateInAppDtoAbTestWinner `json:"winner,omitempty"`
}

// CreateInAppDtoAbTestWinner: CreateInAppDtoAbTestWinner şeması.
type CreateInAppDtoAbTestWinner struct {
	// Zorunlu. Değerler: "manual", "auto".
	Mode           string   `json:"mode"`
	MinImpressions *int64   `json:"min_impressions,omitempty"`
	Confidence     *float64 `json:"confidence,omitempty"`
}

// CreateInAppDtoAudience: CreateInAppDtoAudience şeması.
type CreateInAppDtoAudience struct {
	IncludedSegments []string `json:"included_segments,omitempty"`
	ExcludedSegments []string `json:"excluded_segments,omitempty"`
}

// CreateInAppDtoDisplay: CreateInAppDtoDisplay şeması.
type CreateInAppDtoDisplay struct {
	// Zorunlu. Değerler: "modal", "banner_top", "banner_bottom", "fullscreen".
	Layout             string  `json:"layout"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	// Değerler: "blocks", "html".
	ContentType *string           `json:"content_type,omitempty"`
	HTML        map[string]string `json:"html,omitempty"`
	HTMLHeight  *int64            `json:"html_height,omitempty"`
	// Zorunlu.
	Blocks []any `json:"blocks"`
}

// CreateInAppDtoFrequency: CreateInAppDtoFrequency şeması.
type CreateInAppDtoFrequency struct {
	MaxImpressions       *int64 `json:"max_impressions,omitempty"`
	MinIntervalSeconds   *int64 `json:"min_interval_seconds,omitempty"`
	StopAfterInteraction *bool  `json:"stop_after_interaction,omitempty"`
}

// CreateInAppDtoVariantsItem: CreateInAppDtoVariantsItem şeması.
type CreateInAppDtoVariantsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	Display CreateInAppDtoVariantsItemDisplay `json:"display"`
}

// CreateInAppDtoVariantsItemDisplay: CreateInAppDtoVariantsItemDisplay şeması.
type CreateInAppDtoVariantsItemDisplay struct {
	// Zorunlu. Değerler: "modal", "banner_top", "banner_bottom", "fullscreen".
	Layout             string  `json:"layout"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	// Değerler: "blocks", "html".
	ContentType *string           `json:"content_type,omitempty"`
	HTML        map[string]string `json:"html,omitempty"`
	HTMLHeight  *int64            `json:"html_height,omitempty"`
	// Zorunlu.
	Blocks []any `json:"blocks"`
}

// CreateKeywordDto: CreateKeywordDto şeması.
type CreateKeywordDto struct {
	// Zorunlu.
	Keyword string `json:"keyword"`
	// Değerler: "org", "app".
	Scope *string `json:"scope,omitempty"`
	// Değerler: "sms", "whatsapp", "rcs".
	Channels []string `json:"channels,omitempty"`
	// Zorunlu.
	Actions     CreateKeywordDtoActions `json:"actions"`
	Enabled     *bool                   `json:"enabled,omitempty"`
	Description *string                 `json:"description,omitempty"`
}

// CreateKeywordDtoActions: CreateKeywordDtoActions şeması.
type CreateKeywordDtoActions struct {
	AddTags    map[string]string             `json:"add_tags,omitempty"`
	RemoveTags []string                      `json:"remove_tags,omitempty"`
	Event      *CreateKeywordDtoActionsEvent `json:"event,omitempty"`
	Reply      *CreateKeywordDtoActionsReply `json:"reply,omitempty"`
	OptIn      *bool                         `json:"opt_in,omitempty"`
}

// CreateKeywordDtoActionsEvent: CreateKeywordDtoActionsEvent şeması.
type CreateKeywordDtoActionsEvent struct {
	// Zorunlu.
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties,omitempty"`
}

// CreateKeywordDtoActionsReply: CreateKeywordDtoActionsReply şeması.
type CreateKeywordDtoActionsReply struct {
	// Zorunlu.
	Text string `json:"text"`
}

// CreateLiveActivityTypeDto: CreateLiveActivityTypeDto şeması.
type CreateLiveActivityTypeDto struct {
	// Zorunlu.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Zorunlu.
	ContentStateSchema CreateLiveActivityTypeDtoContentStateSchema `json:"content_state_schema"`
	AttributesSchema   *CreateLiveActivityTypeDtoAttributesSchema  `json:"attributes_schema,omitempty"`
}

// CreateLiveActivityTypeDtoAttributesSchema: CreateLiveActivityTypeDtoAttributesSchema şeması.
type CreateLiveActivityTypeDtoAttributesSchema struct {
	// Zorunlu. Değerler: "object".
	Type string `json:"type"`
	// Zorunlu.
	Properties map[string]CreateLiveActivityTypeDtoAttributesSchemaPropertiesValue `json:"properties"`
	Required   []string                                                            `json:"required,omitempty"`
}

// CreateLiveActivityTypeDtoAttributesSchemaPropertiesValue: CreateLiveActivityTypeDtoAttributesSchemaPropertiesValue şeması.
type CreateLiveActivityTypeDtoAttributesSchemaPropertiesValue struct {
	// Zorunlu. Değerler: "string", "number", "integer", "boolean".
	Type        string   `json:"type"`
	Enum        []any    `json:"enum,omitempty"`
	MaxLength   *int64   `json:"maxLength,omitempty"`
	Minimum     *float64 `json:"minimum,omitempty"`
	Maximum     *float64 `json:"maximum,omitempty"`
	Description *string  `json:"description,omitempty"`
}

// CreateLiveActivityTypeDtoContentStateSchema: CreateLiveActivityTypeDtoContentStateSchema şeması.
type CreateLiveActivityTypeDtoContentStateSchema struct {
	// Zorunlu. Değerler: "object".
	Type string `json:"type"`
	// Zorunlu.
	Properties map[string]CreateLiveActivityTypeDtoContentStateSchemaPropertiesValue `json:"properties"`
	Required   []string                                                              `json:"required,omitempty"`
}

// CreateLiveActivityTypeDtoContentStateSchemaPropertiesValue: CreateLiveActivityTypeDtoContentStateSchemaPropertiesValue şeması.
type CreateLiveActivityTypeDtoContentStateSchemaPropertiesValue struct {
	// Zorunlu. Değerler: "string", "number", "integer", "boolean".
	Type        string   `json:"type"`
	Enum        []any    `json:"enum,omitempty"`
	MaxLength   *int64   `json:"maxLength,omitempty"`
	Minimum     *float64 `json:"minimum,omitempty"`
	Maximum     *float64 `json:"maximum,omitempty"`
	Description *string  `json:"description,omitempty"`
}

// CreateLookalikeDto: CreateLookalikeDto şeması.
type CreateLookalikeDto struct {
	Name *string `json:"name,omitempty"`
	// Zorunlu.
	Country string `json:"country"`
	// Zorunlu.
	Ratio float64 `json:"ratio"`
}

// CreateNotificationDto: CreateNotificationDto şeması.
type CreateNotificationDto struct {
	AppID          *string `json:"app_id,omitempty"`
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	// Zorunlu.
	Target         CreateNotificationDtoTarget        `json:"target"`
	Headings       map[string]string                  `json:"headings,omitempty"`
	Contents       map[string]string                  `json:"contents,omitempty"`
	Subtitle       map[string]string                  `json:"subtitle,omitempty"`
	URL            *string                            `json:"url,omitempty"`
	AppURL         *string                            `json:"app_url,omitempty"`
	Data           map[string]string                  `json:"data,omitempty"`
	Buttons        []CreateNotificationDtoButtonsItem `json:"buttons,omitempty"`
	BigPicture     *string                            `json:"big_picture,omitempty"`
	IOSAttachments map[string]string                  `json:"ios_attachments,omitempty"`
	IOSSound       *string                            `json:"ios_sound,omitempty"`
	// Değerler: "None", "SetTo", "Increase".
	IOSBadgeType  *string `json:"ios_badge_type,omitempty"`
	IOSBadgeCount *int64  `json:"ios_badge_count,omitempty"`
	// Değerler: "passive", "active", "time-sensitive", "critical".
	IOSInterruptionLevel    *string  `json:"ios_interruption_level,omitempty"`
	IOSCategory             *string  `json:"ios_category,omitempty"`
	ThreadID                *string  `json:"thread_id,omitempty"`
	IOSRelevanceScore       *float64 `json:"ios_relevance_score,omitempty"`
	TargetContentIdentifier *string  `json:"target_content_identifier,omitempty"`
	AndroidChannelID        *string  `json:"android_channel_id,omitempty"`
	AndroidGroup            *string  `json:"android_group,omitempty"`
	AndroidSound            *string  `json:"android_sound,omitempty"`
	SmallIcon               *string  `json:"small_icon,omitempty"`
	LargeIcon               *string  `json:"large_icon,omitempty"`
	AndroidAccentColor      *string  `json:"android_accent_color,omitempty"`
	// Değerler: 1, 0, -1.
	AndroidVisibility  *int64                                `json:"android_visibility,omitempty"`
	ChromeWebIcon      *string                               `json:"chrome_web_icon,omitempty"`
	ChromeWebImage     *string                               `json:"chrome_web_image,omitempty"`
	ChromeWebBadge     *string                               `json:"chrome_web_badge,omitempty"`
	WebPushTopic       *string                               `json:"web_push_topic,omitempty"`
	WebButtons         []CreateNotificationDtoWebButtonsItem `json:"web_buttons,omitempty"`
	RequireInteraction *bool                                 `json:"require_interaction,omitempty"`
	ContentAvailable   *bool                                 `json:"content_available,omitempty"`
	MutableContent     *bool                                 `json:"mutable_content,omitempty"`
	CollapseID         *string                               `json:"collapse_id,omitempty"`
	// Değerler: 5, 10.
	Priority  *int64  `json:"priority,omitempty"`
	TTL       *int64  `json:"ttl,omitempty"`
	SendAfter *string `json:"send_after,omitempty"`
	// Değerler: "timezone", "optimized".
	DelayedOption             *string                                   `json:"delayed_option,omitempty"`
	DeliveryTimeOfDay         *string                                   `json:"delivery_time_of_day,omitempty"`
	TemplateID                *string                                   `json:"template_id,omitempty"`
	CustomData                map[string]any                            `json:"custom_data,omitempty"`
	EmailSubject              *string                                   `json:"email_subject,omitempty"`
	EmailBody                 *string                                   `json:"email_body,omitempty"`
	EmailMJML                 *string                                   `json:"email_mjml,omitempty"`
	EmailPreheader            *string                                   `json:"email_preheader,omitempty"`
	EmailFromName             *string                                   `json:"email_from_name,omitempty"`
	EmailFromAddress          *string                                   `json:"email_from_address,omitempty"`
	EmailReplyToAddress       *string                                   `json:"email_reply_to_address,omitempty"`
	DisableEmailClickTracking *bool                                     `json:"disable_email_click_tracking,omitempty"`
	SMSFrom                   *string                                   `json:"sms_from,omitempty"`
	SMSMediaUrls              []string                                  `json:"sms_media_urls,omitempty"`
	DisableSMSLinkShortening  *bool                                     `json:"disable_sms_link_shortening,omitempty"`
	WhatsappTemplate          *CreateNotificationDtoWhatsappTemplate    `json:"whatsapp_template,omitempty"`
	WhatsappMedia             *CreateNotificationDtoWhatsappMedia       `json:"whatsapp_media,omitempty"`
	WhatsappInteractive       *CreateNotificationDtoWhatsappInteractive `json:"whatsapp_interactive,omitempty"`
	RcsCard                   *CreateNotificationDtoRcsCard             `json:"rcs_card,omitempty"`
	RcsSuggestions            []any                                     `json:"rcs_suggestions,omitempty"`
	RcsContentSid             *string                                   `json:"rcs_content_sid,omitempty"`
	RcsContentVariables       map[string]string                         `json:"rcs_content_variables,omitempty"`
	RcsFallbackText           map[string]string                         `json:"rcs_fallback_text,omitempty"`
	AddToInbox                *bool                                     `json:"add_to_inbox,omitempty"`
	InboxTTL                  *int64                                    `json:"inbox_ttl,omitempty"`
	FrequencyCapExempt        *bool                                     `json:"frequency_cap_exempt,omitempty"`
	QuietHoursExempt          *bool                                     `json:"quiet_hours_exempt,omitempty"`
	OnePerUser                *bool                                     `json:"one_per_user,omitempty"`
	ThrottleRatePerMinute     *int64                                    `json:"throttle_rate_per_minute,omitempty"`
	IncludeUnsubscribed       *bool                                     `json:"include_unsubscribed,omitempty"`
	Name                      *string                                   `json:"name,omitempty"`
	Labels                    []string                                  `json:"labels,omitempty"`
	Category                  *string                                   `json:"category,omitempty"`
}

// CreateNotificationDtoButtonsItem: CreateNotificationDtoButtonsItem şeması.
type CreateNotificationDtoButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string  `json:"text"`
	Icon *string `json:"icon,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// CreateNotificationDtoRcsCard: CreateNotificationDtoRcsCard şeması.
type CreateNotificationDtoRcsCard struct {
	Title       map[string]string `json:"title,omitempty"`
	Description map[string]string `json:"description,omitempty"`
	MediaURL    *string           `json:"media_url,omitempty"`
	// Değerler: "SHORT", "MEDIUM", "TALL".
	MediaHeight *string `json:"media_height,omitempty"`
	// Değerler: "VERTICAL", "HORIZONTAL".
	Orientation *string `json:"orientation,omitempty"`
}

// CreateNotificationDtoTarget: CreateNotificationDtoTarget şeması.
type CreateNotificationDtoTarget struct {
	All                    *bool               `json:"all,omitempty"`
	IncludeSubscriptionIds []string            `json:"include_subscription_ids,omitempty"`
	IncludeExternalIds     []string            `json:"include_external_ids,omitempty"`
	IncludeAliases         map[string][]string `json:"include_aliases,omitempty"`
	IncludedSegments       []string            `json:"included_segments,omitempty"`
	ExcludedSegments       []string            `json:"excluded_segments,omitempty"`
	Filters                any                 `json:"filters,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels            []string `json:"channels,omitempty"`
	TestDevicesOnly     *bool    `json:"test_devices_only,omitempty"`
	EmailTo             []string `json:"email_to,omitempty"`
	IncludePhoneNumbers []string `json:"include_phone_numbers,omitempty"`
}

// CreateNotificationDtoWebButtonsItem: CreateNotificationDtoWebButtonsItem şeması.
type CreateNotificationDtoWebButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string  `json:"text"`
	Icon *string `json:"icon,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// CreateNotificationDtoWhatsappInteractive: CreateNotificationDtoWhatsappInteractive şeması.
type CreateNotificationDtoWhatsappInteractive struct {
	// Zorunlu. Değerler: "button", "list".
	Type    string                                                `json:"type"`
	Header  *string                                               `json:"header,omitempty"`
	Footer  *string                                               `json:"footer,omitempty"`
	Buttons []CreateNotificationDtoWhatsappInteractiveButtonsItem `json:"buttons,omitempty"`
	List    *CreateNotificationDtoWhatsappInteractiveList         `json:"list,omitempty"`
}

// CreateNotificationDtoWhatsappInteractiveButtonsItem: CreateNotificationDtoWhatsappInteractiveButtonsItem şeması.
type CreateNotificationDtoWhatsappInteractiveButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title string `json:"title"`
}

// CreateNotificationDtoWhatsappInteractiveList: CreateNotificationDtoWhatsappInteractiveList şeması.
type CreateNotificationDtoWhatsappInteractiveList struct {
	// Zorunlu.
	Button string `json:"button"`
	// Zorunlu.
	Sections []CreateNotificationDtoWhatsappInteractiveListSectionsItem `json:"sections"`
}

// CreateNotificationDtoWhatsappInteractiveListSectionsItem: CreateNotificationDtoWhatsappInteractiveListSectionsItem şeması.
type CreateNotificationDtoWhatsappInteractiveListSectionsItem struct {
	Title *string `json:"title,omitempty"`
	// Zorunlu.
	Rows []CreateNotificationDtoWhatsappInteractiveListSectionsItemRowsItem `json:"rows"`
}

// CreateNotificationDtoWhatsappInteractiveListSectionsItemRowsItem: CreateNotificationDtoWhatsappInteractiveListSectionsItemRowsItem şeması.
type CreateNotificationDtoWhatsappInteractiveListSectionsItemRowsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

// CreateNotificationDtoWhatsappMedia: CreateNotificationDtoWhatsappMedia şeması.
type CreateNotificationDtoWhatsappMedia struct {
	// Zorunlu. Değerler: "image", "video", "document", "audio".
	Type string `json:"type"`
	// Zorunlu.
	URL      string  `json:"url"`
	Filename *string `json:"filename,omitempty"`
}

// CreateNotificationDtoWhatsappTemplate: CreateNotificationDtoWhatsappTemplate şeması.
type CreateNotificationDtoWhatsappTemplate struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Language     string                                                  `json:"language"`
	Header       any                                                     `json:"header,omitempty"`
	BodyParams   []string                                                `json:"body_params,omitempty"`
	ButtonParams []CreateNotificationDtoWhatsappTemplateButtonParamsItem `json:"button_params,omitempty"`
}

// CreateNotificationDtoWhatsappTemplateButtonParamsItem: CreateNotificationDtoWhatsappTemplateButtonParamsItem şeması.
type CreateNotificationDtoWhatsappTemplateButtonParamsItem struct {
	// Zorunlu.
	Index int64 `json:"index"`
	// Zorunlu. Değerler: "url", "quick_reply", "copy_code".
	SubType string `json:"sub_type"`
	// Zorunlu.
	Value string `json:"value"`
}

// CreateNotificationResult: CreateNotificationResult şeması.
type CreateNotificationResult struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "queued", "processing", "completed", "cancelled", "failed".
	Status string `json:"status"`
	// Zorunlu.
	RecipientsEstimate int64 `json:"recipients_estimate"`
	// Zorunlu.
	ExternalIDMissing []string `json:"external_id_missing"`
}

// CreateSuppressionDto: CreateSuppressionDto şeması.
type CreateSuppressionDto struct {
	// Zorunlu. Değerler: "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu.
	Value string `json:"value"`
	// Değerler: "manual", "bounce", "complaint", "unsubscribe_link", "stop".
	Reason    *string `json:"reason,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	Note      *string `json:"note,omitempty"`
}

// CreateSurveyDto: CreateSurveyDto şeması.
type CreateSurveyDto struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Questions []CreateSurveyDtoQuestionsItem `json:"questions"`
	Settings  *CreateSurveyDtoSettings       `json:"settings,omitempty"`
}

// CreateSurveyDtoQuestionsItem: CreateSurveyDtoQuestionsItem şeması.
type CreateSurveyDtoQuestionsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "nps", "csat", "single_choice", "multiple_choice", "text".
	Type string `json:"type"`
	// Zorunlu.
	Text        map[string]string                         `json:"text"`
	Required    *bool                                     `json:"required,omitempty"`
	Options     []CreateSurveyDtoQuestionsItemOptionsItem `json:"options,omitempty"`
	LowLabel    map[string]string                         `json:"low_label,omitempty"`
	HighLabel   map[string]string                         `json:"high_label,omitempty"`
	Placeholder map[string]string                         `json:"placeholder,omitempty"`
}

// CreateSurveyDtoQuestionsItemOptionsItem: CreateSurveyDtoQuestionsItemOptionsItem şeması.
type CreateSurveyDtoQuestionsItemOptionsItem struct {
	// Zorunlu.
	Value string `json:"value"`
	// Zorunlu.
	Label map[string]string `json:"label"`
}

// CreateSurveyDtoSettings: CreateSurveyDtoSettings şeması.
type CreateSurveyDtoSettings struct {
	ThankYou        map[string]string `json:"thank_you,omitempty"`
	SubmitLabel     map[string]string `json:"submit_label,omitempty"`
	WriteProperties *bool             `json:"write_properties,omitempty"`
	AllowMultiple   *bool             `json:"allow_multiple,omitempty"`
	DefaultLanguage *string           `json:"default_language,omitempty"`
}

// CreateTemplateDto: CreateTemplateDto şeması.
type CreateTemplateDto struct {
	// Zorunlu.
	Name string `json:"name"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel *string  `json:"channel,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	// Zorunlu.
	Content CreateTemplateDtoContent `json:"content"`
}

// CreateTemplateDtoContent: CreateTemplateDtoContent şeması.
type CreateTemplateDtoContent struct {
	Headings       map[string]string                     `json:"headings,omitempty"`
	Contents       map[string]string                     `json:"contents,omitempty"`
	Subtitle       map[string]string                     `json:"subtitle,omitempty"`
	URL            *string                               `json:"url,omitempty"`
	AppURL         *string                               `json:"app_url,omitempty"`
	Data           map[string]string                     `json:"data,omitempty"`
	Buttons        []CreateTemplateDtoContentButtonsItem `json:"buttons,omitempty"`
	BigPicture     *string                               `json:"big_picture,omitempty"`
	IOSAttachments map[string]string                     `json:"ios_attachments,omitempty"`
	IOSSound       *string                               `json:"ios_sound,omitempty"`
	// Değerler: "None", "SetTo", "Increase".
	IOSBadgeType  *string `json:"ios_badge_type,omitempty"`
	IOSBadgeCount *int64  `json:"ios_badge_count,omitempty"`
	// Değerler: "passive", "active", "time-sensitive", "critical".
	IOSInterruptionLevel      *string        `json:"ios_interruption_level,omitempty"`
	IOSCategory               *string        `json:"ios_category,omitempty"`
	AndroidChannelID          *string        `json:"android_channel_id,omitempty"`
	AndroidGroup              *string        `json:"android_group,omitempty"`
	AndroidSound              *string        `json:"android_sound,omitempty"`
	ChromeWebIcon             *string        `json:"chrome_web_icon,omitempty"`
	ChromeWebImage            *string        `json:"chrome_web_image,omitempty"`
	ChromeWebBadge            *string        `json:"chrome_web_badge,omitempty"`
	WebPushTopic              *string        `json:"web_push_topic,omitempty"`
	ContentAvailable          *bool          `json:"content_available,omitempty"`
	MutableContent            *bool          `json:"mutable_content,omitempty"`
	CustomData                map[string]any `json:"custom_data,omitempty"`
	EmailSubject              *string        `json:"email_subject,omitempty"`
	EmailBody                 *string        `json:"email_body,omitempty"`
	EmailMJML                 *string        `json:"email_mjml,omitempty"`
	EmailPreheader            *string        `json:"email_preheader,omitempty"`
	EmailFromName             *string        `json:"email_from_name,omitempty"`
	EmailFromAddress          *string        `json:"email_from_address,omitempty"`
	EmailReplyToAddress       *string        `json:"email_reply_to_address,omitempty"`
	DisableEmailClickTracking *bool          `json:"disable_email_click_tracking,omitempty"`
	SMSFrom                   *string        `json:"sms_from,omitempty"`
	DisableSMSLinkShortening  *bool          `json:"disable_sms_link_shortening,omitempty"`
	AddToInbox                *bool          `json:"add_to_inbox,omitempty"`
	InboxTTL                  *int64         `json:"inbox_ttl,omitempty"`
	ThreadID                  *string        `json:"thread_id,omitempty"`
	IOSRelevanceScore         *float64       `json:"ios_relevance_score,omitempty"`
	TargetContentIdentifier   *string        `json:"target_content_identifier,omitempty"`
	SmallIcon                 *string        `json:"small_icon,omitempty"`
	LargeIcon                 *string        `json:"large_icon,omitempty"`
	AndroidAccentColor        *string        `json:"android_accent_color,omitempty"`
	// Değerler: 1, 0, -1.
	AndroidVisibility   *int64                                       `json:"android_visibility,omitempty"`
	WebButtons          []CreateTemplateDtoContentWebButtonsItem     `json:"web_buttons,omitempty"`
	RequireInteraction  *bool                                        `json:"require_interaction,omitempty"`
	SMSMediaUrls        []string                                     `json:"sms_media_urls,omitempty"`
	WhatsappTemplate    *CreateTemplateDtoContentWhatsappTemplate    `json:"whatsapp_template,omitempty"`
	WhatsappMedia       *CreateTemplateDtoContentWhatsappMedia       `json:"whatsapp_media,omitempty"`
	WhatsappInteractive *CreateTemplateDtoContentWhatsappInteractive `json:"whatsapp_interactive,omitempty"`
	RcsCard             *CreateTemplateDtoContentRcsCard             `json:"rcs_card,omitempty"`
	RcsSuggestions      []any                                        `json:"rcs_suggestions,omitempty"`
	RcsContentSid       *string                                      `json:"rcs_content_sid,omitempty"`
	RcsContentVariables map[string]string                            `json:"rcs_content_variables,omitempty"`
	RcsFallbackText     map[string]string                            `json:"rcs_fallback_text,omitempty"`
}

// CreateTemplateDtoContentButtonsItem: CreateTemplateDtoContentButtonsItem şeması.
type CreateTemplateDtoContentButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string  `json:"text"`
	Icon *string `json:"icon,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// CreateTemplateDtoContentRcsCard: CreateTemplateDtoContentRcsCard şeması.
type CreateTemplateDtoContentRcsCard struct {
	Title       map[string]string `json:"title,omitempty"`
	Description map[string]string `json:"description,omitempty"`
	MediaURL    *string           `json:"media_url,omitempty"`
	// Değerler: "SHORT", "MEDIUM", "TALL".
	MediaHeight *string `json:"media_height,omitempty"`
	// Değerler: "VERTICAL", "HORIZONTAL".
	Orientation *string `json:"orientation,omitempty"`
}

// CreateTemplateDtoContentWebButtonsItem: CreateTemplateDtoContentWebButtonsItem şeması.
type CreateTemplateDtoContentWebButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string  `json:"text"`
	Icon *string `json:"icon,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// CreateTemplateDtoContentWhatsappInteractive: CreateTemplateDtoContentWhatsappInteractive şeması.
type CreateTemplateDtoContentWhatsappInteractive struct {
	// Zorunlu. Değerler: "button", "list".
	Type    string                                                   `json:"type"`
	Header  *string                                                  `json:"header,omitempty"`
	Footer  *string                                                  `json:"footer,omitempty"`
	Buttons []CreateTemplateDtoContentWhatsappInteractiveButtonsItem `json:"buttons,omitempty"`
	List    *CreateTemplateDtoContentWhatsappInteractiveList         `json:"list,omitempty"`
}

// CreateTemplateDtoContentWhatsappInteractiveButtonsItem: CreateTemplateDtoContentWhatsappInteractiveButtonsItem şeması.
type CreateTemplateDtoContentWhatsappInteractiveButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title string `json:"title"`
}

// CreateTemplateDtoContentWhatsappInteractiveList: CreateTemplateDtoContentWhatsappInteractiveList şeması.
type CreateTemplateDtoContentWhatsappInteractiveList struct {
	// Zorunlu.
	Button string `json:"button"`
	// Zorunlu.
	Sections []CreateTemplateDtoContentWhatsappInteractiveListSectionsItem `json:"sections"`
}

// CreateTemplateDtoContentWhatsappInteractiveListSectionsItem: CreateTemplateDtoContentWhatsappInteractiveListSectionsItem şeması.
type CreateTemplateDtoContentWhatsappInteractiveListSectionsItem struct {
	Title *string `json:"title,omitempty"`
	// Zorunlu.
	Rows []CreateTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem `json:"rows"`
}

// CreateTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem: CreateTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem şeması.
type CreateTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

// CreateTemplateDtoContentWhatsappMedia: CreateTemplateDtoContentWhatsappMedia şeması.
type CreateTemplateDtoContentWhatsappMedia struct {
	// Zorunlu. Değerler: "image", "video", "document", "audio".
	Type string `json:"type"`
	// Zorunlu.
	URL      string  `json:"url"`
	Filename *string `json:"filename,omitempty"`
}

// CreateTemplateDtoContentWhatsappTemplate: CreateTemplateDtoContentWhatsappTemplate şeması.
type CreateTemplateDtoContentWhatsappTemplate struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Language     string                                                     `json:"language"`
	Header       any                                                        `json:"header,omitempty"`
	BodyParams   []string                                                   `json:"body_params,omitempty"`
	ButtonParams []CreateTemplateDtoContentWhatsappTemplateButtonParamsItem `json:"button_params,omitempty"`
}

// CreateTemplateDtoContentWhatsappTemplateButtonParamsItem: CreateTemplateDtoContentWhatsappTemplateButtonParamsItem şeması.
type CreateTemplateDtoContentWhatsappTemplateButtonParamsItem struct {
	// Zorunlu.
	Index int64 `json:"index"`
	// Zorunlu. Değerler: "url", "quick_reply", "copy_code".
	SubType string `json:"sub_type"`
	// Zorunlu.
	Value string `json:"value"`
}

// CreateWarmupPlanDto: CreateWarmupPlanDto şeması.
type CreateWarmupPlanDto struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "domain", "ip".
	Kind string `json:"kind"`
	// Zorunlu.
	Target         string                             `json:"target"`
	StartDate      *string                            `json:"start_date,omitempty"`
	DailyLimits    []int64                            `json:"daily_limits,omitempty"`
	Curve          *CreateWarmupPlanDtoCurve          `json:"curve,omitempty"`
	ProviderShares *CreateWarmupPlanDtoProviderShares `json:"provider_shares,omitempty"`
	Timezone       *string                            `json:"timezone,omitempty"`
}

// CreateWarmupPlanDtoCurve: CreateWarmupPlanDtoCurve şeması.
type CreateWarmupPlanDtoCurve struct {
	// Zorunlu.
	Start int64 `json:"start"`
	// Zorunlu.
	Growth float64 `json:"growth"`
	// Zorunlu.
	Target int64 `json:"target"`
}

// CreateWarmupPlanDtoProviderShares: CreateWarmupPlanDtoProviderShares şeması.
type CreateWarmupPlanDtoProviderShares struct {
	// Zorunlu.
	Gmail float64 `json:"gmail"`
	// Zorunlu.
	Outlook float64 `json:"outlook"`
	// Zorunlu.
	Yahoo float64 `json:"yahoo"`
	// Zorunlu.
	Other float64 `json:"other"`
}

// CreateWebCampaignDto: CreateWebCampaignDto şeması.
type CreateWebCampaignDto struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Display  CreateWebCampaignDtoDisplay        `json:"display"`
	Variants []CreateWebCampaignDtoVariantsItem `json:"variants,omitempty"`
	// Zorunlu.
	Trigger    any                             `json:"trigger"`
	Conditions *CreateWebCampaignDtoConditions `json:"conditions,omitempty"`
	Targeting  *CreateWebCampaignDtoTargeting  `json:"targeting,omitempty"`
	Frequency  *CreateWebCampaignDtoFrequency  `json:"frequency,omitempty"`
	Priority   *int64                          `json:"priority,omitempty"`
	StartAt    *string                         `json:"start_at,omitempty"`
	EndAt      *string                         `json:"end_at,omitempty"`
}

// CreateWebCampaignDtoConditions: CreateWebCampaignDtoConditions şeması.
type CreateWebCampaignDtoConditions struct {
	URLRules []CreateWebCampaignDtoConditionsURLRulesItem `json:"url_rules,omitempty"`
	UtmRules []CreateWebCampaignDtoConditionsUtmRulesItem `json:"utm_rules,omitempty"`
	Visits   *CreateWebCampaignDtoConditionsVisits        `json:"visits,omitempty"`
}

// CreateWebCampaignDtoConditionsURLRulesItem: CreateWebCampaignDtoConditionsUrlRulesItem şeması.
type CreateWebCampaignDtoConditionsURLRulesItem struct {
	// Zorunlu. Değerler: "equals", "contains", "starts_with", "not_contains", "matches".
	Op string `json:"op"`
	// Zorunlu.
	Value string `json:"value"`
}

// CreateWebCampaignDtoConditionsUtmRulesItem: CreateWebCampaignDtoConditionsUtmRulesItem şeması.
type CreateWebCampaignDtoConditionsUtmRulesItem struct {
	// Zorunlu. Değerler: "utm_source", "utm_medium", "utm_campaign", "utm_term", "utm_content".
	Param string `json:"param"`
	// Zorunlu. Değerler: "equals", "contains", "exists".
	Op    string  `json:"op"`
	Value *string `json:"value,omitempty"`
}

// CreateWebCampaignDtoConditionsVisits: CreateWebCampaignDtoConditionsVisits şeması.
type CreateWebCampaignDtoConditionsVisits struct {
	// Zorunlu. Değerler: "gte", "lte", "eq".
	Op string `json:"op"`
	// Zorunlu.
	Value int64 `json:"value"`
}

// CreateWebCampaignDtoDisplay: CreateWebCampaignDtoDisplay şeması.
type CreateWebCampaignDtoDisplay struct {
	// Zorunlu. Değerler: "popup", "banner_top", "banner_bottom", "slide_in", "fullscreen".
	Layout string `json:"layout"`
	// Değerler: "left", "right".
	Position           *string `json:"position,omitempty"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	MaxWidth           *int64  `json:"max_width,omitempty"`
	// Zorunlu.
	Blocks   []any   `json:"blocks"`
	FormID   *string `json:"form_id,omitempty"`
	SurveyID *string `json:"survey_id,omitempty"`
}

// CreateWebCampaignDtoFrequency: CreateWebCampaignDtoFrequency şeması.
type CreateWebCampaignDtoFrequency struct {
	MaxPerSession       *int64 `json:"max_per_session,omitempty"`
	MaxPerDay           *int64 `json:"max_per_day,omitempty"`
	MaxTotal            *int64 `json:"max_total,omitempty"`
	HideAfterCloseDays  *int64 `json:"hide_after_close_days,omitempty"`
	StopAfterConversion *bool  `json:"stop_after_conversion,omitempty"`
}

// CreateWebCampaignDtoTargeting: CreateWebCampaignDtoTargeting şeması.
type CreateWebCampaignDtoTargeting struct {
	IncludedSegments []string `json:"included_segments,omitempty"`
	ExcludedSegments []string `json:"excluded_segments,omitempty"`
	// Değerler: "all", "new", "returning".
	Visitor *string `json:"visitor,omitempty"`
	// Değerler: "desktop", "mobile", "tablet".
	Devices []string `json:"devices,omitempty"`
}

// CreateWebCampaignDtoVariantsItem: CreateWebCampaignDtoVariantsItem şeması.
type CreateWebCampaignDtoVariantsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	Display CreateWebCampaignDtoVariantsItemDisplay `json:"display"`
}

// CreateWebCampaignDtoVariantsItemDisplay: CreateWebCampaignDtoVariantsItemDisplay şeması.
type CreateWebCampaignDtoVariantsItemDisplay struct {
	// Zorunlu. Değerler: "popup", "banner_top", "banner_bottom", "slide_in", "fullscreen".
	Layout string `json:"layout"`
	// Değerler: "left", "right".
	Position           *string `json:"position,omitempty"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	MaxWidth           *int64  `json:"max_width,omitempty"`
	// Zorunlu.
	Blocks   []any   `json:"blocks"`
	FormID   *string `json:"form_id,omitempty"`
	SurveyID *string `json:"survey_id,omitempty"`
}

// CreateWebhookDto: CreateWebhookDto şeması.
type CreateWebhookDto struct {
	// Zorunlu.
	URL         string  `json:"url"`
	Description *string `json:"description,omitempty"`
	// Zorunlu. Değerler: "message.queued", "message.completed", "message.failed", "delivery.sent", "delivery.delivered", "delivery.failed", "delivery.opened", "delivery.clicked", "delivery.bounced", "delivery.complained", "delivery.unsubscribed", "in_app.impression", "in_app.clicked", "in_app.dismissed", "subscription.created", "subscription.unsubscribed", "subscription.invalidated", "user.created", "user.updated", "user.merged", "user.deleted", "event.received", "conversion.created", "campaign.scheduled", "campaign.started", "campaign.completed", "journey.entered", "journey.step_completed", "journey.completed", "journey.exited", "conversation.created", "conversation.message", "conversation.assigned", "conversation.resolved", "audit.created", "form.submitted", "survey.responded", "ad_audience.synced", "ad_account.revoked", "ad_budget.alert", "ad_insights.imported".
	Events       []string                      `json:"events"`
	EventFilters *CreateWebhookDtoEventFilters `json:"event_filters,omitempty"`
	Batch        *bool                         `json:"batch,omitempty"`
	Headers      map[string]string             `json:"headers,omitempty"`
	BodyTemplate *string                       `json:"body_template,omitempty"`
	// Şablon çıktısının içerik türü (varsayılan application/json). Değerler: "application/json", "text/plain".
	BodyContentType *string                  `json:"body_content_type,omitempty"`
	Filters         *CreateWebhookDtoFilters `json:"filters,omitempty"`
}

// CreateWebhookDtoEventFilters: CreateWebhookDtoEventFilters şeması.
type CreateWebhookDtoEventFilters struct {
	EventNames []string `json:"event_names,omitempty"`
}

// CreateWebhookDtoFilters: Olay filtresi: { match: all|any, conditions: [{ field, op, value? }] }.
type CreateWebhookDtoFilters struct {
	// Zorunlu. Değerler: "all", "any".
	Match string `json:"match"`
	// Zorunlu.
	Conditions []CreateWebhookDtoFiltersConditionsItem `json:"conditions"`
}

// CreateWebhookDtoFiltersConditionsItem: CreateWebhookDtoFiltersConditionsItem şeması.
type CreateWebhookDtoFiltersConditionsItem struct {
	// Zorunlu.
	Field string `json:"field"`
	// Zorunlu. Değerler: "eq", "neq", "in", "not_in", "starts_with", "contains", "exists", "not_exists".
	Op    string `json:"op"`
	Value any    `json:"value,omitempty"`
}

// DataFeed: DataFeed şeması.
type DataFeed struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	URLTemplate string `json:"url_template"`
	// Zorunlu. Değerler: "GET", "POST".
	Method string `json:"method"`
	// Zorunlu.
	BodyTemplate *string `json:"body_template"`
	// Zorunlu.
	Headers map[string]string `json:"headers"`
	// Zorunlu.
	SecretHeaderNames []string `json:"secret_header_names"`
	// Zorunlu.
	CacheTTLSeconds int64 `json:"cache_ttl_seconds"`
	// Zorunlu.
	TimeoutMs int64 `json:"timeout_ms"`
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu.
	Liquid string `json:"liquid"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// DataFeedList: DataFeedList şeması.
type DataFeedList struct {
	// Zorunlu.
	Data []DataFeed `json:"data"`
}

// DataFeedTestResult: DataFeedTestResult şeması.
type DataFeedTestResult struct {
	// Zorunlu. Değerler: "ok", "cached", "timeout", "error", "circuit_open", "blocked", "invalid_json", "too_large", "http_error".
	Status string `json:"status"`
	// Zorunlu.
	HTTPStatus *int64 `json:"http_status"`
	// Zorunlu.
	DurationMs int64 `json:"duration_ms"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	URL *string `json:"url"`
	// Zorunlu.
	Preview any `json:"preview"`
	// Zorunlu.
	Truncated bool `json:"truncated"`
	// Zorunlu.
	User *DataFeedTestResultUser `json:"user"`
}

// DataFeedTestResultUser: DataFeedTestResultUser şeması.
type DataFeedTestResultUser struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
}

// DeliverabilitySettingsDto: DeliverabilitySettingsDto şeması.
type DeliverabilitySettingsDto struct {
	BccAddress *string                              `json:"bcc_address,omitempty"`
	Validation *DeliverabilitySettingsDtoValidation `json:"validation,omitempty"`
	Verifier   *DeliverabilitySettingsDtoVerifier   `json:"verifier,omitempty"`
}

// DeliverabilitySettingsDtoValidation: DeliverabilitySettingsDtoValidation şeması.
type DeliverabilitySettingsDtoValidation struct {
	// Zorunlu.
	Enabled         bool  `json:"enabled"`
	CheckMx         *bool `json:"check_mx,omitempty"`
	BlockDisposable *bool `json:"block_disposable,omitempty"`
	BlockRole       *bool `json:"block_role,omitempty"`
}

// DeliverabilitySettingsDtoVerifier: DeliverabilitySettingsDtoVerifier şeması.
type DeliverabilitySettingsDtoVerifier struct {
	// Zorunlu. Değerler: "zerobounce", "neverbounce", "mock".
	Provider *string `json:"provider"`
	APIKey   *string `json:"api_key,omitempty"`
}

// Delivery: Delivery şeması.
type Delivery struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	MessageID string `json:"message_id"`
	// Zorunlu.
	SubscriptionID string `json:"subscription_id"`
	// Zorunlu.
	UserID *string `json:"user_id"`
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu.
	Provider *string `json:"provider"`
	// Zorunlu. Değerler: "pending", "sent", "delivered", "failed", "bounced", "suppressed", "expired", "cancelled".
	Status string `json:"status"`
	// Zorunlu.
	AttemptCount int64 `json:"attempt_count"`
	// Zorunlu.
	ProviderMessageID *string `json:"provider_message_id"`
	// Zorunlu.
	ErrorCode *string `json:"error_code"`
	// Zorunlu.
	ErrorMessage *string `json:"error_message"`
	// Zorunlu.
	SuppressionReason *string `json:"suppression_reason"`
	// Zorunlu.
	SendAfter *string `json:"send_after"`
	// Zorunlu.
	SentAt *string `json:"sent_at"`
	// Zorunlu.
	DeliveredAt *string `json:"delivered_at"`
	// Zorunlu.
	OpenedAt *string `json:"opened_at"`
	// Zorunlu.
	ClickedAt *string `json:"clicked_at"`
	// Zorunlu.
	FailedAt *string `json:"failed_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// DeliveryList: DeliveryList şeması.
type DeliveryList struct {
	// Zorunlu.
	Data []Delivery `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// EmailAddressCheck: EmailAddressCheck şeması.
type EmailAddressCheck struct {
	// Zorunlu.
	Ok bool `json:"ok"`
	// Zorunlu. Değerler: "syntax", "no_mx", "disposable", "role", "verifier_invalid".
	Reason *string `json:"reason"`
	// Zorunlu.
	Detail *string `json:"detail"`
	// Zorunlu.
	Checks []string `json:"checks"`
	// Zorunlu.
	Unknown []string `json:"unknown"`
}

// EmailDeliverabilitySettings: EmailDeliverabilitySettings şeması.
type EmailDeliverabilitySettings struct {
	// Zorunlu.
	BccAddress *string `json:"bcc_address"`
	// Zorunlu.
	Validation EmailDeliverabilitySettingsValidation `json:"validation"`
	// Zorunlu.
	Verifier EmailDeliverabilitySettingsVerifier `json:"verifier"`
}

// EmailDeliverabilitySettingsValidation: EmailDeliverabilitySettingsValidation şeması.
type EmailDeliverabilitySettingsValidation struct {
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu.
	CheckMx bool `json:"check_mx"`
	// Zorunlu.
	BlockDisposable bool `json:"block_disposable"`
	// Zorunlu.
	BlockRole bool `json:"block_role"`
}

// EmailDeliverabilitySettingsVerifier: EmailDeliverabilitySettingsVerifier şeması.
type EmailDeliverabilitySettingsVerifier struct {
	// Zorunlu. Değerler: "zerobounce", "neverbounce", "mock".
	Provider *string `json:"provider"`
	// Zorunlu.
	HasKey bool `json:"has_key"`
}

// EmailWarmupPlan: EmailWarmupPlan şeması.
type EmailWarmupPlan struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "domain", "ip".
	Kind string `json:"kind"`
	// Zorunlu.
	Target string `json:"target"`
	// Zorunlu.
	StartDate string `json:"start_date"`
	// Zorunlu.
	DailyLimits []int64 `json:"daily_limits"`
	// Zorunlu.
	ProviderShares map[string]float64 `json:"provider_shares"`
	// Zorunlu.
	Timezone string `json:"timezone"`
	// Zorunlu. Değerler: "active", "paused", "completed".
	Status string `json:"status"`
	// Zorunlu.
	Today EmailWarmupPlanToday `json:"today"`
	// Zorunlu.
	ScheduledAhead int64 `json:"scheduled_ahead"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// EmailWarmupPlanList: EmailWarmupPlanList şeması.
type EmailWarmupPlanList struct {
	// Zorunlu.
	Data []EmailWarmupPlan `json:"data"`
}

// EmailWarmupPlanToday: EmailWarmupPlanToday şeması.
type EmailWarmupPlanToday struct {
	// Zorunlu.
	Day string `json:"day"`
	// Zorunlu.
	Index int64 `json:"index"`
	// Zorunlu.
	Limit *int64 `json:"limit"`
	// Zorunlu.
	Used int64 `json:"used"`
	// Zorunlu.
	Buckets map[string]EmailWarmupPlanTodayBucketsValue `json:"buckets"`
}

// EmailWarmupPlanTodayBucketsValue: EmailWarmupPlanTodayBucketsValue şeması.
type EmailWarmupPlanTodayBucketsValue struct {
	// Zorunlu.
	Limit *int64 `json:"limit"`
	// Zorunlu.
	Used int64 `json:"used"`
}

// EmailWarmupUsage: EmailWarmupUsage şeması.
type EmailWarmupUsage struct {
	// Zorunlu.
	Data []EmailWarmupUsageDataItem `json:"data"`
}

// EmailWarmupUsageDataItem: EmailWarmupUsageDataItem şeması.
type EmailWarmupUsageDataItem struct {
	// Zorunlu.
	Day string `json:"day"`
	// Zorunlu.
	Bucket string `json:"bucket"`
	// Zorunlu.
	Used int64 `json:"used"`
	// Zorunlu.
	Limit *int64 `json:"limit"`
}

// EndLiveActivityDto: EndLiveActivityDto şeması.
type EndLiveActivityDto struct {
	// Zorunlu.
	ActivityID     string                   `json:"activity_id"`
	ContentState   map[string]any           `json:"content_state,omitempty"`
	Alert          *EndLiveActivityDtoAlert `json:"alert,omitempty"`
	StaleDate      *int64                   `json:"stale_date,omitempty"`
	DismissalDate  *int64                   `json:"dismissal_date,omitempty"`
	RelevanceScore *float64                 `json:"relevance_score,omitempty"`
	// Değerler: 5, 10.
	Priority *int64  `json:"priority,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// EndLiveActivityDtoAlert: EndLiveActivityDtoAlert şeması.
type EndLiveActivityDtoAlert struct {
	Title map[string]string `json:"title,omitempty"`
	Body  map[string]string `json:"body,omitempty"`
	Sound *string           `json:"sound,omitempty"`
}

// EngagementReport: EngagementReport şeması.
type EngagementReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	Timezone string `json:"timezone"`
	// Zorunlu.
	Totals EngagementReportTotals `json:"totals"`
	// Zorunlu.
	Previous *EngagementReportPrevious `json:"previous"`
	// Zorunlu.
	Cells []EngagementReportCellsItem `json:"cells"`
	// Zorunlu.
	Rows []EngagementReportRowsItem `json:"rows"`
}

// EngagementReportCellsItem: EngagementReportCellsItem şeması.
type EngagementReportCellsItem struct {
	// Zorunlu.
	Weekday int64 `json:"weekday"`
	// Zorunlu.
	Hour int64 `json:"hour"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
}

// EngagementReportPrevious: EngagementReportPrevious şeması.
type EngagementReportPrevious struct {
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
}

// EngagementReportRowsItem: EngagementReportRowsItem şeması.
type EngagementReportRowsItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
}

// EngagementReportTotals: EngagementReportTotals şeması.
type EngagementReportTotals struct {
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
}

// EstimateAudienceDto: EstimateAudienceDto şeması.
type EstimateAudienceDto struct {
	AppID *string `json:"app_id,omitempty"`
	// Zorunlu.
	Target              EstimateAudienceDtoTarget `json:"target"`
	IncludeUnsubscribed *bool                     `json:"include_unsubscribed,omitempty"`
}

// EstimateAudienceDtoTarget: EstimateAudienceDtoTarget şeması.
type EstimateAudienceDtoTarget struct {
	All                    *bool               `json:"all,omitempty"`
	IncludeSubscriptionIds []string            `json:"include_subscription_ids,omitempty"`
	IncludeExternalIds     []string            `json:"include_external_ids,omitempty"`
	IncludeAliases         map[string][]string `json:"include_aliases,omitempty"`
	IncludedSegments       []string            `json:"included_segments,omitempty"`
	ExcludedSegments       []string            `json:"excluded_segments,omitempty"`
	Filters                any                 `json:"filters,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels            []string `json:"channels,omitempty"`
	TestDevicesOnly     *bool    `json:"test_devices_only,omitempty"`
	EmailTo             []string `json:"email_to,omitempty"`
	IncludePhoneNumbers []string `json:"include_phone_numbers,omitempty"`
}

// EstimateAudienceResult: EstimateAudienceResult şeması.
type EstimateAudienceResult struct {
	// Zorunlu.
	Recipients int64 `json:"recipients"`
	// Zorunlu.
	ByChannel map[string]int64 `json:"by_channel"`
	// Zorunlu.
	NewContacts int64 `json:"new_contacts"`
	// Zorunlu.
	ExternalIDMissing []string `json:"external_id_missing"`
}

// EventDefinition: EventDefinition şeması.
type EventDefinition struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	DisplayName *string `json:"display_name"`
	// Zorunlu.
	PropertiesSchema map[string]string `json:"properties_schema"`
	// Zorunlu.
	IsConversion bool `json:"is_conversion"`
	// Zorunlu.
	OutcomeUnique bool `json:"outcome_unique"`
	// Zorunlu.
	OutcomeUnattributed bool `json:"outcome_unattributed"`
	// Zorunlu.
	RetentionDays *int64 `json:"retention_days"`
	// Zorunlu.
	FirstSeenAt string `json:"first_seen_at"`
	// Zorunlu.
	LastSeenAt string `json:"last_seen_at"`
	// Zorunlu.
	Count30d int64 `json:"count_30d"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// EventDefinitionDetail: EventDefinitionDetail şeması.
type EventDefinitionDetail struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	DisplayName *string `json:"display_name"`
	// Zorunlu.
	PropertiesSchema map[string]string `json:"properties_schema"`
	// Zorunlu.
	IsConversion bool `json:"is_conversion"`
	// Zorunlu.
	OutcomeUnique bool `json:"outcome_unique"`
	// Zorunlu.
	OutcomeUnattributed bool `json:"outcome_unattributed"`
	// Zorunlu.
	RetentionDays *int64 `json:"retention_days"`
	// Zorunlu.
	FirstSeenAt string `json:"first_seen_at"`
	// Zorunlu.
	LastSeenAt string `json:"last_seen_at"`
	// Zorunlu.
	Count30d int64 `json:"count_30d"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
	// Zorunlu.
	Daily []EventDefinitionDetailDailyItem `json:"daily"`
	// Zorunlu.
	Properties []EventDefinitionDetailPropertiesItem `json:"properties"`
}

// EventDefinitionDetailDailyItem: EventDefinitionDetailDailyItem şeması.
type EventDefinitionDetailDailyItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	Count int64 `json:"count"`
	// Zorunlu.
	Users int64 `json:"users"`
}

// EventDefinitionDetailPropertiesItem: EventDefinitionDetailPropertiesItem şeması.
type EventDefinitionDetailPropertiesItem struct {
	// Zorunlu.
	Key string `json:"key"`
	// Zorunlu.
	Type string `json:"type"`
	// Zorunlu.
	Values []EventDefinitionDetailPropertiesItemValuesItem `json:"values"`
}

// EventDefinitionDetailPropertiesItemValuesItem: EventDefinitionDetailPropertiesItemValuesItem şeması.
type EventDefinitionDetailPropertiesItemValuesItem struct {
	// Zorunlu.
	Value string `json:"value"`
	// Zorunlu.
	Count int64 `json:"count"`
}

// EventDefinitionList: EventDefinitionList şeması.
type EventDefinitionList struct {
	// Zorunlu.
	Data []EventDefinition `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// EventsAccepted: EventsAccepted şeması.
type EventsAccepted struct {
	// Zorunlu.
	Accepted int64 `json:"accepted"`
}

// EventsReport: EventsReport şeması.
type EventsReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu. Değerler: "day", "name".
	GroupBy string `json:"group_by"`
	// Zorunlu.
	Name *string `json:"name"`
	// Zorunlu.
	Totals EventsReportTotals `json:"totals"`
	// Zorunlu.
	Previous *EventsReportPrevious `json:"previous"`
	// Zorunlu.
	Rows []EventsReportRowsItem `json:"rows"`
}

// EventsReportPrevious: EventsReportPrevious şeması.
type EventsReportPrevious struct {
	// Zorunlu.
	Count int64 `json:"count"`
	// Zorunlu.
	Users int64 `json:"users"`
}

// EventsReportRowsItem: EventsReportRowsItem şeması.
type EventsReportRowsItem struct {
	// Zorunlu.
	Key string `json:"key"`
	// Zorunlu.
	Count int64 `json:"count"`
	// Zorunlu.
	Users int64 `json:"users"`
}

// EventsReportTotals: EventsReportTotals şeması.
type EventsReportTotals struct {
	// Zorunlu.
	Count int64 `json:"count"`
	// Zorunlu.
	Users int64 `json:"users"`
}

// Export: Export şeması.
type Export struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu. Değerler: "users", "subscriptions", "events", "deliveries", "message_events".
	Type string `json:"type"`
	// Zorunlu.
	Filters ExportFilters `json:"filters"`
	// Zorunlu. Değerler: "pending", "running", "completed", "failed".
	Status string `json:"status"`
	// Zorunlu.
	RowCount *int64 `json:"row_count"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	DownloadURL *string `json:"download_url"`
	// Zorunlu.
	ExpiresAt *string `json:"expires_at"`
	// Zorunlu.
	RequestedBy *string `json:"requested_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	CompletedAt *string `json:"completed_at"`
}

// ExportFilters: ExportFilters şeması.
type ExportFilters struct {
	From       *string `json:"from,omitempty"`
	To         *string `json:"to,omitempty"`
	Name       *string `json:"name,omitempty"`
	MessageID  *string `json:"message_id,omitempty"`
	CampaignID *string `json:"campaign_id,omitempty"`
	Status     *string `json:"status,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel   *string `json:"channel,omitempty"`
	SegmentID *string `json:"segment_id,omitempty"`
}

// ExportList: ExportList şeması.
type ExportList struct {
	// Zorunlu.
	Data []Export `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// Form: Form şeması.
type Form struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "active", "archived".
	Status string `json:"status"`
	// Zorunlu.
	Fields []FormFieldsItem `json:"fields"`
	// Zorunlu.
	Settings FormSettings `json:"settings"`
	// Zorunlu.
	Stats FormStats `json:"stats"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// FormFieldsItem: FormFieldsItem şeması.
type FormFieldsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "email", "phone", "first_name", "last_name", "full_name", "text", "textarea", "select", "radio", "checkbox", "consent".
	Type string `json:"type"`
	// Zorunlu.
	Label       map[string]string           `json:"label"`
	Placeholder map[string]string           `json:"placeholder,omitempty"`
	Required    *bool                       `json:"required,omitempty"`
	Options     []FormFieldsItemOptionsItem `json:"options,omitempty"`
	MapsTo      *string                     `json:"maps_to,omitempty"`
	MaxLength   *int64                      `json:"max_length,omitempty"`
	PolicyURL   *string                     `json:"policy_url,omitempty"`
}

// FormFieldsItemOptionsItem: FormFieldsItemOptionsItem şeması.
type FormFieldsItemOptionsItem struct {
	// Zorunlu.
	Value string `json:"value"`
	// Zorunlu.
	Label map[string]string `json:"label"`
}

// FormPage: FormPage şeması.
type FormPage struct {
	// Zorunlu.
	Data []Form `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// FormSettings: FormSettings şeması.
type FormSettings struct {
	SubmitLabel     map[string]string        `json:"submit_label,omitempty"`
	SuccessMessage  map[string]string        `json:"success_message,omitempty"`
	RedirectURL     *string                  `json:"redirect_url,omitempty"`
	Tags            map[string]string        `json:"tags,omitempty"`
	Subscribe       *FormSettingsSubscribe   `json:"subscribe,omitempty"`
	DoubleOptIn     *FormSettingsDoubleOptIn `json:"double_opt_in,omitempty"`
	EventName       *string                  `json:"event_name,omitempty"`
	Captcha         *bool                    `json:"captcha,omitempty"`
	DefaultLanguage *string                  `json:"default_language,omitempty"`
}

// FormSettingsDoubleOptIn: FormSettingsDoubleOptIn şeması.
type FormSettingsDoubleOptIn struct {
	// Zorunlu.
	Enabled bool              `json:"enabled"`
	Subject map[string]string `json:"subject,omitempty"`
	Body    map[string]string `json:"body,omitempty"`
}

// FormSettingsSubscribe: FormSettingsSubscribe şeması.
type FormSettingsSubscribe struct {
	Email *bool `json:"email,omitempty"`
	SMS   *bool `json:"sms,omitempty"`
}

// FormStats: FormStats şeması.
type FormStats struct {
	// Zorunlu.
	Submissions int64 `json:"submissions"`
	// Zorunlu.
	LastSubmissionAt *string `json:"last_submission_at"`
}

// FormSubmission: FormSubmission şeması.
type FormSubmission struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	FormID string `json:"form_id"`
	// Zorunlu.
	UserID *string `json:"user_id"`
	// Zorunlu.
	SubscriptionIds []string `json:"subscription_ids"`
	// Zorunlu. Değerler: "web_campaign", "landing_page", "sdk", "api".
	Source string `json:"source"`
	// Zorunlu.
	SourceID *string `json:"source_id"`
	// Zorunlu. Değerler: "accepted", "pending_confirmation", "confirmed".
	Status string `json:"status"`
	// Zorunlu.
	Data map[string]any `json:"data"`
	// Zorunlu.
	Consent []FormSubmissionConsentItem `json:"consent"`
	// Zorunlu.
	PageURL *string `json:"page_url"`
	// Zorunlu.
	Language *string `json:"language"`
	// Zorunlu.
	ConfirmedAt *string `json:"confirmed_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// FormSubmissionConsentItem: FormSubmissionConsentItem şeması.
type FormSubmissionConsentItem struct {
	// Zorunlu.
	FieldID string `json:"field_id"`
	// Zorunlu.
	Given bool `json:"given"`
	// Zorunlu.
	Text string `json:"text"`
	// Zorunlu.
	TextHash string `json:"text_hash"`
	// Zorunlu.
	PolicyURL *string `json:"policy_url"`
	// Zorunlu.
	Language string `json:"language"`
	// Zorunlu.
	At string `json:"at"`
}

// FormSubmissionPage: FormSubmissionPage şeması.
type FormSubmissionPage struct {
	// Zorunlu.
	Data []FormSubmission `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// FormSubmitResult: FormSubmitResult şeması.
type FormSubmitResult struct {
	// Zorunlu.
	Ok bool `json:"ok"`
	// Zorunlu. Değerler: "accepted", "pending_confirmation".
	Status string `json:"status"`
	// Zorunlu.
	SubmissionID *string `json:"submission_id"`
	// Zorunlu.
	SuccessMessage string `json:"success_message"`
	// Zorunlu.
	RedirectURL *string `json:"redirect_url"`
}

// FunnelBodyDto: FunnelBodyDto şeması.
type FunnelBodyDto struct {
	// Zorunlu.
	Steps       []FunnelBodyDtoStepsItem `json:"steps"`
	From        *string                  `json:"from,omitempty"`
	To          *string                  `json:"to,omitempty"`
	WindowHours *int64                   `json:"window_hours,omitempty"`
	Compare     *bool                    `json:"compare,omitempty"`
	// Değerler: "json", "csv".
	Format *string `json:"format,omitempty"`
}

// FunnelBodyDtoStepsItem: FunnelBodyDtoStepsItem şeması.
type FunnelBodyDtoStepsItem struct {
	// Zorunlu.
	Name  string                       `json:"name"`
	Where *FunnelBodyDtoStepsItemWhere `json:"where,omitempty"`
}

// FunnelBodyDtoStepsItemWhere: FunnelBodyDtoStepsItemWhere şeması.
type FunnelBodyDtoStepsItemWhere struct {
	// Zorunlu.
	Key string `json:"key"`
	// Zorunlu.
	Value string `json:"value"`
}

// FunnelReport: FunnelReport şeması.
type FunnelReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	WindowHours int64 `json:"window_hours"`
	// Zorunlu.
	Steps []FunnelReportStepsItem `json:"steps"`
	// Zorunlu.
	OverallRate *float64 `json:"overall_rate"`
	// Zorunlu.
	Previous *FunnelReportPrevious `json:"previous"`
}

// FunnelReportPrevious: FunnelReportPrevious şeması.
type FunnelReportPrevious struct {
	// Zorunlu.
	Steps []int64 `json:"steps"`
	// Zorunlu.
	OverallRate *float64 `json:"overall_rate"`
}

// FunnelReportStepsItem: FunnelReportStepsItem şeması.
type FunnelReportStepsItem struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Users int64 `json:"users"`
	// Zorunlu.
	ConversionRate *float64 `json:"conversion_rate"`
	// Zorunlu.
	DropOff int64 `json:"drop_off"`
}

// IdentityKey: IdentityKey şeması.
type IdentityKey struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Kid string `json:"kid"`
	// Zorunlu. Değerler: "ES256".
	Alg string `json:"alg"`
	// Zorunlu.
	PublicJwk map[string]string `json:"public_jwk"`
	// Zorunlu.
	Label *string `json:"label"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	RevokedAt *string `json:"revoked_at"`
}

// IdentityKeyList: IdentityKeyList şeması.
type IdentityKeyList struct {
	// Zorunlu.
	Data []IdentityKey `json:"data"`
}

// Import: Import şeması.
type Import struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu. Değerler: "eski sağlayıcı_players", "csv_users".
	Type string `json:"type"`
	// Zorunlu.
	FileName *string `json:"file_name"`
	// Zorunlu. Değerler: "pending", "running", "completed", "failed".
	Status string `json:"status"`
	// Zorunlu.
	Stats ImportStats `json:"stats"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	ErrorReportURL *string `json:"error_report_url"`
	// Zorunlu.
	RequestedBy *string `json:"requested_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	CompletedAt *string `json:"completed_at"`
}

// ImportList: ImportList şeması.
type ImportList struct {
	// Zorunlu.
	Data []Import `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// ImportStats: ImportStats şeması.
type ImportStats struct {
	Processed            *int64 `json:"processed,omitempty"`
	CreatedUsers         *int64 `json:"created_users,omitempty"`
	UpdatedUsers         *int64 `json:"updated_users,omitempty"`
	CreatedSubscriptions *int64 `json:"created_subscriptions,omitempty"`
	Skipped              *int64 `json:"skipped,omitempty"`
	Errors               *int64 `json:"errors,omitempty"`
}

// InAppMessage: InAppMessage şeması.
type InAppMessage struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "draft", "active", "paused", "archived".
	Status string `json:"status"`
	// Zorunlu.
	Priority int64 `json:"priority"`
	// Zorunlu.
	Triggers []any `json:"triggers"`
	// Zorunlu.
	TriggerGroups [][]any `json:"trigger_groups"`
	// Zorunlu.
	Variants []InAppMessageVariantsItem `json:"variants"`
	// Zorunlu.
	Display InAppMessageDisplay `json:"display"`
	// Zorunlu.
	Audience InAppMessageAudience `json:"audience"`
	// Zorunlu.
	Frequency InAppMessageFrequency `json:"frequency"`
	// Zorunlu.
	JourneyOnly bool `json:"journey_only"`
	// Zorunlu.
	StartAt *string `json:"start_at"`
	// Zorunlu.
	EndAt *string `json:"end_at"`
	// Zorunlu.
	Stats InAppMessageStats `json:"stats"`
	// Zorunlu.
	VariantStats []InAppMessageVariantStatsItem `json:"variant_stats"`
	// Zorunlu.
	AbTest InAppMessageAbTest `json:"ab_test"`
	// Zorunlu.
	VariantSignificance *InAppMessageVariantSignificance `json:"variant_significance"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// InAppMessageAbTest: InAppMessageAbTest şeması.
type InAppMessageAbTest struct {
	Winner *InAppMessageAbTestWinner `json:"winner,omitempty"`
	// Zorunlu.
	WinnerVariantID *string `json:"winner_variant_id"`
	// Zorunlu.
	DecidedAt *string `json:"decided_at"`
	// Zorunlu.
	DecidedBy *string `json:"decided_by"`
}

// InAppMessageAbTestWinner: InAppMessageAbTestWinner şeması.
type InAppMessageAbTestWinner struct {
	// Zorunlu. Değerler: "manual", "auto".
	Mode           string   `json:"mode"`
	MinImpressions *int64   `json:"min_impressions,omitempty"`
	Confidence     *float64 `json:"confidence,omitempty"`
}

// InAppMessageAudience: InAppMessageAudience şeması.
type InAppMessageAudience struct {
	IncludedSegments []string `json:"included_segments,omitempty"`
	ExcludedSegments []string `json:"excluded_segments,omitempty"`
}

// InAppMessageDisplay: InAppMessageDisplay şeması.
type InAppMessageDisplay struct {
	// Zorunlu. Değerler: "modal", "banner_top", "banner_bottom", "fullscreen".
	Layout             string  `json:"layout"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	// Değerler: "blocks", "html".
	ContentType *string           `json:"content_type,omitempty"`
	HTML        map[string]string `json:"html,omitempty"`
	HTMLHeight  *int64            `json:"html_height,omitempty"`
	// Zorunlu.
	Blocks []any `json:"blocks"`
}

// InAppMessageFrequency: InAppMessageFrequency şeması.
type InAppMessageFrequency struct {
	MaxImpressions       *int64 `json:"max_impressions,omitempty"`
	MinIntervalSeconds   *int64 `json:"min_interval_seconds,omitempty"`
	StopAfterInteraction *bool  `json:"stop_after_interaction,omitempty"`
}

// InAppMessagePage: InAppMessagePage şeması.
type InAppMessagePage struct {
	// Zorunlu.
	Data []InAppMessage `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// InAppMessageStats: InAppMessageStats şeması.
type InAppMessageStats struct {
	// Zorunlu.
	Impressions int64 `json:"impressions"`
	// Zorunlu.
	Clicks int64 `json:"clicks"`
	// Zorunlu.
	Dismissals int64 `json:"dismissals"`
	// Zorunlu.
	UniqueUsers int64 `json:"unique_users"`
	// Zorunlu.
	Ctr float64 `json:"ctr"`
}

// InAppMessageVariantSignificance: InAppMessageVariantSignificance şeması.
type InAppMessageVariantSignificance struct {
	// Zorunlu.
	Leader *string `json:"leader"`
	// Zorunlu.
	PValue float64 `json:"p_value"`
	// Zorunlu.
	Significant bool `json:"significant"`
}

// InAppMessageVariantsItem: InAppMessageVariantsItem şeması.
type InAppMessageVariantsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	Display InAppMessageVariantsItemDisplay `json:"display"`
}

// InAppMessageVariantsItemDisplay: InAppMessageVariantsItemDisplay şeması.
type InAppMessageVariantsItemDisplay struct {
	// Zorunlu. Değerler: "modal", "banner_top", "banner_bottom", "fullscreen".
	Layout             string  `json:"layout"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	// Değerler: "blocks", "html".
	ContentType *string           `json:"content_type,omitempty"`
	HTML        map[string]string `json:"html,omitempty"`
	HTMLHeight  *int64            `json:"html_height,omitempty"`
	// Zorunlu.
	Blocks []any `json:"blocks"`
}

// InAppMessageVariantStatsItem: InAppMessageVariantStatsItem şeması.
type InAppMessageVariantStatsItem struct {
	// Zorunlu.
	VariantID string `json:"variant_id"`
	// Zorunlu.
	Impressions int64 `json:"impressions"`
	// Zorunlu.
	Clicks int64 `json:"clicks"`
	// Zorunlu.
	Ctr float64 `json:"ctr"`
}

// InAppStatusDto: InAppStatusDto şeması.
type InAppStatusDto struct {
	// Zorunlu. Değerler: "draft", "active", "paused", "archived".
	Status string `json:"status"`
}

// InAppWinnerDto: InAppWinnerDto şeması.
type InAppWinnerDto struct {
	// Zorunlu.
	VariantID string `json:"variant_id"`
}

// InboxAgent: InboxAgent şeması.
type InboxAgent struct {
	// Zorunlu.
	AccountID string `json:"account_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Email string `json:"email"`
	// Zorunlu.
	Role string `json:"role"`
	// Zorunlu. Değerler: "online", "busy", "offline".
	Status string `json:"status"`
	// Zorunlu.
	LastHeartbeatAt *string `json:"last_heartbeat_at"`
	// Zorunlu.
	MaxConcurrent *int64 `json:"max_concurrent"`
	// Zorunlu.
	Languages []string `json:"languages"`
	// Zorunlu.
	OpenConversations int64 `json:"open_conversations"`
	// Zorunlu.
	TeamIds []string `json:"team_ids"`
}

// InboxAgentList: InboxAgentList şeması.
type InboxAgentList struct {
	// Zorunlu.
	Data []InboxAgent `json:"data"`
}

// InboxItem: InboxItem şeması.
type InboxItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	MessageID *string `json:"message_id"`
	// Zorunlu.
	Title *string `json:"title"`
	// Zorunlu.
	Body *string `json:"body"`
	// Zorunlu.
	URL *string `json:"url"`
	// Zorunlu.
	ImageURL *string `json:"image_url"`
	// Zorunlu.
	Data map[string]string `json:"data"`
	// Zorunlu.
	Read bool `json:"read"`
	// Zorunlu.
	ReadAt *string `json:"read_at"`
	// Zorunlu.
	ArchivedAt *string `json:"archived_at"`
	// Zorunlu.
	ExpiresAt *string `json:"expires_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// InboxMetrics: InboxMetrics şeması.
type InboxMetrics struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	FirstResponse InboxMetricsFirstResponse `json:"first_response"`
	// Zorunlu.
	Resolution InboxMetricsResolution `json:"resolution"`
	// Zorunlu.
	Sla InboxMetricsSla `json:"sla"`
	// Zorunlu.
	Volume InboxMetricsVolume `json:"volume"`
	// Zorunlu.
	Agents []InboxMetricsAgentsItem `json:"agents"`
	// Zorunlu.
	Daily []InboxMetricsDailyItem `json:"daily"`
}

// InboxMetricsAgentsItem: InboxMetricsAgentsItem şeması.
type InboxMetricsAgentsItem struct {
	// Zorunlu.
	AccountID string `json:"account_id"`
	// Zorunlu.
	Name *string `json:"name"`
	// Zorunlu.
	FirstResponse InboxMetricsAgentsItemFirstResponse `json:"first_response"`
	// Zorunlu.
	Resolution InboxMetricsAgentsItemResolution `json:"resolution"`
	// Zorunlu.
	Resolved int64 `json:"resolved"`
	// Zorunlu.
	OpenNow int64 `json:"open_now"`
}

// InboxMetricsAgentsItemFirstResponse: InboxMetricsAgentsItemFirstResponse şeması.
type InboxMetricsAgentsItemFirstResponse struct {
	// Zorunlu.
	Count int64 `json:"count"`
	// Zorunlu.
	AvgSeconds *float64 `json:"avg_seconds"`
	// Zorunlu.
	P90Seconds *float64 `json:"p90_seconds"`
}

// InboxMetricsAgentsItemResolution: InboxMetricsAgentsItemResolution şeması.
type InboxMetricsAgentsItemResolution struct {
	// Zorunlu.
	Count int64 `json:"count"`
	// Zorunlu.
	AvgSeconds *float64 `json:"avg_seconds"`
	// Zorunlu.
	P90Seconds *float64 `json:"p90_seconds"`
}

// InboxMetricsDailyItem: InboxMetricsDailyItem şeması.
type InboxMetricsDailyItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	Created int64 `json:"created"`
	// Zorunlu.
	Resolved int64 `json:"resolved"`
	// Zorunlu.
	FirstResponseAvgSeconds *float64 `json:"first_response_avg_seconds"`
}

// InboxMetricsFirstResponse: InboxMetricsFirstResponse şeması.
type InboxMetricsFirstResponse struct {
	// Zorunlu.
	Count int64 `json:"count"`
	// Zorunlu.
	AvgSeconds *float64 `json:"avg_seconds"`
	// Zorunlu.
	P90Seconds *float64 `json:"p90_seconds"`
}

// InboxMetricsResolution: InboxMetricsResolution şeması.
type InboxMetricsResolution struct {
	// Zorunlu.
	Count int64 `json:"count"`
	// Zorunlu.
	AvgSeconds *float64 `json:"avg_seconds"`
	// Zorunlu.
	P90Seconds *float64 `json:"p90_seconds"`
}

// InboxMetricsSla: InboxMetricsSla şeması.
type InboxMetricsSla struct {
	// Zorunlu.
	TargetSeconds int64 `json:"target_seconds"`
	// Zorunlu.
	Met int64 `json:"met"`
	// Zorunlu.
	Breached int64 `json:"breached"`
	// Zorunlu.
	Rate *float64 `json:"rate"`
}

// InboxMetricsVolume: InboxMetricsVolume şeması.
type InboxMetricsVolume struct {
	// Zorunlu.
	Created int64 `json:"created"`
	// Zorunlu.
	Resolved int64 `json:"resolved"`
	// Zorunlu.
	OpenNow int64 `json:"open_now"`
	// Zorunlu.
	UnassignedNow int64 `json:"unassigned_now"`
	// Zorunlu.
	ByChannel map[string]int64 `json:"by_channel"`
}

// InboxPage: InboxPage şeması.
type InboxPage struct {
	// Zorunlu.
	Data []InboxItem `json:"data"`
	// Zorunlu.
	UnreadCount int64 `json:"unread_count"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// InboxReadAllResult: InboxReadAllResult şeması.
type InboxReadAllResult struct {
	// Zorunlu.
	Updated int64 `json:"updated"`
	// Zorunlu.
	UnreadCount int64 `json:"unread_count"`
}

// InboxSettings: InboxSettings şeması.
type InboxSettings struct {
	// Zorunlu.
	Widget InboxSettingsWidget `json:"widget"`
	// Zorunlu.
	BusinessHours InboxSettingsBusinessHours `json:"business_hours"`
	// Zorunlu.
	Routing InboxSettingsRouting `json:"routing"`
	// Zorunlu.
	Email InboxSettingsEmail `json:"email"`
	// Zorunlu.
	Timezone string `json:"timezone"`
	// Zorunlu.
	WithinBusinessHours bool `json:"within_business_hours"`
	// Zorunlu.
	Install InboxSettingsInstall `json:"install"`
}

// InboxSettingsBusinessHours: InboxSettingsBusinessHours şeması.
type InboxSettingsBusinessHours struct {
	Enabled  *bool                                `json:"enabled,omitempty"`
	Timezone *string                              `json:"timezone,omitempty"`
	Days     []InboxSettingsBusinessHoursDaysItem `json:"days,omitempty"`
	Holidays []string                             `json:"holidays,omitempty"`
}

// InboxSettingsBusinessHoursDaysItem: InboxSettingsBusinessHoursDaysItem şeması.
type InboxSettingsBusinessHoursDaysItem struct {
	// Zorunlu.
	Day int64 `json:"day"`
	// Zorunlu.
	Start string `json:"start"`
	// Zorunlu.
	End string `json:"end"`
}

// InboxSettingsEmail: InboxSettingsEmail şeması.
type InboxSettingsEmail struct {
	// Zorunlu.
	CaptureReplies bool `json:"capture_replies"`
	// Zorunlu.
	ReplyDomain *string `json:"reply_domain"`
}

// InboxSettingsInstall: InboxSettingsInstall şeması.
type InboxSettingsInstall struct {
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	SDKURL string `json:"sdk_url"`
	// Zorunlu.
	APIURL string `json:"api_url"`
}

// InboxSettingsRouting: InboxSettingsRouting şeması.
type InboxSettingsRouting struct {
	// Zorunlu.
	AutoAssign bool `json:"auto_assign"`
	// Zorunlu. Değerler: "all", "team", "assigned".
	Visibility string `json:"visibility"`
	// Zorunlu.
	ShowUnassigned bool `json:"show_unassigned"`
	// Zorunlu.
	FirstResponseSlaSeconds int64 `json:"first_response_sla_seconds"`
	// Zorunlu.
	MaxConcurrentDefault *int64 `json:"max_concurrent_default"`
}

// InboxSettingsUpdateDto: InboxSettingsUpdateDto şeması.
type InboxSettingsUpdateDto struct {
	Widget        *InboxSettingsUpdateDtoWidget        `json:"widget,omitempty"`
	BusinessHours *InboxSettingsUpdateDtoBusinessHours `json:"business_hours,omitempty"`
	Routing       *InboxSettingsUpdateDtoRouting       `json:"routing,omitempty"`
	Email         *InboxSettingsUpdateDtoEmail         `json:"email,omitempty"`
}

// InboxSettingsUpdateDtoBusinessHours: InboxSettingsUpdateDtoBusinessHours şeması.
type InboxSettingsUpdateDtoBusinessHours struct {
	Enabled  *bool                                         `json:"enabled,omitempty"`
	Timezone *string                                       `json:"timezone,omitempty"`
	Days     []InboxSettingsUpdateDtoBusinessHoursDaysItem `json:"days,omitempty"`
	Holidays []string                                      `json:"holidays,omitempty"`
}

// InboxSettingsUpdateDtoBusinessHoursDaysItem: InboxSettingsUpdateDtoBusinessHoursDaysItem şeması.
type InboxSettingsUpdateDtoBusinessHoursDaysItem struct {
	// Zorunlu.
	Day int64 `json:"day"`
	// Zorunlu.
	Start string `json:"start"`
	// Zorunlu.
	End string `json:"end"`
}

// InboxSettingsUpdateDtoEmail: InboxSettingsUpdateDtoEmail şeması.
type InboxSettingsUpdateDtoEmail struct {
	CaptureReplies *bool `json:"capture_replies,omitempty"`
}

// InboxSettingsUpdateDtoRouting: InboxSettingsUpdateDtoRouting şeması.
type InboxSettingsUpdateDtoRouting struct {
	AutoAssign *bool `json:"auto_assign,omitempty"`
	// Değerler: "all", "team", "assigned".
	Visibility              *string `json:"visibility,omitempty"`
	ShowUnassigned          *bool   `json:"show_unassigned,omitempty"`
	FirstResponseSlaSeconds *int64  `json:"first_response_sla_seconds,omitempty"`
	MaxConcurrentDefault    *int64  `json:"max_concurrent_default,omitempty"`
}

// InboxSettingsUpdateDtoWidget: InboxSettingsUpdateDtoWidget şeması.
type InboxSettingsUpdateDtoWidget struct {
	Enabled   *bool   `json:"enabled,omitempty"`
	Color     *string `json:"color,omitempty"`
	TextColor *string `json:"text_color,omitempty"`
	// Değerler: "right", "left".
	Position *string `json:"position,omitempty"`
	// Değerler: "light", "dark", "auto".
	Theme               *string           `json:"theme,omitempty"`
	Title               map[string]string `json:"title,omitempty"`
	Greeting            map[string]string `json:"greeting,omitempty"`
	OfflineMessage      map[string]string `json:"offline_message,omitempty"`
	CollectEmailOffline *bool             `json:"collect_email_offline,omitempty"`
	Attachments         *bool             `json:"attachments,omitempty"`
	RequireEmail        *bool             `json:"require_email,omitempty"`
}

// InboxSettingsWidget: InboxSettingsWidget şeması.
type InboxSettingsWidget struct {
	Enabled   *bool   `json:"enabled,omitempty"`
	Color     *string `json:"color,omitempty"`
	TextColor *string `json:"text_color,omitempty"`
	// Değerler: "right", "left".
	Position *string `json:"position,omitempty"`
	// Değerler: "light", "dark", "auto".
	Theme               *string           `json:"theme,omitempty"`
	Title               map[string]string `json:"title,omitempty"`
	Greeting            map[string]string `json:"greeting,omitempty"`
	OfflineMessage      map[string]string `json:"offline_message,omitempty"`
	CollectEmailOffline *bool             `json:"collect_email_offline,omitempty"`
	Attachments         *bool             `json:"attachments,omitempty"`
	RequireEmail        *bool             `json:"require_email,omitempty"`
}

// InboxTeam: InboxTeam şeması.
type InboxTeam struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	MemberIds []string `json:"member_ids"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// InboxTeamList: InboxTeamList şeması.
type InboxTeamList struct {
	// Zorunlu.
	Data []InboxTeam `json:"data"`
}

// IntegrationOk: IntegrationOk şeması.
type IntegrationOk struct {
	// Zorunlu.
	Ok bool `json:"ok"`
}

// Journey: Journey şeması.
type Journey struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu. Değerler: "draft", "active", "paused", "archived".
	Status string `json:"status"`
	// Zorunlu.
	CurrentVersion *int64 `json:"current_version"`
	// Zorunlu.
	DraftVersion *int64 `json:"draft_version"`
	// Zorunlu.
	TriggerKind *string `json:"trigger_kind"`
	// Zorunlu.
	PublishedAt *string `json:"published_at"`
	// Zorunlu.
	CreatedBy *string `json:"created_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
	// Zorunlu.
	Summary JourneySummary `json:"summary"`
}

// JourneyDetail: JourneyDetail şeması.
type JourneyDetail struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu. Değerler: "draft", "active", "paused", "archived".
	Status string `json:"status"`
	// Zorunlu.
	CurrentVersion *int64 `json:"current_version"`
	// Zorunlu.
	DraftVersion *int64 `json:"draft_version"`
	// Zorunlu.
	TriggerKind *string `json:"trigger_kind"`
	// Zorunlu.
	PublishedAt *string `json:"published_at"`
	// Zorunlu.
	CreatedBy *string `json:"created_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
	// Zorunlu.
	Summary JourneyDetailSummary `json:"summary"`
	// Zorunlu.
	Editing JourneyDetailEditing `json:"editing"`
	// Zorunlu.
	Issues []JourneyDetailIssuesItem `json:"issues"`
}

// JourneyDetailEditing: JourneyDetailEditing şeması.
type JourneyDetailEditing struct {
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	Graph map[string]any `json:"graph"`
	// Zorunlu.
	Settings map[string]any `json:"settings"`
	// Zorunlu.
	PublishedAt *string `json:"published_at"`
	// Zorunlu.
	PublishedBy *string `json:"published_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// JourneyDetailIssuesItem: JourneyDetailIssuesItem şeması.
type JourneyDetailIssuesItem struct {
	// Zorunlu.
	Code   string  `json:"code"`
	NodeID *string `json:"node_id,omitempty"`
	EdgeID *string `json:"edge_id,omitempty"`
	// Zorunlu.
	Message string `json:"message"`
}

// JourneyDetailSummary: JourneyDetailSummary şeması.
type JourneyDetailSummary struct {
	// Zorunlu.
	RunsTotal int64 `json:"runs_total"`
	// Zorunlu.
	Active int64 `json:"active"`
	// Zorunlu.
	Completed int64 `json:"completed"`
	// Zorunlu.
	Exited int64 `json:"exited"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	GoalReached int64 `json:"goal_reached"`
}

// JourneyPage: JourneyPage şeması.
type JourneyPage struct {
	// Zorunlu.
	Data []Journey `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// JourneyPublishResult: JourneyPublishResult şeması.
type JourneyPublishResult struct {
	// Zorunlu.
	Journey Journey `json:"journey"`
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	Issues []JourneyPublishResultIssuesItem `json:"issues"`
}

// JourneyPublishResultIssuesItem: JourneyPublishResultIssuesItem şeması.
type JourneyPublishResultIssuesItem struct {
	// Zorunlu.
	Code   string  `json:"code"`
	NodeID *string `json:"node_id,omitempty"`
	EdgeID *string `json:"edge_id,omitempty"`
	// Zorunlu.
	Message string `json:"message"`
}

// JourneyRun: JourneyRun şeması.
type JourneyRun struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	JourneyID string `json:"journey_id"`
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	UserID string `json:"user_id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu. Değerler: "active", "waiting", "completed", "exited", "failed".
	Status string `json:"status"`
	// Zorunlu.
	CurrentNodeID *string `json:"current_node_id"`
	// Zorunlu.
	NextActionAt *string `json:"next_action_at"`
	// Zorunlu.
	EnteredAt string `json:"entered_at"`
	// Zorunlu.
	CompletedAt *string `json:"completed_at"`
	// Zorunlu.
	ExitReason *string `json:"exit_reason"`
	// Zorunlu.
	GoalID *string `json:"goal_id"`
	// Zorunlu.
	GoalReachedAt *string `json:"goal_reached_at"`
	// Zorunlu.
	Context map[string]any `json:"context"`
}

// JourneyRunDetail: JourneyRunDetail şeması.
type JourneyRunDetail struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	JourneyID string `json:"journey_id"`
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	UserID string `json:"user_id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu. Değerler: "active", "waiting", "completed", "exited", "failed".
	Status string `json:"status"`
	// Zorunlu.
	CurrentNodeID *string `json:"current_node_id"`
	// Zorunlu.
	NextActionAt *string `json:"next_action_at"`
	// Zorunlu.
	EnteredAt string `json:"entered_at"`
	// Zorunlu.
	CompletedAt *string `json:"completed_at"`
	// Zorunlu.
	ExitReason *string `json:"exit_reason"`
	// Zorunlu.
	GoalID *string `json:"goal_id"`
	// Zorunlu.
	GoalReachedAt *string `json:"goal_reached_at"`
	// Zorunlu.
	Context map[string]any `json:"context"`
	// Zorunlu.
	JourneyName string `json:"journey_name"`
	// Zorunlu.
	Steps []JourneyRunDetailStepsItem `json:"steps"`
}

// JourneyRunDetailStepsItem: JourneyRunDetailStepsItem şeması.
type JourneyRunDetailStepsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	NodeID string `json:"node_id"`
	// Zorunlu.
	NodeType string `json:"node_type"`
	// Zorunlu. Değerler: "completed", "waiting", "failed", "skipped".
	Status string `json:"status"`
	// Zorunlu.
	Branch *string `json:"branch"`
	// Zorunlu.
	Result map[string]any `json:"result"`
	// Zorunlu.
	StartedAt string `json:"started_at"`
	// Zorunlu.
	FinishedAt *string `json:"finished_at"`
}

// JourneyRunPage: JourneyRunPage şeması.
type JourneyRunPage struct {
	// Zorunlu.
	Data []JourneyRun `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// JourneysCreateBody: JourneysCreateBody şeması.
type JourneysCreateBody struct {
	// Zorunlu.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// React Flow uyumlu journey grafı: { nodes: [{ id, type, position, data }], edges: [{ id, source, target, sourceHandle }] } (docs/journeys.md)
	Graph *JourneysCreateBodyGraph `json:"graph,omitempty"`
	// Yeniden giriş, yalnızca yeni girenler, çıkış kuralları, çıkış tag'leri, hedefler, takvim (docs/journeys.md)
	Settings map[string]any `json:"settings,omitempty"`
}

// JourneysCreateBodyGraph: React Flow uyumlu journey grafı: { nodes: [{ id, type, position, data }], edges: [{ id, source, target, sourceHandle }] } (docs/journeys.md)
type JourneysCreateBodyGraph struct {
	// Zorunlu.
	Nodes []map[string]any `json:"nodes"`
	// Zorunlu.
	Edges []map[string]any `json:"edges"`
}

// JourneysSaveDraftBody: JourneysSaveDraftBody şeması.
type JourneysSaveDraftBody struct {
	// Zorunlu. React Flow uyumlu journey grafı: { nodes: [{ id, type, position, data }], edges: [{ id, source, target, sourceHandle }] } (docs/journeys.md)
	Graph JourneysSaveDraftBodyGraph `json:"graph"`
	// Yeniden giriş, yalnızca yeni girenler, çıkış kuralları, çıkış tag'leri, hedefler, takvim (docs/journeys.md)
	Settings map[string]any `json:"settings,omitempty"`
}

// JourneysSaveDraftBodyGraph: React Flow uyumlu journey grafı: { nodes: [{ id, type, position, data }], edges: [{ id, source, target, sourceHandle }] } (docs/journeys.md)
type JourneysSaveDraftBodyGraph struct {
	// Zorunlu.
	Nodes []map[string]any `json:"nodes"`
	// Zorunlu.
	Edges []map[string]any `json:"edges"`
}

// JourneyStats: JourneyStats şeması.
type JourneyStats struct {
	// Zorunlu.
	Version *int64 `json:"version"`
	// Zorunlu.
	Summary JourneyStatsSummary `json:"summary"`
	// Zorunlu.
	Nodes []JourneyStatsNodesItem `json:"nodes"`
	// Zorunlu.
	Goals []JourneyStatsGoalsItem `json:"goals"`
	// Zorunlu.
	ExitReasons map[string]int64 `json:"exit_reasons"`
}

// JourneyStatsGoalsItem: JourneyStatsGoalsItem şeması.
type JourneyStatsGoalsItem struct {
	// Zorunlu.
	GoalID string `json:"goal_id"`
	// Zorunlu.
	Reached int64 `json:"reached"`
}

// JourneyStatsNodesItem: JourneyStatsNodesItem şeması.
type JourneyStatsNodesItem struct {
	// Zorunlu.
	NodeID string `json:"node_id"`
	// Zorunlu.
	NodeType string `json:"node_type"`
	// Zorunlu.
	Entered int64 `json:"entered"`
	// Zorunlu.
	Completed int64 `json:"completed"`
	// Zorunlu.
	WaitingNow int64 `json:"waiting_now"`
	// Zorunlu.
	Skipped int64 `json:"skipped"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Branches map[string]int64                    `json:"branches"`
	Messages map[string]int64                    `json:"messages,omitempty"`
	Variants []JourneyStatsNodesItemVariantsItem `json:"variants,omitempty"`
}

// JourneyStatsNodesItemVariantsItem: JourneyStatsNodesItemVariantsItem şeması.
type JourneyStatsNodesItemVariantsItem struct {
	// Zorunlu.
	VariantID string `json:"variant_id"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
	// Zorunlu.
	Converted int64 `json:"converted"`
	// Zorunlu.
	Winner bool `json:"winner"`
}

// JourneyStatsSummary: JourneyStatsSummary şeması.
type JourneyStatsSummary struct {
	// Zorunlu.
	RunsTotal int64 `json:"runs_total"`
	// Zorunlu.
	Active int64 `json:"active"`
	// Zorunlu.
	Completed int64 `json:"completed"`
	// Zorunlu.
	Exited int64 `json:"exited"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	GoalReached int64 `json:"goal_reached"`
	// Zorunlu.
	Waiting int64 `json:"waiting"`
	// Zorunlu.
	GoalConversionRate float64 `json:"goal_conversion_rate"`
}

// JourneySummary: JourneySummary şeması.
type JourneySummary struct {
	// Zorunlu.
	RunsTotal int64 `json:"runs_total"`
	// Zorunlu.
	Active int64 `json:"active"`
	// Zorunlu.
	Completed int64 `json:"completed"`
	// Zorunlu.
	Exited int64 `json:"exited"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	GoalReached int64 `json:"goal_reached"`
}

// JourneyTriggerResult: JourneyTriggerResult şeması.
type JourneyTriggerResult struct {
	// Zorunlu.
	Entered bool `json:"entered"`
	// Zorunlu.
	Run *JourneyRun `json:"run"`
	// Zorunlu. Değerler: "inactive", "outside_schedule", "re_entry", "already_active", "unknown_user".
	Reason *string `json:"reason"`
}

// JourneyVersionList: JourneyVersionList şeması.
type JourneyVersionList struct {
	// Zorunlu.
	Data []JourneyVersionListDataItem `json:"data"`
}

// JourneyVersionListDataItem: JourneyVersionListDataItem şeması.
type JourneyVersionListDataItem struct {
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	Graph map[string]any `json:"graph"`
	// Zorunlu.
	Settings map[string]any `json:"settings"`
	// Zorunlu.
	PublishedAt *string `json:"published_at"`
	// Zorunlu.
	PublishedBy *string `json:"published_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// LandingPage: LandingPage şeması.
type LandingPage struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Slug string `json:"slug"`
	// Zorunlu. Değerler: "draft", "published", "archived".
	Status string `json:"status"`
	// Zorunlu.
	Seo LandingPageSeo `json:"seo"`
	// Zorunlu.
	Settings LandingPageSettings `json:"settings"`
	// Zorunlu.
	CustomDomainID *string `json:"custom_domain_id"`
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	PublishedVersion *int64 `json:"published_version"`
	// Zorunlu.
	PublishedAt *string `json:"published_at"`
	// Zorunlu.
	HasUnpublishedChanges bool `json:"has_unpublished_changes"`
	// Zorunlu.
	URL string `json:"url"`
	// Zorunlu.
	DomainURL *string `json:"domain_url"`
	// Zorunlu.
	Stats LandingPageStats `json:"stats"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
	// Zorunlu.
	Document LandingPageDocument `json:"document"`
}

// LandingPageDocument: LandingPageDocument şeması.
type LandingPageDocument struct {
	// Zorunlu. Değerler: 1.
	Version float64 `json:"version"`
	// Zorunlu.
	Settings LandingPageDocumentSettings `json:"settings"`
	// Zorunlu.
	Sections []LandingPageDocumentSectionsItem `json:"sections"`
}

// LandingPageDocumentSectionsItem: LandingPageDocumentSectionsItem şeması.
type LandingPageDocumentSectionsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "section".
	Type            string                                  `json:"type"`
	BackgroundColor *string                                 `json:"background_color,omitempty"`
	BackgroundImage *string                                 `json:"background_image,omitempty"`
	Padding         *LandingPageDocumentSectionsItemPadding `json:"padding,omitempty"`
	StackOnMobile   *bool                                   `json:"stack_on_mobile,omitempty"`
	// Zorunlu.
	Columns []LandingPageDocumentSectionsItemColumnsItem `json:"columns"`
}

// LandingPageDocumentSectionsItemColumnsItem: LandingPageDocumentSectionsItemColumnsItem şeması.
type LandingPageDocumentSectionsItemColumnsItem struct {
	// Zorunlu.
	ID              string  `json:"id"`
	WidthPercent    *int64  `json:"width_percent,omitempty"`
	BackgroundColor *string `json:"background_color,omitempty"`
	// Değerler: "top", "middle", "bottom".
	VerticalAlign *string                                            `json:"vertical_align,omitempty"`
	Padding       *LandingPageDocumentSectionsItemColumnsItemPadding `json:"padding,omitempty"`
	// Zorunlu.
	Blocks []any `json:"blocks"`
}

// LandingPageDocumentSectionsItemColumnsItemPadding: LandingPageDocumentSectionsItemColumnsItemPadding şeması.
type LandingPageDocumentSectionsItemColumnsItemPadding struct {
	// Zorunlu.
	Top int64 `json:"top"`
	// Zorunlu.
	Right int64 `json:"right"`
	// Zorunlu.
	Bottom int64 `json:"bottom"`
	// Zorunlu.
	Left int64 `json:"left"`
}

// LandingPageDocumentSectionsItemPadding: LandingPageDocumentSectionsItemPadding şeması.
type LandingPageDocumentSectionsItemPadding struct {
	// Zorunlu.
	Top int64 `json:"top"`
	// Zorunlu.
	Right int64 `json:"right"`
	// Zorunlu.
	Bottom int64 `json:"bottom"`
	// Zorunlu.
	Left int64 `json:"left"`
}

// LandingPageDocumentSettings: LandingPageDocumentSettings şeması.
type LandingPageDocumentSettings struct {
	Width                  *int64  `json:"width,omitempty"`
	BackgroundColor        *string `json:"background_color,omitempty"`
	ContentBackgroundColor *string `json:"content_background_color,omitempty"`
	TextColor              *string `json:"text_color,omitempty"`
	LinkColor              *string `json:"link_color,omitempty"`
	// Değerler: "system", "arial", "georgia", "verdana", "trebuchet", "courier".
	FontFamily *string                              `json:"font_family,omitempty"`
	DarkMode   *LandingPageDocumentSettingsDarkMode `json:"dark_mode,omitempty"`
}

// LandingPageDocumentSettingsDarkMode: LandingPageDocumentSettingsDarkMode şeması.
type LandingPageDocumentSettingsDarkMode struct {
	BackgroundColor        *string `json:"background_color,omitempty"`
	ContentBackgroundColor *string `json:"content_background_color,omitempty"`
	TextColor              *string `json:"text_color,omitempty"`
	LinkColor              *string `json:"link_color,omitempty"`
}

// LandingPageList: LandingPageList şeması.
type LandingPageList struct {
	// Zorunlu.
	Data []LandingPageListDataItem `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// LandingPageListDataItem: LandingPageListDataItem şeması.
type LandingPageListDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Slug string `json:"slug"`
	// Zorunlu. Değerler: "draft", "published", "archived".
	Status string `json:"status"`
	// Zorunlu.
	Seo LandingPageListDataItemSeo `json:"seo"`
	// Zorunlu.
	Settings LandingPageListDataItemSettings `json:"settings"`
	// Zorunlu.
	CustomDomainID *string `json:"custom_domain_id"`
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	PublishedVersion *int64 `json:"published_version"`
	// Zorunlu.
	PublishedAt *string `json:"published_at"`
	// Zorunlu.
	HasUnpublishedChanges bool `json:"has_unpublished_changes"`
	// Zorunlu.
	URL string `json:"url"`
	// Zorunlu.
	DomainURL *string `json:"domain_url"`
	// Zorunlu.
	Stats LandingPageListDataItemStats `json:"stats"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// LandingPageListDataItemSeo: LandingPageListDataItemSeo şeması.
type LandingPageListDataItemSeo struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	OgImage      *string `json:"og_image,omitempty"`
	CanonicalURL *string `json:"canonical_url,omitempty"`
	Noindex      *bool   `json:"noindex,omitempty"`
}

// LandingPageListDataItemSettings: LandingPageListDataItemSettings şeması.
type LandingPageListDataItemSettings struct {
	Language      *string                                       `json:"language,omitempty"`
	FormID        *string                                       `json:"form_id,omitempty"`
	FormSectionID *string                                       `json:"form_section_id,omitempty"`
	CookieConsent *LandingPageListDataItemSettingsCookieConsent `json:"cookie_consent,omitempty"`
	WebSDK        *bool                                         `json:"web_sdk,omitempty"`
	FaviconURL    *string                                       `json:"favicon_url,omitempty"`
}

// LandingPageListDataItemSettingsCookieConsent: LandingPageListDataItemSettingsCookieConsent şeması.
type LandingPageListDataItemSettingsCookieConsent struct {
	// Zorunlu.
	Enabled     bool    `json:"enabled"`
	Text        *string `json:"text,omitempty"`
	AcceptLabel *string `json:"accept_label,omitempty"`
	RejectLabel *string `json:"reject_label,omitempty"`
	PolicyURL   *string `json:"policy_url,omitempty"`
}

// LandingPageListDataItemStats: LandingPageListDataItemStats şeması.
type LandingPageListDataItemStats struct {
	// Zorunlu.
	Submissions int64 `json:"submissions"`
}

// LandingPageSeo: LandingPageSeo şeması.
type LandingPageSeo struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	OgImage      *string `json:"og_image,omitempty"`
	CanonicalURL *string `json:"canonical_url,omitempty"`
	Noindex      *bool   `json:"noindex,omitempty"`
}

// LandingPageSettings: LandingPageSettings şeması.
type LandingPageSettings struct {
	Language      *string                           `json:"language,omitempty"`
	FormID        *string                           `json:"form_id,omitempty"`
	FormSectionID *string                           `json:"form_section_id,omitempty"`
	CookieConsent *LandingPageSettingsCookieConsent `json:"cookie_consent,omitempty"`
	WebSDK        *bool                             `json:"web_sdk,omitempty"`
	FaviconURL    *string                           `json:"favicon_url,omitempty"`
}

// LandingPageSettingsCookieConsent: LandingPageSettingsCookieConsent şeması.
type LandingPageSettingsCookieConsent struct {
	// Zorunlu.
	Enabled     bool    `json:"enabled"`
	Text        *string `json:"text,omitempty"`
	AcceptLabel *string `json:"accept_label,omitempty"`
	RejectLabel *string `json:"reject_label,omitempty"`
	PolicyURL   *string `json:"policy_url,omitempty"`
}

// LandingPageStats: LandingPageStats şeması.
type LandingPageStats struct {
	// Zorunlu.
	Submissions int64 `json:"submissions"`
}

// LiveActivitySendResult: LiveActivitySendResult şeması.
type LiveActivitySendResult struct {
	// Zorunlu.
	MessageID string `json:"message_id"`
	// Zorunlu.
	ActivityID string `json:"activity_id"`
	// Zorunlu. Değerler: "start", "update", "end".
	Event string `json:"event"`
	// Zorunlu.
	RecipientsEstimate int64 `json:"recipients_estimate"`
}

// LiveActivityTokenResult: LiveActivityTokenResult şeması.
type LiveActivityTokenResult struct {
	// Zorunlu.
	Ok bool `json:"ok"`
}

// LiveActivityType: LiveActivityType şeması.
type LiveActivityType struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	ContentStateSchema LiveActivityTypeContentStateSchema `json:"content_state_schema"`
	// Zorunlu.
	AttributesSchema LiveActivityTypeAttributesSchema `json:"attributes_schema"`
	// Zorunlu.
	PushToStartTokens int64 `json:"push_to_start_tokens"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// LiveActivityTypeAttributesSchema: LiveActivityTypeAttributesSchema şeması.
type LiveActivityTypeAttributesSchema struct {
	// Zorunlu. Değerler: "object".
	Type string `json:"type"`
	// Zorunlu.
	Properties map[string]LiveActivityTypeAttributesSchemaPropertiesValue `json:"properties"`
	Required   []string                                                   `json:"required,omitempty"`
}

// LiveActivityTypeAttributesSchemaPropertiesValue: LiveActivityTypeAttributesSchemaPropertiesValue şeması.
type LiveActivityTypeAttributesSchemaPropertiesValue struct {
	// Zorunlu. Değerler: "string", "number", "integer", "boolean".
	Type        string   `json:"type"`
	Enum        []any    `json:"enum,omitempty"`
	MaxLength   *int64   `json:"maxLength,omitempty"`
	Minimum     *float64 `json:"minimum,omitempty"`
	Maximum     *float64 `json:"maximum,omitempty"`
	Description *string  `json:"description,omitempty"`
}

// LiveActivityTypeContentStateSchema: LiveActivityTypeContentStateSchema şeması.
type LiveActivityTypeContentStateSchema struct {
	// Zorunlu. Değerler: "object".
	Type string `json:"type"`
	// Zorunlu.
	Properties map[string]LiveActivityTypeContentStateSchemaPropertiesValue `json:"properties"`
	Required   []string                                                     `json:"required,omitempty"`
}

// LiveActivityTypeContentStateSchemaPropertiesValue: LiveActivityTypeContentStateSchemaPropertiesValue şeması.
type LiveActivityTypeContentStateSchemaPropertiesValue struct {
	// Zorunlu. Değerler: "string", "number", "integer", "boolean".
	Type        string   `json:"type"`
	Enum        []any    `json:"enum,omitempty"`
	MaxLength   *int64   `json:"maxLength,omitempty"`
	Minimum     *float64 `json:"minimum,omitempty"`
	Maximum     *float64 `json:"maximum,omitempty"`
	Description *string  `json:"description,omitempty"`
}

// LiveActivityTypeList: LiveActivityTypeList şeması.
type LiveActivityTypeList struct {
	// Zorunlu.
	Data []LiveActivityType `json:"data"`
}

// McpCreateCampaignDraftBody: McpCreateCampaignDraftBody şeması.
type McpCreateCampaignDraftBody struct {
	Name *string `json:"name,omitempty"`
	// Değerler: "one_time", "scheduled", "recurring".
	Type     *string                             `json:"type,omitempty"`
	Audience *McpCreateCampaignDraftBodyAudience `json:"audience,omitempty"`
	// messages.content ile aynı şema (headings, contents, url, buttons…)
	Content map[string]any `json:"content,omitempty"`
	Labels  []string       `json:"labels,omitempty"`
	// Tercih kategorisi anahtarı; kapalı kullanıcılar bastırılır
	Category *string                           `json:"category,omitempty"`
	AbTest   *McpCreateCampaignDraftBodyAbTest `json:"ab_test,omitempty"`
	// A/B varyantları (ağırlıklar toplamı 100; en fazla bir is_control)
	Variants        []McpCreateCampaignDraftBodyVariantsItem `json:"variants,omitempty"`
	TemplateID      *string                                  `json:"template_id,omitempty"`
	TemplateVersion *int64                                   `json:"template_version,omitempty"`
	Delivery        *McpCreateCampaignDraftBodyDelivery      `json:"delivery,omitempty"`
	Recurrence      *McpCreateCampaignDraftBodyRecurrence    `json:"recurrence,omitempty"`
	SendAt          *string                                  `json:"send_at,omitempty"`
}

// McpCreateCampaignDraftBodyAbTest: McpCreateCampaignDraftBodyAbTest şeması.
type McpCreateCampaignDraftBodyAbTest struct {
	Enabled        *bool                                   `json:"enabled,omitempty"`
	TestPercentage *int64                                  `json:"test_percentage,omitempty"`
	Winner         *McpCreateCampaignDraftBodyAbTestWinner `json:"winner,omitempty"`
}

// McpCreateCampaignDraftBodyAbTestWinner: McpCreateCampaignDraftBodyAbTestWinner şeması.
type McpCreateCampaignDraftBodyAbTestWinner struct {
	// Değerler: "manual", "auto".
	Mode *string `json:"mode,omitempty"`
	// Değerler: "open_rate", "click_rate", "conversion_rate".
	Metric               *string `json:"metric,omitempty"`
	EvaluateAfterMinutes *int64  `json:"evaluate_after_minutes,omitempty"`
}

// McpCreateCampaignDraftBodyAudience: McpCreateCampaignDraftBodyAudience şeması.
type McpCreateCampaignDraftBodyAudience struct {
	All             *bool    `json:"all,omitempty"`
	Segments        []string `json:"segments,omitempty"`
	ExcludeSegments []string `json:"exclude_segments,omitempty"`
	// Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
	Filters *McpCreateCampaignDraftBodyAudienceFilters `json:"filters,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels        []string `json:"channels,omitempty"`
	TestDevicesOnly *bool    `json:"test_devices_only,omitempty"`
}

// McpCreateCampaignDraftBodyAudienceFilters: Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
type McpCreateCampaignDraftBodyAudienceFilters struct {
	All  []map[string]any `json:"all,omitempty"`
	Any  []map[string]any `json:"any,omitempty"`
	None []map[string]any `json:"none,omitempty"`
}

// McpCreateCampaignDraftBodyDelivery: McpCreateCampaignDraftBodyDelivery şeması.
type McpCreateCampaignDraftBodyDelivery struct {
	// Değerler: "immediate", "per_user_timezone", "optimized".
	TimezoneMode      *string `json:"timezone_mode,omitempty"`
	DeliveryTimeOfDay *string `json:"delivery_time_of_day,omitempty"`
	ThrottlePerMinute *int64  `json:"throttle_per_minute,omitempty"`
	TTL               *int64  `json:"ttl,omitempty"`
	// Değerler: 5, 10.
	Priority           *int64  `json:"priority,omitempty"`
	CollapseID         *string `json:"collapse_id,omitempty"`
	FrequencyCapExempt *bool   `json:"frequency_cap_exempt,omitempty"`
	QuietHoursExempt   *bool   `json:"quiet_hours_exempt,omitempty"`
	OnePerUser         *bool   `json:"one_per_user,omitempty"`
}

// McpCreateCampaignDraftBodyRecurrence: McpCreateCampaignDraftBodyRecurrence şeması.
type McpCreateCampaignDraftBodyRecurrence struct {
	Cron     *string `json:"cron,omitempty"`
	Until    *string `json:"until,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// McpCreateCampaignDraftBodyVariantsItem: McpCreateCampaignDraftBodyVariantsItem şeması.
type McpCreateCampaignDraftBodyVariantsItem struct {
	ID *string `json:"id,omitempty"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64          `json:"weight_percent"`
	Content       map[string]any `json:"content,omitempty"`
	IsControl     *bool          `json:"is_control,omitempty"`
}

// McpCreateSegmentDraftBody: McpCreateSegmentDraftBody şeması.
type McpCreateSegmentDraftBody struct {
	// Zorunlu.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Zorunlu. Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
	Definition McpCreateSegmentDraftBodyDefinition `json:"definition"`
	IsDynamic  *bool                               `json:"is_dynamic,omitempty"`
}

// McpCreateSegmentDraftBodyDefinition: Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
type McpCreateSegmentDraftBodyDefinition struct {
	All  []map[string]any `json:"all,omitempty"`
	Any  []map[string]any `json:"any,omitempty"`
	None []map[string]any `json:"none,omitempty"`
}

// Member: Member şeması.
type Member struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AccountID string `json:"account_id"`
	// Zorunlu.
	Email string `json:"email"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "owner", "admin", "developer", "marketer", "analyst", "billing", "viewer", "composer", "agent".
	Role string `json:"role"`
	// Zorunlu.
	AppScopes []string `json:"app_scopes"`
	// Zorunlu.
	AcceptedAt *string `json:"accepted_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	LastLoginAt *string `json:"last_login_at"`
}

// MemberList: MemberList şeması.
type MemberList struct {
	// Zorunlu.
	Data []Member `json:"data"`
}

// Message: Message şeması.
type Message struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu. Değerler: "campaign", "api", "journey", "test".
	Source string `json:"source"`
	// Zorunlu. Değerler: "queued", "processing", "completed", "cancelled", "failed".
	Status string `json:"status"`
	// Zorunlu.
	IdempotencyKey *string `json:"idempotency_key"`
	// Zorunlu.
	Name *string `json:"name"`
	// Zorunlu.
	Content map[string]any `json:"content"`
	// Zorunlu.
	Target map[string]any `json:"target"`
	// Zorunlu.
	Options map[string]any `json:"options"`
	// Zorunlu.
	Labels []string `json:"labels"`
	// Zorunlu.
	Category *string `json:"category"`
	// Zorunlu.
	TemplateID *string `json:"template_id"`
	// Zorunlu.
	CampaignID *string `json:"campaign_id"`
	// Zorunlu.
	Variants []MessageVariantsItem `json:"variants"`
	// Zorunlu.
	Counts MessageCounts `json:"counts"`
	// Zorunlu.
	RequestedBy MessageRequestedBy `json:"requested_by"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	ScheduledFor *string `json:"scheduled_for"`
	// Zorunlu.
	StartedAt *string `json:"started_at"`
	// Zorunlu.
	CompletedAt *string `json:"completed_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// MessageCategory: MessageCategory şeması.
type MessageCategory struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Key string `json:"key"`
	// Zorunlu.
	Name map[string]string `json:"name"`
	// Zorunlu.
	Description map[string]string `json:"description"`
	// Zorunlu.
	Channels []string `json:"channels"`
	// Zorunlu.
	DefaultOptedIn bool `json:"default_opted_in"`
	// Zorunlu.
	Position int64 `json:"position"`
	// Zorunlu.
	OptedOutUsers int64 `json:"opted_out_users"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// MessageCategoryList: MessageCategoryList şeması.
type MessageCategoryList struct {
	// Zorunlu.
	Data []MessageCategory `json:"data"`
}

// MessageCounts: MessageCounts şeması.
type MessageCounts struct {
	Targeted     *int64 `json:"targeted,omitempty"`
	Suppressed   *int64 `json:"suppressed,omitempty"`
	Sent         *int64 `json:"sent,omitempty"`
	Delivered    *int64 `json:"delivered,omitempty"`
	Failed       *int64 `json:"failed,omitempty"`
	Errored      *int64 `json:"errored,omitempty"`
	Expired      *int64 `json:"expired,omitempty"`
	Opened       *int64 `json:"opened,omitempty"`
	Clicked      *int64 `json:"clicked,omitempty"`
	Converted    *int64 `json:"converted,omitempty"`
	Bounced      *int64 `json:"bounced,omitempty"`
	Complained   *int64 `json:"complained,omitempty"`
	Unsubscribed *int64 `json:"unsubscribed,omitempty"`
	SMSSegments  *int64 `json:"sms_segments,omitempty"`
}

// MessageList: MessageList şeması.
type MessageList struct {
	// Zorunlu.
	Data []Message `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// MessageRequestedBy: MessageRequestedBy şeması.
type MessageRequestedBy struct {
	// Zorunlu. Değerler: "api_key", "account", "system".
	Type string  `json:"type"`
	ID   *string `json:"id,omitempty"`
}

// MessagesReport: MessagesReport şeması.
type MessagesReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu. Değerler: "day", "channel", "campaign".
	GroupBy string `json:"group_by"`
	// Zorunlu.
	Totals MessagesReportTotals `json:"totals"`
	// Zorunlu.
	Previous *MessagesReportPrevious `json:"previous"`
	// Zorunlu.
	Rows []MessagesReportRowsItem `json:"rows"`
}

// MessagesReportPrevious: MessagesReportPrevious şeması.
type MessagesReportPrevious struct {
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Delivered int64 `json:"delivered"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
	// Zorunlu.
	Converted int64 `json:"converted"`
	// Zorunlu.
	DeliveryRate *float64 `json:"delivery_rate"`
	// Zorunlu.
	OpenRate *float64 `json:"open_rate"`
	// Zorunlu.
	ClickRate *float64 `json:"click_rate"`
}

// MessagesReportRowsItem: MessagesReportRowsItem şeması.
type MessagesReportRowsItem struct {
	// Zorunlu.
	Key string `json:"key"`
	// Zorunlu.
	Label string `json:"label"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Delivered int64 `json:"delivered"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
	// Zorunlu.
	Converted int64 `json:"converted"`
}

// MessagesReportTotals: MessagesReportTotals şeması.
type MessagesReportTotals struct {
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Delivered int64 `json:"delivered"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
	// Zorunlu.
	Converted int64 `json:"converted"`
	// Zorunlu.
	DeliveryRate *float64 `json:"delivery_rate"`
	// Zorunlu.
	OpenRate *float64 `json:"open_rate"`
	// Zorunlu.
	ClickRate *float64 `json:"click_rate"`
}

// MessageVariantsItem: MessageVariantsItem şeması.
type MessageVariantsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	IsControl bool `json:"is_control"`
}

// Ok: Ok şeması.
type Ok struct {
	// Zorunlu.
	Ok bool `json:"ok"`
}

// Org: Org şeması.
type Org struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Slug string `json:"slug"`
	// Zorunlu.
	PlanCode string `json:"plan_code"`
	// Zorunlu. Değerler: "active", "suspended", "deleted".
	Status string `json:"status"`
	// Zorunlu.
	Settings OrgSettings `json:"settings"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	Region *string `json:"region"`
	// Zorunlu.
	ManagedByReseller bool `json:"managed_by_reseller"`
}

// OrgSettings: OrgSettings şeması.
type OrgSettings struct {
	MfaRequired     *bool                   `json:"mfaRequired,omitempty"`
	IPAllowlist     []string                `json:"ipAllowlist,omitempty"`
	RetentionDays   *int64                  `json:"retentionDays,omitempty"`
	SessionMaxHours *int64                  `json:"sessionMaxHours,omitempty"`
	AuditStream     *OrgSettingsAuditStream `json:"auditStream,omitempty"`
}

// OrgSettingsAuditStream: OrgSettingsAuditStream şeması.
type OrgSettingsAuditStream struct {
	S3           *bool   `json:"s3,omitempty"`
	WebhookAppID *string `json:"webhookAppId,omitempty"`
}

// OverviewStats: OverviewStats şeması.
type OverviewStats struct {
	// Zorunlu. Değerler: "7d", "30d", "custom".
	Range string `json:"range"`
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	Kpis OverviewStatsKpis `json:"kpis"`
	// Zorunlu.
	Series []OverviewStatsSeriesItem `json:"series"`
	// Zorunlu.
	RecentMessages []Message `json:"recent_messages"`
	// Zorunlu.
	RecentCampaigns []OverviewStatsRecentCampaignsItem `json:"recent_campaigns"`
	// Zorunlu.
	TopErrors []OverviewStatsTopErrorsItem `json:"top_errors"`
	// Zorunlu.
	Channels []OverviewStatsChannelsItem `json:"channels"`
}

// OverviewStatsChannelsItem: OverviewStatsChannelsItem şeması.
type OverviewStatsChannelsItem struct {
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu. Değerler: "apns", "fcm", "hms", "webpush", "smtp", "ses", "resend", "twilio", "netgsm", "meta_whatsapp", "google_rbm", "mock".
	Provider *string `json:"provider"`
	// Zorunlu. Değerler: "unconfigured", "active", "invalid", "suspended".
	Status string `json:"status"`
	// Zorunlu.
	LastError *string `json:"last_error"`
	// Zorunlu.
	ActiveSubscriptions int64 `json:"active_subscriptions"`
}

// OverviewStatsKpis: OverviewStatsKpis şeması.
type OverviewStatsKpis struct {
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Delivered int64 `json:"delivered"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
	// Zorunlu.
	DeliveryRate *float64 `json:"delivery_rate"`
	// Zorunlu.
	OpenRate *float64 `json:"open_rate"`
	// Zorunlu.
	ClickRate *float64 `json:"click_rate"`
	// Zorunlu.
	ActiveSubscriptions int64 `json:"active_subscriptions"`
	// Zorunlu.
	NewSubscriptions int64 `json:"new_subscriptions"`
	// Zorunlu.
	OptOuts int64 `json:"opt_outs"`
}

// OverviewStatsRecentCampaignsItem: OverviewStatsRecentCampaignsItem şeması.
type OverviewStatsRecentCampaignsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Type string `json:"type"`
	// Zorunlu.
	Status string `json:"status"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// OverviewStatsSeriesItem: OverviewStatsSeriesItem şeması.
type OverviewStatsSeriesItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Delivered int64 `json:"delivered"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	Opened int64 `json:"opened"`
}

// OverviewStatsTopErrorsItem: OverviewStatsTopErrorsItem şeması.
type OverviewStatsTopErrorsItem struct {
	// Zorunlu.
	Code string `json:"code"`
	// Zorunlu.
	Count int64 `json:"count"`
}

// PatchEventDefinitionDto: PatchEventDefinitionDto şeması.
type PatchEventDefinitionDto struct {
	DisplayName         *string `json:"display_name,omitempty"`
	IsConversion        *bool   `json:"is_conversion,omitempty"`
	OutcomeUnique       *bool   `json:"outcome_unique,omitempty"`
	OutcomeUnattributed *bool   `json:"outcome_unattributed,omitempty"`
	RetentionDays       *int64  `json:"retention_days,omitempty"`
}

// PatchSubscriptionDto: PatchSubscriptionDto şeması.
type PatchSubscriptionDto struct {
	Token  *string                     `json:"token,omitempty"`
	Device *PatchSubscriptionDtoDevice `json:"device,omitempty"`
	// Değerler: "subscribed", "unsubscribed".
	Status         *string `json:"status,omitempty"`
	TestDevice     *bool   `json:"test_device,omitempty"`
	TestDeviceName *string `json:"test_device_name,omitempty"`
}

// PatchSubscriptionDtoDevice: PatchSubscriptionDtoDevice şeması.
type PatchSubscriptionDtoDevice struct {
	Model       *string `json:"model,omitempty"`
	Os          *string `json:"os,omitempty"`
	OsVersion   *string `json:"os_version,omitempty"`
	AppVersion  *string `json:"app_version,omitempty"`
	SDKVersion  *string `json:"sdk_version,omitempty"`
	SDKPlatform *string `json:"sdk_platform,omitempty"`
	Carrier     *string `json:"carrier,omitempty"`
	Locale      *string `json:"locale,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
	PushCapable *bool   `json:"push_capable,omitempty"`
	// Değerler: "sandbox", "production".
	PushEnvironment *string `json:"push_environment,omitempty"`
}

// PatchUserDto: PatchUserDto şeması.
type PatchUserDto struct {
	Properties map[string]any      `json:"properties,omitempty"`
	Tags       map[string]string   `json:"tags,omitempty"`
	Deltas     *PatchUserDtoDeltas `json:"deltas,omitempty"`
}

// PatchUserDtoDeltas: PatchUserDtoDeltas şeması.
type PatchUserDtoDeltas struct {
	SessionCount *int64                            `json:"session_count,omitempty"`
	SessionTime  *int64                            `json:"session_time,omitempty"`
	Purchases    []PatchUserDtoDeltasPurchasesItem `json:"purchases,omitempty"`
}

// PatchUserDtoDeltasPurchasesItem: PatchUserDtoDeltasPurchasesItem şeması.
type PatchUserDtoDeltasPurchasesItem struct {
	// Zorunlu.
	Sku string `json:"sku"`
	// Zorunlu.
	Amount float64 `json:"amount"`
	Iso    *string `json:"iso,omitempty"`
	Count  *int64  `json:"count,omitempty"`
}

// ProvidersReport: ProvidersReport şeması.
type ProvidersReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	Totals ProvidersReportTotals `json:"totals"`
	// Zorunlu.
	Previous *ProvidersReportPrevious `json:"previous"`
	// Zorunlu.
	Rows []ProvidersReportRowsItem `json:"rows"`
}

// ProvidersReportPrevious: ProvidersReportPrevious şeması.
type ProvidersReportPrevious struct {
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	FailureRate *float64 `json:"failure_rate"`
}

// ProvidersReportRowsItem: ProvidersReportRowsItem şeması.
type ProvidersReportRowsItem struct {
	// Zorunlu.
	Channel string `json:"channel"`
	// Zorunlu.
	Provider *string `json:"provider"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	FailureRate *float64 `json:"failure_rate"`
	// Zorunlu.
	LatencyP50Ms *float64 `json:"latency_p50_ms"`
	// Zorunlu.
	LatencyP95Ms *float64 `json:"latency_p95_ms"`
	// Zorunlu.
	Errors []ProvidersReportRowsItemErrorsItem `json:"errors"`
}

// ProvidersReportRowsItemErrorsItem: ProvidersReportRowsItemErrorsItem şeması.
type ProvidersReportRowsItemErrorsItem struct {
	// Zorunlu.
	Code string `json:"code"`
	// Zorunlu.
	Count int64 `json:"count"`
}

// ProvidersReportTotals: ProvidersReportTotals şeması.
type ProvidersReportTotals struct {
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Failed int64 `json:"failed"`
	// Zorunlu.
	FailureRate *float64 `json:"failure_rate"`
}

// PublicFormSubmitDto: PublicFormSubmitDto şeması.
type PublicFormSubmitDto struct {
	// Zorunlu.
	Values     map[string]any `json:"values"`
	ExternalID *string        `json:"external_id,omitempty"`
	PageURL    *string        `json:"page_url,omitempty"`
	Language   *string        `json:"language,omitempty"`
}

// PublicSurveyResponseDto: PublicSurveyResponseDto şeması.
type PublicSurveyResponseDto struct {
	// Zorunlu.
	Answers    map[string]any `json:"answers"`
	ExternalID *string        `json:"external_id,omitempty"`
}

// RenderPreview: RenderPreview şeması.
type RenderPreview struct {
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu.
	Language string `json:"language"`
	// Zorunlu.
	Title *string `json:"title"`
	// Zorunlu.
	Subtitle *string `json:"subtitle"`
	// Zorunlu.
	Body *string `json:"body"`
	// Zorunlu.
	URL *string `json:"url"`
	// Zorunlu.
	ImageURL *string `json:"image_url"`
	// Zorunlu.
	IconURL *string `json:"icon_url"`
	// Zorunlu.
	Buttons []RenderPreviewButtonsItem `json:"buttons"`
	// Zorunlu.
	Data map[string]string `json:"data"`
	// Zorunlu.
	Warnings []string            `json:"warnings"`
	Email    *RenderPreviewEmail `json:"email,omitempty"`
	SMS      *RenderPreviewSMS   `json:"sms,omitempty"`
}

// RenderPreviewButtonsItem: RenderPreviewButtonsItem şeması.
type RenderPreviewButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string `json:"text"`
	// Zorunlu.
	URL *string `json:"url"`
}

// RenderPreviewEmail: RenderPreviewEmail şeması.
type RenderPreviewEmail struct {
	// Zorunlu.
	Subject string `json:"subject"`
	// Zorunlu.
	Preheader *string `json:"preheader"`
	// Zorunlu.
	HTML string `json:"html"`
	// Zorunlu.
	Text string `json:"text"`
}

// RenderPreviewSMS: RenderPreviewSms şeması.
type RenderPreviewSMS struct {
	// Zorunlu.
	Body string `json:"body"`
	// Zorunlu. Değerler: "gsm7", "gsm7_tr", "ucs2".
	Encoding string `json:"encoding"`
	// Zorunlu.
	Characters int64 `json:"characters"`
	// Zorunlu.
	Segments int64 `json:"segments"`
	// Zorunlu.
	PerSegment int64 `json:"per_segment"`
	// Zorunlu.
	NonGsmCharacters []string `json:"non_gsm_characters"`
	// Zorunlu.
	TurkishCharacters []string `json:"turkish_characters"`
}

// RenderTemplateDto: RenderTemplateDto şeması.
type RenderTemplateDto struct {
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel    *string                   `json:"channel,omitempty"`
	Language   *string                   `json:"language,omitempty"`
	User       *RenderTemplateDtoUser    `json:"user,omitempty"`
	CustomData map[string]any            `json:"custom_data,omitempty"`
	Content    *RenderTemplateDtoContent `json:"content,omitempty"`
	Version    *int64                    `json:"version,omitempty"`
}

// RenderTemplateDtoContent: RenderTemplateDtoContent şeması.
type RenderTemplateDtoContent struct {
	Headings       map[string]string                     `json:"headings,omitempty"`
	Contents       map[string]string                     `json:"contents,omitempty"`
	Subtitle       map[string]string                     `json:"subtitle,omitempty"`
	URL            *string                               `json:"url,omitempty"`
	AppURL         *string                               `json:"app_url,omitempty"`
	Data           map[string]string                     `json:"data,omitempty"`
	Buttons        []RenderTemplateDtoContentButtonsItem `json:"buttons,omitempty"`
	BigPicture     *string                               `json:"big_picture,omitempty"`
	IOSAttachments map[string]string                     `json:"ios_attachments,omitempty"`
	IOSSound       *string                               `json:"ios_sound,omitempty"`
	// Değerler: "None", "SetTo", "Increase".
	IOSBadgeType  *string `json:"ios_badge_type,omitempty"`
	IOSBadgeCount *int64  `json:"ios_badge_count,omitempty"`
	// Değerler: "passive", "active", "time-sensitive", "critical".
	IOSInterruptionLevel      *string        `json:"ios_interruption_level,omitempty"`
	IOSCategory               *string        `json:"ios_category,omitempty"`
	AndroidChannelID          *string        `json:"android_channel_id,omitempty"`
	AndroidGroup              *string        `json:"android_group,omitempty"`
	AndroidSound              *string        `json:"android_sound,omitempty"`
	ChromeWebIcon             *string        `json:"chrome_web_icon,omitempty"`
	ChromeWebImage            *string        `json:"chrome_web_image,omitempty"`
	ChromeWebBadge            *string        `json:"chrome_web_badge,omitempty"`
	WebPushTopic              *string        `json:"web_push_topic,omitempty"`
	ContentAvailable          *bool          `json:"content_available,omitempty"`
	MutableContent            *bool          `json:"mutable_content,omitempty"`
	CustomData                map[string]any `json:"custom_data,omitempty"`
	EmailSubject              *string        `json:"email_subject,omitempty"`
	EmailBody                 *string        `json:"email_body,omitempty"`
	EmailMJML                 *string        `json:"email_mjml,omitempty"`
	EmailPreheader            *string        `json:"email_preheader,omitempty"`
	EmailFromName             *string        `json:"email_from_name,omitempty"`
	EmailFromAddress          *string        `json:"email_from_address,omitempty"`
	EmailReplyToAddress       *string        `json:"email_reply_to_address,omitempty"`
	DisableEmailClickTracking *bool          `json:"disable_email_click_tracking,omitempty"`
	SMSFrom                   *string        `json:"sms_from,omitempty"`
	DisableSMSLinkShortening  *bool          `json:"disable_sms_link_shortening,omitempty"`
	AddToInbox                *bool          `json:"add_to_inbox,omitempty"`
	InboxTTL                  *int64         `json:"inbox_ttl,omitempty"`
	ThreadID                  *string        `json:"thread_id,omitempty"`
	IOSRelevanceScore         *float64       `json:"ios_relevance_score,omitempty"`
	TargetContentIdentifier   *string        `json:"target_content_identifier,omitempty"`
	SmallIcon                 *string        `json:"small_icon,omitempty"`
	LargeIcon                 *string        `json:"large_icon,omitempty"`
	AndroidAccentColor        *string        `json:"android_accent_color,omitempty"`
	// Değerler: 1, 0, -1.
	AndroidVisibility   *int64                                       `json:"android_visibility,omitempty"`
	WebButtons          []RenderTemplateDtoContentWebButtonsItem     `json:"web_buttons,omitempty"`
	RequireInteraction  *bool                                        `json:"require_interaction,omitempty"`
	SMSMediaUrls        []string                                     `json:"sms_media_urls,omitempty"`
	WhatsappTemplate    *RenderTemplateDtoContentWhatsappTemplate    `json:"whatsapp_template,omitempty"`
	WhatsappMedia       *RenderTemplateDtoContentWhatsappMedia       `json:"whatsapp_media,omitempty"`
	WhatsappInteractive *RenderTemplateDtoContentWhatsappInteractive `json:"whatsapp_interactive,omitempty"`
	RcsCard             *RenderTemplateDtoContentRcsCard             `json:"rcs_card,omitempty"`
	RcsSuggestions      []any                                        `json:"rcs_suggestions,omitempty"`
	RcsContentSid       *string                                      `json:"rcs_content_sid,omitempty"`
	RcsContentVariables map[string]string                            `json:"rcs_content_variables,omitempty"`
	RcsFallbackText     map[string]string                            `json:"rcs_fallback_text,omitempty"`
}

// RenderTemplateDtoContentButtonsItem: RenderTemplateDtoContentButtonsItem şeması.
type RenderTemplateDtoContentButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string  `json:"text"`
	Icon *string `json:"icon,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// RenderTemplateDtoContentRcsCard: RenderTemplateDtoContentRcsCard şeması.
type RenderTemplateDtoContentRcsCard struct {
	Title       map[string]string `json:"title,omitempty"`
	Description map[string]string `json:"description,omitempty"`
	MediaURL    *string           `json:"media_url,omitempty"`
	// Değerler: "SHORT", "MEDIUM", "TALL".
	MediaHeight *string `json:"media_height,omitempty"`
	// Değerler: "VERTICAL", "HORIZONTAL".
	Orientation *string `json:"orientation,omitempty"`
}

// RenderTemplateDtoContentWebButtonsItem: RenderTemplateDtoContentWebButtonsItem şeması.
type RenderTemplateDtoContentWebButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string  `json:"text"`
	Icon *string `json:"icon,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// RenderTemplateDtoContentWhatsappInteractive: RenderTemplateDtoContentWhatsappInteractive şeması.
type RenderTemplateDtoContentWhatsappInteractive struct {
	// Zorunlu. Değerler: "button", "list".
	Type    string                                                   `json:"type"`
	Header  *string                                                  `json:"header,omitempty"`
	Footer  *string                                                  `json:"footer,omitempty"`
	Buttons []RenderTemplateDtoContentWhatsappInteractiveButtonsItem `json:"buttons,omitempty"`
	List    *RenderTemplateDtoContentWhatsappInteractiveList         `json:"list,omitempty"`
}

// RenderTemplateDtoContentWhatsappInteractiveButtonsItem: RenderTemplateDtoContentWhatsappInteractiveButtonsItem şeması.
type RenderTemplateDtoContentWhatsappInteractiveButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title string `json:"title"`
}

// RenderTemplateDtoContentWhatsappInteractiveList: RenderTemplateDtoContentWhatsappInteractiveList şeması.
type RenderTemplateDtoContentWhatsappInteractiveList struct {
	// Zorunlu.
	Button string `json:"button"`
	// Zorunlu.
	Sections []RenderTemplateDtoContentWhatsappInteractiveListSectionsItem `json:"sections"`
}

// RenderTemplateDtoContentWhatsappInteractiveListSectionsItem: RenderTemplateDtoContentWhatsappInteractiveListSectionsItem şeması.
type RenderTemplateDtoContentWhatsappInteractiveListSectionsItem struct {
	Title *string `json:"title,omitempty"`
	// Zorunlu.
	Rows []RenderTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem `json:"rows"`
}

// RenderTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem: RenderTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem şeması.
type RenderTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

// RenderTemplateDtoContentWhatsappMedia: RenderTemplateDtoContentWhatsappMedia şeması.
type RenderTemplateDtoContentWhatsappMedia struct {
	// Zorunlu. Değerler: "image", "video", "document", "audio".
	Type string `json:"type"`
	// Zorunlu.
	URL      string  `json:"url"`
	Filename *string `json:"filename,omitempty"`
}

// RenderTemplateDtoContentWhatsappTemplate: RenderTemplateDtoContentWhatsappTemplate şeması.
type RenderTemplateDtoContentWhatsappTemplate struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Language     string                                                     `json:"language"`
	Header       any                                                        `json:"header,omitempty"`
	BodyParams   []string                                                   `json:"body_params,omitempty"`
	ButtonParams []RenderTemplateDtoContentWhatsappTemplateButtonParamsItem `json:"button_params,omitempty"`
}

// RenderTemplateDtoContentWhatsappTemplateButtonParamsItem: RenderTemplateDtoContentWhatsappTemplateButtonParamsItem şeması.
type RenderTemplateDtoContentWhatsappTemplateButtonParamsItem struct {
	// Zorunlu.
	Index int64 `json:"index"`
	// Zorunlu. Değerler: "url", "quick_reply", "copy_code".
	SubType string `json:"sub_type"`
	// Zorunlu.
	Value string `json:"value"`
}

// RenderTemplateDtoUser: RenderTemplateDtoUser şeması.
type RenderTemplateDtoUser struct {
	ExternalID *string           `json:"external_id,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
	Properties map[string]any    `json:"properties,omitempty"`
}

// RetentionReport: RetentionReport şeması.
type RetentionReport struct {
	// Zorunlu.
	Weeks int64 `json:"weeks"`
	// Zorunlu.
	Cohorts []RetentionReportCohortsItem `json:"cohorts"`
	// Zorunlu.
	Average []float64 `json:"average"`
}

// RetentionReportCohortsItem: RetentionReportCohortsItem şeması.
type RetentionReportCohortsItem struct {
	// Zorunlu.
	Week string `json:"week"`
	// Zorunlu.
	Users int64 `json:"users"`
	// Zorunlu.
	Retention []float64 `json:"retention"`
}

// ScheduleCampaignDto: ScheduleCampaignDto şeması.
type ScheduleCampaignDto struct {
	SendAt     *string                        `json:"send_at,omitempty"`
	Recurrence *ScheduleCampaignDtoRecurrence `json:"recurrence,omitempty"`
}

// ScheduleCampaignDtoRecurrence: ScheduleCampaignDtoRecurrence şeması.
type ScheduleCampaignDtoRecurrence struct {
	// Zorunlu.
	Cron     string  `json:"cron"`
	Until    *string `json:"until,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// Segment: Segment şeması.
type Segment struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	Definition map[string]any `json:"definition"`
	// Zorunlu.
	IsDynamic bool `json:"is_dynamic"`
	// Zorunlu.
	EstimatedCount *int64 `json:"estimated_count"`
	// Zorunlu.
	EstimatedAt *string `json:"estimated_at"`
	// Zorunlu.
	CreatedBy *string `json:"created_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// SegmentCampaigns: SegmentCampaigns şeması.
type SegmentCampaigns struct {
	// Zorunlu.
	Data []SegmentCampaignsDataItem `json:"data"`
}

// SegmentCampaignsDataItem: SegmentCampaignsDataItem şeması.
type SegmentCampaignsDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Status string `json:"status"`
	// Zorunlu.
	Type string `json:"type"`
	// Zorunlu. Değerler: "include", "exclude".
	Usage string `json:"usage"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// SegmentEstimate: SegmentEstimate şeması.
type SegmentEstimate struct {
	// Zorunlu.
	SegmentID *string `json:"segment_id"`
	// Zorunlu.
	Subscriptions int64 `json:"subscriptions"`
	// Zorunlu.
	Users int64 `json:"users"`
	// Zorunlu.
	EstimatedAt string `json:"estimated_at"`
	// Zorunlu.
	Cached bool `json:"cached"`
	// Zorunlu.
	Sampled bool `json:"sampled"`
}

// SegmentPage: SegmentPage şeması.
type SegmentPage struct {
	// Zorunlu.
	Data []Segment `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// SegmentPreview: SegmentPreview şeması.
type SegmentPreview struct {
	// Zorunlu.
	Subscriptions int64 `json:"subscriptions"`
	// Zorunlu.
	Users int64 `json:"users"`
	// Zorunlu.
	Sampled bool `json:"sampled"`
	// Zorunlu.
	Sample []SegmentPreviewSampleItem `json:"sample"`
}

// SegmentPreviewSampleItem: SegmentPreviewSampleItem şeması.
type SegmentPreviewSampleItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu.
	Tags map[string]string `json:"tags"`
	// Zorunlu.
	LastActiveAt string `json:"last_active_at"`
}

// SegmentsCreateBody: SegmentsCreateBody şeması.
type SegmentsCreateBody struct {
	// Zorunlu.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Zorunlu. Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
	Definition SegmentsCreateBodyDefinition `json:"definition"`
	IsDynamic  *bool                        `json:"is_dynamic,omitempty"`
}

// SegmentsCreateBodyDefinition: Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
type SegmentsCreateBodyDefinition struct {
	All  []map[string]any `json:"all,omitempty"`
	Any  []map[string]any `json:"any,omitempty"`
	None []map[string]any `json:"none,omitempty"`
}

// SegmentsPreviewBody: SegmentsPreviewBody şeması.
type SegmentsPreviewBody struct {
	// Zorunlu. Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
	Definition SegmentsPreviewBodyDefinition `json:"definition"`
	Limit      *int64                        `json:"limit,omitempty"`
}

// SegmentsPreviewBodyDefinition: Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
type SegmentsPreviewBodyDefinition struct {
	All  []map[string]any `json:"all,omitempty"`
	Any  []map[string]any `json:"any,omitempty"`
	None []map[string]any `json:"none,omitempty"`
}

// SegmentsUpdateBody: SegmentsUpdateBody şeması.
type SegmentsUpdateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
	Definition *SegmentsUpdateBodyDefinition `json:"definition,omitempty"`
	IsDynamic  *bool                         `json:"is_dynamic,omitempty"`
}

// SegmentsUpdateBodyDefinition: Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
type SegmentsUpdateBodyDefinition struct {
	All  []map[string]any `json:"all,omitempty"`
	Any  []map[string]any `json:"any,omitempty"`
	None []map[string]any `json:"none,omitempty"`
}

// SelectWinnerDto: SelectWinnerDto şeması.
type SelectWinnerDto struct {
	// Zorunlu.
	VariantID string `json:"variant_id"`
}

// SetOwnerDto: SetOwnerDto şeması.
type SetOwnerDto struct {
	// Zorunlu.
	ExternalID string `json:"external_id"`
}

// SetPreferencesDto: SetPreferencesDto şeması.
type SetPreferencesDto struct {
	// Zorunlu.
	Preferences []SetPreferencesDtoPreferencesItem `json:"preferences"`
}

// SetPreferencesDtoPreferencesItem: SetPreferencesDtoPreferencesItem şeması.
type SetPreferencesDtoPreferencesItem struct {
	// Zorunlu.
	Category string `json:"category"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel *string `json:"channel,omitempty"`
	// Zorunlu.
	OptedIn bool `json:"opted_in"`
}

// SMSKeyword: SmsKeyword şeması.
type SMSKeyword struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Keyword string `json:"keyword"`
	// Zorunlu. Değerler: "org", "app".
	Scope string `json:"scope"`
	// Zorunlu.
	AppID *string `json:"app_id"`
	// Zorunlu. Değerler: "sms", "whatsapp", "rcs".
	Channels []string `json:"channels"`
	// Zorunlu.
	Actions map[string]any `json:"actions"`
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// SMSKeywordList: SmsKeywordList şeması.
type SMSKeywordList struct {
	// Zorunlu.
	Data []SMSKeyword `json:"data"`
	// Zorunlu.
	Reserved []string `json:"reserved"`
}

// StartLiveActivityDto: StartLiveActivityDto şeması.
type StartLiveActivityDto struct {
	ActivityID *string        `json:"activity_id,omitempty"`
	Attributes map[string]any `json:"attributes,omitempty"`
	// Zorunlu.
	Target StartLiveActivityDtoTarget `json:"target"`
	// Zorunlu.
	ContentState   map[string]any             `json:"content_state"`
	Alert          *StartLiveActivityDtoAlert `json:"alert,omitempty"`
	StaleDate      *int64                     `json:"stale_date,omitempty"`
	DismissalDate  *int64                     `json:"dismissal_date,omitempty"`
	RelevanceScore *float64                   `json:"relevance_score,omitempty"`
	// Değerler: 5, 10.
	Priority *int64  `json:"priority,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// StartLiveActivityDtoAlert: StartLiveActivityDtoAlert şeması.
type StartLiveActivityDtoAlert struct {
	Title map[string]string `json:"title,omitempty"`
	Body  map[string]string `json:"body,omitempty"`
	Sound *string           `json:"sound,omitempty"`
}

// StartLiveActivityDtoTarget: StartLiveActivityDtoTarget şeması.
type StartLiveActivityDtoTarget struct {
	IncludeExternalIds     []string `json:"include_external_ids,omitempty"`
	IncludeSubscriptionIds []string `json:"include_subscription_ids,omitempty"`
	IncludedSegments       []string `json:"included_segments,omitempty"`
	All                    *bool    `json:"all,omitempty"`
}

// Subscription: Subscription şeması.
type Subscription struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	UserID *string `json:"user_id"`
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu.
	Token string `json:"token"`
	// Zorunlu. Değerler: "subscribed", "unsubscribed", "invalid", "pending".
	Status string `json:"status"`
	// Zorunlu. Değerler: "provider_rejected", "user_opt_out", "api", "uninstall".
	UnsubscribeReason *string `json:"unsubscribe_reason"`
	// Zorunlu.
	Device SubscriptionDevice `json:"device"`
	// Zorunlu.
	TestDevice bool `json:"test_device"`
	// Zorunlu.
	TestDeviceName *string `json:"test_device_name"`
	// Zorunlu.
	LastActiveAt string `json:"last_active_at"`
	// Zorunlu.
	LastSentAt *string `json:"last_sent_at"`
	// Zorunlu.
	LastDeliveredAt *string `json:"last_delivered_at"`
	// Zorunlu.
	OptedInAt *string `json:"opted_in_at"`
	// Zorunlu.
	OptInSource *string `json:"opt_in_source"`
	// Zorunlu.
	UnsubscribedAt *string `json:"unsubscribed_at"`
	// Zorunlu.
	Location *SubscriptionLocation `json:"location"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// SubscriptionDevice: SubscriptionDevice şeması.
type SubscriptionDevice struct {
	Model       *string `json:"model,omitempty"`
	Os          *string `json:"os,omitempty"`
	OsVersion   *string `json:"os_version,omitempty"`
	AppVersion  *string `json:"app_version,omitempty"`
	SDKVersion  *string `json:"sdk_version,omitempty"`
	SDKPlatform *string `json:"sdk_platform,omitempty"`
	Carrier     *string `json:"carrier,omitempty"`
	Locale      *string `json:"locale,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
	PushCapable *bool   `json:"push_capable,omitempty"`
	// Değerler: "sandbox", "production".
	PushEnvironment *string `json:"push_environment,omitempty"`
}

// SubscriptionLocation: SubscriptionLocation şeması.
type SubscriptionLocation struct {
	// Zorunlu.
	Lat float64 `json:"lat"`
	// Zorunlu.
	Lng float64 `json:"lng"`
	// Zorunlu.
	AccuracyM *float64 `json:"accuracy_m"`
	// Zorunlu.
	Geohash *string `json:"geohash"`
	// Zorunlu.
	UpdatedAt *string `json:"updated_at"`
	// Zorunlu.
	ConsentAt *string `json:"consent_at"`
}

// SubscriptionsReport: SubscriptionsReport şeması.
type SubscriptionsReport struct {
	// Zorunlu.
	From string `json:"from"`
	// Zorunlu.
	To string `json:"to"`
	// Zorunlu.
	Totals SubscriptionsReportTotals `json:"totals"`
	// Zorunlu.
	Previous *SubscriptionsReportPrevious `json:"previous"`
	// Zorunlu.
	Rows []SubscriptionsReportRowsItem `json:"rows"`
	// Zorunlu.
	ByChannel []SubscriptionsReportByChannelItem `json:"by_channel"`
}

// SubscriptionsReportByChannelItem: SubscriptionsReportByChannelItem şeması.
type SubscriptionsReportByChannelItem struct {
	// Zorunlu.
	Channel string `json:"channel"`
	// Zorunlu.
	Active int64 `json:"active"`
	// Zorunlu.
	New int64 `json:"new"`
	// Zorunlu.
	Unsubscribed int64 `json:"unsubscribed"`
}

// SubscriptionsReportPrevious: SubscriptionsReportPrevious şeması.
type SubscriptionsReportPrevious struct {
	// Zorunlu.
	Active int64 `json:"active"`
	// Zorunlu.
	New int64 `json:"new"`
	// Zorunlu.
	Unsubscribed int64 `json:"unsubscribed"`
	// Zorunlu.
	OptOuts int64 `json:"opt_outs"`
	// Zorunlu.
	Net int64 `json:"net"`
}

// SubscriptionsReportRowsItem: SubscriptionsReportRowsItem şeması.
type SubscriptionsReportRowsItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	New int64 `json:"new"`
	// Zorunlu.
	Unsubscribed int64 `json:"unsubscribed"`
	// Zorunlu.
	OptOuts int64 `json:"opt_outs"`
	// Zorunlu.
	Net int64 `json:"net"`
}

// SubscriptionsReportTotals: SubscriptionsReportTotals şeması.
type SubscriptionsReportTotals struct {
	// Zorunlu.
	Active int64 `json:"active"`
	// Zorunlu.
	New int64 `json:"new"`
	// Zorunlu.
	Unsubscribed int64 `json:"unsubscribed"`
	// Zorunlu.
	OptOuts int64 `json:"opt_outs"`
	// Zorunlu.
	Net int64 `json:"net"`
}

// Suppression: Suppression şeması.
type Suppression struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu. Değerler: "email", "sms", "whatsapp", "rcs".
	Channel string `json:"channel"`
	// Zorunlu.
	ValueMasked string `json:"value_masked"`
	// Zorunlu. Değerler: "bounce", "complaint", "manual", "unsubscribe_link", "stop".
	Reason string `json:"reason"`
	// Zorunlu.
	Source string `json:"source"`
	// Zorunlu.
	Detail map[string]any `json:"detail"`
	// Zorunlu.
	ExpiresAt *string `json:"expires_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// SuppressionPage: SuppressionPage şeması.
type SuppressionPage struct {
	// Zorunlu.
	Data []Suppression `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// Survey: Survey şeması.
type Survey struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "draft", "active", "closed".
	Status string `json:"status"`
	// Zorunlu.
	Questions []SurveyQuestionsItem `json:"questions"`
	// Zorunlu.
	Settings SurveySettings `json:"settings"`
	// Zorunlu.
	InAppMessageID *string `json:"in_app_message_id"`
	// Zorunlu.
	Stats SurveyStats `json:"stats"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// SurveyPage: SurveyPage şeması.
type SurveyPage struct {
	// Zorunlu.
	Data []Survey `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// SurveyQuestionsItem: SurveyQuestionsItem şeması.
type SurveyQuestionsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "nps", "csat", "single_choice", "multiple_choice", "text".
	Type string `json:"type"`
	// Zorunlu.
	Text        map[string]string                `json:"text"`
	Required    *bool                            `json:"required,omitempty"`
	Options     []SurveyQuestionsItemOptionsItem `json:"options,omitempty"`
	LowLabel    map[string]string                `json:"low_label,omitempty"`
	HighLabel   map[string]string                `json:"high_label,omitempty"`
	Placeholder map[string]string                `json:"placeholder,omitempty"`
}

// SurveyQuestionsItemOptionsItem: SurveyQuestionsItemOptionsItem şeması.
type SurveyQuestionsItemOptionsItem struct {
	// Zorunlu.
	Value string `json:"value"`
	// Zorunlu.
	Label map[string]string `json:"label"`
}

// SurveyReport: SurveyReport şeması.
type SurveyReport struct {
	// Zorunlu.
	SurveyID string `json:"survey_id"`
	// Zorunlu.
	Days int64 `json:"days"`
	// Zorunlu.
	Responses int64 `json:"responses"`
	// Zorunlu.
	Nps SurveyReportNps `json:"nps"`
	// Zorunlu.
	Csat SurveyReportCsat `json:"csat"`
	// Zorunlu.
	Trend []SurveyReportTrendItem `json:"trend"`
	// Zorunlu.
	Channels map[string]int64 `json:"channels"`
	// Zorunlu.
	Questions []SurveyReportQuestionsItem `json:"questions"`
}

// SurveyReportCsat: SurveyReportCsat şeması.
type SurveyReportCsat struct {
	// Zorunlu.
	Responses int64 `json:"responses"`
	// Zorunlu.
	Average *float64 `json:"average"`
	// Zorunlu.
	Distribution map[string]int64 `json:"distribution"`
}

// SurveyReportNps: SurveyReportNps şeması.
type SurveyReportNps struct {
	// Zorunlu.
	Responses int64 `json:"responses"`
	// Zorunlu.
	Promoters int64 `json:"promoters"`
	// Zorunlu.
	Passives int64 `json:"passives"`
	// Zorunlu.
	Detractors int64 `json:"detractors"`
	// Zorunlu.
	Nps *float64 `json:"nps"`
}

// SurveyReportQuestionsItem: SurveyReportQuestionsItem şeması.
type SurveyReportQuestionsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu. Değerler: "nps", "csat", "single_choice", "multiple_choice", "text".
	Type string `json:"type"`
	// Zorunlu.
	Counts map[string]int64 `json:"counts"`
}

// SurveyReportTrendItem: SurveyReportTrendItem şeması.
type SurveyReportTrendItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	Responses int64 `json:"responses"`
	// Zorunlu.
	Promoters int64 `json:"promoters"`
	// Zorunlu.
	Passives int64 `json:"passives"`
	// Zorunlu.
	Detractors int64 `json:"detractors"`
	// Zorunlu.
	Nps *float64 `json:"nps"`
}

// SurveyResponse: SurveyResponse şeması.
type SurveyResponse struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	SurveyID string `json:"survey_id"`
	// Zorunlu.
	UserID *string `json:"user_id"`
	// Zorunlu.
	VisitorID *string `json:"visitor_id"`
	// Zorunlu. Değerler: "in_app", "web", "email", "api", "sdk".
	Channel string `json:"channel"`
	// Zorunlu.
	Answers map[string]any `json:"answers"`
	// Zorunlu.
	NpsScore *int64 `json:"nps_score"`
	// Zorunlu. Değerler: "promoter", "passive", "detractor".
	NpsCategory *string `json:"nps_category"`
	// Zorunlu.
	CsatScore *int64 `json:"csat_score"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// SurveyResponsePage: SurveyResponsePage şeması.
type SurveyResponsePage struct {
	// Zorunlu.
	Data []SurveyResponse `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// SurveyResponseResult: SurveyResponseResult şeması.
type SurveyResponseResult struct {
	// Zorunlu.
	Ok bool `json:"ok"`
	// Zorunlu.
	ResponseID *string `json:"response_id"`
	// Zorunlu.
	Duplicate bool `json:"duplicate"`
	// Zorunlu.
	ThankYou string `json:"thank_you"`
}

// SurveySettings: SurveySettings şeması.
type SurveySettings struct {
	ThankYou        map[string]string `json:"thank_you,omitempty"`
	SubmitLabel     map[string]string `json:"submit_label,omitempty"`
	WriteProperties *bool             `json:"write_properties,omitempty"`
	AllowMultiple   *bool             `json:"allow_multiple,omitempty"`
	DefaultLanguage *string           `json:"default_language,omitempty"`
}

// SurveyStats: SurveyStats şeması.
type SurveyStats struct {
	// Zorunlu.
	Responses int64 `json:"responses"`
	// Zorunlu.
	Nps *float64 `json:"nps"`
	// Zorunlu.
	LastResponseAt *string `json:"last_response_at"`
}

// SurveyStatusDto: SurveyStatusDto şeması.
type SurveyStatusDto struct {
	// Zorunlu. Değerler: "draft", "active", "closed".
	Status string `json:"status"`
}

// SyncAdAudienceDto: SyncAdAudienceDto şeması.
type SyncAdAudienceDto struct {
	Full *bool `json:"full,omitempty"`
}

// Template: Template şeması.
type Template struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel *string `json:"channel"`
	// Zorunlu.
	CurrentVersion int64 `json:"current_version"`
	// Zorunlu.
	Tags []string `json:"tags"`
	// Zorunlu.
	Content map[string]any `json:"content"`
	// Zorunlu.
	CreatedBy *string `json:"created_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// TemplatePage: TemplatePage şeması.
type TemplatePage struct {
	// Zorunlu.
	Data []Template `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// TemplateStats: TemplateStats şeması.
type TemplateStats struct {
	// Zorunlu.
	TemplateID string `json:"template_id"`
	// Zorunlu.
	Messages int64 `json:"messages"`
	// Zorunlu.
	Campaigns int64 `json:"campaigns"`
	// Zorunlu.
	Targeted int64 `json:"targeted"`
	// Zorunlu.
	Sent int64 `json:"sent"`
	// Zorunlu.
	Delivered int64 `json:"delivered"`
	// Zorunlu.
	Opened int64 `json:"opened"`
	// Zorunlu.
	Clicked int64 `json:"clicked"`
	// Zorunlu.
	Converted int64 `json:"converted"`
	// Zorunlu.
	OpenRate float64 `json:"open_rate"`
	// Zorunlu.
	ClickRate float64 `json:"click_rate"`
	// Zorunlu.
	ConversionRate float64 `json:"conversion_rate"`
	// Zorunlu.
	LastUsedAt *string `json:"last_used_at"`
}

// TemplateVersionList: TemplateVersionList şeması.
type TemplateVersionList struct {
	// Zorunlu.
	Data []TemplateVersionListDataItem `json:"data"`
}

// TemplateVersionListDataItem: TemplateVersionListDataItem şeması.
type TemplateVersionListDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	Content map[string]any `json:"content"`
	// Zorunlu.
	CreatedBy *string `json:"created_by"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// TestDataFeedDto: TestDataFeedDto şeması.
type TestDataFeedDto struct {
	ExternalID *string `json:"external_id,omitempty"`
	UserID     *string `json:"user_id,omitempty"`
}

// TriggerRunDto: TriggerRunDto şeması.
type TriggerRunDto struct {
	ExternalID *string             `json:"external_id,omitempty"`
	UserID     *string             `json:"user_id,omitempty"`
	Alias      *TriggerRunDtoAlias `json:"alias,omitempty"`
	Context    map[string]any      `json:"context,omitempty"`
}

// TriggerRunDtoAlias: TriggerRunDtoAlias şeması.
type TriggerRunDtoAlias struct {
	// Zorunlu.
	Label string `json:"label"`
	// Zorunlu.
	Value string `json:"value"`
}

// UpdateAdAudienceDto: UpdateAdAudienceDto şeması.
type UpdateAdAudienceDto struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	SegmentID   *string `json:"segment_id,omitempty"`
	// Değerler: "email", "phone", "maid".
	Identifiers         []string `json:"identifiers,omitempty"`
	SyncIntervalMinutes *int64   `json:"sync_interval_minutes,omitempty"`
	// Değerler: "active", "paused".
	Status *string `json:"status,omitempty"`
}

// UpdateAdCampaignDto: UpdateAdCampaignDto şeması.
type UpdateAdCampaignDto struct {
	// Değerler: "active", "paused".
	Status      *string `json:"status,omitempty"`
	DailyBudget any     `json:"daily_budget,omitempty"`
}

// UpdateAPIKeyDto: UpdateApiKeyDto şeması.
type UpdateAPIKeyDto struct {
	Name *string `json:"name,omitempty"`
	// Değerler: "admin", "notifications:write", "users:read", "users:write", "events:write", "segments:read", "analytics:read", "mcp:write".
	Scopes      []string `json:"scopes,omitempty"`
	IPAllowlist []string `json:"ip_allowlist,omitempty"`
	ExpiresAt   *string  `json:"expires_at,omitempty"`
}

// UpdateAppDto: UpdateAppDto şeması.
type UpdateAppDto struct {
	Name            *string                    `json:"name,omitempty"`
	PlatformFlags   *UpdateAppDtoPlatformFlags `json:"platform_flags,omitempty"`
	IconURL         *string                    `json:"icon_url,omitempty"`
	DefaultLanguage *string                    `json:"default_language,omitempty"`
	Timezone        *string                    `json:"timezone,omitempty"`
	Settings        *UpdateAppDtoSettings      `json:"settings,omitempty"`
	// Değerler: "hmac", "jwt", "both".
	IdentityVerificationMode *string `json:"identity_verification_mode,omitempty"`
}

// UpdateAppDtoPlatformFlags: UpdateAppDtoPlatformFlags şeması.
type UpdateAppDtoPlatformFlags struct {
	IOS     *bool `json:"ios,omitempty"`
	Android *bool `json:"android,omitempty"`
	Web     *bool `json:"web,omitempty"`
	Email   *bool `json:"email,omitempty"`
	SMS     *bool `json:"sms,omitempty"`
}

// UpdateAppDtoSettings: UpdateAppDtoSettings şeması.
type UpdateAppDtoSettings struct {
	QuietHours                   *UpdateAppDtoSettingsQuietHours           `json:"quietHours,omitempty"`
	FrequencyCap                 *UpdateAppDtoSettingsFrequencyCap         `json:"frequencyCap,omitempty"`
	FrequencyCaps                []UpdateAppDtoSettingsFrequencyCapsItem   `json:"frequencyCaps,omitempty"`
	InApp                        *UpdateAppDtoSettingsInApp                `json:"inApp,omitempty"`
	AndroidChannels              []UpdateAppDtoSettingsAndroidChannelsItem `json:"androidChannels,omitempty"`
	SMS                          *UpdateAppDtoSettingsSMS                  `json:"sms,omitempty"`
	PreferenceCenter             *UpdateAppDtoSettingsPreferenceCenter     `json:"preferenceCenter,omitempty"`
	DefaultTTLSeconds            *int64                                    `json:"defaultTtlSeconds,omitempty"`
	IdentityVerificationRequired *bool                                     `json:"identityVerificationRequired,omitempty"`
	AllowedOrigins               []string                                  `json:"allowedOrigins,omitempty"`
	CampaignApprovalRequired     *bool                                     `json:"campaignApprovalRequired,omitempty"`
	WebPush                      *UpdateAppDtoSettingsWebPush              `json:"webPush,omitempty"`
	Onboarding                   *UpdateAppDtoSettingsOnboarding           `json:"onboarding,omitempty"`
	SDKSetup                     *UpdateAppDtoSettingsSDKSetup             `json:"sdkSetup,omitempty"`
}

// UpdateAppDtoSettingsAndroidChannelsItem: UpdateAppDtoSettingsAndroidChannelsItem şeması.
type UpdateAppDtoSettingsAndroidChannelsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Zorunlu. Değerler: "low", "default", "high", "urgent".
	Importance string  `json:"importance"`
	Sound      *string `json:"sound,omitempty"`
	Vibration  *bool   `json:"vibration,omitempty"`
	Group      *string `json:"group,omitempty"`
}

// UpdateAppDtoSettingsFrequencyCap: UpdateAppDtoSettingsFrequencyCap şeması.
type UpdateAppDtoSettingsFrequencyCap struct {
	// Zorunlu.
	Max int64 `json:"max"`
	// Zorunlu.
	WindowHours int64 `json:"windowHours"`
}

// UpdateAppDtoSettingsFrequencyCapsItem: UpdateAppDtoSettingsFrequencyCapsItem şeması.
type UpdateAppDtoSettingsFrequencyCapsItem struct {
	// Zorunlu.
	Max int64 `json:"max"`
	// Zorunlu.
	Window UpdateAppDtoSettingsFrequencyCapsItemWindow `json:"window"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels []string `json:"channels,omitempty"`
}

// UpdateAppDtoSettingsFrequencyCapsItemWindow: UpdateAppDtoSettingsFrequencyCapsItemWindow şeması.
type UpdateAppDtoSettingsFrequencyCapsItemWindow struct {
	// Zorunlu.
	Value int64 `json:"value"`
	// Zorunlu. Değerler: "minutes", "hours", "days", "weeks".
	Unit string `json:"unit"`
}

// UpdateAppDtoSettingsInApp: UpdateAppDtoSettingsInApp şeması.
type UpdateAppDtoSettingsInApp struct {
	MinIntervalSeconds *int64 `json:"minIntervalSeconds,omitempty"`
}

// UpdateAppDtoSettingsOnboarding: UpdateAppDtoSettingsOnboarding şeması.
type UpdateAppDtoSettingsOnboarding struct {
	// Zorunlu.
	Step        int64   `json:"step"`
	Simulation  *bool   `json:"simulation,omitempty"`
	CompletedAt *string `json:"completed_at,omitempty"`
	Version     *int64  `json:"version,omitempty"`
	// Değerler: "ios", "android", "huawei", "web".
	Platforms []string `json:"platforms,omitempty"`
	// Değerler: "native", "flutter", "react-native", "unity", "capacitor", "cordova", "maui".
	Framework *string `json:"framework,omitempty"`
	// Değerler: "code", "wordpress".
	WebIntegration *string `json:"web_integration,omitempty"`
}

// UpdateAppDtoSettingsPreferenceCenter: UpdateAppDtoSettingsPreferenceCenter şeması.
type UpdateAppDtoSettingsPreferenceCenter struct {
	Enabled              *bool                                                     `json:"enabled,omitempty"`
	BrandName            *string                                                   `json:"brandName,omitempty"`
	LogoURL              *string                                                   `json:"logoUrl,omitempty"`
	AccentColor          *string                                                   `json:"accentColor,omitempty"`
	CustomUnsubscribeURL *string                                                   `json:"customUnsubscribeUrl,omitempty"`
	Texts                map[string]UpdateAppDtoSettingsPreferenceCenterTextsValue `json:"texts,omitempty"`
}

// UpdateAppDtoSettingsPreferenceCenterTextsValue: UpdateAppDtoSettingsPreferenceCenterTextsValue şeması.
type UpdateAppDtoSettingsPreferenceCenterTextsValue struct {
	Title          *string `json:"title,omitempty"`
	Description    *string `json:"description,omitempty"`
	UnsubscribeAll *string `json:"unsubscribeAll,omitempty"`
	Saved          *string `json:"saved,omitempty"`
}

// UpdateAppDtoSettingsQuietHours: UpdateAppDtoSettingsQuietHours şeması.
type UpdateAppDtoSettingsQuietHours struct {
	// Zorunlu.
	Start string `json:"start"`
	// Zorunlu.
	End string `json:"end"`
	// Zorunlu. Değerler: "user", "app".
	TimezoneMode string `json:"timezoneMode"`
}

// UpdateAppDtoSettingsSDKSetup: UpdateAppDtoSettingsSdkSetup şeması.
type UpdateAppDtoSettingsSDKSetup struct {
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	Platform *string `json:"platform"`
	// Zorunlu.
	Step string `json:"step"`
	// Zorunlu.
	Furthest               string  `json:"furthest"`
	Variant                *string `json:"variant,omitempty"`
	DetectedSubscriptionID *string `json:"detectedSubscriptionId,omitempty"`
	TestMessageID          *string `json:"testMessageId,omitempty"`
	CompletedAt            *string `json:"completedAt,omitempty"`
}

// UpdateAppDtoSettingsSMS: UpdateAppDtoSettingsSms şeması.
type UpdateAppDtoSettingsSMS struct {
	HelpReply *string `json:"helpReply,omitempty"`
	StopReply *string `json:"stopReply,omitempty"`
	AutoReply *bool   `json:"autoReply,omitempty"`
}

// UpdateAppDtoSettingsWebPush: UpdateAppDtoSettingsWebPush şeması.
type UpdateAppDtoSettingsWebPush struct {
	ServiceWorkerPath   *string                                         `json:"serviceWorkerPath,omitempty"`
	ServiceWorkerScope  *string                                         `json:"serviceWorkerScope,omitempty"`
	Prompt              *UpdateAppDtoSettingsWebPushPrompt              `json:"prompt,omitempty"`
	WelcomeNotification *UpdateAppDtoSettingsWebPushWelcomeNotification `json:"welcomeNotification,omitempty"`
	ClickBehavior       *UpdateAppDtoSettingsWebPushClickBehavior       `json:"clickBehavior,omitempty"`
	CustomLink          *UpdateAppDtoSettingsWebPushCustomLink          `json:"customLink,omitempty"`
	PromptBackoffDays   []int64                                         `json:"promptBackoffDays,omitempty"`
}

// UpdateAppDtoSettingsWebPushClickBehavior: UpdateAppDtoSettingsWebPushClickBehavior şeması.
type UpdateAppDtoSettingsWebPushClickBehavior struct {
	// Zorunlu. Değerler: "exact", "origin".
	Match string `json:"match"`
	// Zorunlu. Değerler: "navigate", "focus".
	Action string `json:"action"`
}

// UpdateAppDtoSettingsWebPushCustomLink: UpdateAppDtoSettingsWebPushCustomLink şeması.
type UpdateAppDtoSettingsWebPushCustomLink struct {
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu. Değerler: "button", "link".
	Style string `json:"style"`
	// Zorunlu. Değerler: "small", "medium", "large".
	Size               string                                                     `json:"size"`
	Color              *UpdateAppDtoSettingsWebPushCustomLinkColor                `json:"color,omitempty"`
	Texts              map[string]UpdateAppDtoSettingsWebPushCustomLinkTextsValue `json:"texts,omitempty"`
	UnsubscribeEnabled *bool                                                      `json:"unsubscribeEnabled,omitempty"`
}

// UpdateAppDtoSettingsWebPushCustomLinkColor: UpdateAppDtoSettingsWebPushCustomLinkColor şeması.
type UpdateAppDtoSettingsWebPushCustomLinkColor struct {
	Button *string `json:"button,omitempty"`
	Text   *string `json:"text,omitempty"`
}

// UpdateAppDtoSettingsWebPushCustomLinkTextsValue: UpdateAppDtoSettingsWebPushCustomLinkTextsValue şeması.
type UpdateAppDtoSettingsWebPushCustomLinkTextsValue struct {
	Subscribe   *string `json:"subscribe,omitempty"`
	Unsubscribe *string `json:"unsubscribe,omitempty"`
	Explanation *string `json:"explanation,omitempty"`
}

// UpdateAppDtoSettingsWebPushPrompt: UpdateAppDtoSettingsWebPushPrompt şeması.
type UpdateAppDtoSettingsWebPushPrompt struct {
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu. Değerler: "slidedown", "native".
	Type string `json:"type"`
	// Değerler: "push", "email", "sms", "email_sms".
	Collect       *string                                                `json:"collect,omitempty"`
	DelaySeconds  *int64                                                 `json:"delaySeconds,omitempty"`
	PageViews     *int64                                                 `json:"pageViews,omitempty"`
	ScrollPercent *int64                                                 `json:"scrollPercent,omitempty"`
	Texts         map[string]UpdateAppDtoSettingsWebPushPromptTextsValue `json:"texts,omitempty"`
	Categories    []UpdateAppDtoSettingsWebPushPromptCategoriesItem      `json:"categories,omitempty"`
	AccentColor   *string                                                `json:"accentColor,omitempty"`
}

// UpdateAppDtoSettingsWebPushPromptCategoriesItem: UpdateAppDtoSettingsWebPushPromptCategoriesItem şeması.
type UpdateAppDtoSettingsWebPushPromptCategoriesItem struct {
	// Zorunlu.
	Tag string `json:"tag"`
	// Zorunlu.
	Label map[string]string `json:"label"`
}

// UpdateAppDtoSettingsWebPushPromptTextsValue: UpdateAppDtoSettingsWebPushPromptTextsValue şeması.
type UpdateAppDtoSettingsWebPushPromptTextsValue struct {
	Title   *string `json:"title,omitempty"`
	Message *string `json:"message,omitempty"`
	Accept  *string `json:"accept,omitempty"`
	Cancel  *string `json:"cancel,omitempty"`
}

// UpdateAppDtoSettingsWebPushWelcomeNotification: UpdateAppDtoSettingsWebPushWelcomeNotification şeması.
type UpdateAppDtoSettingsWebPushWelcomeNotification struct {
	// Zorunlu.
	Enabled bool                                                                `json:"enabled"`
	URL     *string                                                             `json:"url,omitempty"`
	Texts   map[string]UpdateAppDtoSettingsWebPushWelcomeNotificationTextsValue `json:"texts,omitempty"`
}

// UpdateAppDtoSettingsWebPushWelcomeNotificationTextsValue: UpdateAppDtoSettingsWebPushWelcomeNotificationTextsValue şeması.
type UpdateAppDtoSettingsWebPushWelcomeNotificationTextsValue struct {
	Title   *string `json:"title,omitempty"`
	Message *string `json:"message,omitempty"`
}

// UpdateBudgetAlertRuleDto: UpdateBudgetAlertRuleDto şeması.
type UpdateBudgetAlertRuleDto struct {
	Name               *string  `json:"name,omitempty"`
	AdAccountID        *string  `json:"ad_account_id,omitempty"`
	CampaignExternalID *string  `json:"campaign_external_id,omitempty"`
	Threshold          *float64 `json:"threshold,omitempty"`
	WindowDays         *int64   `json:"window_days,omitempty"`
	Recipients         []string `json:"recipients,omitempty"`
	Enabled            *bool    `json:"enabled,omitempty"`
}

// UpdateCategoryDto: UpdateCategoryDto şeması.
type UpdateCategoryDto struct {
	Name        map[string]string `json:"name,omitempty"`
	Description map[string]string `json:"description,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels       []string `json:"channels,omitempty"`
	DefaultOptedIn *bool    `json:"default_opted_in,omitempty"`
	Position       *int64   `json:"position,omitempty"`
}

// UpdateContentTableDto: UpdateContentTableDto şeması.
type UpdateContentTableDto struct {
	Description *string `json:"description,omitempty"`
}

// UpdateDataFeedDto: UpdateDataFeedDto şeması.
type UpdateDataFeedDto struct {
	Description *string `json:"description,omitempty"`
	URLTemplate *string `json:"url_template,omitempty"`
	// Değerler: "GET", "POST".
	Method          *string           `json:"method,omitempty"`
	BodyTemplate    *string           `json:"body_template,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	SecretHeaders   map[string]string `json:"secret_headers,omitempty"`
	CacheTTLSeconds *int64            `json:"cache_ttl_seconds,omitempty"`
	TimeoutMs       *int64            `json:"timeout_ms,omitempty"`
	Enabled         *bool             `json:"enabled,omitempty"`
}

// UpdateInAppDto: UpdateInAppDto şeması.
type UpdateInAppDto struct {
	Name          *string                      `json:"name,omitempty"`
	Triggers      []any                        `json:"triggers,omitempty"`
	TriggerGroups [][]any                      `json:"trigger_groups,omitempty"`
	Variants      []UpdateInAppDtoVariantsItem `json:"variants,omitempty"`
	AbTest        *UpdateInAppDtoAbTest        `json:"ab_test,omitempty"`
	Display       *UpdateInAppDtoDisplay       `json:"display,omitempty"`
	Audience      *UpdateInAppDtoAudience      `json:"audience,omitempty"`
	Frequency     *UpdateInAppDtoFrequency     `json:"frequency,omitempty"`
	JourneyOnly   *bool                        `json:"journey_only,omitempty"`
	Priority      *int64                       `json:"priority,omitempty"`
	StartAt       *string                      `json:"start_at,omitempty"`
	EndAt         *string                      `json:"end_at,omitempty"`
}

// UpdateInAppDtoAbTest: UpdateInAppDtoAbTest şeması.
type UpdateInAppDtoAbTest struct {
	Winner *UpdateInAppDtoAbTestWinner `json:"winner,omitempty"`
}

// UpdateInAppDtoAbTestWinner: UpdateInAppDtoAbTestWinner şeması.
type UpdateInAppDtoAbTestWinner struct {
	// Zorunlu. Değerler: "manual", "auto".
	Mode           string   `json:"mode"`
	MinImpressions *int64   `json:"min_impressions,omitempty"`
	Confidence     *float64 `json:"confidence,omitempty"`
}

// UpdateInAppDtoAudience: UpdateInAppDtoAudience şeması.
type UpdateInAppDtoAudience struct {
	IncludedSegments []string `json:"included_segments,omitempty"`
	ExcludedSegments []string `json:"excluded_segments,omitempty"`
}

// UpdateInAppDtoDisplay: UpdateInAppDtoDisplay şeması.
type UpdateInAppDtoDisplay struct {
	// Zorunlu. Değerler: "modal", "banner_top", "banner_bottom", "fullscreen".
	Layout             string  `json:"layout"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	// Değerler: "blocks", "html".
	ContentType *string           `json:"content_type,omitempty"`
	HTML        map[string]string `json:"html,omitempty"`
	HTMLHeight  *int64            `json:"html_height,omitempty"`
	// Zorunlu.
	Blocks []any `json:"blocks"`
}

// UpdateInAppDtoFrequency: UpdateInAppDtoFrequency şeması.
type UpdateInAppDtoFrequency struct {
	MaxImpressions       *int64 `json:"max_impressions,omitempty"`
	MinIntervalSeconds   *int64 `json:"min_interval_seconds,omitempty"`
	StopAfterInteraction *bool  `json:"stop_after_interaction,omitempty"`
}

// UpdateInAppDtoVariantsItem: UpdateInAppDtoVariantsItem şeması.
type UpdateInAppDtoVariantsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	Display UpdateInAppDtoVariantsItemDisplay `json:"display"`
}

// UpdateInAppDtoVariantsItemDisplay: UpdateInAppDtoVariantsItemDisplay şeması.
type UpdateInAppDtoVariantsItemDisplay struct {
	// Zorunlu. Değerler: "modal", "banner_top", "banner_bottom", "fullscreen".
	Layout             string  `json:"layout"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	// Değerler: "blocks", "html".
	ContentType *string           `json:"content_type,omitempty"`
	HTML        map[string]string `json:"html,omitempty"`
	HTMLHeight  *int64            `json:"html_height,omitempty"`
	// Zorunlu.
	Blocks []any `json:"blocks"`
}

// UpdateJourneyDto: UpdateJourneyDto şeması.
type UpdateJourneyDto struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateKeywordDto: UpdateKeywordDto şeması.
type UpdateKeywordDto struct {
	Keyword *string `json:"keyword,omitempty"`
	// Değerler: "sms", "whatsapp", "rcs".
	Channels    []string                 `json:"channels,omitempty"`
	Actions     *UpdateKeywordDtoActions `json:"actions,omitempty"`
	Enabled     *bool                    `json:"enabled,omitempty"`
	Description *string                  `json:"description,omitempty"`
}

// UpdateKeywordDtoActions: UpdateKeywordDtoActions şeması.
type UpdateKeywordDtoActions struct {
	AddTags    map[string]string             `json:"add_tags,omitempty"`
	RemoveTags []string                      `json:"remove_tags,omitempty"`
	Event      *UpdateKeywordDtoActionsEvent `json:"event,omitempty"`
	Reply      *UpdateKeywordDtoActionsReply `json:"reply,omitempty"`
	OptIn      *bool                         `json:"opt_in,omitempty"`
}

// UpdateKeywordDtoActionsEvent: UpdateKeywordDtoActionsEvent şeması.
type UpdateKeywordDtoActionsEvent struct {
	// Zorunlu.
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties,omitempty"`
}

// UpdateKeywordDtoActionsReply: UpdateKeywordDtoActionsReply şeması.
type UpdateKeywordDtoActionsReply struct {
	// Zorunlu.
	Text string `json:"text"`
}

// UpdateLiveActivityDto: UpdateLiveActivityDto şeması.
type UpdateLiveActivityDto struct {
	// Zorunlu.
	ActivityID string `json:"activity_id"`
	// Zorunlu.
	ContentState   map[string]any              `json:"content_state"`
	Alert          *UpdateLiveActivityDtoAlert `json:"alert,omitempty"`
	StaleDate      *int64                      `json:"stale_date,omitempty"`
	DismissalDate  *int64                      `json:"dismissal_date,omitempty"`
	RelevanceScore *float64                    `json:"relevance_score,omitempty"`
	// Değerler: 5, 10.
	Priority *int64  `json:"priority,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// UpdateLiveActivityDtoAlert: UpdateLiveActivityDtoAlert şeması.
type UpdateLiveActivityDtoAlert struct {
	Title map[string]string `json:"title,omitempty"`
	Body  map[string]string `json:"body,omitempty"`
	Sound *string           `json:"sound,omitempty"`
}

// UpdateLiveActivityTypeDto: UpdateLiveActivityTypeDto şeması.
type UpdateLiveActivityTypeDto struct {
	Description        *string                                      `json:"description,omitempty"`
	ContentStateSchema *UpdateLiveActivityTypeDtoContentStateSchema `json:"content_state_schema,omitempty"`
	AttributesSchema   *UpdateLiveActivityTypeDtoAttributesSchema   `json:"attributes_schema,omitempty"`
}

// UpdateLiveActivityTypeDtoAttributesSchema: UpdateLiveActivityTypeDtoAttributesSchema şeması.
type UpdateLiveActivityTypeDtoAttributesSchema struct {
	// Zorunlu. Değerler: "object".
	Type string `json:"type"`
	// Zorunlu.
	Properties map[string]UpdateLiveActivityTypeDtoAttributesSchemaPropertiesValue `json:"properties"`
	Required   []string                                                            `json:"required,omitempty"`
}

// UpdateLiveActivityTypeDtoAttributesSchemaPropertiesValue: UpdateLiveActivityTypeDtoAttributesSchemaPropertiesValue şeması.
type UpdateLiveActivityTypeDtoAttributesSchemaPropertiesValue struct {
	// Zorunlu. Değerler: "string", "number", "integer", "boolean".
	Type        string   `json:"type"`
	Enum        []any    `json:"enum,omitempty"`
	MaxLength   *int64   `json:"maxLength,omitempty"`
	Minimum     *float64 `json:"minimum,omitempty"`
	Maximum     *float64 `json:"maximum,omitempty"`
	Description *string  `json:"description,omitempty"`
}

// UpdateLiveActivityTypeDtoContentStateSchema: UpdateLiveActivityTypeDtoContentStateSchema şeması.
type UpdateLiveActivityTypeDtoContentStateSchema struct {
	// Zorunlu. Değerler: "object".
	Type string `json:"type"`
	// Zorunlu.
	Properties map[string]UpdateLiveActivityTypeDtoContentStateSchemaPropertiesValue `json:"properties"`
	Required   []string                                                              `json:"required,omitempty"`
}

// UpdateLiveActivityTypeDtoContentStateSchemaPropertiesValue: UpdateLiveActivityTypeDtoContentStateSchemaPropertiesValue şeması.
type UpdateLiveActivityTypeDtoContentStateSchemaPropertiesValue struct {
	// Zorunlu. Değerler: "string", "number", "integer", "boolean".
	Type        string   `json:"type"`
	Enum        []any    `json:"enum,omitempty"`
	MaxLength   *int64   `json:"maxLength,omitempty"`
	Minimum     *float64 `json:"minimum,omitempty"`
	Maximum     *float64 `json:"maximum,omitempty"`
	Description *string  `json:"description,omitempty"`
}

// UpdateNodeDto: UpdateNodeDto şeması.
type UpdateNodeDto struct {
	Data     map[string]any         `json:"data,omitempty"`
	Position *UpdateNodeDtoPosition `json:"position,omitempty"`
}

// UpdateNodeDtoPosition: UpdateNodeDtoPosition şeması.
type UpdateNodeDtoPosition struct {
	// Zorunlu.
	X float64 `json:"x"`
	// Zorunlu.
	Y float64 `json:"y"`
}

// UpdateTemplateDto: UpdateTemplateDto şeması.
type UpdateTemplateDto struct {
	Name *string `json:"name,omitempty"`
	// Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channel *string                   `json:"channel,omitempty"`
	Tags    []string                  `json:"tags,omitempty"`
	Content *UpdateTemplateDtoContent `json:"content,omitempty"`
}

// UpdateTemplateDtoContent: UpdateTemplateDtoContent şeması.
type UpdateTemplateDtoContent struct {
	Headings       map[string]string                     `json:"headings,omitempty"`
	Contents       map[string]string                     `json:"contents,omitempty"`
	Subtitle       map[string]string                     `json:"subtitle,omitempty"`
	URL            *string                               `json:"url,omitempty"`
	AppURL         *string                               `json:"app_url,omitempty"`
	Data           map[string]string                     `json:"data,omitempty"`
	Buttons        []UpdateTemplateDtoContentButtonsItem `json:"buttons,omitempty"`
	BigPicture     *string                               `json:"big_picture,omitempty"`
	IOSAttachments map[string]string                     `json:"ios_attachments,omitempty"`
	IOSSound       *string                               `json:"ios_sound,omitempty"`
	// Değerler: "None", "SetTo", "Increase".
	IOSBadgeType  *string `json:"ios_badge_type,omitempty"`
	IOSBadgeCount *int64  `json:"ios_badge_count,omitempty"`
	// Değerler: "passive", "active", "time-sensitive", "critical".
	IOSInterruptionLevel      *string        `json:"ios_interruption_level,omitempty"`
	IOSCategory               *string        `json:"ios_category,omitempty"`
	AndroidChannelID          *string        `json:"android_channel_id,omitempty"`
	AndroidGroup              *string        `json:"android_group,omitempty"`
	AndroidSound              *string        `json:"android_sound,omitempty"`
	ChromeWebIcon             *string        `json:"chrome_web_icon,omitempty"`
	ChromeWebImage            *string        `json:"chrome_web_image,omitempty"`
	ChromeWebBadge            *string        `json:"chrome_web_badge,omitempty"`
	WebPushTopic              *string        `json:"web_push_topic,omitempty"`
	ContentAvailable          *bool          `json:"content_available,omitempty"`
	MutableContent            *bool          `json:"mutable_content,omitempty"`
	CustomData                map[string]any `json:"custom_data,omitempty"`
	EmailSubject              *string        `json:"email_subject,omitempty"`
	EmailBody                 *string        `json:"email_body,omitempty"`
	EmailMJML                 *string        `json:"email_mjml,omitempty"`
	EmailPreheader            *string        `json:"email_preheader,omitempty"`
	EmailFromName             *string        `json:"email_from_name,omitempty"`
	EmailFromAddress          *string        `json:"email_from_address,omitempty"`
	EmailReplyToAddress       *string        `json:"email_reply_to_address,omitempty"`
	DisableEmailClickTracking *bool          `json:"disable_email_click_tracking,omitempty"`
	SMSFrom                   *string        `json:"sms_from,omitempty"`
	DisableSMSLinkShortening  *bool          `json:"disable_sms_link_shortening,omitempty"`
	AddToInbox                *bool          `json:"add_to_inbox,omitempty"`
	InboxTTL                  *int64         `json:"inbox_ttl,omitempty"`
	ThreadID                  *string        `json:"thread_id,omitempty"`
	IOSRelevanceScore         *float64       `json:"ios_relevance_score,omitempty"`
	TargetContentIdentifier   *string        `json:"target_content_identifier,omitempty"`
	SmallIcon                 *string        `json:"small_icon,omitempty"`
	LargeIcon                 *string        `json:"large_icon,omitempty"`
	AndroidAccentColor        *string        `json:"android_accent_color,omitempty"`
	// Değerler: 1, 0, -1.
	AndroidVisibility   *int64                                       `json:"android_visibility,omitempty"`
	WebButtons          []UpdateTemplateDtoContentWebButtonsItem     `json:"web_buttons,omitempty"`
	RequireInteraction  *bool                                        `json:"require_interaction,omitempty"`
	SMSMediaUrls        []string                                     `json:"sms_media_urls,omitempty"`
	WhatsappTemplate    *UpdateTemplateDtoContentWhatsappTemplate    `json:"whatsapp_template,omitempty"`
	WhatsappMedia       *UpdateTemplateDtoContentWhatsappMedia       `json:"whatsapp_media,omitempty"`
	WhatsappInteractive *UpdateTemplateDtoContentWhatsappInteractive `json:"whatsapp_interactive,omitempty"`
	RcsCard             *UpdateTemplateDtoContentRcsCard             `json:"rcs_card,omitempty"`
	RcsSuggestions      []any                                        `json:"rcs_suggestions,omitempty"`
	RcsContentSid       *string                                      `json:"rcs_content_sid,omitempty"`
	RcsContentVariables map[string]string                            `json:"rcs_content_variables,omitempty"`
	RcsFallbackText     map[string]string                            `json:"rcs_fallback_text,omitempty"`
}

// UpdateTemplateDtoContentButtonsItem: UpdateTemplateDtoContentButtonsItem şeması.
type UpdateTemplateDtoContentButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string  `json:"text"`
	Icon *string `json:"icon,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// UpdateTemplateDtoContentRcsCard: UpdateTemplateDtoContentRcsCard şeması.
type UpdateTemplateDtoContentRcsCard struct {
	Title       map[string]string `json:"title,omitempty"`
	Description map[string]string `json:"description,omitempty"`
	MediaURL    *string           `json:"media_url,omitempty"`
	// Değerler: "SHORT", "MEDIUM", "TALL".
	MediaHeight *string `json:"media_height,omitempty"`
	// Değerler: "VERTICAL", "HORIZONTAL".
	Orientation *string `json:"orientation,omitempty"`
}

// UpdateTemplateDtoContentWebButtonsItem: UpdateTemplateDtoContentWebButtonsItem şeması.
type UpdateTemplateDtoContentWebButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Text string  `json:"text"`
	Icon *string `json:"icon,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// UpdateTemplateDtoContentWhatsappInteractive: UpdateTemplateDtoContentWhatsappInteractive şeması.
type UpdateTemplateDtoContentWhatsappInteractive struct {
	// Zorunlu. Değerler: "button", "list".
	Type    string                                                   `json:"type"`
	Header  *string                                                  `json:"header,omitempty"`
	Footer  *string                                                  `json:"footer,omitempty"`
	Buttons []UpdateTemplateDtoContentWhatsappInteractiveButtonsItem `json:"buttons,omitempty"`
	List    *UpdateTemplateDtoContentWhatsappInteractiveList         `json:"list,omitempty"`
}

// UpdateTemplateDtoContentWhatsappInteractiveButtonsItem: UpdateTemplateDtoContentWhatsappInteractiveButtonsItem şeması.
type UpdateTemplateDtoContentWhatsappInteractiveButtonsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title string `json:"title"`
}

// UpdateTemplateDtoContentWhatsappInteractiveList: UpdateTemplateDtoContentWhatsappInteractiveList şeması.
type UpdateTemplateDtoContentWhatsappInteractiveList struct {
	// Zorunlu.
	Button string `json:"button"`
	// Zorunlu.
	Sections []UpdateTemplateDtoContentWhatsappInteractiveListSectionsItem `json:"sections"`
}

// UpdateTemplateDtoContentWhatsappInteractiveListSectionsItem: UpdateTemplateDtoContentWhatsappInteractiveListSectionsItem şeması.
type UpdateTemplateDtoContentWhatsappInteractiveListSectionsItem struct {
	Title *string `json:"title,omitempty"`
	// Zorunlu.
	Rows []UpdateTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem `json:"rows"`
}

// UpdateTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem: UpdateTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem şeması.
type UpdateTemplateDtoContentWhatsappInteractiveListSectionsItemRowsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

// UpdateTemplateDtoContentWhatsappMedia: UpdateTemplateDtoContentWhatsappMedia şeması.
type UpdateTemplateDtoContentWhatsappMedia struct {
	// Zorunlu. Değerler: "image", "video", "document", "audio".
	Type string `json:"type"`
	// Zorunlu.
	URL      string  `json:"url"`
	Filename *string `json:"filename,omitempty"`
}

// UpdateTemplateDtoContentWhatsappTemplate: UpdateTemplateDtoContentWhatsappTemplate şeması.
type UpdateTemplateDtoContentWhatsappTemplate struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Language     string                                                     `json:"language"`
	Header       any                                                        `json:"header,omitempty"`
	BodyParams   []string                                                   `json:"body_params,omitempty"`
	ButtonParams []UpdateTemplateDtoContentWhatsappTemplateButtonParamsItem `json:"button_params,omitempty"`
}

// UpdateTemplateDtoContentWhatsappTemplateButtonParamsItem: UpdateTemplateDtoContentWhatsappTemplateButtonParamsItem şeması.
type UpdateTemplateDtoContentWhatsappTemplateButtonParamsItem struct {
	// Zorunlu.
	Index int64 `json:"index"`
	// Zorunlu. Değerler: "url", "quick_reply", "copy_code".
	SubType string `json:"sub_type"`
	// Zorunlu.
	Value string `json:"value"`
}

// UpdateWarehouseExportDto: UpdateWarehouseExportDto şeması.
type UpdateWarehouseExportDto struct {
	Name *string `json:"name,omitempty"`
	// Değerler: "deliveries", "delivery_events", "events", "conversions", "users", "subscriptions".
	Tables []string `json:"tables,omitempty"`
	// Değerler: "hourly", "daily".
	Schedule    *string        `json:"schedule,omitempty"`
	Enabled     *bool          `json:"enabled,omitempty"`
	Config      map[string]any `json:"config,omitempty"`
	Credentials map[string]any `json:"credentials,omitempty"`
}

// UpdateWarmupPlanDto: UpdateWarmupPlanDto şeması.
type UpdateWarmupPlanDto struct {
	Name *string `json:"name,omitempty"`
	// Değerler: "active", "paused".
	Status         *string                            `json:"status,omitempty"`
	DailyLimits    []int64                            `json:"daily_limits,omitempty"`
	ProviderShares *UpdateWarmupPlanDtoProviderShares `json:"provider_shares,omitempty"`
}

// UpdateWarmupPlanDtoProviderShares: UpdateWarmupPlanDtoProviderShares şeması.
type UpdateWarmupPlanDtoProviderShares struct {
	// Zorunlu.
	Gmail float64 `json:"gmail"`
	// Zorunlu.
	Outlook float64 `json:"outlook"`
	// Zorunlu.
	Yahoo float64 `json:"yahoo"`
	// Zorunlu.
	Other float64 `json:"other"`
}

// UpdateWebCampaignDto: UpdateWebCampaignDto şeması.
type UpdateWebCampaignDto struct {
	Name       *string                            `json:"name,omitempty"`
	Display    *UpdateWebCampaignDtoDisplay       `json:"display,omitempty"`
	Variants   []UpdateWebCampaignDtoVariantsItem `json:"variants,omitempty"`
	Trigger    any                                `json:"trigger,omitempty"`
	Conditions *UpdateWebCampaignDtoConditions    `json:"conditions,omitempty"`
	Targeting  *UpdateWebCampaignDtoTargeting     `json:"targeting,omitempty"`
	Frequency  *UpdateWebCampaignDtoFrequency     `json:"frequency,omitempty"`
	Priority   *int64                             `json:"priority,omitempty"`
	StartAt    *string                            `json:"start_at,omitempty"`
	EndAt      *string                            `json:"end_at,omitempty"`
}

// UpdateWebCampaignDtoConditions: UpdateWebCampaignDtoConditions şeması.
type UpdateWebCampaignDtoConditions struct {
	URLRules []UpdateWebCampaignDtoConditionsURLRulesItem `json:"url_rules,omitempty"`
	UtmRules []UpdateWebCampaignDtoConditionsUtmRulesItem `json:"utm_rules,omitempty"`
	Visits   *UpdateWebCampaignDtoConditionsVisits        `json:"visits,omitempty"`
}

// UpdateWebCampaignDtoConditionsURLRulesItem: UpdateWebCampaignDtoConditionsUrlRulesItem şeması.
type UpdateWebCampaignDtoConditionsURLRulesItem struct {
	// Zorunlu. Değerler: "equals", "contains", "starts_with", "not_contains", "matches".
	Op string `json:"op"`
	// Zorunlu.
	Value string `json:"value"`
}

// UpdateWebCampaignDtoConditionsUtmRulesItem: UpdateWebCampaignDtoConditionsUtmRulesItem şeması.
type UpdateWebCampaignDtoConditionsUtmRulesItem struct {
	// Zorunlu. Değerler: "utm_source", "utm_medium", "utm_campaign", "utm_term", "utm_content".
	Param string `json:"param"`
	// Zorunlu. Değerler: "equals", "contains", "exists".
	Op    string  `json:"op"`
	Value *string `json:"value,omitempty"`
}

// UpdateWebCampaignDtoConditionsVisits: UpdateWebCampaignDtoConditionsVisits şeması.
type UpdateWebCampaignDtoConditionsVisits struct {
	// Zorunlu. Değerler: "gte", "lte", "eq".
	Op string `json:"op"`
	// Zorunlu.
	Value int64 `json:"value"`
}

// UpdateWebCampaignDtoDisplay: UpdateWebCampaignDtoDisplay şeması.
type UpdateWebCampaignDtoDisplay struct {
	// Zorunlu. Değerler: "popup", "banner_top", "banner_bottom", "slide_in", "fullscreen".
	Layout string `json:"layout"`
	// Değerler: "left", "right".
	Position           *string `json:"position,omitempty"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	MaxWidth           *int64  `json:"max_width,omitempty"`
	// Zorunlu.
	Blocks   []any   `json:"blocks"`
	FormID   *string `json:"form_id,omitempty"`
	SurveyID *string `json:"survey_id,omitempty"`
}

// UpdateWebCampaignDtoFrequency: UpdateWebCampaignDtoFrequency şeması.
type UpdateWebCampaignDtoFrequency struct {
	MaxPerSession       *int64 `json:"max_per_session,omitempty"`
	MaxPerDay           *int64 `json:"max_per_day,omitempty"`
	MaxTotal            *int64 `json:"max_total,omitempty"`
	HideAfterCloseDays  *int64 `json:"hide_after_close_days,omitempty"`
	StopAfterConversion *bool  `json:"stop_after_conversion,omitempty"`
}

// UpdateWebCampaignDtoTargeting: UpdateWebCampaignDtoTargeting şeması.
type UpdateWebCampaignDtoTargeting struct {
	IncludedSegments []string `json:"included_segments,omitempty"`
	ExcludedSegments []string `json:"excluded_segments,omitempty"`
	// Değerler: "all", "new", "returning".
	Visitor *string `json:"visitor,omitempty"`
	// Değerler: "desktop", "mobile", "tablet".
	Devices []string `json:"devices,omitempty"`
}

// UpdateWebCampaignDtoVariantsItem: UpdateWebCampaignDtoVariantsItem şeması.
type UpdateWebCampaignDtoVariantsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	Display UpdateWebCampaignDtoVariantsItemDisplay `json:"display"`
}

// UpdateWebCampaignDtoVariantsItemDisplay: UpdateWebCampaignDtoVariantsItemDisplay şeması.
type UpdateWebCampaignDtoVariantsItemDisplay struct {
	// Zorunlu. Değerler: "popup", "banner_top", "banner_bottom", "slide_in", "fullscreen".
	Layout string `json:"layout"`
	// Değerler: "left", "right".
	Position           *string `json:"position,omitempty"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	MaxWidth           *int64  `json:"max_width,omitempty"`
	// Zorunlu.
	Blocks   []any   `json:"blocks"`
	FormID   *string `json:"form_id,omitempty"`
	SurveyID *string `json:"survey_id,omitempty"`
}

// UpdateWebhookDto: UpdateWebhookDto şeması.
type UpdateWebhookDto struct {
	URL         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	// Değerler: "message.queued", "message.completed", "message.failed", "delivery.sent", "delivery.delivered", "delivery.failed", "delivery.opened", "delivery.clicked", "delivery.bounced", "delivery.complained", "delivery.unsubscribed", "in_app.impression", "in_app.clicked", "in_app.dismissed", "subscription.created", "subscription.unsubscribed", "subscription.invalidated", "user.created", "user.updated", "user.merged", "user.deleted", "event.received", "conversion.created", "campaign.scheduled", "campaign.started", "campaign.completed", "journey.entered", "journey.step_completed", "journey.completed", "journey.exited", "conversation.created", "conversation.message", "conversation.assigned", "conversation.resolved", "audit.created", "form.submitted", "survey.responded", "ad_audience.synced", "ad_account.revoked", "ad_budget.alert", "ad_insights.imported".
	Events       []string                      `json:"events,omitempty"`
	EventFilters *UpdateWebhookDtoEventFilters `json:"event_filters,omitempty"`
	Batch        *bool                         `json:"batch,omitempty"`
	Headers      map[string]string             `json:"headers,omitempty"`
	// Değerler: "active", "disabled".
	Status       *string `json:"status,omitempty"`
	BodyTemplate *string `json:"body_template,omitempty"`
	// Şablon çıktısının içerik türü (varsayılan application/json). Değerler: "application/json", "text/plain".
	BodyContentType *string                  `json:"body_content_type,omitempty"`
	Filters         *UpdateWebhookDtoFilters `json:"filters,omitempty"`
}

// UpdateWebhookDtoEventFilters: UpdateWebhookDtoEventFilters şeması.
type UpdateWebhookDtoEventFilters struct {
	EventNames []string `json:"event_names,omitempty"`
}

// UpdateWebhookDtoFilters: Olay filtresi: { match: all|any, conditions: [{ field, op, value? }] }.
type UpdateWebhookDtoFilters struct {
	// Zorunlu. Değerler: "all", "any".
	Match string `json:"match"`
	// Zorunlu.
	Conditions []UpdateWebhookDtoFiltersConditionsItem `json:"conditions"`
}

// UpdateWebhookDtoFiltersConditionsItem: UpdateWebhookDtoFiltersConditionsItem şeması.
type UpdateWebhookDtoFiltersConditionsItem struct {
	// Zorunlu.
	Field string `json:"field"`
	// Zorunlu. Değerler: "eq", "neq", "in", "not_in", "starts_with", "contains", "exists", "not_exists".
	Op    string `json:"op"`
	Value any    `json:"value,omitempty"`
}

// UpsertUserDto: UpsertUserDto şeması.
type UpsertUserDto struct {
	Identity      any                  `json:"identity,omitempty"`
	ExternalID    *string              `json:"external_id,omitempty"`
	Properties    map[string]any       `json:"properties,omitempty"`
	Tags          map[string]string    `json:"tags,omitempty"`
	Subscriptions []AddSubscriptionDto `json:"subscriptions,omitempty"`
}

// UsageReport: UsageReport şeması.
type UsageReport struct {
	// Zorunlu.
	OrgID string `json:"org_id"`
	// Zorunlu.
	PlanCode string `json:"plan_code"`
	// Zorunlu.
	Period UsageReportPeriod `json:"period"`
	// Zorunlu.
	Metrics []UsageReportMetricsItem `json:"metrics"`
	// Zorunlu.
	Apps []UsageReportAppsItem `json:"apps"`
	// Zorunlu.
	Daily []UsageReportDailyItem `json:"daily"`
	// Zorunlu.
	Trialing bool `json:"trialing"`
	// Zorunlu.
	TrialEndsAt *string             `json:"trial_ends_at"`
	Overage     *UsageReportOverage `json:"overage,omitempty"`
}

// UsageReportAppsItem: UsageReportAppsItem şeması.
type UsageReportAppsItem struct {
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	MessagesSent float64 `json:"messages_sent"`
	// Zorunlu.
	ActiveSubscriptions float64 `json:"active_subscriptions"`
	// Zorunlu.
	Mau float64 `json:"mau"`
}

// UsageReportDailyItem: UsageReportDailyItem şeması.
type UsageReportDailyItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	MessagesSent float64 `json:"messages_sent"`
	// Zorunlu.
	APIRequests float64 `json:"api_requests"`
	// Zorunlu.
	EventsIngested float64 `json:"events_ingested"`
}

// UsageReportMetricsItem: UsageReportMetricsItem şeması.
type UsageReportMetricsItem struct {
	// Zorunlu.
	Metric string `json:"metric"`
	// Zorunlu.
	Used float64 `json:"used"`
	// Zorunlu.
	Limit *float64 `json:"limit"`
	// Zorunlu.
	Percent *float64 `json:"percent"`
	// Zorunlu.
	Hard bool `json:"hard"`
}

// UsageReportOverage: UsageReportOverage şeması.
type UsageReportOverage struct {
	// Zorunlu.
	Lines []UsageReportOverageLinesItem `json:"lines"`
	// Zorunlu.
	Totals UsageReportOverageTotals `json:"totals"`
}

// UsageReportOverageLinesItem: UsageReportOverageLinesItem şeması.
type UsageReportOverageLinesItem struct {
	// Zorunlu.
	Metric string `json:"metric"`
	// Zorunlu.
	Included float64 `json:"included"`
	// Zorunlu.
	Used float64 `json:"used"`
	// Zorunlu.
	Billable float64 `json:"billable"`
	// Zorunlu.
	Amounts UsageReportOverageLinesItemAmounts `json:"amounts"`
}

// UsageReportOverageLinesItemAmounts: UsageReportOverageLinesItemAmounts şeması.
type UsageReportOverageLinesItemAmounts struct {
	// Zorunlu.
	Usd float64 `json:"USD"`
	// Zorunlu.
	Try float64 `json:"TRY"`
}

// UsageReportOverageTotals: UsageReportOverageTotals şeması.
type UsageReportOverageTotals struct {
	// Zorunlu.
	Usd float64 `json:"USD"`
	// Zorunlu.
	Try float64 `json:"TRY"`
}

// UsageReportPeriod: UsageReportPeriod şeması.
type UsageReportPeriod struct {
	// Zorunlu.
	Start string `json:"start"`
	// Zorunlu.
	End string `json:"end"`
}

// User: User şeması.
type User struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu.
	Properties map[string]any `json:"properties"`
	// Zorunlu.
	Tags map[string]string `json:"tags"`
	// Zorunlu.
	FirstSeenAt string `json:"first_seen_at"`
	// Zorunlu.
	LastActiveAt string `json:"last_active_at"`
	// Zorunlu.
	SessionCount int64 `json:"session_count"`
	// Zorunlu.
	UsageSeconds int64 `json:"usage_seconds"`
	// Zorunlu.
	MergedInto *string `json:"merged_into"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
	// Zorunlu.
	SubscriptionCount int64 `json:"subscription_count"`
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels []string `json:"channels"`
}

// UserDetail: UserDetail şeması.
type UserDetail struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu.
	Properties map[string]any `json:"properties"`
	// Zorunlu.
	Tags map[string]string `json:"tags"`
	// Zorunlu.
	FirstSeenAt string `json:"first_seen_at"`
	// Zorunlu.
	LastActiveAt string `json:"last_active_at"`
	// Zorunlu.
	SessionCount int64 `json:"session_count"`
	// Zorunlu.
	UsageSeconds int64 `json:"usage_seconds"`
	// Zorunlu.
	MergedInto *string `json:"merged_into"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
	// Zorunlu.
	SubscriptionCount int64 `json:"subscription_count"`
	// Zorunlu. Değerler: "ios_push", "android_push", "hms_push", "web_push", "email", "sms", "whatsapp", "rcs".
	Channels []string `json:"channels"`
	// Zorunlu.
	Aliases []UserDetailAliasesItem `json:"aliases"`
	// Zorunlu.
	Subscriptions []Subscription `json:"subscriptions"`
}

// UserDetailAliasesItem: UserDetailAliasesItem şeması.
type UserDetailAliasesItem struct {
	// Zorunlu.
	Label string `json:"label"`
	// Zorunlu.
	Value string `json:"value"`
}

// UserExport: UserExport şeması.
type UserExport struct {
	// Zorunlu.
	ExportedAt string `json:"exported_at"`
	// Zorunlu.
	User UserDetail `json:"user"`
	// Zorunlu.
	Events []UserExportEventsItem `json:"events"`
}

// UserExportEventsItem: UserExportEventsItem şeması.
type UserExportEventsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	UserID *string `json:"user_id"`
	// Zorunlu.
	SubscriptionID *string `json:"subscription_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Properties map[string]any `json:"properties"`
	// Zorunlu.
	Value *float64 `json:"value"`
	// Zorunlu. Değerler: "sdk", "api", "system", "import".
	Source string `json:"source"`
	// Zorunlu.
	OccurredAt string `json:"occurred_at"`
}

// UserIdentity: UserIdentity şeması.
type UserIdentity struct {
	// Zorunlu.
	Identity map[string]string `json:"identity"`
}

// UserJourneyRuns: UserJourneyRuns şeması.
type UserJourneyRuns struct {
	// Zorunlu.
	Data []UserJourneyRunsDataItem `json:"data"`
}

// UserJourneyRunsDataItem: UserJourneyRunsDataItem şeması.
type UserJourneyRunsDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	JourneyID string `json:"journey_id"`
	// Zorunlu.
	Version int64 `json:"version"`
	// Zorunlu.
	UserID string `json:"user_id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu. Değerler: "active", "waiting", "completed", "exited", "failed".
	Status string `json:"status"`
	// Zorunlu.
	CurrentNodeID *string `json:"current_node_id"`
	// Zorunlu.
	NextActionAt *string `json:"next_action_at"`
	// Zorunlu.
	EnteredAt string `json:"entered_at"`
	// Zorunlu.
	CompletedAt *string `json:"completed_at"`
	// Zorunlu.
	ExitReason *string `json:"exit_reason"`
	// Zorunlu.
	GoalID *string `json:"goal_id"`
	// Zorunlu.
	GoalReachedAt *string `json:"goal_reached_at"`
	// Zorunlu.
	Context map[string]any `json:"context"`
	// Zorunlu.
	JourneyName string `json:"journey_name"`
}

// UserList: UserList şeması.
type UserList struct {
	// Zorunlu.
	Data []User `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// UserPreferences: UserPreferences şeması.
type UserPreferences struct {
	// Zorunlu.
	UserID string `json:"user_id"`
	// Zorunlu.
	Data []UserPreferencesDataItem `json:"data"`
	// Zorunlu.
	Categories []MessageCategory `json:"categories"`
	// Zorunlu.
	Consents []UserPreferencesConsentsItem `json:"consents"`
}

// UserPreferencesConsentsItem: UserPreferencesConsentsItem şeması.
type UserPreferencesConsentsItem struct {
	// Zorunlu. Değerler: "ads_personalization".
	Purpose string `json:"purpose"`
	// Zorunlu.
	Granted bool `json:"granted"`
	// Zorunlu.
	Source *string `json:"source"`
	// Zorunlu.
	UpdatedAt *string `json:"updated_at"`
}

// UserPreferencesDataItem: UserPreferencesDataItem şeması.
type UserPreferencesDataItem struct {
	// Zorunlu.
	Category string `json:"category"`
	// Zorunlu.
	Channel *string `json:"channel"`
	// Zorunlu.
	OptedIn bool `json:"opted_in"`
	// Zorunlu.
	Explicit bool `json:"explicit"`
	// Zorunlu.
	Source *string `json:"source"`
	// Zorunlu.
	UpdatedAt *string `json:"updated_at"`
}

// UsersSearchBody: UsersSearchBody şeması.
type UsersSearchBody struct {
	// Zorunlu. Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
	Definition UsersSearchBodyDefinition `json:"definition"`
	Limit      *int64                    `json:"limit,omitempty"`
	Cursor     *string                   `json:"cursor,omitempty"`
}

// UsersSearchBodyDefinition: Filtre AST: all/any/none grupları ve { field, op, value } koşulları (docs/plan/05 §5.7).
type UsersSearchBodyDefinition struct {
	All  []map[string]any `json:"all,omitempty"`
	Any  []map[string]any `json:"any,omitempty"`
	None []map[string]any `json:"none,omitempty"`
}

// VerifyAddressDto: VerifyAddressDto şeması.
type VerifyAddressDto struct {
	// Zorunlu.
	Address string `json:"address"`
}

// WarehouseConnectionResult: WarehouseConnectionResult şeması.
type WarehouseConnectionResult struct {
	// Zorunlu.
	Ok bool `json:"ok"`
	// Zorunlu.
	Detail *string `json:"detail"`
	// Zorunlu.
	Error *WarehouseConnectionResultError `json:"error"`
}

// WarehouseConnectionResultError: WarehouseConnectionResultError şeması.
type WarehouseConnectionResultError struct {
	// Zorunlu.
	Code string `json:"code"`
	// Zorunlu.
	Message string `json:"message"`
}

// WarehouseExport: WarehouseExport şeması.
type WarehouseExport struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "bigquery", "snowflake", "s3_parquet".
	Destination string `json:"destination"`
	// Zorunlu.
	Config map[string]any `json:"config"`
	// Zorunlu.
	CredentialsHint string `json:"credentials_hint"`
	// Zorunlu. Değerler: "deliveries", "delivery_events", "events", "conversions", "users", "subscriptions".
	Tables []string `json:"tables"`
	// Zorunlu. Değerler: "hourly", "daily".
	Schedule string `json:"schedule"`
	// Zorunlu.
	Enabled bool `json:"enabled"`
	// Zorunlu. Değerler: "active", "paused", "failing".
	Status string `json:"status"`
	// Zorunlu.
	ConsecutiveFailures int64 `json:"consecutive_failures"`
	// Zorunlu.
	LastError *string `json:"last_error"`
	// Zorunlu.
	Cursors map[string]WarehouseExportCursorsValue `json:"cursors"`
	// Zorunlu.
	SchemaVersions map[string]int64 `json:"schema_versions"`
	// Zorunlu.
	LastRunAt *string `json:"last_run_at"`
	// Zorunlu.
	LastSuccessAt *string `json:"last_success_at"`
	// Zorunlu.
	NextRunAt *string `json:"next_run_at"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// WarehouseExportCursorsValue: WarehouseExportCursorsValue şeması.
type WarehouseExportCursorsValue struct {
	// Zorunlu.
	Ts string `json:"ts"`
	// Zorunlu.
	ID string `json:"id"`
}

// WarehouseExportList: WarehouseExportList şeması.
type WarehouseExportList struct {
	// Zorunlu.
	Data []WarehouseExport `json:"data"`
}

// WarehouseExportRun: WarehouseExportRun şeması.
type WarehouseExportRun struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	ExportID string `json:"export_id"`
	// Zorunlu. Değerler: "schedule", "manual".
	Trigger string `json:"trigger"`
	// Zorunlu. Değerler: "queued", "running", "succeeded", "failed".
	Status string `json:"status"`
	// Zorunlu.
	Tables map[string]WarehouseExportRunTablesValue `json:"tables"`
	// Zorunlu.
	RowCount int64 `json:"row_count"`
	// Zorunlu.
	Attempt int64 `json:"attempt"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	StartedAt *string `json:"started_at"`
	// Zorunlu.
	FinishedAt *string `json:"finished_at"`
	// Zorunlu.
	DurationMs *int64 `json:"duration_ms"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// WarehouseExportRunPage: WarehouseExportRunPage şeması.
type WarehouseExportRunPage struct {
	// Zorunlu.
	Data []WarehouseExportRun `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// WarehouseExportRunTablesValue: WarehouseExportRunTablesValue şeması.
type WarehouseExportRunTablesValue struct {
	// Zorunlu.
	Rows int64 `json:"rows"`
	// Zorunlu.
	SchemaVersion int64 `json:"schema_version"`
	// Zorunlu.
	From *WarehouseExportRunTablesValueFrom `json:"from"`
	// Zorunlu.
	To *WarehouseExportRunTablesValueTo `json:"to"`
	// Zorunlu.
	Objects int64 `json:"objects"`
}

// WarehouseExportRunTablesValueFrom: WarehouseExportRunTablesValueFrom şeması.
type WarehouseExportRunTablesValueFrom struct {
	// Zorunlu.
	Ts string `json:"ts"`
	// Zorunlu.
	ID string `json:"id"`
}

// WarehouseExportRunTablesValueTo: WarehouseExportRunTablesValueTo şeması.
type WarehouseExportRunTablesValueTo struct {
	// Zorunlu.
	Ts string `json:"ts"`
	// Zorunlu.
	ID string `json:"id"`
}

// WarehouseTables: WarehouseTables şeması.
type WarehouseTables struct {
	// Zorunlu.
	Data []WarehouseTablesDataItem `json:"data"`
}

// WarehouseTablesDataItem: WarehouseTablesDataItem şeması.
type WarehouseTablesDataItem struct {
	// Zorunlu. Değerler: "deliveries", "delivery_events", "events", "conversions", "users", "subscriptions".
	Name string `json:"name"`
	// Zorunlu. Değerler: "incremental", "snapshot".
	Mode string `json:"mode"`
	// Zorunlu.
	SchemaVersion int64 `json:"schema_version"`
	// Zorunlu.
	CursorColumn *string `json:"cursor_column"`
	// Zorunlu.
	Columns []WarehouseTablesDataItemColumnsItem `json:"columns"`
}

// WarehouseTablesDataItemColumnsItem: WarehouseTablesDataItemColumnsItem şeması.
type WarehouseTablesDataItemColumnsItem struct {
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Type string `json:"type"`
	// Zorunlu.
	Nullable bool `json:"nullable"`
}

// WebCampaign: WebCampaign şeması.
type WebCampaign struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu. Değerler: "draft", "active", "paused", "archived".
	Status string `json:"status"`
	// Zorunlu.
	Priority int64 `json:"priority"`
	// Zorunlu.
	Display WebCampaignDisplay `json:"display"`
	// Zorunlu.
	Variants []WebCampaignVariantsItem `json:"variants"`
	// Zorunlu.
	AbTest WebCampaignAbTest `json:"ab_test"`
	// Zorunlu.
	Trigger any `json:"trigger"`
	// Zorunlu.
	Conditions WebCampaignConditions `json:"conditions"`
	// Zorunlu.
	Targeting WebCampaignTargeting `json:"targeting"`
	// Zorunlu.
	Frequency WebCampaignFrequency `json:"frequency"`
	// Zorunlu.
	StartAt *string `json:"start_at"`
	// Zorunlu.
	EndAt *string `json:"end_at"`
	// Zorunlu.
	Stats WebCampaignStats `json:"stats"`
	// Zorunlu.
	VariantStats []WebCampaignVariantStatsItem `json:"variant_stats"`
	// Zorunlu.
	VariantSignificance *WebCampaignVariantSignificance `json:"variant_significance"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// WebCampaignAbTest: WebCampaignAbTest şeması.
type WebCampaignAbTest struct {
	// Zorunlu.
	WinnerVariantID *string `json:"winner_variant_id"`
	// Zorunlu.
	DecidedAt *string `json:"decided_at"`
	// Zorunlu.
	DecidedBy *string `json:"decided_by"`
}

// WebCampaignConditions: WebCampaignConditions şeması.
type WebCampaignConditions struct {
	URLRules []WebCampaignConditionsURLRulesItem `json:"url_rules,omitempty"`
	UtmRules []WebCampaignConditionsUtmRulesItem `json:"utm_rules,omitempty"`
	Visits   *WebCampaignConditionsVisits        `json:"visits,omitempty"`
}

// WebCampaignConditionsURLRulesItem: WebCampaignConditionsUrlRulesItem şeması.
type WebCampaignConditionsURLRulesItem struct {
	// Zorunlu. Değerler: "equals", "contains", "starts_with", "not_contains", "matches".
	Op string `json:"op"`
	// Zorunlu.
	Value string `json:"value"`
}

// WebCampaignConditionsUtmRulesItem: WebCampaignConditionsUtmRulesItem şeması.
type WebCampaignConditionsUtmRulesItem struct {
	// Zorunlu. Değerler: "utm_source", "utm_medium", "utm_campaign", "utm_term", "utm_content".
	Param string `json:"param"`
	// Zorunlu. Değerler: "equals", "contains", "exists".
	Op    string  `json:"op"`
	Value *string `json:"value,omitempty"`
}

// WebCampaignConditionsVisits: WebCampaignConditionsVisits şeması.
type WebCampaignConditionsVisits struct {
	// Zorunlu. Değerler: "gte", "lte", "eq".
	Op string `json:"op"`
	// Zorunlu.
	Value int64 `json:"value"`
}

// WebCampaignDisplay: WebCampaignDisplay şeması.
type WebCampaignDisplay struct {
	// Zorunlu. Değerler: "popup", "banner_top", "banner_bottom", "slide_in", "fullscreen".
	Layout string `json:"layout"`
	// Değerler: "left", "right".
	Position           *string `json:"position,omitempty"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	MaxWidth           *int64  `json:"max_width,omitempty"`
	// Zorunlu.
	Blocks   []any   `json:"blocks"`
	FormID   *string `json:"form_id,omitempty"`
	SurveyID *string `json:"survey_id,omitempty"`
}

// WebCampaignFrequency: WebCampaignFrequency şeması.
type WebCampaignFrequency struct {
	MaxPerSession       *int64 `json:"max_per_session,omitempty"`
	MaxPerDay           *int64 `json:"max_per_day,omitempty"`
	MaxTotal            *int64 `json:"max_total,omitempty"`
	HideAfterCloseDays  *int64 `json:"hide_after_close_days,omitempty"`
	StopAfterConversion *bool  `json:"stop_after_conversion,omitempty"`
}

// WebCampaignPage: WebCampaignPage şeması.
type WebCampaignPage struct {
	// Zorunlu.
	Data []WebCampaign `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// WebCampaignReport: WebCampaignReport şeması.
type WebCampaignReport struct {
	// Zorunlu.
	CampaignID string `json:"campaign_id"`
	// Zorunlu.
	Days int64 `json:"days"`
	// Zorunlu.
	Daily []WebCampaignReportDailyItem `json:"daily"`
	// Zorunlu.
	Totals WebCampaignReportTotals `json:"totals"`
}

// WebCampaignReportDailyItem: WebCampaignReportDailyItem şeması.
type WebCampaignReportDailyItem struct {
	// Zorunlu.
	Date string `json:"date"`
	// Zorunlu.
	Impressions int64 `json:"impressions"`
	// Zorunlu.
	Clicks int64 `json:"clicks"`
	// Zorunlu.
	Closes int64 `json:"closes"`
	// Zorunlu.
	Conversions int64 `json:"conversions"`
}

// WebCampaignReportTotals: WebCampaignReportTotals şeması.
type WebCampaignReportTotals struct {
	// Zorunlu.
	Impressions int64 `json:"impressions"`
	// Zorunlu.
	Clicks int64 `json:"clicks"`
	// Zorunlu.
	Closes int64 `json:"closes"`
	// Zorunlu.
	Conversions int64 `json:"conversions"`
	// Zorunlu.
	UniqueVisitors int64 `json:"unique_visitors"`
	// Zorunlu.
	Ctr float64 `json:"ctr"`
	// Zorunlu.
	ConversionRate float64 `json:"conversion_rate"`
}

// WebCampaignStats: WebCampaignStats şeması.
type WebCampaignStats struct {
	// Zorunlu.
	Impressions int64 `json:"impressions"`
	// Zorunlu.
	Clicks int64 `json:"clicks"`
	// Zorunlu.
	Closes int64 `json:"closes"`
	// Zorunlu.
	Conversions int64 `json:"conversions"`
	// Zorunlu.
	UniqueVisitors int64 `json:"unique_visitors"`
	// Zorunlu.
	Ctr float64 `json:"ctr"`
	// Zorunlu.
	ConversionRate float64 `json:"conversion_rate"`
}

// WebCampaignStatusDto: WebCampaignStatusDto şeması.
type WebCampaignStatusDto struct {
	// Zorunlu. Değerler: "draft", "active", "paused", "archived".
	Status string `json:"status"`
}

// WebCampaignTargeting: WebCampaignTargeting şeması.
type WebCampaignTargeting struct {
	IncludedSegments []string `json:"included_segments,omitempty"`
	ExcludedSegments []string `json:"excluded_segments,omitempty"`
	// Değerler: "all", "new", "returning".
	Visitor *string `json:"visitor,omitempty"`
	// Değerler: "desktop", "mobile", "tablet".
	Devices []string `json:"devices,omitempty"`
}

// WebCampaignVariantSignificance: WebCampaignVariantSignificance şeması.
type WebCampaignVariantSignificance struct {
	// Zorunlu.
	Leader *string `json:"leader"`
	// Zorunlu.
	PValue float64 `json:"p_value"`
	// Zorunlu.
	Significant bool `json:"significant"`
}

// WebCampaignVariantsItem: WebCampaignVariantsItem şeması.
type WebCampaignVariantsItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	WeightPercent int64 `json:"weight_percent"`
	// Zorunlu.
	Display WebCampaignVariantsItemDisplay `json:"display"`
}

// WebCampaignVariantsItemDisplay: WebCampaignVariantsItemDisplay şeması.
type WebCampaignVariantsItemDisplay struct {
	// Zorunlu. Değerler: "popup", "banner_top", "banner_bottom", "slide_in", "fullscreen".
	Layout string `json:"layout"`
	// Değerler: "left", "right".
	Position           *string `json:"position,omitempty"`
	BackgroundColor    *string `json:"background_color,omitempty"`
	TextColor          *string `json:"text_color,omitempty"`
	Dismissible        *bool   `json:"dismissible,omitempty"`
	AutoDismissSeconds *int64  `json:"auto_dismiss_seconds,omitempty"`
	MaxWidth           *int64  `json:"max_width,omitempty"`
	// Zorunlu.
	Blocks   []any   `json:"blocks"`
	FormID   *string `json:"form_id,omitempty"`
	SurveyID *string `json:"survey_id,omitempty"`
}

// WebCampaignVariantStatsItem: WebCampaignVariantStatsItem şeması.
type WebCampaignVariantStatsItem struct {
	// Zorunlu.
	VariantID string `json:"variant_id"`
	// Zorunlu.
	Impressions int64 `json:"impressions"`
	// Zorunlu.
	Clicks int64 `json:"clicks"`
	// Zorunlu.
	Conversions int64 `json:"conversions"`
	// Zorunlu.
	Ctr float64 `json:"ctr"`
	// Zorunlu.
	ConversionRate float64 `json:"conversion_rate"`
}

// Webhook: Webhook şeması.
type Webhook struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	URL string `json:"url"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	Events []string `json:"events"`
	// Zorunlu.
	EventFilters WebhookEventFilters `json:"event_filters"`
	// Zorunlu.
	Batch bool `json:"batch"`
	// Zorunlu. Değerler: "active", "failing", "disabled".
	Status string `json:"status"`
	// Zorunlu.
	DisabledReason *string `json:"disabled_reason"`
	// Zorunlu.
	FailureCount int64 `json:"failure_count"`
	// Zorunlu.
	FailingSince *string `json:"failing_since"`
	// Zorunlu.
	LastSuccessAt *string `json:"last_success_at"`
	// Zorunlu.
	LastFailureAt *string `json:"last_failure_at"`
	// Zorunlu.
	HeaderNames []string `json:"header_names"`
	// Zorunlu.
	PreviousSecretExpiresAt *string `json:"previous_secret_expires_at"`
	// Zorunlu.
	BodyTemplate *string `json:"body_template"`
	// Zorunlu. Değerler: "application/json", "text/plain".
	BodyContentType string `json:"body_content_type"`
	// Zorunlu.
	Filters *WebhookFilters `json:"filters"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
}

// WebhookDeliveryDetail: WebhookDeliveryDetail şeması.
type WebhookDeliveryDetail struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	EndpointID string `json:"endpoint_id"`
	// Zorunlu.
	EventType string `json:"event_type"`
	// Zorunlu.
	EventID string `json:"event_id"`
	// Zorunlu.
	SubjectID *string `json:"subject_id"`
	// Zorunlu.
	Attempt int64 `json:"attempt"`
	// Zorunlu. Değerler: "pending", "success", "failed", "dead".
	Status string `json:"status"`
	// Zorunlu.
	ResponseStatus *int64 `json:"response_status"`
	// Zorunlu.
	ResponseBodyExcerpt *string `json:"response_body_excerpt"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	DurationMs *int64 `json:"duration_ms"`
	// Zorunlu.
	NextRetryAt *string `json:"next_retry_at"`
	// Zorunlu.
	DeliveredAt *string `json:"delivered_at"`
	// Zorunlu.
	ReplayOf *string `json:"replay_of"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	Payload map[string]any `json:"payload"`
}

// WebhookDeliveryPage: WebhookDeliveryPage şeması.
type WebhookDeliveryPage struct {
	// Zorunlu.
	Data []WebhookDeliveryPageDataItem `json:"data"`
	// Zorunlu.
	NextCursor *string `json:"next_cursor"`
	// Zorunlu.
	HasMore bool `json:"has_more"`
}

// WebhookDeliveryPageDataItem: WebhookDeliveryPageDataItem şeması.
type WebhookDeliveryPageDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	EndpointID string `json:"endpoint_id"`
	// Zorunlu.
	EventType string `json:"event_type"`
	// Zorunlu.
	EventID string `json:"event_id"`
	// Zorunlu.
	SubjectID *string `json:"subject_id"`
	// Zorunlu.
	Attempt int64 `json:"attempt"`
	// Zorunlu. Değerler: "pending", "success", "failed", "dead".
	Status string `json:"status"`
	// Zorunlu.
	ResponseStatus *int64 `json:"response_status"`
	// Zorunlu.
	ResponseBodyExcerpt *string `json:"response_body_excerpt"`
	// Zorunlu.
	Error *string `json:"error"`
	// Zorunlu.
	DurationMs *int64 `json:"duration_ms"`
	// Zorunlu.
	NextRetryAt *string `json:"next_retry_at"`
	// Zorunlu.
	DeliveredAt *string `json:"delivered_at"`
	// Zorunlu.
	ReplayOf *string `json:"replay_of"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// WebhookEventFilters: WebhookEventFilters şeması.
type WebhookEventFilters struct {
	EventNames []string `json:"event_names,omitempty"`
}

// WebhookFilters: WebhookFilters şeması.
type WebhookFilters struct {
	// Zorunlu. Değerler: "all", "any".
	Match string `json:"match"`
	// Zorunlu.
	Conditions []WebhookFiltersConditionsItem `json:"conditions"`
}

// WebhookFiltersConditionsItem: WebhookFiltersConditionsItem şeması.
type WebhookFiltersConditionsItem struct {
	// Zorunlu.
	Field string `json:"field"`
	// Zorunlu. Değerler: "eq", "neq", "in", "not_in", "starts_with", "contains", "exists", "not_exists".
	Op    string `json:"op"`
	Value any    `json:"value,omitempty"`
}

// WebhookList: WebhookList şeması.
type WebhookList struct {
	// Zorunlu.
	Data []Webhook `json:"data"`
}

// WebhookPreviewDto: WebhookPreviewDto şeması.
type WebhookPreviewDto struct {
	// Zorunlu. Değerler: "message.queued", "message.completed", "message.failed", "delivery.sent", "delivery.delivered", "delivery.failed", "delivery.opened", "delivery.clicked", "delivery.bounced", "delivery.complained", "delivery.unsubscribed", "in_app.impression", "in_app.clicked", "in_app.dismissed", "subscription.created", "subscription.unsubscribed", "subscription.invalidated", "user.created", "user.updated", "user.merged", "user.deleted", "event.received", "conversion.created", "campaign.scheduled", "campaign.started", "campaign.completed", "journey.entered", "journey.step_completed", "journey.completed", "journey.exited", "conversation.created", "conversation.message", "conversation.assigned", "conversation.resolved", "audit.created", "form.submitted", "survey.responded", "ad_audience.synced", "ad_account.revoked", "ad_budget.alert", "ad_insights.imported", "webhook.test".
	EventType    string         `json:"event_type"`
	Data         map[string]any `json:"data,omitempty"`
	BodyTemplate *string        `json:"body_template,omitempty"`
	// Şablon çıktısının içerik türü (varsayılan application/json). Değerler: "application/json", "text/plain".
	BodyContentType *string                   `json:"body_content_type,omitempty"`
	Filters         *WebhookPreviewDtoFilters `json:"filters,omitempty"`
	Batch           *bool                     `json:"batch,omitempty"`
}

// WebhookPreviewDtoFilters: Olay filtresi: { match: all|any, conditions: [{ field, op, value? }] }.
type WebhookPreviewDtoFilters struct {
	// Zorunlu. Değerler: "all", "any".
	Match string `json:"match"`
	// Zorunlu.
	Conditions []WebhookPreviewDtoFiltersConditionsItem `json:"conditions"`
}

// WebhookPreviewDtoFiltersConditionsItem: WebhookPreviewDtoFiltersConditionsItem şeması.
type WebhookPreviewDtoFiltersConditionsItem struct {
	// Zorunlu.
	Field string `json:"field"`
	// Zorunlu. Değerler: "eq", "neq", "in", "not_in", "starts_with", "contains", "exists", "not_exists".
	Op    string `json:"op"`
	Value any    `json:"value,omitempty"`
}

// WebhookPreviewResult: WebhookPreviewResult şeması.
type WebhookPreviewResult struct {
	// Zorunlu.
	Payload map[string]any `json:"payload"`
	// Zorunlu. Olay filtreden geçer mi (geçmezse gönderilmez).
	Matched bool `json:"matched"`
	// Zorunlu.
	FailedConditions []WebhookPreviewResultFailedConditionsItem `json:"failed_conditions"`
	// Zorunlu.
	Body *string `json:"body"`
	// Zorunlu.
	ContentType string `json:"content_type"`
	// Zorunlu.
	Error *WebhookPreviewResultError `json:"error"`
}

// WebhookPreviewResultError: WebhookPreviewResultError şeması.
type WebhookPreviewResultError struct {
	// Zorunlu.
	Code string `json:"code"`
	// Zorunlu.
	Message string `json:"message"`
}

// WebhookPreviewResultFailedConditionsItem: WebhookPreviewResultFailedConditionsItem şeması.
type WebhookPreviewResultFailedConditionsItem struct {
	// Zorunlu.
	Field string `json:"field"`
	// Zorunlu.
	Op    string `json:"op"`
	Value any    `json:"value,omitempty"`
}

// WebhookWithSecret: WebhookWithSecret şeması.
type WebhookWithSecret struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	URL string `json:"url"`
	// Zorunlu.
	Description *string `json:"description"`
	// Zorunlu.
	Events []string `json:"events"`
	// Zorunlu.
	EventFilters WebhookWithSecretEventFilters `json:"event_filters"`
	// Zorunlu.
	Batch bool `json:"batch"`
	// Zorunlu. Değerler: "active", "failing", "disabled".
	Status string `json:"status"`
	// Zorunlu.
	DisabledReason *string `json:"disabled_reason"`
	// Zorunlu.
	FailureCount int64 `json:"failure_count"`
	// Zorunlu.
	FailingSince *string `json:"failing_since"`
	// Zorunlu.
	LastSuccessAt *string `json:"last_success_at"`
	// Zorunlu.
	LastFailureAt *string `json:"last_failure_at"`
	// Zorunlu.
	HeaderNames []string `json:"header_names"`
	// Zorunlu.
	PreviousSecretExpiresAt *string `json:"previous_secret_expires_at"`
	// Zorunlu.
	BodyTemplate *string `json:"body_template"`
	// Zorunlu. Değerler: "application/json", "text/plain".
	BodyContentType string `json:"body_content_type"`
	// Zorunlu.
	Filters *WebhookWithSecretFilters `json:"filters"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	UpdatedAt string `json:"updated_at"`
	// Zorunlu. Yalnızca oluşturma ve rotasyon yanıtında bir kez döner.
	Secret string `json:"secret"`
}

// WebhookWithSecretEventFilters: WebhookWithSecretEventFilters şeması.
type WebhookWithSecretEventFilters struct {
	EventNames []string `json:"event_names,omitempty"`
}

// WebhookWithSecretFilters: WebhookWithSecretFilters şeması.
type WebhookWithSecretFilters struct {
	// Zorunlu. Değerler: "all", "any".
	Match string `json:"match"`
	// Zorunlu.
	Conditions []WebhookWithSecretFiltersConditionsItem `json:"conditions"`
}

// WebhookWithSecretFiltersConditionsItem: WebhookWithSecretFiltersConditionsItem şeması.
type WebhookWithSecretFiltersConditionsItem struct {
	// Zorunlu.
	Field string `json:"field"`
	// Zorunlu. Değerler: "eq", "neq", "in", "not_in", "starts_with", "contains", "exists", "not_exists".
	Op    string `json:"op"`
	Value any    `json:"value,omitempty"`
}

// WhatsAppTemplateList: WhatsAppTemplateList şeması.
type WhatsAppTemplateList struct {
	// Zorunlu.
	Data []WhatsAppTemplateListDataItem `json:"data"`
}

// WhatsAppTemplateListDataItem: WhatsAppTemplateListDataItem şeması.
type WhatsAppTemplateListDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Language string `json:"language"`
	// Zorunlu.
	Category string `json:"category"`
	// Zorunlu.
	Status string `json:"status"`
	// Zorunlu.
	Components []map[string]any `json:"components"`
	// Zorunlu.
	BodyParams int64 `json:"body_params"`
	// Zorunlu.
	RejectedReason *string `json:"rejected_reason"`
	// Zorunlu.
	QualityScore *string `json:"quality_score"`
	// Zorunlu.
	SyncedAt string `json:"synced_at"`
}

// WhatsAppTemplateSync: WhatsAppTemplateSync şeması.
type WhatsAppTemplateSync struct {
	// Zorunlu.
	Total int64 `json:"total"`
	// Zorunlu.
	Removed int64 `json:"removed"`
	// Zorunlu.
	Data []WhatsAppTemplateSyncDataItem `json:"data"`
}

// WhatsAppTemplateSyncDataItem: WhatsAppTemplateSyncDataItem şeması.
type WhatsAppTemplateSyncDataItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	ExternalID *string `json:"external_id"`
	// Zorunlu.
	Name string `json:"name"`
	// Zorunlu.
	Language string `json:"language"`
	// Zorunlu.
	Category string `json:"category"`
	// Zorunlu.
	Status string `json:"status"`
	// Zorunlu.
	Components []map[string]any `json:"components"`
	// Zorunlu.
	BodyParams int64 `json:"body_params"`
	// Zorunlu.
	RejectedReason *string `json:"rejected_reason"`
	// Zorunlu.
	QualityScore *string `json:"quality_score"`
	// Zorunlu.
	SyncedAt string `json:"synced_at"`
}

// ZapierHook: ZapierHook şeması.
type ZapierHook struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Event string `json:"event"`
	// Zorunlu.
	EventName *string `json:"event_name"`
	// Zorunlu.
	TargetURL string `json:"target_url"`
	// Zorunlu.
	LastDeliveredAt *string `json:"last_delivered_at"`
	// Zorunlu.
	FailureCount int64 `json:"failure_count"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
}

// ZapierMe: ZapierMe şeması.
type ZapierMe struct {
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	AppName string `json:"app_name"`
	// Zorunlu.
	OrgID string `json:"org_id"`
	// Zorunlu.
	Scopes []string `json:"scopes"`
}

// ZapierSamplesItem: ZapierSamplesItem şeması.
type ZapierSamplesItem struct {
	// Zorunlu.
	ID string `json:"id"`
	// Zorunlu.
	Type string `json:"type"`
	// Zorunlu.
	CreatedAt string `json:"created_at"`
	// Zorunlu.
	AppID string `json:"app_id"`
	// Zorunlu.
	Data map[string]any `json:"data"`
}

// ZapierSubscribeDto: ZapierSubscribeDto şeması.
type ZapierSubscribeDto struct {
	// Zorunlu.
	TargetURL string `json:"target_url"`
	// Zorunlu. Değerler: "subscription.created", "subscription.unsubscribed", "event.received", "delivery.delivered", "delivery.opened", "delivery.clicked", "user.created", "user.updated", "conversion.created", "campaign.completed".
	Event     string  `json:"event"`
	EventName *string `json:"event_name,omitempty"`
}
