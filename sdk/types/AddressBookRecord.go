package types

type AddressBookRecord struct {
	AccountIDs *[]string `json:"accountIds,omitempty"`
	ActivatedAt *string `json:"activatedAt,omitempty"`
	Address *string `json:"address,omitempty"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy,omitempty"`
	Email *string `json:"email,omitempty"`
	ExternalID string `json:"externalId"`
	ImageID *string `json:"imageId,omitempty"`
	IsFavorite *bool `json:"isFavorite,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	LastUsedAt *string `json:"lastUsedAt,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Name string `json:"name"`
	NetworkID *string `json:"networkId,omitempty"`
	RecordID string `json:"recordId"`
	RecordType RecordType `json:"recordType"`
	Status RecordStatus `json:"status"`
	Tag *string `json:"tag,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
	WorkspaceID string `json:"workspaceId"`
}
