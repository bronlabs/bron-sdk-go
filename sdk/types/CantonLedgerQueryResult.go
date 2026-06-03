package types

type CantonLedgerQueryResult struct {
	Result interface{} `json:"result,omitempty"`
	Status *string `json:"status,omitempty"`
}
