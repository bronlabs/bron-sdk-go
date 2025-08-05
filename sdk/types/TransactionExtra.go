package types

type TransactionExtra struct {
	Approvers *TransactionApprovers `json:"approvers"`
	BlockchainDetails *[]BlockchainTxDetails `json:"blockchainDetails"`
	BlockchainRequest *BlockchainRequest `json:"blockchainRequest"`
	Confirmations *string `json:"confirmations"`
	DepositTransactionId *string `json:"depositTransactionId"`
	Description *string `json:"description"`
	ExternalBroadcast *bool `json:"externalBroadcast"`
	FromAccountId *string `json:"fromAccountId"`
	FromAddress *string `json:"fromAddress"`
	Memo *string `json:"memo"`
	SigningRequestId *string `json:"signingRequestId"`
	ToAccountId *string `json:"toAccountId"`
	ToAddress *string `json:"toAddress"`
	WithdrawTransactionId *string `json:"withdrawTransactionId"`
}
