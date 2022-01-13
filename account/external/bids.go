package external

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrExternalGetBids = errors.New("Can't get user bids")
var ErrExternalDeleteBids = errors.New("Can't delete user bid")

type Bid struct {
	UserID   string  `json:"user_id"`
	BidID    string  `json:"bid_id"`
	AdId     string  `json:"ad_id"`
	BidPrice float32 `json:"bid_price"`
	Message  string  `json:"message"`
	Status   string  `json:"status"`
}

type Bids []Bid

func GetAllUserBids(token string) ([]Bid, error) {
	var url = "http://localhost:8082/bids"

	var bearer = "Bearer " + token

	req, errReq := http.NewRequest("GET", url, nil)

	if errReq != nil {
		return nil, errReq
	}

	req.Header.Add("Authorization", bearer)

	httpClient := &http.Client{}
	response, errCall := httpClient.Do(req)

	if errCall != nil {
		return nil, errCall
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, ErrExternalGetBids
	}

	bids := Bids{}

	json.NewDecoder(response.Body).Decode(&bids)

	return bids, nil
}

func DeleteUserBid(token string, bidId string) error {
	var url = "http://localhost:8082/bid/delete/" + bidId

	var bearer = "Bearer " + token

	req, errReq := http.NewRequest("DELETE", url, nil)

	if errReq != nil {
		return errReq
	}

	req.Header.Add("Authorization", bearer)

	httpClient := &http.Client{}
	response, errCall := httpClient.Do(req)

	if errCall != nil {
		return errCall
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return ErrExternalDeleteBids
	}

	return nil
}
