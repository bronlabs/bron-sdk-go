package types

type Address struct {
	Status AddressStatus `json:"status"`
	RequiresAssetsActivation bool `json:"requiresAssetsActivation"`
	AccountId *string `json:"accountId"`
	AddressId string `json:"addressId"`
	ExternalId string `json:"externalId"`
	Memo *string `json:"memo"`
	Metadata *map[string]interface{} `json:"metadata"`
	UpdatedAt string `json:"updatedAt"`
	WorkspaceId *string `json:"workspaceId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
	AcceptsAllAssets bool `json:"acceptsAllAssets"`
	ActivatedAssets *[]ActivatedAsset `json:"activatedAssets"`
	Address *string `json:"address"`
	AccountType AccountType `json:"accountType"`
	NetworkId string `json:"networkId"`
}
