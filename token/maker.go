package token

import (
	"context"
	"crypto/rsa"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/pkg/e"
	"golang-server/pkg/logger"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type IJWTMaker interface {
	CreateToken(ctx context.Context, data interface{}) (string, error)
	VerifyToken(ctx context.Context, token string) (interface{}, error)
}

type jwtMaker struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWTMaker(ctx context.Context, cnf config.Token) (IJWTMaker, error) {
	privateKeyFile, err := os.ReadFile(cnf.PrivateKeyPath)
	if err != nil {
		logger.Error(ctx, err, "get private key error")
	}
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		logger.Error(ctx, err, "parse private key error")
	}

	publicKeyFile, err := os.ReadFile(cnf.PublicKeyPath)
	if err != nil {
		logger.Error(ctx, err, "get public key error")
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		logger.Error(ctx, err, "parse public key error")
	}
	return jwtMaker{
		privateKey: priKey,
		publicKey:  pubKey,
	}, nil
}

// CreateToken implements IJWTMaker.
func (j jwtMaker) CreateToken(ctx context.Context, payload interface{}) (string, error) {
	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload.(jwt.Claims))

	// Sign the token with the RSA private key
	signedToken, err := token.SignedString(j.privateKey)
	if err != nil {
		logger.Error(ctx, err, "Failed to sign JWT token")
		return "", err
	}
	return signedToken, nil
}

// VerifyToken implements IJWTMaker.
func (j jwtMaker) VerifyToken(ctx context.Context, tokenString string) (interface{}, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &dto.UserPayload{}, func(token *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	})

	if err != nil {
		logger.Error(ctx, err, "Failed to parse JWT token:")
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(*dto.UserPayload); ok && token.Valid {
		return claims, nil
	} else {
		logger.Info(ctx, "Token is invalid")
		return nil, e.ErrUnauthorized
	}
}
