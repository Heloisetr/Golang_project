package domain

import "errors"

var ErrAdNotFound error = errors.New("ad has not been found")
var ErrCreate error = errors.New("Can't create new Ad")
var ErrCantDelete error = errors.New("Can't delete Ad")
var ErrCantUpdate error = errors.New("Can't update Ad")
var ErrUnauthorized error = errors.New("Unauthorized")

//var ErrAdAlreadyExist error = errors.New("ad already exist")

type Ad struct {
	UserID      string  `json:"user_id"`
	AdID        string  `json:"ad_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Picture     Picture `json:"picture"`
}

type UpdateAd struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Picture     Picture `json:"picture,omitempty"`
}
