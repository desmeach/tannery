package user

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	errs "tanneryApp/internal/domain/errors"
	"tanneryApp/internal/domain/model"
	"tanneryApp/internal/infrastructure/repository/mongo/user/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	mu     sync.Mutex
	logger *slog.Logger
	db     *mongo.Collection
}

func NewRepository(logger *slog.Logger, db *mongo.Database) *Repository {
	return &Repository{
		logger: logger,
		db:     db.Collection("user"),
	}
}

func (r *Repository) Create(ctx context.Context, user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	record, err := entity.ToUserRecord(user)
	if err != nil {
		return err
	}

	_, err = r.db.InsertOne(ctx, record)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, userId string) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO: implement

	return model.User{}, errors.New("not implemented")
}

func (r *Repository) FindBy(ctx context.Context, column string, value interface{}) (model.User, error) {
	result := r.db.FindOne(ctx, bson.M{column: value})
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return model.User{}, errs.ErrUserNotFound
	}

	userRecord := entity.UserRecord{}
	if err := result.Decode(userRecord); err != nil {
		return model.User{}, err
	}

	user, err := entity.ToUser(userRecord)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) FindByNickname(ctx context.Context, nickname string) (model.User, error) {
	return r.FindBy(ctx, "nickname", nickname)
}

func (r *Repository) FindById(ctx context.Context, userId string) (model.User, error) {
	return r.FindBy(ctx, "_id", userId)
}
