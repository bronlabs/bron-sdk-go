package types

type AddressStatus string

const (
	AddressStatus_NEW AddressStatus = "new"
	AddressStatus_PENDING AddressStatus = "pending"
	AddressStatus_ADDRESS_ACTIVATION_REQUIRED AddressStatus = "address-activation-required"
	AddressStatus_ADDRESS_CREATION_REQUIRED AddressStatus = "address-creation-required"
	AddressStatus_APPROVAL_PENDING AddressStatus = "approval-pending"
	AddressStatus_ENABLED AddressStatus = "enabled"
	AddressStatus_DISABLED AddressStatus = "disabled"
	AddressStatus_ERROR AddressStatus = "error"
	AddressStatus_ACCOUNT_ARCHIVED AddressStatus = "account-archived"
)
