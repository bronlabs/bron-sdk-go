package types

type SwapParams struct {
	ToAmount *string `json:"toAmount"`
	ToAssetId string `json:"toAssetId"`
	FromAmount *string `json:"fromAmount"`
	FromAssetId string `json:"fromAssetId"`
	QuoteId string `json:"quoteId"`
}
