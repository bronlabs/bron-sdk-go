package types

type Address struct {
	AcceptsAllAssets bool `json:"acceptsAllAssets"`
	AccountId *string `json:"accountId"`
	Address *string `json:"address"`
	Metadata *map[string]interface{} `json:"metadata"`
	NetworkId string `json:"networkId"`
	AccountType AccountType `json:"accountType"`
	ActivatedAssets *[]ActivatedAsset `json:"activatedAssets"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	ExternalId string `json:"externalId"`
	Memo *string `json:"memo"`
	UpdatedAt string `json:"updatedAt"`
	AddressId string `json:"addressId"`
	RequiresAssetsActivation bool `json:"requiresAssetsActivation"`
	Status AddressStatus `json:"status"`
	WorkspaceId *string `json:"workspaceId"`
	UpdatedBy string `json:"updatedBy"`
}
