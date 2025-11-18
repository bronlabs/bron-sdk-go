package types

type TransactionType string

const (
	TransactionType_DEPOSIT TransactionType = "deposit"
	TransactionType_WITHDRAWAL TransactionType = "withdrawal"
	TransactionType_MULTI_WITHDRAWAL TransactionType = "multi-withdrawal"
	TransactionType_NEGATIVE_DEPOSIT TransactionType = "negative-deposit"
	TransactionType_AUTO_WITHDRAWAL TransactionType = "auto-withdrawal"
	TransactionType_ALLOWANCE TransactionType = "allowance"
	TransactionType_DEFI TransactionType = "defi"
	TransactionType_DEFI_MESSAGE TransactionType = "defi-message"
	TransactionType_ADDRESS_ACTIVATION TransactionType = "address-activation"
	TransactionType_ADDRESS_CREATION TransactionType = "address-creation"
	TransactionType_SWAP_LIFI TransactionType = "swap-lifi"
	TransactionType_INTENTS TransactionType = "intents"
	TransactionType_LOYALTY_LOCK TransactionType = "loyalty-lock"
	TransactionType_LOYALTY_UNLOCK TransactionType = "loyalty-unlock"
	TransactionType_LOYALTY_COLLECT_REWARDS TransactionType = "loyalty-collect-rewards"
	TransactionType_CANTON_REWARD TransactionType = "canton-reward"
	TransactionType_NFT_DEPOSIT TransactionType = "nft-deposit"
	TransactionType_NFT_WITHDRAWAL TransactionType = "nft-withdrawal"
	TransactionType_NFT_ALLOWANCE TransactionType = "nft-allowance"
)
