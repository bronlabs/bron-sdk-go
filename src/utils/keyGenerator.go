package utils

import (
	"fmt"
	"os"
)

func GenerateKeys() error {
	keyPair, err := GenerateBronKeyPair()
	if err != nil {
		return fmt.Errorf("failed to generate key pair: %w", err)
	}

	fmt.Println("Public JWK (send to Bron):")
	fmt.Println(keyPair.PublicJwk)
	fmt.Println()
	fmt.Println("Private JWK (keep safe):")
	fmt.Println(keyPair.PrivateJwk)

	return nil
}

func ValidateJwk(jwkString string) error {
	valid := ValidateBronJwk(jwkString)
	keyId := ExtractKeyId(jwkString)
	isPrivate := IsPrivateKey(jwkString)

	fmt.Printf("Valid ES256 JWK: %t\n", valid)
	fmt.Printf("Key ID: %s\n", keyId)
	fmt.Printf("Is private key: %t\n", isPrivate)

	return nil
}

func Main() {
	if len(os.Args) > 1 && os.Args[1] == "--validate" && len(os.Args) > 2 {
		if err := ValidateJwk(os.Args[2]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := GenerateKeys(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
} 