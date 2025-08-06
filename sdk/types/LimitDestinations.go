package types

type LimitDestinations struct {
	AccountIds *[]string `json:"accountIds,omitempty"`
	AddressBookRecordIds *[]string `json:"addressBookRecordIds,omitempty"`
	ToAccounts *bool `json:"toAccounts,omitempty"`
	ToAddressBook *bool `json:"toAddressBook,omitempty"`
	ToExternalAddresses *bool `json:"toExternalAddresses,omitempty"`
}
