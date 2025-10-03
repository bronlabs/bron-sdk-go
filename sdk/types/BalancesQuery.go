package types

type BalancesQuery struct {
	AccountIDs *[]string `json:"accountIds,omitempty"`
	BalanceIDs *[]string `json:"balanceIds,omitempty"`
	AssetIDs *[]string `json:"assetIds,omitempty"`
	NetworkIDs *[]string `json:"networkIds,omitempty"`
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	ExcludedAccountTypes *[]AccountType `json:"excludedAccountTypes,omitempty"`
	NonEmpty *bool `json:"nonEmpty,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
