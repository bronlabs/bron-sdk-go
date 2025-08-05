package types

type Account struct {
	UpdatedAt *string `json:"updatedAt"`
	WorkspaceId string `json:"workspaceId"`
	AccountId string `json:"accountId"`
	CreatedAt string `json:"createdAt"`
	ExternalId string `json:"externalId"`
	Status AccountStatus `json:"status"`
	AccountName string `json:"accountName"`
	AccountType AccountType `json:"accountType"`
	CreatedBy *string `json:"createdBy"`
	Extra *AccountExtra `json:"extra"`
	IsTestnet *bool `json:"isTestnet"`
	ParentAccountId *string `json:"parentAccountId"`
}
