package types

type BankAddress struct {
	Address *string `json:"address,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	State *string `json:"state,omitempty"`
}
