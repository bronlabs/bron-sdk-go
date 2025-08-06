package types

type CreateTransaction struct {
	AccountId string `json:"accountId"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	ExternalId string `json:"externalId"`
	Params interface{} `json:"params,omitempty"`
	TransactionType TransactionType `json:"transactionType"`
}
