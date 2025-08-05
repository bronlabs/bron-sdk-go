package types

type MessageForSigning struct {
	HashFunction *HashFunction `json:"hashFunction"`
	KeyType *KeyType `json:"keyType"`
	Message *string `json:"message"`
	SignatureScheme *SignatureScheme `json:"signatureScheme"`
	SignatureVariant *SignatureVariant `json:"signatureVariant"`
}
