package utils

import (
	"api/configs"
	"errors"
	"strings"
)

// CheckRegNumber Check RegistrationNumber
func CheckRegNumber(RegistrationNumber string) (err error) {
	if len(RegistrationNumber) > 8 || len(RegistrationNumber) < 5 {
		err = errors.New(configs.ErrorRegistrationNumber)
	}
	return
}

func FixRegNum(registrationNumber string) (s string) {
	if len(registrationNumber) > 0 {
		s = strings.ReplaceAll(registrationNumber, " ", "")
		s = strings.ToUpper(s)
	}

	return s
}
