package request

import (
	"tanneryApp/internal/application/auth/command"
)

type SignUpRequest struct {
	Login    string `json:"login" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

func (r *SignUpRequest) ToSignUpCommand() command.SignUpCommand {
	return command.SignUpCommand{
		Login:    r.Login,
		Password: r.Password,
	}
}
