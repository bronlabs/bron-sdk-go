package types

type LimitRule struct {
	Approve *LimitRuleApprove `json:"approve"`
	SecurityDelay *LimitRuleSecurityDelay `json:"securityDelay"`
	SkipApproval *bool `json:"skipApproval"`
}
