package model

// User is entity to store public user data
type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     UserEmail
	Phone     UserPhone
	Nickname  UserNickname
}

func NewUser(id, firstName, lastName string, nickname UserNickname, phone UserPhone, email UserEmail) User {
	return User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Nickname:  nickname,
	}
}
