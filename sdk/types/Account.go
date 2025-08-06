package types

type Account struct {
	AccountId string `json:"accountId"`
	AccountName string `json:"accountName"`
	AccountType AccountType `json:"accountType"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy,omitempty"`
	ExternalId string `json:"externalId"`
	Extra *AccountExtra `json:"extra,omitempty"`
	IsTestnet *bool `json:"isTestnet,omitempty"`
	ParentAccountId *string `json:"parentAccountId,omitempty"`
	Status AccountStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	WorkspaceId string `json:"workspaceId"`
}
