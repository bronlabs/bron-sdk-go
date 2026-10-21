package types

type CreateAddressBookRecord struct {
	AccountIDs *[]string `json:"accountIds,omitempty"`
	Address *string `json:"address,omitempty"`
	BankDetails *BankDetails `json:"bankDetails,omitempty"`
	Email *string `json:"email,omitempty"`
	ExternalID string `json:"externalId"`
	ImageID *string `json:"imageId,omitempty"`
	IsFavorite *bool `json:"isFavorite,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Name string `json:"name"`
	NetworkID *string `json:"networkId,omitempty"`
	RecordType *RecordType `json:"recordType,omitempty"`
	Tag *string `json:"tag,omitempty"`
}
