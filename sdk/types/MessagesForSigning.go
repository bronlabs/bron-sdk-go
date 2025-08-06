package types

type MessagesForSigning struct {
	Messages *[]MessageForSigning `json:"messages,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
	UseBackupPrimitive *bool `json:"useBackupPrimitive,omitempty"`
}
