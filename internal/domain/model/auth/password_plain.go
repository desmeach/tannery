package auth

import (
	"fmt"
	"regexp"
	"tanneryApp/internal/domain/errors"

	"golang.org/x/crypto/bcrypt"
)

// PasswordPlain is value object to store password which can be hashed
type PasswordPlain string

const MinPasswordLength = 8
const MaxPasswordLength = 40

// NewPasswordPlain create plain password value object
func NewPasswordPlain(value string) (*PasswordPlain, error) {
	if len(value) < MinPasswordLength {
		return nil, fmt.Errorf("%w: minimum %d characters required", errors.ErrPasswordTooShort, MinPasswordLength)
	}

	if len(value) > MaxPasswordLength {
		return nil, fmt.Errorf("%w: password can be only %d characters long", errors.ErrPasswordTooLong, MaxPasswordLength)
	}

	pattern := "^[a-zA-Z0-9]+$"
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.ErrPasswordInvalidCharacters
	}

	plain := PasswordPlain(value)
	return &plain, nil
}

// Hash convert PasswordPlain to PasswordHash
func (p PasswordPlain) Hash() (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return PasswordHash(hash), nil
}

// String type cast from PasswordPlain to string
func (p PasswordPlain) String() string {
	return string(p)
}
