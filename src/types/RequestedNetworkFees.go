package types

type RequestedNetworkFees struct {
	MaxFeePerGas *string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *string `json:"maxPriorityFeePerGas"`
	FeePerByte *string `json:"feePerByte"`
	GasLimit *string `json:"gasLimit"`
	GasPriceGwei *string `json:"gasPriceGwei"`
}
