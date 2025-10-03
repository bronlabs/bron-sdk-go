package types

type TransactionExtra struct {
	Approvers *TransactionApprovers `json:"approvers,omitempty"`
	BlockchainDetails *[]BlockchainTxDetails `json:"blockchainDetails,omitempty"`
	BlockchainRequest *BlockchainRequest `json:"blockchainRequest,omitempty"`
	Confirmations *string `json:"confirmations,omitempty"`
	DepositTransactionID *string `json:"depositTransactionId,omitempty"`
	Description *string `json:"description,omitempty"`
	FromAccountID *string `json:"fromAccountId,omitempty"`
	FromAddress *string `json:"fromAddress,omitempty"`
	Memo *string `json:"memo,omitempty"`
	SigningRequestID *string `json:"signingRequestId,omitempty"`
	ToAccountID *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	WithdrawTransactionID *string `json:"withdrawTransactionId,omitempty"`
}
