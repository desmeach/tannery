package response

import (
	"tanneryApp/internal/application/user/command"

	"github.com/google/uuid"
)

type CreateProfileResponse struct {
	Id uuid.UUID `json:"id"`
}

func ToCreateProfileResponse(cmd command.CreateProfileCommandResult) *CreateProfileResponse {
	return &CreateProfileResponse{
		Id: cmd.Id,
	}
}
