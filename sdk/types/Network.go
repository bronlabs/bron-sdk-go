package types

type Network struct {
	AddressExplorerUrl *string `json:"addressExplorerUrl,omitempty"`
	Confirmations *string `json:"confirmations,omitempty"`
	ExplorerUrl *string `json:"explorerUrl,omitempty"`
	IsTestnet *bool `json:"isTestnet,omitempty"`
	Name *string `json:"name,omitempty"`
	NativeAssetId *string `json:"nativeAssetId,omitempty"`
	NativeAssetSymbol *string `json:"nativeAssetSymbol,omitempty"`
	NetworkId *string `json:"networkId,omitempty"`
	Tags *[]NetworkTag `json:"tags,omitempty"`
}
