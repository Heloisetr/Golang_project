package domain

import "errors"

var ErrUserNotFound = errors.New("user has not been found")
var ErrUserAlreadyExist = errors.New("user already exist")

type User struct {
	UserID   string  `json:"user_id"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Address  Address `json:"address"`
}
