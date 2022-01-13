package external

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrExternalGetAd = errors.New("Can't get ad information")

type Ad struct {
	AdID  string `json:"ad_id"`
	Title string `json:"title"`
	Owner string `json:"user_id"`
}

func GetAdInfos(adID string) <-chan string {
	owner := make(chan string)
	var url = "http://localhost:8080/ad/get/" + adID

	req, errReq := http.NewRequest("GET", url, nil)

	if errReq != nil {
		owner <- ""
		return owner
	}

	go func(req *http.Request) {
		defer close(owner)
		httpClient := &http.Client{}
		response, errCall := httpClient.Do(req)

		if errCall != nil {
			owner <- ""
		}

		defer response.Body.Close()

		if response.StatusCode != 200 {
			owner <- ""
		}

		ad := Ad{}

		json.NewDecoder(response.Body).Decode(&ad)

		owner <- ad.Owner
	}(req)

	return owner
}
