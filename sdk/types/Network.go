package types

type Network struct {
	AddressExplorerURL *string `json:"addressExplorerUrl,omitempty"`
	Confirmations *string `json:"confirmations,omitempty"`
	ExplorerURL *string `json:"explorerUrl,omitempty"`
	IsTestnet *bool `json:"isTestnet,omitempty"`
	Name *string `json:"name,omitempty"`
	NativeAssetID *string `json:"nativeAssetId,omitempty"`
	NativeAssetSymbol *string `json:"nativeAssetSymbol,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	Tags *[]NetworkTag `json:"tags,omitempty"`
}
