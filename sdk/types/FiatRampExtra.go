package types

type FiatRampExtra struct {
	Amount *string `json:"amount,omitempty"`
	BankDetails *FiatInBankDetails `json:"bankDetails,omitempty"`
	FiatAmount *string `json:"fiatAmount,omitempty"`
	FiatInReceived *FiatInReceived `json:"fiatInReceived,omitempty"`
	PlatformFee *string `json:"platformFee,omitempty"`
	Provider *string `json:"provider,omitempty"`
	ProviderFee *string `json:"providerFee,omitempty"`
	Rate *string `json:"rate,omitempty"`
}
