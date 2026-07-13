package user

import (
	"regexp"
	"strings"
	"tanneryApp/internal/domain/errors"
)

// Email is value object to store and validate profile email
type Email string

const MaxEmailLength = 100

// NewEmail create profile email value object
func NewEmail(value string) (*Email, error) {
	value = strings.TrimSpace(value)

	if len(value) > MaxEmailLength {
		return nil, errors.ErrEmailTooLong
	}

	if !validateEmailFormat(value) {
		return nil, errors.ErrInvalidEmailFormat
	}

	email := Email(value)
	return &email, nil
}

// validateEmailFormat validate profile email's format
func validateEmailFormat(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// String type cast from Email to string
func (ue Email) String() string {
	return string(ue)
}

// Mask return masked Email as string
func (ue Email) Mask() string {
	email := string(ue)

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	name := parts[0]
	domain := parts[1]

	maskedName := string(name[0]) + "***"
	if len(name) > 1 {
		maskedName = string(name[0]) + "***" + string(name[len(name)-1])
	}

	return maskedName + "@" + domain
}
