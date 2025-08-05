package types

type SwapParams struct {
	FromAmount *string `json:"fromAmount"`
	FromAssetId string `json:"fromAssetId"`
	QuoteId string `json:"quoteId"`
	ToAmount *string `json:"toAmount"`
	ToAssetId string `json:"toAssetId"`
}
