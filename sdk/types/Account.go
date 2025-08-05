package types

type Account struct {
	Extra *AccountExtra `json:"extra"`
	WorkspaceId string `json:"workspaceId"`
	AccountName string `json:"accountName"`
	AccountType AccountType `json:"accountType"`
	CreatedBy *string `json:"createdBy"`
	ExternalId string `json:"externalId"`
	IsTestnet *bool `json:"isTestnet"`
	ParentAccountId *string `json:"parentAccountId"`
	Status AccountStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt"`
	AccountId string `json:"accountId"`
	CreatedAt string `json:"createdAt"`
}
