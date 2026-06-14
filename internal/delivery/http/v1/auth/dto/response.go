package dto

type UserSignInResponse struct {
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
}
