package types

type Network struct {
	AddressExplorerUrl *string `json:"addressExplorerUrl"`
	Confirmations *string `json:"confirmations"`
	ExplorerUrl *string `json:"explorerUrl"`
	IsTestnet *bool `json:"isTestnet"`
	Name *string `json:"name"`
	NativeAssetId *string `json:"nativeAssetId"`
	NativeAssetSymbol *string `json:"nativeAssetSymbol"`
	NetworkId *string `json:"networkId"`
	Tags *[]NetworkTag `json:"tags"`
}
