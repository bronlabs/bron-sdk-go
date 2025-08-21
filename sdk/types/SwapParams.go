package types

type SwapParams struct {
	FromAmount *string `json:"fromAmount,omitempty"`
	FromAssetID string `json:"fromAssetId"`
	QuoteID string `json:"quoteId"`
	ToAmount *string `json:"toAmount,omitempty"`
	ToAssetID string `json:"toAssetId"`
}
