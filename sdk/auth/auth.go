package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type BronJwtOptions struct {
	Method     string
	Path       string
	Body       string
	Kid        string
	PrivateKey string
	Iat        *int64
}

func GenerateBronJwt(options BronJwtOptions) (string, error) {
	var iat int64
	if options.Iat != nil {
		iat = *options.Iat
	} else {
		iat = time.Now().Unix()
	}
	messageString := fmt.Sprintf("%d%s%s", iat, options.Method, options.Path+options.Body)

	hash := sha256.Sum256([]byte(messageString))
	hashHex := fmt.Sprintf("%x", hash)

	// Parse JWK to get private key
	privateKey, _, err := ParseJwkEcPrivateKey(options.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse JWK: %w", err)
	}

	// Create JWT claims
	claims := jwt.MapClaims{
		"iat":     iat,
		"message": hashHex,
	}

	// Create JWT header
	header := map[string]interface{}{
		"alg": "ES256",
		"kid": options.Kid,
	}

	// Sign JWT
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header = header

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return tokenString, nil
}

func ParseJwkEcPrivateKey(jwkString string) (privateKey *ecdsa.PrivateKey, kid string, err error) {
	var jwk map[string]interface{}
	if err := json.Unmarshal([]byte(jwkString), &jwk); err != nil {
		return nil, "", fmt.Errorf("failed to parse JWK: %w", err)
	}

	// Validate JWK format
	if jwk["kty"] != "EC" || jwk["crv"] != "P-256" {
		return nil, "", fmt.Errorf("unsupported JWK format")
	}

	// Extract components
	xStr, _ := jwk["x"].(string)
	yStr, _ := jwk["y"].(string)
	dStr, _ := jwk["d"].(string)
	kid, _ = jwk["kid"].(string)

	// Decode base64url components
	xBytes, err := base64.RawURLEncoding.DecodeString(xStr)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode x: %w", err)
	}
	yBytes, err := base64.RawURLEncoding.DecodeString(yStr)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode y: %w", err)
	}
	dBytes, err := base64.RawURLEncoding.DecodeString(dStr)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode d: %w", err)
	}

	// Create ECDSA private key
	x := new(big.Int).SetBytes(xBytes)
	y := new(big.Int).SetBytes(yBytes)
	d := new(big.Int).SetBytes(dBytes)

	privateKey = &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     x,
			Y:     y,
		},
		D: d,
	}

	return privateKey, kid, nil
}
