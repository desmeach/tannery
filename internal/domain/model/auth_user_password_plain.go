package model

import (
	"fmt"
	"regexp"
	"tanneryApp/internal/domain/errors"

	"golang.org/x/crypto/bcrypt"
)

// UserPasswordPlain is value object to store password which can be hashed
type UserPasswordPlain string

const MinPasswordLength = 8

// NewUserPasswordPlain create plain password value object
func NewUserPasswordPlain(value string) (UserPasswordPlain, error) {
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

	return UserPasswordPlain(value), nil
}

// Hash create plain password's hash
func (p UserPasswordPlain) Hash() (UserPasswordHashed, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return UserPasswordHashed(hash), nil
}

// String type cast from NewUserPasswordPlain to string
func (p UserPasswordPlain) String() string {
	return string(p)
}
