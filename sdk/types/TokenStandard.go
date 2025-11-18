package types

type TokenStandard string

const (
	TokenStandard_ERC20 TokenStandard = "erc20"
	TokenStandard_ERC721 TokenStandard = "erc721"
	TokenStandard_ERC1155 TokenStandard = "erc1155"
	TokenStandard_SPL TokenStandard = "spl"
	TokenStandard_TRC20 TokenStandard = "trc20"
)
