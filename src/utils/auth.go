package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type BronJwtOptions struct {
	Method     string
	Path       string
	Body       string
	Kid        string
	PrivateKey string
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

func GenerateBronJwt(options BronJwtOptions) (string, error) {
	iat := time.Now().Unix()
	messageString := fmt.Sprintf("%d%s%s%s", iat, strings.ToUpper(options.Method), options.Path, options.Body)

	hash := sha256.Sum256([]byte(messageString))
	hashHex := fmt.Sprintf("%x", hash)

	claims := jwt.MapClaims{
		"iat":     iat,
		"message": hashHex,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = options.Kid

	privateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(options.PrivateKey))
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func ParseJwkEcPrivateKey(jwkString string) (string, string, error) {
	// Parse JWK using the jwk library (equivalent to jwk-to-pem in TypeScript)
	key, err := jwk.ParseKey([]byte(jwkString))
	if err != nil {
		return "", "", fmt.Errorf("failed to parse JWK: %w", err)
	}

	// Check if it's an EC key
	if key.KeyType() != "EC" {
		return "", "", fmt.Errorf("key is not an EC key")
	}

	// Get the key ID
	kid := key.KeyID()
	if kid == "" {
		return "", "", fmt.Errorf("key ID not found")
	}

	// Convert to PEM format
	pemBytes, err := jwk.Pem(key)
	if err != nil {
		return "", "", fmt.Errorf("failed to convert to PEM: %w", err)
	}

	return string(pemBytes), kid, nil
}

func GenerateBronKeyPair() (*GeneratedKeyPair, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %w", err)
	}

	kid := generateKid()

	publicJwk := JWK{
		Kty: "EC",
		Crv: "P-256",
		X:   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.X.Bytes()),
		Y:   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.Y.Bytes()),
		Kid: kid,
	}

	privateJwk := JWK{
		Kty: "EC",
		Crv: "P-256",
		X:   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.X.Bytes()),
		Y:   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.Y.Bytes()),
		D:   base64.RawURLEncoding.EncodeToString(privateKey.D.Bytes()),
		Kid: kid,
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
		Kid:        kid,
	}, nil
}

func ValidateBronJwk(jwkString string) bool {
	var jwk JWK
	if err := json.Unmarshal([]byte(jwkString), &jwk); err != nil {
		return false
	}

	return jwk.Kty == "EC" &&
		jwk.Crv == "P-256" &&
		jwk.X != "" &&
		jwk.Y != "" &&
		jwk.Kid != ""
}

func ExtractKeyId(jwkString string) string {
	var jwk JWK
	if err := json.Unmarshal([]byte(jwkString), &jwk); err != nil {
		return ""
	}
	return jwk.Kid
}

func IsPrivateKey(jwkString string) bool {
	var jwk JWK
	if err := json.Unmarshal([]byte(jwkString), &jwk); err != nil {
		return false
	}
	return jwk.D != ""
}

func generateKid() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return base64.RawURLEncoding.EncodeToString(bytes)
}
