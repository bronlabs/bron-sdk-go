package types

type TransactionEstimation struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	CreatedAt string `json:"createdAt"`
	EstimationID string `json:"estimationId"`
	EventType EventType `json:"eventType"`
	Extra *EventExtra `json:"extra,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TransactionID string `json:"transactionId"`
	UsdAmount *string `json:"usdAmount,omitempty"`
}
