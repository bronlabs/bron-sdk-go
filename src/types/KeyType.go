package types

type KeyType string

const (
	KeyType_SECP256K1 KeyType = "secp256k1"
	KeyType_EDWARDS25519 KeyType = "edwards25519"
	KeyType_BLS12381G1 KeyType = "BLS12381G1"
	KeyType_PALLAS KeyType = "pallas"
	KeyType_RSA4096 KeyType = "RSA4096"
)
