package types

type CreateAddressBookRecord struct {
	AccountIDs *[]string `json:"accountIds,omitempty"`
	Address *string `json:"address,omitempty"`
	ExternalID string `json:"externalId"`
	ImageID *string `json:"imageId,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Name string `json:"name"`
	NetworkID *string `json:"networkId,omitempty"`
	RecordType *RecordType `json:"recordType,omitempty"`
	Tag *string `json:"tag,omitempty"`
}
