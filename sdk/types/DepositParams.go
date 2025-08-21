package types

type DepositParams struct {
	Amount string `json:"amount"`
	AssetID string `json:"assetId"`
	NetworkID string `json:"networkId"`
}
