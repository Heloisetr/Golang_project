package domain

import "errors"

var ErrAccountNotFound = errors.New("account not found")
var ErrEmailAlreadyUsed = errors.New("email or login already used")
var ErrAccessUnauthorized = errors.New("access unauthorized")

type Account struct {
	AccountID string  `json:"account_id"`
	Email     string  `json:"email"`
	Login     string  `json:"login"`
	Password  string  `json:"-"`
	Balance   float32 `json:"balance,omitempty"`
}

type UpdateAccount struct {
	Email string `json:"email,omitempty"`
	Login string `json:"login,omitempty"`
}
