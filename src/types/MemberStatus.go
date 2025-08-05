package types

type MemberStatus string

const (
	MemberStatus_NEW MemberStatus = "new"
	MemberStatus_ACTIVE MemberStatus = "active"
	MemberStatus_REJECTED MemberStatus = "rejected"
	MemberStatus_DEACTIVATED MemberStatus = "deactivated"
)
