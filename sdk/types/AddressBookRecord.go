package types

type AddressBookRecord struct {
	AccountIds *[]string `json:"accountIds"`
	Address string `json:"address"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	ExternalId string `json:"externalId"`
	LastUsedAt *string `json:"lastUsedAt"`
	Memo *string `json:"memo"`
	Name string `json:"name"`
	NetworkId string `json:"networkId"`
	RecordId string `json:"recordId"`
	Status RecordStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt"`
	UpdatedBy *string `json:"updatedBy"`
	WorkspaceId string `json:"workspaceId"`
}
