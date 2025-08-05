package types

type CreateTransaction struct {
	AccountId string `json:"accountId"`
	ExpiresAt *string `json:"expiresAt"`
	ExternalId string `json:"externalId"`
	Params interface{} `json:"params"`
	TransactionType TransactionType `json:"transactionType"`
}
