package types

type LimitRuleApprove struct {
	AuthorisedApproversUserIDs *[]string `json:"authorisedApproversUserIds,omitempty"`
	NumberOfApprovals string `json:"numberOfApprovals"`
}
