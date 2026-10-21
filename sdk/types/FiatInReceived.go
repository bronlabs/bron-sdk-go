package types

type FiatInReceived struct {
	FiatAmount *string `json:"fiatAmount,omitempty"`
	Reference *string `json:"reference,omitempty"`
	SenderAccountNumber *string `json:"senderAccountNumber,omitempty"`
	SenderName *string `json:"senderName,omitempty"`
}
