package auth

import (
	"time"

	"github.com/google/uuid"
)

// User is entity to store auth user data with credentials
type User struct {
	id        string
	login     Login
	password  PasswordHash
	createdAt time.Time
	updatedAt time.Time
}

// NewUser create auth user entity
func NewUser(login Login, password PasswordHash) *User {
	return &User{
		id:        uuid.NewString(),
		login:     login,
		password:  password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

// RestoreUser build User from all fields
func RestoreUser(id string, login Login, password string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		login:     login,
		password:  PasswordHash(password),
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (au *User) Id() string {
	return au.id
}

func (au *User) Login() *Login {
	return &au.login
}

func (au *User) UpdateLogin(new string) error {
	newLogin, err := NewLogin(new)
	if err != nil {
		return err
	}

	au.login = *newLogin
	return nil
}

func (au *User) Password() *PasswordHash {
	return &au.password
}

func (au *User) PasswordCompare(plain PasswordPlain) bool {
	return au.password.Compare(plain)
}

func (au *User) CreatedAt() time.Time {
	return au.createdAt
}

func (au *User) UpdatedAt() time.Time {
	return au.updatedAt
}
