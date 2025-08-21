package types

type AddressBookRecord struct {
	AccountIDs *[]string `json:"accountIds,omitempty"`
	Address string `json:"address"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy,omitempty"`
	ExternalID string `json:"externalId"`
	LastUsedAt *string `json:"lastUsedAt,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Name string `json:"name"`
	NetworkID string `json:"networkId"`
	RecordID string `json:"recordId"`
	Status RecordStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
	WorkspaceID string `json:"workspaceId"`
}
