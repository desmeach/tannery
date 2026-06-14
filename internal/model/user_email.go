package model

import "tanneryApp/internal/errors"

type UserEmail string

const MaxEmailLength = 100

func NewUserEmail(value string) (UserEmail, error) {
	if len(value) > MaxEmailLength {
		return "", errors.ErrEmailTooLong
	}

	return UserEmail(value), nil
}

func (u UserEmail) String() string {
	return string(u)
}
