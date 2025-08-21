package types

type CreateAddressBookRecord struct {
	AccountIDs *[]string `json:"accountIds,omitempty"`
	Address string `json:"address"`
	ExternalID string `json:"externalId"`
	Memo *string `json:"memo,omitempty"`
	Name string `json:"name"`
	NetworkID string `json:"networkId"`
}
