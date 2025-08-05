package types

type TransactionType string

const (
	TransactionType_DEPOSIT TransactionType = "deposit"
	TransactionType_WITHDRAWAL TransactionType = "withdrawal"
	TransactionType_MULTI_WITHDRAWAL TransactionType = "multi-withdrawal"
	TransactionType_NEGATIVE_DEPOSIT TransactionType = "negative-deposit"
	TransactionType_AUTO_WITHDRAWAL TransactionType = "auto-withdrawal"
	TransactionType_ALLOWANCE TransactionType = "allowance"
	TransactionType_RAW_TRANSACTION TransactionType = "raw-transaction"
	TransactionType_ADDRESS_ACTIVATION TransactionType = "address-activation"
	TransactionType_ADDRESS_CREATION TransactionType = "address-creation"
	TransactionType_SWAP_LIFI TransactionType = "swap-lifi"
	TransactionType_INTENTS TransactionType = "intents"
)
