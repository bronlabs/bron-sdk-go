package types

type SigningRequestStatus string

const (
	SigningRequestStatus_NEW SigningRequestStatus = "new"
	SigningRequestStatus_SIGNED SigningRequestStatus = "signed"
	SigningRequestStatus_BROADCASTED SigningRequestStatus = "broadcasted"
	SigningRequestStatus_UNDER_RBF SigningRequestStatus = "under-rbf"
	SigningRequestStatus_COMPLETED SigningRequestStatus = "completed"
	SigningRequestStatus_MANUAL_RESOLVING SigningRequestStatus = "manual-resolving"
	SigningRequestStatus_CANCELED SigningRequestStatus = "canceled"
	SigningRequestStatus_ERROR_ON_BROADCAST SigningRequestStatus = "error-on-broadcast"
	SigningRequestStatus_FAILED_ON_CHAIN SigningRequestStatus = "failed-on-chain"
	SigningRequestStatus_MARKED_AS_ERROR SigningRequestStatus = "marked-as-error"
)
