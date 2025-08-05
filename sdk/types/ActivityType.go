package types

type ActivityType string

const (
	ActivityType_LOGIN ActivityType = "login"
	ActivityType_NEW_DEVICE_LOGIN ActivityType = "new-device-login"
	ActivityType_NEW_ADDRESS_BOOK_RECORD ActivityType = "new-address-book-record"
	ActivityType_UPDATE_ADDRESS_BOOK_RECORD ActivityType = "update-address-book-record"
	ActivityType_DELETE_ADDRESS_BOOK_RECORD ActivityType = "delete-address-book-record"
	ActivityType_WORKSPACE_CREATION ActivityType = "workspace-creation"
	ActivityType_WORKSPACE_SETTINGS_APPROVAL_UPDATE ActivityType = "workspace-settings-approval-update"
	ActivityType_WORKSPACE_NAME_CHANGE ActivityType = "workspace-name-change"
	ActivityType_WORKSPACE_TAG_CHANGE ActivityType = "workspace-tag-change"
	ActivityType_ONLY_ADDRESS_BOOK_WITHDRAWALS_SETTINGS_UPDATE ActivityType = "only-address-book-withdrawals-settings-update"
	ActivityType_TRANSACTION_APPROVAL_SETTINGS_UPDATE ActivityType = "transaction-approval-settings-update"
	ActivityType_WORKSPACE_MEMBERS_APPROVAL_UPDATE ActivityType = "workspace-members-approval-update"
	ActivityType_ADDRESS_BOOK_SETTINGS_APPROVAL_UPDATE ActivityType = "address-book-settings-approval-update"
	ActivityType_TRANSACTION_LIMITS_APPROVAL_SETTINGS_UPDATE ActivityType = "transaction-limits-approval-settings-update"
	ActivityType_MEMBER_CREATION ActivityType = "member-creation"
	ActivityType_SERVICE_MEMBER_CREATION ActivityType = "service-member-creation"
	ActivityType_TRANSACTION_LIMIT_CREATION ActivityType = "transaction-limit-creation"
	ActivityType_TRANSACTION_LIMIT_UPDATE ActivityType = "transaction-limit-update"
	ActivityType_ACCOUNT_CREATION ActivityType = "account-creation"
	ActivityType_PASSKEY_CREATION ActivityType = "passkey-creation"
	ActivityType_PASSKEY_DELETION ActivityType = "passkey-deletion"
	ActivityType_TYPE_2FA_CHANGE ActivityType = "2fa-change"
	ActivityType_ACCOUNT_PINS_RESET ActivityType = "account-pins-reset"
	ActivityType_SHARD_ACCESS_REQUEST_CREATED ActivityType = "shard-access-request-created"
	ActivityType_SHARD_ACCESS_REQUEST_FINISHED ActivityType = "shard-access-request-finished"
	ActivityType_TRANSACTION_COMPLETED ActivityType = "transaction-completed"
)
