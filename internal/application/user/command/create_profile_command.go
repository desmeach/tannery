package command

import "github.com/google/uuid"

type CreateProfileCommand struct {
	UserId    uuid.UUID
	FirstName string
	LastName  string
	Nickname  string
	Email     *string
	Phone     *string
}

type CreateProfileCommandResult struct {
	Id uuid.UUID
}
