package types

type HashFunction string

const (
	HashFunction_NONE HashFunction = "none"
	HashFunction_SHA256D HashFunction = "sha256d"
	HashFunction_KECCAK256 HashFunction = "keccak256"
	HashFunction_BLAKE2B256 HashFunction = "blake2b256"
	HashFunction_SHA256 HashFunction = "sha256"
	HashFunction_SHA512 HashFunction = "sha512"
	HashFunction_SHA512_HALF HashFunction = "sha512_half"
	HashFunction_SHA512_256 HashFunction = "sha512_256"
	HashFunction_POSEIDON HashFunction = "poseidon"
)
