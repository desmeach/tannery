package user

import (
	"context"
	"errors"
	"log/slog"
	errs "tanneryApp/internal/errors"
	"tanneryApp/internal/lib/logger/sl"
	"tanneryApp/internal/model"
	authEntity "tanneryApp/internal/repository/mongo/auth/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	logger *slog.Logger
	db     *mongo.Collection
}

func NewRepository(logger *slog.Logger, db *mongo.Database) *Repository {
	return &Repository{
		logger: logger,
		db:     db.Collection("user"),
	}
}

func (r *Repository) Create(ctx context.Context, user model.AuthUser) error {
	record, err := authEntity.ToAuthUserRecord(user)
	if err != nil {
		return err
	}
	_, err = r.db.InsertOne(ctx, record)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindById(ctx context.Context, id string) (model.AuthUser, error) {
	return model.AuthUser{}, errors.New("not implemented")
}

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

	userRecord := authEntity.AuthUserRecord{}
	if err := result.Decode(&userRecord); err != nil {
		log.Error("failed to decode user", op, sl.Err(err))
		return model.AuthUser{}, err
	}

	userModel, err := authEntity.ToAuthUser(userRecord)
	if err != nil {
		log.Error("failed to convert user record to entity", op, sl.Err(err))
		return model.AuthUser{}, err
	}

	return userModel, nil
}
