package user

import (
	"context"
	"errors"
	"log/slog"
	errs "tanneryApp/internal/domain/errors"
	"tanneryApp/internal/domain/model"
	"tanneryApp/internal/infrastructure/logger/sl"
	"tanneryApp/internal/infrastructure/repository/mongo/auth/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository of auth user collection
type Repository struct {
	logger *slog.Logger
	db     *mongo.Collection
}

// NewRepository create auth user repo
func NewRepository(logger *slog.Logger, db *mongo.Database) *Repository {
	return &Repository{
		logger: logger,
		db:     db.Collection("auth_user"),
	}
}

// Create auth user in db
func (r *Repository) Create(ctx context.Context, user model.AuthUser) error {
	record, err := entity.ToAuthUserRecord(user)
	if err != nil {
		return err
	}
	_, err = r.db.InsertOne(ctx, record)
	if err != nil {
		return err
	}

	return nil
}

// FindById returns auth user by id
func (r *Repository) FindById(ctx context.Context, id string) (model.AuthUser, error) {
	return model.AuthUser{}, errors.New("not implemented")
}

// FindByLogin returns auth user by login
func (r *Repository) FindByLogin(ctx context.Context, login string) (model.AuthUser, error) {
	op := "auth.repo.findByLogin"

	log := r.logger.With(
		slog.String("op", op),
	)

	result := r.db.FindOne(ctx, bson.M{"login": login})
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		log.Error("user not found", op, sl.Err(result.Err()))
		return model.AuthUser{}, errs.ErrUserNotFound
	}

	userRecord := entity.AuthUserRecord{}
	if err := result.Decode(&userRecord); err != nil {
		log.Error("failed to decode user", op, sl.Err(err))
		return model.AuthUser{}, err
	}

	userModel := entity.ToAuthUser(userRecord)

	return userModel, nil
}
