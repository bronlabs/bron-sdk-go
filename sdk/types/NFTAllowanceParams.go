package types

type NFTAllowanceParams struct {
	Amount *string `json:"amount,omitempty"`
	ApprovalForAll *bool `json:"approvalForAll,omitempty"`
	AssetID string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	NetworkFees *RequestedNetworkFees `json:"networkFees,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	TokenID *string `json:"tokenId,omitempty"`
}
