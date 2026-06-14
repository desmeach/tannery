package model

import (
	"fmt"
	"regexp"
	"tanneryApp/internal/errors"

	"golang.org/x/crypto/bcrypt"
)

type UserPasswordPlain string

const MinPasswordLength = 8

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

func (p UserPasswordPlain) Hash() (UserPasswordHashed, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return UserPasswordHashed(hash), nil
}

func (p UserPasswordPlain) String() string {
	return string(p)
}
