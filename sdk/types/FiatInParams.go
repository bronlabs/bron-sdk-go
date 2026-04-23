package types

type FiatInParams struct {
	Amount string `json:"amount"`
	AssetID string `json:"assetId"`
	FiatAmount *string `json:"fiatAmount,omitempty"`
	FiatAssetID string `json:"fiatAssetId"`
}
