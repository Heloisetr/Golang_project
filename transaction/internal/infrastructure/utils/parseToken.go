package utils

import (
	"os"
	"transaction/domain"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string) (string, error) {
	claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTKEY")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", jwt.ErrSignatureInvalid
		}
		return "", domain.ErrToken
	}
	if !tkn.Valid {
		return "", domain.ErrToken
	}

	return claims["account_id"].(string), nil
}
