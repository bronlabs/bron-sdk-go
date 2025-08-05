package types

type TransactionLimit struct {
	CreatedAt string `json:"createdAt"`
	LimitRule LimitRule `json:"limitRule"`
	LimitType TransactionLimitType `json:"limitType"`
	Status TransactionLimitStatus `json:"status"`
	TransactionParams LimitTransactionParams `json:"transactionParams"`
	UpdatedAt *string `json:"updatedAt"`
	AppliesTo LimitAppliesTo `json:"appliesTo"`
	CreatedBy *string `json:"createdBy"`
	Destinations LimitDestinations `json:"destinations"`
	ExternalId string `json:"externalId"`
	LimitId string `json:"limitId"`
	Sources LimitSources `json:"sources"`
	UpdatedBy *string `json:"updatedBy"`
	WorkspaceId string `json:"workspaceId"`
}
