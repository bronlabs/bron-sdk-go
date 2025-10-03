package types

type EventOutput struct {
	Address *string `json:"address,omitempty"`
	Amount *string `json:"amount,omitempty"`
	Memo *string `json:"memo,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	ToAccountID *string `json:"toAccountId,omitempty"`
	ToAddressBookRecordID *string `json:"toAddressBookRecordId,omitempty"`
}
