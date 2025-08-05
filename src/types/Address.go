package types

type Address struct {
	AccountType AccountType `json:"accountType"`
	CreatedAt string `json:"createdAt"`
	Memo *string `json:"memo"`
	RequiresAssetsActivation bool `json:"requiresAssetsActivation"`
	Status AddressStatus `json:"status"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
	WorkspaceId *string `json:"workspaceId"`
	AccountId *string `json:"accountId"`
	ActivatedAssets *[]ActivatedAsset `json:"activatedAssets"`
	Address *string `json:"address"`
	AddressId string `json:"addressId"`
	ExternalId string `json:"externalId"`
	AcceptsAllAssets bool `json:"acceptsAllAssets"`
	CreatedBy string `json:"createdBy"`
	Metadata *map[string]interface{} `json:"metadata"`
	NetworkId string `json:"networkId"`
}
