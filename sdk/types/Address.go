package types

type Address struct {
	AcceptsAllAssets bool `json:"acceptsAllAssets"`
	AccountId *string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	ActivatedAssets *[]ActivatedAsset `json:"activatedAssets"`
	Address *string `json:"address"`
	AddressId string `json:"addressId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	ExternalId string `json:"externalId"`
	Memo *string `json:"memo"`
	Metadata *map[string]interface{} `json:"metadata"`
	NetworkId string `json:"networkId"`
	RequiresAssetsActivation bool `json:"requiresAssetsActivation"`
	Status AddressStatus `json:"status"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
	WorkspaceId *string `json:"workspaceId"`
}
