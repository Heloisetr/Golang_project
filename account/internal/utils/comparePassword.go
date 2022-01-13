package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))

	if err != nil {
		return false
	}

	return true
}
