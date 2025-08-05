package types

type TransactionExtra struct {
	Confirmations *string `json:"confirmations"`
	DepositTransactionId *string `json:"depositTransactionId"`
	Description *string `json:"description"`
	FromAccountId *string `json:"fromAccountId"`
	FromAddress *string `json:"fromAddress"`
	Memo *string `json:"memo"`
	Approvers *TransactionApprovers `json:"approvers"`
	ExternalBroadcast *bool `json:"externalBroadcast"`
	SigningRequestId *string `json:"signingRequestId"`
	ToAccountId *string `json:"toAccountId"`
	ToAddress *string `json:"toAddress"`
	WithdrawTransactionId *string `json:"withdrawTransactionId"`
	BlockchainDetails *[]BlockchainTxDetails `json:"blockchainDetails"`
	BlockchainRequest *BlockchainRequest `json:"blockchainRequest"`
}
