package command

import "github.com/google/uuid"

type UpdateProfileCommand struct {
	Id       uuid.UUID
	Nickname string
}
