package auth

import (
	"regexp"

	"github.com/nyaruka/phonenumbers"
)

func isPhoneValid(phone string) bool {
	if len(phone) != 10 {
		return false
	}
	_, err := phonenumbers.Parse(phone, "RU")
	return err == nil
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
