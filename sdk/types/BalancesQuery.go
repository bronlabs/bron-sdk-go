package types

type BalancesQuery struct {
	AccountIds *[]string `json:"accountIds,omitempty"`
	BalanceIds *[]string `json:"balanceIds,omitempty"`
	AssetIds *[]string `json:"assetIds,omitempty"`
	NetworkId *string `json:"networkId,omitempty"`
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	ExcludedAccountTypes *[]AccountType `json:"excludedAccountTypes,omitempty"`
	NonEmpty *bool `json:"nonEmpty,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
