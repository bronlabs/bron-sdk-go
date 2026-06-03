package types

type BankChannelType string

const (
	BankChannelType_SEPA BankChannelType = "sepa"
	BankChannelType_ACH BankChannelType = "ach"
	BankChannelType_SWIFT BankChannelType = "swift"
	BankChannelType_FEDWIRE BankChannelType = "fedwire"
	BankChannelType_LOCAL BankChannelType = "local"
)
