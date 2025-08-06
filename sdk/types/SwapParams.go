package types

type SwapParams struct {
	FromAmount *string `json:"fromAmount,omitempty"`
	FromAssetId string `json:"fromAssetId"`
	QuoteId string `json:"quoteId"`
	ToAmount *string `json:"toAmount,omitempty"`
	ToAssetId string `json:"toAssetId"`
}
