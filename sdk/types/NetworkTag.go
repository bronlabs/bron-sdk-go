package types

type NetworkTag string

const (
	NetworkTag_SHOW_VAULT NetworkTag = "show-vault"
	NetworkTag_SUPPORTS_RBF NetworkTag = "supports-rbf"
	NetworkTag_SUPPORTS_MEMO NetworkTag = "supports-memo"
	NetworkTag_SWAP NetworkTag = "swap"
	NetworkTag_SUPPORTS_PARALLEL_SIGNING NetworkTag = "supports-parallel-signing"
	NetworkTag_SUPPORTS_CHAINED_SIGNING NetworkTag = "supports-chained-signing"
)
