package domain

import "errors"

var ErrAdNotFound error = errors.New("ad has not been found")

//var ErrAdAlreadyExist error = errors.New("ad already exist")

type Ad struct {
	UserID      string  `json:"user_id"`
	AdID        string  `json:"ad_id"`
	Title       string  `json:"title"`
	Description string  `jason:"description"`
	Price       float32 `json:"price"`
	Picture     Picture `json:"picture"`
}
