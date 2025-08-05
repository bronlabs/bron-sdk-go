package types

type AddressBookRecord struct {
	AccountIds *[]string `json:"accountIds"`
	CreatedAt string `json:"createdAt"`
	Memo *string `json:"memo"`
	RecordId string `json:"recordId"`
	Status RecordStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt"`
	WorkspaceId string `json:"workspaceId"`
	Address string `json:"address"`
	CreatedBy *string `json:"createdBy"`
	ExternalId string `json:"externalId"`
	LastUsedAt *string `json:"lastUsedAt"`
	Name string `json:"name"`
	NetworkId string `json:"networkId"`
	UpdatedBy *string `json:"updatedBy"`
}
