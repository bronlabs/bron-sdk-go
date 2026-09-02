package types

type MessagesForSigning struct {
	Messages *[]MessageForSigning `json:"messages,omitempty"`
	ParallelSigning *bool `json:"parallelSigning,omitempty"`
	PrimitivesVersion *string `json:"primitivesVersion,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
	UseBackupPrimitive *bool `json:"useBackupPrimitive,omitempty"`
}
