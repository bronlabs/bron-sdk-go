package types

type TransactionExtra struct {
	Approvers *TransactionApprovers `json:"approvers,omitempty"`
	BlockchainDetails *[]BlockchainTxDetails `json:"blockchainDetails,omitempty"`
	BlockchainRequest *BlockchainRequest `json:"blockchainRequest,omitempty"`
	BronLock *BronLockExtra `json:"bronLock,omitempty"`
	Confirmations *string `json:"confirmations,omitempty"`
	DepositTransactionID *string `json:"depositTransactionId,omitempty"`
	Description *string `json:"description,omitempty"`
	EstimatedPrice *string `json:"estimatedPrice,omitempty"`
	FiatRamp *FiatRampExtra `json:"fiatRamp,omitempty"`
	FromAccountID *string `json:"fromAccountId,omitempty"`
	FromAddress *string `json:"fromAddress,omitempty"`
	FromWorkspaceIcon *string `json:"fromWorkspaceIcon,omitempty"`
	FromWorkspaceName *string `json:"fromWorkspaceName,omitempty"`
	FromWorkspaceTag *string `json:"fromWorkspaceTag,omitempty"`
	Loan *LoanExtra `json:"loan,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Price *string `json:"price,omitempty"`
	QuoteDetails *QuoteDetails `json:"quoteDetails,omitempty"`
	QuoteID *string `json:"quoteId,omitempty"`
	SigningRequestID *string `json:"signingRequestId,omitempty"`
	ToAccountID *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	ToWorkspaceIcon *string `json:"toWorkspaceIcon,omitempty"`
	ToWorkspaceName *string `json:"toWorkspaceName,omitempty"`
	ToWorkspaceTag *string `json:"toWorkspaceTag,omitempty"`
	WithdrawTransactionID *string `json:"withdrawTransactionId,omitempty"`
}
