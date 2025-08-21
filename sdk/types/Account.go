package types

type Account struct {
	AccountID string `json:"accountId"`
	AccountName string `json:"accountName"`
	AccountType AccountType `json:"accountType"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy,omitempty"`
	ExternalID string `json:"externalId"`
	Extra *AccountExtra `json:"extra,omitempty"`
	IsTestnet *bool `json:"isTestnet,omitempty"`
	ParentAccountID *string `json:"parentAccountId,omitempty"`
	Status AccountStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	WorkspaceID string `json:"workspaceId"`
}
