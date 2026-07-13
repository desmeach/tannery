package auth

import (
	"fmt"
	"regexp"
	"tanneryApp/internal/domain/errors"
)

// Login is value object to store profile login
type Login string

const MinLoginLength = 5
const MaxLoginLength = 20

// NewLogin create plain password value object
func NewLogin(value string) (*Login, error) {
	if len(value) < MinLoginLength {
		return nil, fmt.Errorf("%w: minimum %d characters required", errors.ErrLoginTooShort, MinLoginLength)
	}

	if len(value) > MaxLoginLength {
		return nil, fmt.Errorf("%w: login can be only %d characters long", errors.ErrLoginTooLong, MaxLoginLength)
	}

	pattern := "^[a-zA-Z0-9_]+$"
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.ErrLoginInvalidCharacters
	}

	login := Login(value)
	return &login, nil
}

// Login type cast from Login to string
func (l Login) String() string {
	return string(l)
}
