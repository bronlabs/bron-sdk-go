package types

type SwapRoute struct {
	AccountID string `json:"accountId"`
	ExternalID *string `json:"externalId,omitempty"`
	FromAmount *string `json:"fromAmount,omitempty"`
	FromAssetID string `json:"fromAssetId"`
	FromTokenID *string `json:"fromTokenId,omitempty"`
	ToAmount *string `json:"toAmount,omitempty"`
	ToAssetID string `json:"toAssetId"`
	ToTokenID *string `json:"toTokenId,omitempty"`
}
