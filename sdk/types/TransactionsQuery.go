package types

type TransactionsQuery struct {
	TransactionIDs *[]string `json:"transactionIds,omitempty"`
	TransactionTypes *[]TransactionType `json:"transactionTypes,omitempty"`
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	AccountIDs *[]string `json:"accountIds,omitempty"`
	TransactionStatuses *[]TransactionStatus `json:"transactionStatuses,omitempty"`
	TransactionStatusNotIn *[]TransactionStatus `json:"transactionStatusNotIn,omitempty"`
	BlockchainTxID *string `json:"blockchainTxId,omitempty"`
	ToAccountID *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	IsTerminated *bool `json:"isTerminated,omitempty"`
	TerminatedAtFrom *string `json:"terminatedAtFrom,omitempty"`
	TerminatedAtTo *string `json:"terminatedAtTo,omitempty"`
	CreatedAtFrom *string `json:"createdAtFrom,omitempty"`
	CreatedAtTo *string `json:"createdAtTo,omitempty"`
	UpdatedAtFrom *string `json:"updatedAtFrom,omitempty"`
	UpdatedAtTo *string `json:"updatedAtTo,omitempty"`
	CanSignWithDeviceID *string `json:"canSignWithDeviceId,omitempty"`
	SortDirection *SortingDirection `json:"sortDirection,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
	ExternalID *string `json:"externalId,omitempty"`
	IncludeEvents *bool `json:"includeEvents,omitempty"`
}
