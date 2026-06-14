package dto

type UserSignUpRequest struct {
	Login    string `json:"login" binding:"required,max=32"`
	Password string `json:"password" binding:"required,max=64"`
}

type UserSignInRequest struct {
	Login    string `json:"login" binding:"required,max=32"`
	Password string `json:"password" binding:"required,max=64"`
}
