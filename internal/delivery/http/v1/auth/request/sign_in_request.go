package request

import "tanneryApp/internal/application/auth/query"

type SignInRequest struct {
	Login    string `json:"login" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

func (r *SignInRequest) ToSignInQuery() *query.SignInQuery {
	return &query.SignInQuery{
		Login:    r.Login,
		Password: r.Password,
	}
}
