package types

type LimitDestinations struct {
	AccountIDs *[]string `json:"accountIds,omitempty"`
	AddressBookRecordIDs *[]string `json:"addressBookRecordIds,omitempty"`
	ToAccounts *bool `json:"toAccounts,omitempty"`
	ToAddressBook *bool `json:"toAddressBook,omitempty"`
	ToExternalAddresses *bool `json:"toExternalAddresses,omitempty"`
}
