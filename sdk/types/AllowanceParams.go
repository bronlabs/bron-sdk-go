package types

type AllowanceParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	NetworkFees *RequestedNetworkFees `json:"networkFees,omitempty"`
	ToAddress string `json:"toAddress"`
	Unlimited *bool `json:"unlimited,omitempty"`
}
