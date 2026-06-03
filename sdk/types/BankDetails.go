package types

type BankDetails struct {
	AccountHolderType AccountHolderType `json:"accountHolderType"`
	AccountNumber string `json:"accountNumber"`
	AccountType *BankAccountType `json:"accountType,omitempty"`
	BankAddress *BankAddress `json:"bankAddress,omitempty"`
	BankCode *string `json:"bankCode,omitempty"`
	BusinessName *string `json:"businessName,omitempty"`
	BusinessRegistrationNumber *string `json:"businessRegistrationNumber,omitempty"`
	ChannelType BankChannelType `json:"channelType"`
	CorrespondentBankCode *string `json:"correspondentBankCode,omitempty"`
	Country string `json:"country"`
	FiatCurrency string `json:"fiatCurrency"`
	FirstName *string `json:"firstName,omitempty"`
	Issuer string `json:"issuer"`
	LastName *string `json:"lastName,omitempty"`
	PaymentPurpose *string `json:"paymentPurpose,omitempty"`
	Reference *string `json:"reference,omitempty"`
	RegisteredAddress *string `json:"registeredAddress,omitempty"`
}
