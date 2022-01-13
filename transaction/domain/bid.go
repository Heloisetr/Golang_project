package domain

import "errors"

var ErrBidNotFound error = errors.New("Bid has not been found")
var ErrToken = errors.New("Error while parsing token")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrCantDelete = errors.New("Can't delete Bid")
var ErrCantUpdate = errors.New("Can't update Bid")

type Bid struct {
	UserID   string  `json:"user_id"`
	Owner    string  `json:"owner"`
	BidID    string  `json:"bid_id"`
	AdId     string  `json:"ad_id"`
	BidPrice float32 `json:"bid_price"`
	Message  string  `json:"message"`
	Status   string  `json:"status"`
}
