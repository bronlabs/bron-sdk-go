package types

type TransactionEvent struct {
	AccountID string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	CreatedAt string `json:"createdAt"`
	EventID string `json:"eventId"`
	EventType EventType `json:"eventType"`
	Extra *EventExtra `json:"extra,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TransactionID string `json:"transactionId"`
	UsdAmount *string `json:"usdAmount,omitempty"`
	WorkspaceID string `json:"workspaceId"`
}
