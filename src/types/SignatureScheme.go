package types

type SignatureScheme string

const (
	SignatureScheme_ECDSA SignatureScheme = "ecdsa"
	SignatureScheme_EDDSA SignatureScheme = "eddsa"
	SignatureScheme_BLS SignatureScheme = "bls"
	SignatureScheme_SCHNORR SignatureScheme = "schnorr"
	SignatureScheme_RSA_PSS SignatureScheme = "rsa-pss"
)
