package types

type AddressBookRecord struct {
	AccountIds *[]string `json:"accountIds,omitempty"`
	Address string `json:"address"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy,omitempty"`
	ExternalId string `json:"externalId"`
	LastUsedAt *string `json:"lastUsedAt,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Name string `json:"name"`
	NetworkId string `json:"networkId"`
	RecordId string `json:"recordId"`
	Status RecordStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
	WorkspaceId string `json:"workspaceId"`
}
