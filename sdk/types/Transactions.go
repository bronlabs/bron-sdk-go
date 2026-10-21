package types

type Transactions struct {
	Embedded *TransactionEmbedded `json:"_embedded,omitempty"`
	Transactions []Transaction `json:"transactions"`
}
