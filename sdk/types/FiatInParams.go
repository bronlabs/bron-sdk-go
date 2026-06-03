package types

type FiatInParams struct {
	Amount string `json:"amount"`
	AssetID string `json:"assetId"`
	FiatAssetID string `json:"fiatAssetId"`
}
