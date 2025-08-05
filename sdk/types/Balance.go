package types

type Balance struct {
	BalanceId string `json:"balanceId"`
	CreatedAt *string `json:"createdAt"`
	NetworkId *string `json:"networkId"`
	Symbol *string `json:"symbol"`
	TotalBalance *string `json:"totalBalance"`
	WorkspaceId string `json:"workspaceId"`
	AccountType AccountType `json:"accountType"`
	AssetId string `json:"assetId"`
	UpdatedAt *string `json:"updatedAt"`
	AccountId string `json:"accountId"`
}
