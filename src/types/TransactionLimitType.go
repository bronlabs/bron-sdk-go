package types

type TransactionLimitType string

const (
	TransactionLimitType_TRANSACTIONS_VOLUME TransactionLimitType = "transactions-volume"
	TransactionLimitType_TRANSACTION_AMOUNT TransactionLimitType = "transaction-amount"
)
