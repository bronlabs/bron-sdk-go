package types

type Account struct {
	AccountType AccountType `json:"accountType"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	Extra *AccountExtra `json:"extra"`
	IsTestnet *bool `json:"isTestnet"`
	ParentAccountId *string `json:"parentAccountId"`
	Status AccountStatus `json:"status"`
	WorkspaceId string `json:"workspaceId"`
	AccountId string `json:"accountId"`
	AccountName string `json:"accountName"`
	ExternalId string `json:"externalId"`
	UpdatedAt *string `json:"updatedAt"`
}
