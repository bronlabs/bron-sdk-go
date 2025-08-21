package types

type Balance struct {
	AccountID string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	AssetID string `json:"assetId"`
	BalanceID string `json:"balanceId"`
	CreatedAt *string `json:"createdAt,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TotalBalance *string `json:"totalBalance,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	WorkspaceID string `json:"workspaceId"`
}
