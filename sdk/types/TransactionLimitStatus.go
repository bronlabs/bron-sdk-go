package types

type TransactionLimitStatus string

const (
	TransactionLimitStatus_NEW TransactionLimitStatus = "new"
	TransactionLimitStatus_ACTIVE TransactionLimitStatus = "active"
	TransactionLimitStatus_DEACTIVATED TransactionLimitStatus = "deactivated"
	TransactionLimitStatus_DECLINED TransactionLimitStatus = "declined"
)
