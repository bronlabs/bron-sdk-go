package types

type TransactionsQuery struct {
	TransactionIds *[]string `json:"transactionIds,omitempty"`
	TransactionTypes *[]TransactionType `json:"transactionTypes,omitempty"`
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	AccountIds *[]string `json:"accountIds,omitempty"`
	TransactionStatuses *[]TransactionStatus `json:"transactionStatuses,omitempty"`
	TransactionStatusNotIn *[]TransactionStatus `json:"transactionStatusNotIn,omitempty"`
	BlockchainTxId *string `json:"blockchainTxId,omitempty"`
	ToAccountId *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	IsTerminated *bool `json:"isTerminated,omitempty"`
	TerminatedAtFrom *string `json:"terminatedAtFrom,omitempty"`
	TerminatedAtTo *string `json:"terminatedAtTo,omitempty"`
	CreatedAtFrom *string `json:"createdAtFrom,omitempty"`
	CreatedAtTo *string `json:"createdAtTo,omitempty"`
	UpdatedAtFrom *string `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo *string `json:"updatedAtTo,omitempty"`
	CanSignWithDeviceId *string `json:"canSignWithDeviceId,omitempty"`
	SortDirection *SortingDirection `json:"sortDirection,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
}
