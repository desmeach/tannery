package auth

import (
	"context"
	"log/slog"
	"tanneryApp/internal/application/auth/command"
	"tanneryApp/internal/application/auth/query"
	errs "tanneryApp/internal/domain/errors"
	"tanneryApp/internal/domain/model/auth"
	"tanneryApp/internal/infrastructure/jwt"
	"tanneryApp/internal/infrastructure/logger/sl"
)

type Repository interface {
	Create(ctx context.Context, user auth.User) error
	FindById(ctx context.Context, id string) (*auth.User, error)
	FindByLogin(ctx context.Context, login string) (*auth.User, error)
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

func (s *Service) SignUp(ctx context.Context, cmd command.SignUpCommand) error {
	op := "auth.service.signUp"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("login", cmd.Login),
	)

	plainPass, err := auth.NewPasswordPlain(cmd.Password)
	if err != nil {
		log.Error("failed to create plain password", op, sl.Err(err))
		return err
	}

	passHash, err := plainPass.Hash()
	if err != nil {
		log.Error("failed to hash plain password", op, sl.Err(err))
		return err
	}

	login, err := auth.NewLogin(cmd.Login)
	if err != nil {
		log.Error("failed to create login", op, sl.Err(err))
		return err
	}

	user := auth.NewUser(*login, passHash)

	err = s.repo.Create(ctx, *user)
	if err != nil {
		log.Error("failed to create user", op, sl.Err(err))
		return err
	}

	return nil
}

func (s *Service) SignIn(ctx context.Context, q query.SignInQuery) (*query.SignInQueryResult, error) {
	op := "auth.service.signIn"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("login", q.Login),
	)

	user, err := s.repo.FindByLogin(ctx, q.Login)
	if err != nil {
		log.Error("failed to find user", op, sl.Err(err))
		return nil, err
	}

	plainPass, err := auth.NewPasswordPlain(q.Password)
	if err != nil {
		log.Error("failed to create plain password", op, sl.Err(err))
		return nil, err
	}

	if !user.PasswordCompare(*plainPass) {
		return nil, errs.ErrSignInInvalidData
	}

	userId := user.Id()
	token, err := s.jwtManager.Build(userId)
	if err != nil {
		log.Error("failed to create jwt token", op, sl.Err(err))
		return nil, err
	}

	res := &query.SignInQueryResult{AccessToken: token}

	return res, nil
}
