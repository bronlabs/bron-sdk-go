package types

type WebhookEvent struct {
	SubscriptionId string `json:"subscriptionId"`
	Event string `json:"event"`
	EventId string `json:"eventId"`
	Payload map[string]interface{} `json:"payload"`
}
