package types

type EventInput struct {
	Address *string `json:"address,omitempty"`
	Amount *string `json:"amount,omitempty"`
	FromAccountID *string `json:"fromAccountId,omitempty"`
	FromAddressBookRecordID *string `json:"fromAddressBookRecordId,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
}
