package types

type Intent struct {
	CreatedAt string `json:"createdAt"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	FromAmount *string `json:"fromAmount,omitempty"`
	FromAssetID string `json:"fromAssetId"`
	IntentID string `json:"intentId"`
	Price *string `json:"price,omitempty"`
	Status IntentOrderStatus `json:"status"`
	ToAmount *string `json:"toAmount,omitempty"`
	ToAssetID string `json:"toAssetId"`
	UpdatedAt string `json:"updatedAt"`
	UserSettlementDeadline *string `json:"userSettlementDeadline,omitempty"`
}
