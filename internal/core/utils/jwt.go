package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID    string `json:"user_id"`
	CSRFToken string `json:"csrf_token"`
	jwt.RegisteredClaims
}

func getJWTKey() []byte {
	key := os.Getenv("JWT_KEY")
	if key == "" {
		panic("JWT_KEY environment variable is required")
	}
	return []byte(key)
}

func GenerateCSRFToken() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func GenerateJWT(userId string) (string, string, error) {
	csrfToken, err := GenerateCSRFToken()
	if err != nil {
		return "", "", fmt.Errorf("error generating CSRF token: %w", err)
	}

	claims := CustomClaims{
		UserID:    userId,
		CSRFToken: csrfToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "rest-crud-go",
			Subject:   userId,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(getJWTKey()))
	if err != nil {
		return "", "", fmt.Errorf("error generating JWT: %w", err)
	}

	return signedString, csrfToken, nil
}

func ParseJWT(token string) (*CustomClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(getJWTKey()), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing JWT: %w", err)
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := parsedToken.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("error parsing JWT claims: %w", err)
	}

	return claims, nil
}
