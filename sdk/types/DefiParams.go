package types

type DefiParams struct {
	Data *string `json:"data,omitempty"`
	ExternalBroadcast *bool `json:"externalBroadcast,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	Method string `json:"method"`
	NetworkID string `json:"networkId"`
	Origin string `json:"origin"`
	RawTransaction *string `json:"rawTransaction,omitempty"`
	RawTransactions *[]string `json:"rawTransactions,omitempty"`
	To *string `json:"to,omitempty"`
	Value *string `json:"value,omitempty"`
}
