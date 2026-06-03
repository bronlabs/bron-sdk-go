package types

type AccountStatus string

const (
	AccountStatus_ACTIVE AccountStatus = "active"
	AccountStatus_ARCHIVED AccountStatus = "archived"
	AccountStatus_SHARD_GENERATING AccountStatus = "shard-generating"
	AccountStatus_ERROR_ON_GENERATING AccountStatus = "error-on-generating"
)
