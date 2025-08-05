package types

type RequestedNetworkFees struct {
	FeePerByte *string `json:"feePerByte"`
	GasLimit *string `json:"gasLimit"`
	GasPriceGwei *string `json:"gasPriceGwei"`
	MaxFeePerGas *string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *string `json:"maxPriorityFeePerGas"`
}
