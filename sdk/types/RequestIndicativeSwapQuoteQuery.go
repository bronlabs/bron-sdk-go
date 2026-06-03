package types

type RequestIndicativeSwapQuoteQuery struct {
	FromAssetID string `json:"fromAssetId"`
	ToAssetID string `json:"toAssetId"`
	FromAmount *string `json:"fromAmount,omitempty"`
	ToAmount *string `json:"toAmount,omitempty"`
}
