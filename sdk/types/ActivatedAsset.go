package types

type ActivatedAsset struct {
	ActivationID *string `json:"activationId,omitempty"`
	Address *string `json:"address,omitempty"`
	AssetID *string `json:"assetId,omitempty"`
	Status *AddressStatus `json:"status,omitempty"`
}
