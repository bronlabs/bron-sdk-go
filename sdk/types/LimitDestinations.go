package types

type LimitDestinations struct {
	AccountIds *[]string `json:"accountIds"`
	AddressBookRecordIds *[]string `json:"addressBookRecordIds"`
	ToAccounts *bool `json:"toAccounts"`
	ToAddressBook *bool `json:"toAddressBook"`
	ToExternalAddresses *bool `json:"toExternalAddresses"`
}
