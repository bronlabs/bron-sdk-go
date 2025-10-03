package types

type EventAllowance struct {
	Address *string `json:"address,omitempty"`
	Amount *string `json:"amount,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	Unlimited *bool `json:"unlimited,omitempty"`
}
