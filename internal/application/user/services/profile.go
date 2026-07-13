package ProfileServices

import (
	"context"
	"log/slog"
	"tanneryApp/internal/application/user/command"
	"tanneryApp/internal/application/user/query"
	"tanneryApp/internal/domain/model/user"
	"tanneryApp/internal/infrastructure/jwt"
	"tanneryApp/internal/infrastructure/logger/sl"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, profile user.Profile) (*uuid.UUID, error)
	Update(ctx context.Context, profile user.Profile) error
	FindById(ctx context.Context, id uuid.UUID) (*user.Profile, error)
	FindByUserId(ctx context.Context, userId uuid.UUID) (*user.Profile, error)
	FindByPhone(ctx context.Context, phone string) (*user.Profile, error)
	FindByNickname(ctx context.Context, nickname string) (*user.Profile, error)
}

type ProfileService struct {
	logger     slog.Logger
	repo       Repository
	jwtManager jwt.Manager
}

func NewProfileService(logger slog.Logger, repo Repository, jwtManager jwt.Manager) *ProfileService {
	return &ProfileService{
		logger:     logger,
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (s *ProfileService) Create(ctx context.Context, cmd command.CreateProfileCommand) (*command.CreateProfileCommandResult, error) {
	op := "user.profileService.create"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("firstname", cmd.FirstName),
		slog.String("lastname", cmd.LastName),
		slog.String("nickname", cmd.Nickname),
	)

	var email *user.Email
	if cmd.Email != nil {
		emailVO, err := user.NewEmail(*cmd.Email)
		if err != nil {
			log.Error("failed to create profile email", sl.Err(err))
			return nil, err
		}
		email = emailVO

		log.With(slog.String("email", email.Mask()))
	}

	var phone *user.Phone
	if cmd.Phone != nil {
		phoneVO, err := user.NewPhone(*cmd.Phone)
		if err != nil {
			log.Error("failed to create profile phone", sl.Err(err))
			return nil, err
		}
		phone = phoneVO

		log.With(slog.String("phone", phone.Mask()))
	}

	nickname, err := user.NewNickname(cmd.Nickname)
	if err != nil {
		log.Error("failed to create profile nickname", sl.Err(err))
		return nil, err
	}

	profile, err := user.NewProfile(cmd.UserId, cmd.FirstName, cmd.LastName, *nickname, email, phone)
	if err != nil {
		log.Error("failed to create profile entity", sl.Err(err))
		return nil, err
	}

	res, err := s.repo.Create(ctx, *profile)
	if err != nil {
		log.Error("failed to create profile in repo", sl.Err(err))
		return nil, err
	}

	result := &command.CreateProfileCommandResult{
		Id: *res,
	}

	return result, err
}

func (s *ProfileService) FindById(ctx context.Context, q query.FindByIdQuery) (*query.FindByIdQueryResult, error) {
	op := "user.profileService.findById"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("id", q.Id.String()),
	)

	profile, err := s.repo.FindById(ctx, q.Id)
	if err != nil {
		log.Error("failed to find profile", sl.Err(err))
		return nil, err
	}

	result := query.FindByIdQueryResultFromEntity(profile)

	return result, nil
}

func (s *ProfileService) FindByUserId(ctx context.Context, q query.FindByUserIdQuery) (*query.FindByUserIdQueryResult, error) {
	op := "user.profileService.findByUserId"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("userId", q.UserId.String()),
	)

	profile, err := s.repo.FindByUserId(ctx, q.UserId)
	if err != nil {
		log.Error("failed to find profile", sl.Err(err))
		return nil, err
	}

	result := query.FindByUserIdQueryResultFromEntity(profile)

	return result, nil
}

func (s *ProfileService) UpdateNickname(ctx context.Context, cmd command.UpdateProfileCommand) error {
	op := "user.profileService.updateNickname"

	log := s.logger.With(
		slog.String("op", op),
		slog.String("id", cmd.Id.String()),
		slog.String("newUsername", cmd.Nickname),
	)

	newNickname, err := user.NewNickname(cmd.Nickname)
	if err != nil {
		log.Error("failed to create profile nickname", sl.Err(err))
		return err
	}

	profile, err := s.repo.FindById(ctx, cmd.Id)
	if err != nil {
		log.Error("failed to find profile", sl.Err(err))
		return err
	}

	if profile.Nickname().Equals(*newNickname) {
		return nil
	}

	profile.SetNickname(*newNickname)

	err = s.repo.Update(ctx, *profile)
	if err != nil {
		log.Error("failed to update profile nickname", sl.Err(err))
		return err
	}

	return nil
}
