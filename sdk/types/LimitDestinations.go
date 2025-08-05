package types

type LimitDestinations struct {
	AddressBookRecordIds *[]string `json:"addressBookRecordIds"`
	ToAccounts *bool `json:"toAccounts"`
	ToAddressBook *bool `json:"toAddressBook"`
	ToExternalAddresses *bool `json:"toExternalAddresses"`
	AccountIds *[]string `json:"accountIds"`
}
