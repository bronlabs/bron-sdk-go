package types

type AllowanceParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	ToAddress string `json:"toAddress"`
	Unlimited *bool `json:"unlimited,omitempty"`
}
