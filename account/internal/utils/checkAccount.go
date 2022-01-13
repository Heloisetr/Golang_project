package utils

import "account/domain"

func CheckAccount(account domain.Account) bool {
	if account.Email == "" || account.Login == "" || account.Password == "" {
		return false
	} else {
		return true
	}
}
