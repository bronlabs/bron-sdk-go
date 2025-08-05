package types

type MessagesForSigning struct {
	Messages *[]MessageForSigning `json:"messages"`
	PublicKey *string `json:"publicKey"`
	UseBackupPrimitive *bool `json:"useBackupPrimitive"`
}
