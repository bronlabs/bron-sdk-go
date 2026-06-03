package types

type MessagesForSigning struct {
	Messages *[]MessageForSigning `json:"messages,omitempty"`
	PrimitivesVersion *string `json:"primitivesVersion,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
	UseBackupPrimitive *bool `json:"useBackupPrimitive,omitempty"`
}
