package types

type TransactionLimitType string

const (
	TransactionLimitType_TRANSACTIONS_VOLUME TransactionLimitType = "transactions-volume"
	TransactionLimitType_TRANSACTION_AMOUNT TransactionLimitType = "transaction-amount"
	TransactionLimitType_HYPERLIQUID_ORDERS_VOLUME TransactionLimitType = "hyperliquid-orders-volume"
	TransactionLimitType_HYPERLIQUID_ORDER_AMOUNT TransactionLimitType = "hyperliquid-order-amount"
)
