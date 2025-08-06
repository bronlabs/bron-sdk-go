package types

type CreateAddressBookRecord struct {
	AccountIds *[]string `json:"accountIds,omitempty"`
	Address string `json:"address"`
	ExternalId string `json:"externalId"`
	Memo *string `json:"memo,omitempty"`
	Name string `json:"name"`
	NetworkId string `json:"networkId"`
}
