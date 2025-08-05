package utils

import (
	"crypto/sha256"
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
