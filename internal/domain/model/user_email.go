package model

import (
	"regexp"
	"strings"
	"tanneryApp/internal/domain/errors"
)

// UserEmail is value object to store and validate user email
type UserEmail string

const MaxEmailLength = 100

// NewUserEmail create user email value object
func NewUserEmail(value string) (UserEmail, error) {
	value = strings.TrimSpace(value)

	if len(value) > MaxEmailLength {
		return "", errors.ErrEmailTooLong
	}

	if !validateEmailFormat(value) {
		return "", errors.ErrInvalidEmailFormat
	}

	return UserEmail(value), nil
}

// validateEmailFormat validate user email's format
func validateEmailFormat(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// String type cast from UserEmail to string
func (u UserEmail) String() string {
	return string(u)
}
