package types

type TransactionApprovers struct {
	SecurityDelayDuration *string `json:"securityDelayDuration"`
	SecurityDelayExpiresAt *string `json:"securityDelayExpiresAt"`
	SkipApproval *bool `json:"skipApproval"`
	ApprovedBy *[]string `json:"approvedBy"`
	AvailableApprovers *[]string `json:"availableApprovers"`
	LimitId *string `json:"limitId"`
	Number *string `json:"number"`
}
