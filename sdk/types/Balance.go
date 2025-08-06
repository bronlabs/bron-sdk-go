package types

type Balance struct {
	AccountId string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	AssetId string `json:"assetId"`
	BalanceId string `json:"balanceId"`
	CreatedAt *string `json:"createdAt,omitempty"`
	NetworkId *string `json:"networkId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TotalBalance *string `json:"totalBalance,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	WorkspaceId string `json:"workspaceId"`
}
