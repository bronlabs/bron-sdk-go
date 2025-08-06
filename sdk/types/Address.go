package types

type Address struct {
	AcceptsAllAssets bool `json:"acceptsAllAssets"`
	AccountId *string `json:"accountId,omitempty"`
	AccountType AccountType `json:"accountType"`
	ActivatedAssets *[]ActivatedAsset `json:"activatedAssets,omitempty"`
	Address *string `json:"address,omitempty"`
	AddressId string `json:"addressId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	ExternalId string `json:"externalId"`
	Memo *string `json:"memo,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	NetworkId string `json:"networkId"`
	RequiresAssetsActivation bool `json:"requiresAssetsActivation"`
	Status AddressStatus `json:"status"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
	WorkspaceId *string `json:"workspaceId,omitempty"`
}
