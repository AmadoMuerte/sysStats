package validator

import (
	"errors"
	"regexp"
)

func isUpperCase(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func ValidatePassword(pass string) error {
	if len(pass) < 8 || len(pass) > 20 {
		return errors.New("invalid password: must be 8-20 characters long")
	}

	// Проверка, что первая буква заглавная
	if len(pass) == 0 || !isUpperCase(pass[0]) {
		return errors.New("invalid password: must start with an uppercase letter")
	}

	var hasLetter = regexp.MustCompile(`[a-zA-Z]`).MatchString
	var hasDigitOrSpecial = regexp.MustCompile(`[\d!@#$%^&*()\-_=+]`).MatchString

	if !hasLetter(pass) || !hasDigitOrSpecial(pass) {
		return errors.New("invalid password: must contain at least one letter and one number or special character")
	}

	return nil
}

func ValidateEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return emailRegex.MatchString(email)
}
