package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
)

func main() {
	keyPair, err := GenerateKeyPair()
	if err != nil {
		panic(fmt.Errorf("failed to generate key pair: %w", err))
	}

	fmt.Print("\n-------------------------------------\n")

	fmt.Print("\n✅ Public JWK (send to Bron):\n\n")
	fmt.Println(keyPair.PublicJwk)

	fmt.Print("\n-------------------------------------\n")

	fmt.Print("\n🔒 Private JWK (keep safe):\n\n")
	fmt.Println(keyPair.PrivateJwk)

	fmt.Print("\n-------------------------------------\n\n")
}

type JWK struct {
	Kty string `json:"kty"`
	Crv string `json:"crv"`
	X   string `json:"x"`
	Y   string `json:"y"`
	D   string `json:"d,omitempty"`
	Kid string `json:"kid"`
}

type GeneratedKeyPair struct {
	PublicJwk  string `json:"publicJwk"`
	PrivateJwk string `json:"privateJwk"`
	Kid        string `json:"kid"`
}

func GenerateKeyPair() (*GeneratedKeyPair, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), crand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %w", err)
	}

	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	kid := make([]byte, 24)
	for i := range kid {
		kid[i] = charset[rand.Intn(len(charset))]
	}

	publicJwk := JWK{
		Kty: "EC",
		Crv: "P-256",
		X:   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.X.Bytes()),
		Y:   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.Y.Bytes()),
		Kid: string(kid),
	}

	privateJwk := JWK{
		Kty: "EC",
		Crv: "P-256",
		X:   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.X.Bytes()),
		Y:   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.Y.Bytes()),
		D:   base64.RawURLEncoding.EncodeToString(privateKey.D.Bytes()),
		Kid: string(kid),
	}

	publicJwkBytes, err := json.MarshalIndent(publicJwk, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public JWK: %w", err)
	}

	privateJwkBytes, err := json.Marshal(privateJwk)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private JWK: %w", err)
	}

	return &GeneratedKeyPair{
		PublicJwk:  string(publicJwkBytes),
		PrivateJwk: string(privateJwkBytes),
		Kid:        string(kid),
	}, nil
}
