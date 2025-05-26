package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key string = os.Getenv("JWT_KEY")

func GenerateJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		Issuer:    "rest-crud-go",
		Subject:   userId,
	})

	signedString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", fmt.Errorf("error generating a json web token")
	}

	return signedString, nil
}

func ParseJWT(token string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(key), nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing JWT: %w", err)
	}

	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)

	if !ok {
		return "", fmt.Errorf("error parsing JWT claims: %w", err)
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return "", fmt.Errorf("token expired")
	}

	return claims.Subject, nil
}
