package response

import "tanneryApp/internal/application/auth/query"

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}

func ToSignInResponse(q query.SignInQueryResult) *SignInResponse {
	return &SignInResponse{
		AccessToken: q.AccessToken,
	}
}
