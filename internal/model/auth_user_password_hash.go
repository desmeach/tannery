package model

import (
	"golang.org/x/crypto/bcrypt"
)

type UserPasswordHashed string

func NewUserPasswordHashed(value string) UserPasswordHashed {
	return UserPasswordHashed(value)
}

func (h UserPasswordHashed) String() string {
	return string(h)
}

func (h UserPasswordHashed) Compare(plain UserPasswordPlain) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(plain))
	return err == nil
}
