package types

type TransactionLimit struct {
	LimitType TransactionLimitType `json:"limitType"`
	TransactionParams LimitTransactionParams `json:"transactionParams"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	LimitId string `json:"limitId"`
	Sources LimitSources `json:"sources"`
	Status TransactionLimitStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt"`
	UpdatedBy *string `json:"updatedBy"`
	WorkspaceId string `json:"workspaceId"`
	AppliesTo LimitAppliesTo `json:"appliesTo"`
	Destinations LimitDestinations `json:"destinations"`
	ExternalId string `json:"externalId"`
	LimitRule LimitRule `json:"limitRule"`
}
