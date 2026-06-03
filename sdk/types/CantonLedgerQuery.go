package types

type CantonLedgerQuery struct {
	AccountID string `json:"accountId"`
	Params interface{} `json:"params"`
	SessionTopic string `json:"sessionTopic"`
}
