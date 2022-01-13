package domain

import "errors"

var ErrAdNotFound error = errors.New("ad has not been found")

type Ad struct {
	UserID string `json:"user_id"`
	AdID   string `json:"ad_id"`
}
