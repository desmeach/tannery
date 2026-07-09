package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash is value object to store password's hash
type PasswordHash string

// String type cast from PasswordHash to string
func (h PasswordHash) String() string {
	return string(h)
}

// Compare hash password with plain
func (h PasswordHash) Compare(plain PasswordPlain) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(plain))
	return err == nil
}
