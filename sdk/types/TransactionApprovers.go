package types

type TransactionApprovers struct {
	SkipApproval *bool `json:"skipApproval"`
	ApprovedBy *[]string `json:"approvedBy"`
	AvailableApprovers *[]string `json:"availableApprovers"`
	LimitId *string `json:"limitId"`
	Number *string `json:"number"`
	SecurityDelayDuration *string `json:"securityDelayDuration"`
	SecurityDelayExpiresAt *string `json:"securityDelayExpiresAt"`
}
