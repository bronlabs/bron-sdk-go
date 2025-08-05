package types

type Transaction struct {
	AccountId string `json:"accountId"`
	CreatedBy *string `json:"createdBy"`
	TransactionId string `json:"transactionId"`
	CreatedAt string `json:"createdAt"`
	Embedded *TransactionEmbedded `json:"embedded"`
	ExternalId string `json:"externalId"`
	Extra *TransactionExtra `json:"extra"`
	WorkspaceId string `json:"workspaceId"`
	ExpiresAt *string `json:"expiresAt"`
	Status TransactionStatus `json:"status"`
	TerminatedAt *string `json:"terminatedAt"`
	TransactionType TransactionType `json:"transactionType"`
	UpdatedAt *string `json:"updatedAt"`
	AccountType AccountType `json:"accountType"`
	Params interface{} `json:"params"`
}
