package model

// AuthUser is entity to store auth user data with credentials
type AuthUser struct {
	Id       string
	Login    string
	Password UserPasswordHashed
}

// NewAuthUser create auth user entity
func NewAuthUser(id, login string, password UserPasswordHashed) AuthUser {
	return AuthUser{
		Id:       id,
		Login:    login,
		Password: password,
	}
}
