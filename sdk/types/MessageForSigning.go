package types

type MessageForSigning struct {
	KeyType *KeyType `json:"keyType"`
	Message *string `json:"message"`
	SignatureScheme *SignatureScheme `json:"signatureScheme"`
	SignatureVariant *SignatureVariant `json:"signatureVariant"`
	HashFunction *HashFunction `json:"hashFunction"`
}
