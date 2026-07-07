package model

import (
	"github.com/google/uuid"
)

// AuthUser is entity to store auth user data with credentials
type AuthUser struct {
	id       string
	login    string
	password AuthUserPasswordHashed
}

// NewAuthUser create auth user entity
func NewAuthUser(login string, password AuthUserPasswordHashed) AuthUser {
	return AuthUser{
		id:       uuid.NewString(),
		login:    login,
		password: password,
	}
}

func RestoreAuthUser(id string, login string, password string) AuthUser {
	return AuthUser{
		id:       id,
		login:    login,
		password: AuthUserPasswordHashed(password),
	}
}

func (au *AuthUser) Id() string {
	return au.id
}

func (au *AuthUser) Login() string {
	return au.login
}

func (au *AuthUser) Password() AuthUserPasswordHashed {
	return au.password
}

func (au *AuthUser) PasswordCompare(plain AuthUserPasswordPlain) bool {
	return au.password.Compare(plain)
}
