package types

type AccountsQuery struct {
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	ExcludedAccountTypes *[]AccountType `json:"excludedAccountTypes,omitempty"`
	Statuses *[]AccountStatus `json:"statuses,omitempty"`
	AccountIds *[]string `json:"accountIds,omitempty"`
	IsDefiVault *bool `json:"isDefiVault,omitempty"`
	Offset *string `json:"offset,omitempty"`
	Limit *string `json:"limit,omitempty"`
	IsTestnet *bool `json:"isTestnet,omitempty"`
}
