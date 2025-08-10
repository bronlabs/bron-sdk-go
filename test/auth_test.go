package test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mrand "math/rand"
	"strings"
	"testing"

	"github.com/bronlabs/bron-sdk-go/sdk/auth"
)

// genMockJwk creates a syntactically valid P-256 JWK for tests.
func genMockJwk() string {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	enc := base64.RawURLEncoding.EncodeToString
	x := enc(priv.X.Bytes())
	y := enc(priv.Y.Bytes())
	d := enc(priv.D.Bytes())
	kid := fmt.Sprintf("kid-%d", mrand.Int())
	return fmt.Sprintf(`{"kty":"EC","crv":"P-256","x":"%s","y":"%s","d":"%s","kid":"%s"}`,
		x, y, d, kid)
}

func TestParseJwk(t *testing.T) {
	jwk := genMockJwk()
	priv, kid, err := auth.ParseJwkEcPrivateKey(jwk)
	if err != nil {
		t.Fatal(err)
	}
	if priv == nil {
		t.Fatalf("nil private key")
	}
	if kid == "" {
		t.Fatalf("empty kid")
	}
}

func TestGenerateJwt(t *testing.T) {
	jwk := genMockJwk()
	_, kid, _ := auth.ParseJwkEcPrivateKey(jwk)
	jwt, err := auth.GenerateBronJwt(auth.BronJwtOptions{
		Method:     "GET",
		Path:       "/api/v1/workspaces",
		Kid:        kid,
		PrivateKey: jwk,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(strings.Split(jwt, ".")) != 3 {
		t.Fatalf("not a JWT")
	}
}
