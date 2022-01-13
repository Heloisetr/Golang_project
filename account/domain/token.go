package domain

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var ErrTokenCreation = errors.New("Can't create token")

type Claims struct {
	AccountID string `json:"account_id"`
	jwt.StandardClaims
}
