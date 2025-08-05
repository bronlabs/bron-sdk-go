package types

type RecordStatus string

const (
	RecordStatus_NEW RecordStatus = "new"
	RecordStatus_ACTIVE RecordStatus = "active"
	RecordStatus_REJECTED RecordStatus = "rejected"
	RecordStatus_DELETED RecordStatus = "deleted"
)
