package types

type Transaction struct {
	TransactionId string `json:"transactionId"`
	WorkspaceId string `json:"workspaceId"`
	Status TransactionStatus `json:"status"`
	CreatedAt string `json:"createdAt"`
	ExpiresAt *string `json:"expiresAt"`
	Extra *TransactionExtra `json:"extra"`
	TerminatedAt *string `json:"terminatedAt"`
	AccountId string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	Embedded *TransactionEmbedded `json:"embedded"`
	Params interface{} `json:"params"`
	TransactionType TransactionType `json:"transactionType"`
	UpdatedAt *string `json:"updatedAt"`
	CreatedBy *string `json:"createdBy"`
	ExternalId string `json:"externalId"`
}
