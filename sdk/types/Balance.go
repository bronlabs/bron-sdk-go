package types

type Balance struct {
	AssetId string `json:"assetId"`
	TotalBalance *string `json:"totalBalance"`
	UpdatedAt *string `json:"updatedAt"`
	WorkspaceId string `json:"workspaceId"`
	AccountId string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	BalanceId string `json:"balanceId"`
	CreatedAt *string `json:"createdAt"`
	NetworkId *string `json:"networkId"`
	Symbol *string `json:"symbol"`
}
