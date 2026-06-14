package user

import (
	"context"
	"log/slog"
	"tanneryApp/internal/contracts"
	"tanneryApp/internal/errors"
	"tanneryApp/internal/model"
	"tanneryApp/pkg/jwt"
)

type SignInOut struct {
	UserId      string
	AccessToken string
}

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

func (s *Service) SignIn(ctx context.Context, login, password string) (contracts.SignInResult, error) {
	user, err := s.repo.FindByLogin(ctx, login)
	if err != nil {
		return contracts.SignInResult{}, err
	}

	plainPass, err := model.NewUserPasswordPlain(password)
	if err != nil {
		return contracts.SignInResult{}, err
	}

	if !user.Password.Compare(plainPass) {
		return contracts.SignInResult{}, errors.ErrSignInInvalidData
	}

	token, err := s.jwtManager.Build(user.Id)
	if err != nil {
		return contracts.SignInResult{}, err
	}

	return contracts.SignInResult{UserId: user.Id, AccessToken: token}, nil
}
