package types

type AddressBookRecord struct {
	Name string `json:"name"`
	RecordId string `json:"recordId"`
	Status RecordStatus `json:"status"`
	WorkspaceId string `json:"workspaceId"`
	AccountIds *[]string `json:"accountIds"`
	Address string `json:"address"`
	ExternalId string `json:"externalId"`
	LastUsedAt *string `json:"lastUsedAt"`
	NetworkId string `json:"networkId"`
	UpdatedAt *string `json:"updatedAt"`
	UpdatedBy *string `json:"updatedBy"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	Memo *string `json:"memo"`
}
