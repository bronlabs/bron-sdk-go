package types

type QuoteDetails struct {
	ApprovalAddress *string `json:"approvalAddress,omitempty"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	FromAmount *string `json:"fromAmount,omitempty"`
	FromAmountUsd *string `json:"fromAmountUSD,omitempty"`
	IsCrossChain *bool `json:"isCrossChain,omitempty"`
	NetworkFee *string `json:"networkFee,omitempty"`
	NetworkFeeAssetID *string `json:"networkFeeAssetId,omitempty"`
	NetworkFeeUsd *string `json:"networkFeeUSD,omitempty"`
	ProcessingTime *string `json:"processingTime,omitempty"`
	Provider *string `json:"provider,omitempty"`
	QuoteID *string `json:"quoteId,omitempty"`
	ServiceFeeUsd *string `json:"serviceFeeUSD,omitempty"`
	Slippage *string `json:"slippage,omitempty"`
	ToAmount *string `json:"toAmount,omitempty"`
	ToAmountMin *string `json:"toAmountMin,omitempty"`
	ToAmountUsd *string `json:"toAmountUSD,omitempty"`
}
