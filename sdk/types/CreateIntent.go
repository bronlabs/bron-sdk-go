package types

type CreateIntent struct {
	AccountID string `json:"accountId"`
	FromAmount *string `json:"fromAmount,omitempty"`
	FromAssetID string `json:"fromAssetId"`
	IntentID string `json:"intentId"`
	ToAmount *string `json:"toAmount,omitempty"`
	ToAssetID string `json:"toAssetId"`
}
