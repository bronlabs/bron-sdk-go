package types

type Account struct {
	AccountId string `json:"accountId"`
	AccountName string `json:"accountName"`
	AccountType AccountType `json:"accountType"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	ExternalId string `json:"externalId"`
	Extra *AccountExtra `json:"extra"`
	IsTestnet *bool `json:"isTestnet"`
	ParentAccountId *string `json:"parentAccountId"`
	Status AccountStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt"`
	WorkspaceId string `json:"workspaceId"`
}
