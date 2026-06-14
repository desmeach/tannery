package model

type AuthUser struct {
	Id       string
	Login    string
	Password UserPasswordHashed
}

func NewAuthUser(id, login string, password UserPasswordHashed) AuthUser {
	return AuthUser{
		Id:       id,
		Login:    login,
		Password: password,
	}
}
