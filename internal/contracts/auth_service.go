package contracts

import (
	"context"
)

type AuthService interface {
	SignUp(ctx context.Context, login, password string) error
	SignIn(ctx context.Context, login, password string) (SignInResult, error)
}

type SignInResult struct {
	UserId      string
	AccessToken string
}
