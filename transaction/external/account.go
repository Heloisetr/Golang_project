package external

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrExternalGetAccount = errors.New("Can't get user account")
var ErrExternalNotEnoughBalance = errors.New("Not enough balance on account")

type Account struct {
	UserID  string  `json:"accountId"`
	Balance float32 `json:"balance"`
}

func GetAccountCompareBalance(token string, accountId string, price float32) error {
	var url = "http://localhost:8081/account/" + accountId

	var bearer = "Bearer " + token

	req, errReq := http.NewRequest("GET", url, nil)

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
		return ErrExternalGetAccount
	}

	account := Account{}

	json.NewDecoder(response.Body).Decode(&account)

	if account.Balance < price {
		return ErrExternalNotEnoughBalance
	} else {
		return nil
	}
}
