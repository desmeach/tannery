package user

import (
	"regexp"
	"strings"
	errs "tanneryApp/internal/domain/errors"
)

// Phone is value object to store and validate profile phone
type Phone string

// NewPhone create profile phone value object
func NewPhone(value string) (*Phone, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, errs.ErrPhoneEmpty
	}

	if !strings.HasPrefix(value, "+") {
		return nil, errs.ErrInvalidPhoneFormat
	}

	if !isNumeric(value[1:]) {
		return nil, errs.ErrPhoneContainsLetters
	}

	if err := validatePhoneFormat(value); err != nil {
		return nil, err
	}

	phone := Phone(value)
	return &phone, nil
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

// String type cast from Phone to string
func (up Phone) String() string {
	return string(up)
}

// Mask return masked Phone as string
func (up Phone) Mask() string {
	return "****" + up.String()[len(up.String())-4:]
}
