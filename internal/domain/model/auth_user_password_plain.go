package model

import (
	"fmt"
	"regexp"
	"tanneryApp/internal/domain/errors"

	"golang.org/x/crypto/bcrypt"
)

// AuthUserPasswordPlain is value object to store password which can be hashed
type AuthUserPasswordPlain string

const MinPasswordLength = 8

// NewAuthUserPasswordPlain create plain password value object
func NewAuthUserPasswordPlain(value string) (AuthUserPasswordPlain, error) {
	if len(value) < MinPasswordLength {
		return "", fmt.Errorf("%w: minimum %d characters required", errors.ErrPasswordTooShort, MinPasswordLength)
	}

	pattern := "^[a-zA-Z0-9]+$"
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", errors.ErrPasswordInvalidCharacters
	}

	return AuthUserPasswordPlain(value), nil
}

// Hash create plain password's hash
func (p AuthUserPasswordPlain) Hash() (AuthUserPasswordHashed, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return AuthUserPasswordHashed(hash), nil
}

// String type cast from AuthUserPasswordPlain to string
func (p AuthUserPasswordPlain) String() string {
	return string(p)
}
