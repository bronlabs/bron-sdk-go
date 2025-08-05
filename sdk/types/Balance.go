package types

type Balance struct {
	AccountId string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	AssetId string `json:"assetId"`
	BalanceId string `json:"balanceId"`
	CreatedAt *string `json:"createdAt"`
	NetworkId *string `json:"networkId"`
	Symbol *string `json:"symbol"`
	TotalBalance *string `json:"totalBalance"`
	UpdatedAt *string `json:"updatedAt"`
	WorkspaceId string `json:"workspaceId"`
}
