package types

type TransactionExtra struct {
	Approvers *TransactionApprovers `json:"approvers,omitempty"`
	BlockchainDetails *[]BlockchainTxDetails `json:"blockchainDetails,omitempty"`
	BlockchainRequest *BlockchainRequest `json:"blockchainRequest,omitempty"`
	Confirmations *string `json:"confirmations,omitempty"`
	DepositTransactionId *string `json:"depositTransactionId,omitempty"`
	Description *string `json:"description,omitempty"`
	ExternalBroadcast *bool `json:"externalBroadcast,omitempty"`
	FromAccountId *string `json:"fromAccountId,omitempty"`
	FromAddress *string `json:"fromAddress,omitempty"`
	Memo *string `json:"memo,omitempty"`
	SigningRequestId *string `json:"signingRequestId,omitempty"`
	ToAccountId *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	WithdrawTransactionId *string `json:"withdrawTransactionId,omitempty"`
}
