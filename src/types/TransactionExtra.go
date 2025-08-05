package types

type TransactionExtra struct {
	FromAddress *string `json:"fromAddress"`
	ToAccountId *string `json:"toAccountId"`
	WithdrawTransactionId *string `json:"withdrawTransactionId"`
	BlockchainRequest *BlockchainRequest `json:"blockchainRequest"`
	DepositTransactionId *string `json:"depositTransactionId"`
	Description *string `json:"description"`
	Memo *string `json:"memo"`
	SigningRequestId *string `json:"signingRequestId"`
	ToAddress *string `json:"toAddress"`
	Approvers *TransactionApprovers `json:"approvers"`
	BlockchainDetails *[]BlockchainTxDetails `json:"blockchainDetails"`
	Confirmations *string `json:"confirmations"`
	ExternalBroadcast *bool `json:"externalBroadcast"`
	FromAccountId *string `json:"fromAccountId"`
}
