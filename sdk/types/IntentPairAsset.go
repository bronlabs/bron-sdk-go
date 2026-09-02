package types

type IntentPairAsset struct {
	AssetID string `json:"assetId"`
	ContractAddress *string `json:"contractAddress,omitempty"`
	Decimals string `json:"decimals"`
	Name *string `json:"name,omitempty"`
	NetworkID string `json:"networkId"`
	Symbol string `json:"symbol"`
}
