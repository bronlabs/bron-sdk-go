package types

type CreateAddressBookRecord struct {
	Memo *string `json:"memo"`
	Name string `json:"name"`
	NetworkId string `json:"networkId"`
	AccountIds *[]string `json:"accountIds"`
	Address string `json:"address"`
	ExternalId string `json:"externalId"`
}
