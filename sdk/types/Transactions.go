package types

type Transactions struct {
	Embedded *TransactionEmbedded `json:"embedded,omitempty"`
	Transactions []Transaction `json:"transactions"`
}
