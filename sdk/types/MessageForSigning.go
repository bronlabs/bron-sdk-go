package types

type MessageForSigning struct {
	HashFunction *HashFunction `json:"hashFunction,omitempty"`
	HashPersonalization *string `json:"hashPersonalization,omitempty"`
	KeyType *KeyType `json:"keyType,omitempty"`
	Message *string `json:"message,omitempty"`
	SignatureScheme *SignatureScheme `json:"signatureScheme,omitempty"`
	SignatureVariant *SignatureVariant `json:"signatureVariant,omitempty"`
}
