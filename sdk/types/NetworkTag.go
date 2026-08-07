package types

type NetworkTag string

const (
	NetworkTag_SHOW_VAULT NetworkTag = "show-vault"
	NetworkTag_SUPPORTS_RBF NetworkTag = "supports-rbf"
	NetworkTag_SUPPORTS_RBF_CANCEL NetworkTag = "supports-rbf-cancel"
	NetworkTag_SUPPORTS_MEMO NetworkTag = "supports-memo"
	NetworkTag_SWAP NetworkTag = "swap"
	NetworkTag_SUPPORTS_PARALLEL_SIGNING NetworkTag = "supports-parallel-signing"
	NetworkTag_SUPPORTS_CHAINED_SIGNING NetworkTag = "supports-chained-signing"
	NetworkTag_SUPPORTS_FEE_LEVELS NetworkTag = "supports-fee-levels"
	NetworkTag_SUPPORTS_SPONSORED_TRANSACTIONS NetworkTag = "supports-sponsored-transactions"
	NetworkTag_HAS_TOKEN_ACTIVATION NetworkTag = "has-token-activation"
	NetworkTag_HAS_ADDRESS_CREATION NetworkTag = "has-address-creation"
	NetworkTag_HAS_ADDRESS_SERVER_CREATION NetworkTag = "has-address-server-creation"
	NetworkTag_HAS_ADDRESS_ACTIVATION NetworkTag = "has-address-activation"
	NetworkTag_EVM NetworkTag = "evm"
	NetworkTag_BRIDGE NetworkTag = "bridge"
)
