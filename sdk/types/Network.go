package types

type Network struct {
	ExplorerUrl *string `json:"explorerUrl"`
	IsTestnet *bool `json:"isTestnet"`
	Name *string `json:"name"`
	NetworkId *string `json:"networkId"`
	Tags *[]NetworkTag `json:"tags"`
	AddressExplorerUrl *string `json:"addressExplorerUrl"`
	Confirmations *string `json:"confirmations"`
}
