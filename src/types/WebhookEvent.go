package types

type WebhookEvent struct {
	Event string `json:"event"`
	EventId string `json:"eventId"`
	Payload map[string]interface{} `json:"payload"`
	SubscriptionId string `json:"subscriptionId"`
}
