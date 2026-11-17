package types

type IndicativeQuoteForAnIntentSwapPairQuery struct {
	FromAmount *string `json:"fromAmount,omitempty"`
	ToAmount *string `json:"toAmount,omitempty"`
	BroadcasterFeePercent *string `json:"broadcasterFeePercent,omitempty"`
}
