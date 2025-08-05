package types

type CreateTransaction struct {
	TransactionType TransactionType `json:"transactionType"`
	AccountId string `json:"accountId"`
	ExpiresAt *string `json:"expiresAt"`
	ExternalId string `json:"externalId"`
	Params interface{} `json:"params"`
}
