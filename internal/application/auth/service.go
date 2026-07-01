package auth

import (
	"context"
	"log/slog"
	"tanneryApp/internal/domain/errors"
	"tanneryApp/internal/domain/model"
	"tanneryApp/internal/infrastructure/jwt"
)

type Repository interface {
	Create(ctx context.Context, user model.AuthUser) error
	FindById(ctx context.Context, id string) (model.AuthUser, error)
	FindByLogin(ctx context.Context, login string) (model.AuthUser, error)
}

type Service struct {
	logger     *slog.Logger
	repo       Repository
	jwtManager *jwt.Manager
}

func NewUserService(logger *slog.Logger, repo Repository, jwtManager *jwt.Manager) *Service {
	return &Service{
		logger:     logger,
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (s *Service) SignUp(ctx context.Context, login, password string) error {
	plainPass, err := model.NewUserPasswordPlain(password)
	if err != nil {
		return err
	}

	passHash, err := plainPass.Hash()
	if err != nil {
		return err
	}

	user := model.NewAuthUser("", login, passHash)

	err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SignIn(ctx context.Context, login, password string) (SignInResult, error) {
	user, err := s.repo.FindByLogin(ctx, login)
	if err != nil {
		return SignInResult{}, err
	}

	plainPass, err := model.NewUserPasswordPlain(password)
	if err != nil {
		return SignInResult{}, err
	}

	if !user.Password.Compare(plainPass) {
		return SignInResult{}, errors.ErrSignInInvalidData
	}

	token, err := s.jwtManager.Build(user.Id)
	if err != nil {
		return SignInResult{}, err
	}

	return SignInResult{UserId: user.Id, AccessToken: token}, nil
}
