package types

type MessageForSigning struct {
	SignatureVariant *SignatureVariant `json:"signatureVariant"`
	HashFunction *HashFunction `json:"hashFunction"`
	KeyType *KeyType `json:"keyType"`
	Message *string `json:"message"`
	SignatureScheme *SignatureScheme `json:"signatureScheme"`
}
