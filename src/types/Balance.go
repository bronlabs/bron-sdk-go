package types

type Balance struct {
	AccountId string `json:"accountId"`
	AssetId string `json:"assetId"`
	BalanceId string `json:"balanceId"`
	AccountType AccountType `json:"accountType"`
	CreatedAt *string `json:"createdAt"`
	NetworkId *string `json:"networkId"`
	Symbol *string `json:"symbol"`
	TotalBalance *string `json:"totalBalance"`
	UpdatedAt *string `json:"updatedAt"`
	WorkspaceId string `json:"workspaceId"`
}
