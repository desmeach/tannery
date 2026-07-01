package model

import (
	"regexp"
	"strings"
	errs "tanneryApp/internal/domain/errors"
)

// UserPhone is value object to store and validate user phone
type UserPhone string

// NewUserPhone create user phone value object
func NewUserPhone(value string) (UserPhone, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return "", errs.ErrPhoneEmpty
	}

	if !strings.HasPrefix(value, "+") {
		return "", errs.ErrInvalidPhoneFormat
	}

	if !isNumeric(value[1:]) {
		return "", errs.ErrPhoneContainsLetters
	}

	if err := validatePhoneFormat(value); err != nil {
		return "", err
	}

	return UserPhone(value), nil
}

// validatePhoneFormat validate phone's format
func validatePhoneFormat(phone string) error {
	regex := regexp.MustCompile(`^\+\d{10,15}$`)
	if !regex.MatchString(phone) {
		return errs.ErrInvalidPhoneFormat
	}

	if len(phone) < 11 || len(phone) > 13 {
		return errs.ErrInvalidPhoneLength
	}

	return nil
}

// isNumeric check phone contains only numbers
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// String type cast from UserPhone to string
func (u UserPhone) String() string {
	return string(u)
}
