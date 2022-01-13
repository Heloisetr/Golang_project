package external

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrExternalGetAds = errors.New("Can't get user ads")
var ErrExternalDeleteAd = errors.New("Can't delete user ad")

type Ad struct {
	UserID      string `json:"user_id"`
	AdID        string `json:"ad_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type Ads []Ad

func GetAllUserAds(token string, id string) ([]Ad, error) {
	var url = "http://localhost:8080/ad/get_all/" + id

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
		return nil, ErrExternalGetAds
	}

	ads := Ads{}

	json.NewDecoder(response.Body).Decode(&ads)

	return ads, nil
}

func DeleteUserAd(token string, adId string) error {
	var url = "http://localhost:8080/ad/delete/" + adId

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
		return ErrExternalDeleteAd
	}

	return nil
}
