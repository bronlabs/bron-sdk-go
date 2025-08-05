package types

type CreateAddressBookRecord struct {
	AccountIds *[]string `json:"accountIds"`
	Address string `json:"address"`
	ExternalId string `json:"externalId"`
	Memo *string `json:"memo"`
	Name string `json:"name"`
	NetworkId string `json:"networkId"`
}
