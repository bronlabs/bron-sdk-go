package types

import "encoding/json"

func NewWithdrawalTx(accountID, externalID string, params WithdrawalParams) CreateTransaction {
	return CreateTransaction{
		AccountID:       accountID,
		ExternalID:      externalID,
		TransactionType: TransactionType_WITHDRAWAL,
		Params:          params,
	}
}

func NewCustomTx(accountID, externalID string, txType TransactionType, raw json.RawMessage) CreateTransaction {
	return CreateTransaction{
		AccountID:       accountID,
		ExternalID:      externalID,
		TransactionType: txType,
		Params:          raw,
	}
}
