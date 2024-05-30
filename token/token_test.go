package token

import (
	"fmt"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

type Payload struct {
	UserID string `json:"user_id"`
}

func (uc Payload) Valid() error {
	return nil
}

func CreateToken(payload interface{}) string {
	privateKeyFile, err := os.ReadFile("../cert/private.key")
	if err != nil {
		fmt.Println("errrr", err)
	}
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		fmt.Println("err", err)
	}
	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload.(jwt.Claims))

	// Sign the token with the RSA private key
	signedToken, err := token.SignedString(priKey)
	if err != nil {
		fmt.Println("Failed to sign JWT token:", err)
		return ""
	}
	return signedToken
}

func VerifyToken(tokenString string) interface{} {
	publicKeyFile, err := os.ReadFile("../cert/public.key")
	if err != nil {
		fmt.Println("get public key error")
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		fmt.Println("parse public key error")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})

	if err != nil {
		fmt.Println("Failed to parse JWT token:", err)
		return nil
	}

	// Validate the token
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		fmt.Println("Token is valid")
		fmt.Println("User ID:", claims.UserID)
		return claims
	} else {
		fmt.Println("Token is invalid")
		return nil
	}
}

func TestToken(t *testing.T) {
	payload := Payload{
		UserID: "abc",
	}
	signedToken := CreateToken(payload)
	fmt.Println("Signed JWT token:", signedToken)

	VerifyToken(signedToken)

	panic("done")
}
