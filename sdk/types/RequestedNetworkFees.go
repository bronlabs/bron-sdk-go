package types

type RequestedNetworkFees struct {
	FeePerByte *string `json:"feePerByte,omitempty"`
	GasLimit *string `json:"gasLimit,omitempty"`
	GasPriceGwei *string `json:"gasPriceGwei,omitempty"`
	MaxFeePerGas *string `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *string `json:"maxPriorityFeePerGas,omitempty"`
}
