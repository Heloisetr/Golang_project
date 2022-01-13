package domain

import "errors"

var ErrEmailNotFound = errors.New("email not found")
var ErrWrongPassword = errors.New("wrong password")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	AccountID string `json:"account_id"`
}
