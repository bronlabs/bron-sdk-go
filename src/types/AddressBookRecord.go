package types

type AddressBookRecord struct {
	UpdatedBy *string `json:"updatedBy"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	ExternalId string `json:"externalId"`
	Name string `json:"name"`
	Status RecordStatus `json:"status"`
	WorkspaceId string `json:"workspaceId"`
	AccountIds *[]string `json:"accountIds"`
	Address string `json:"address"`
	LastUsedAt *string `json:"lastUsedAt"`
	Memo *string `json:"memo"`
	NetworkId string `json:"networkId"`
	RecordId string `json:"recordId"`
	UpdatedAt *string `json:"updatedAt"`
}
