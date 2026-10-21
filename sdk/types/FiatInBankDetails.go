package types

type FiatInBankDetails struct {
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	BankAddress *BankAddress `json:"bankAddress,omitempty"`
	BankCode *string `json:"bankCode,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ChannelType *BankChannelType `json:"channelType,omitempty"`
}
