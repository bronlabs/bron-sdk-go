package types

type WebhookEvent struct {
	Event string `json:"event"`
	EventID string `json:"eventId"`
	Payload map[string]interface{} `json:"payload"`
	SubscriptionID string `json:"subscriptionId"`
}
