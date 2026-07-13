package request

import (
	"tanneryApp/internal/application/user/command"

	"github.com/google/uuid"
)

type CreateProfileRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     *string `json:"email,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Nickname  string  `json:"nickname"`
}

func (r *CreateProfileRequest) ToCreateProfileCommand(userId uuid.UUID) *command.CreateProfileCommand {
	return &command.CreateProfileCommand{
		UserId:    userId,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Phone:     r.Phone,
		Nickname:  r.Nickname,
	}
}
