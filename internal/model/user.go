package model

// User - user domain model
type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     UserEmail
	Phone     string
	Nickname  string
}

func NewUser(id, firstName, lastName, phone, nickname string, email UserEmail) User {
	return User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Nickname:  nickname,
	}
}
