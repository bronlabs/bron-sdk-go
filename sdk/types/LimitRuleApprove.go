package types

type LimitRuleApprove struct {
	AuthorisedApproversUserIds *[]string `json:"authorisedApproversUserIds,omitempty"`
	NumberOfApprovals string `json:"numberOfApprovals"`
}
