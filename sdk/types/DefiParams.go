package types

type DefiParams struct {
	Data *string `json:"data,omitempty"`
	ExternalBroadcast *bool `json:"externalBroadcast,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	NetworkID string `json:"networkId"`
	Origin string `json:"origin"`
	RawTransactions *[]string `json:"rawTransactions,omitempty"`
	To *string `json:"to,omitempty"`
	Value *string `json:"value,omitempty"`
}
