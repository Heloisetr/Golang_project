package utils

import (
	"account/domain"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(account_id string) (string, error) {
	expiration := time.Now().Add(15 * time.Minute)

	claims := &domain.Claims{
		AccountID: account_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))

	if err != nil {
		return "", domain.ErrTokenCreation
	}

	return tokenString, nil
}

func ParseToken(token string) (string, error) {
	claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTKEY")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", jwt.ErrSignatureInvalid
		}
		return "", domain.ErrTokenCreation
	}
	if !tkn.Valid {
		return "", domain.ErrTokenCreation
	}

	return claims["account_id"].(string), nil
}
