package model

import (
	"golang.org/x/crypto/bcrypt"
)

// UserPasswordHashed is value object to store password's hash
type UserPasswordHashed string

// NewUserPasswordHashed create password hash value object
func NewUserPasswordHashed(value string) UserPasswordHashed {
	return UserPasswordHashed(value)
}

// String type cast from UserPasswordHashed to string
func (h UserPasswordHashed) String() string {
	return string(h)
}

// Compare hash password with plain
func (h UserPasswordHashed) Compare(plain UserPasswordPlain) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(plain))
	return err == nil
}
