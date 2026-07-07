package model

import (
	"golang.org/x/crypto/bcrypt"
)

// AuthUserPasswordHashed is value object to store password's hash
type AuthUserPasswordHashed string

// NewAuthUserPasswordHashed create password hash value object
func NewAuthUserPasswordHashed(value string) AuthUserPasswordHashed {
	return AuthUserPasswordHashed(value)
}

// String type cast from AuthUserPasswordHashed to string
func (h AuthUserPasswordHashed) String() string {
	return string(h)
}

// Compare hash password with plain
func (h AuthUserPasswordHashed) Compare(plain AuthUserPasswordPlain) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(plain))
	return err == nil
}
