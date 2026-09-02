package types

type PublicIntentPair struct {
	AssetA IntentPairAsset `json:"assetA"`
	AssetB IntentPairAsset `json:"assetB"`
	IsBidirectional bool `json:"isBidirectional"`
}
