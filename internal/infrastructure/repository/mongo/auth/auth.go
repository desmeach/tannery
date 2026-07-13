package user

import (
	"context"
	"errors"
	"sync"
	errs "tanneryApp/internal/domain/errors"
	"tanneryApp/internal/domain/model/auth"
	"tanneryApp/internal/infrastructure/repository/mongo/auth/entity"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository of auth profile collection
type Repository struct {
	mu sync.Mutex
	db *mongo.Collection
}

// NewRepository create auth profile repo
func NewRepository(db mongo.Database) *Repository {
	return &Repository{
		db: db.Collection("auth"),
	}
}

// Create auth profile in db
func (r *Repository) Create(ctx context.Context, user auth.User) (*uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	record, err := entity.ToUserRecord(user)
	if err != nil {
		return nil, err
	}

	result, err := r.db.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(result.InsertedID.(string))
	if err != nil {
		return nil, err
	}

	return &id, nil
}

// findBy find auth profile by column
func (r *Repository) findBy(ctx context.Context, column string, value interface{}) (*auth.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.FindOne(ctx, bson.M{column: value})

	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, errs.ErrUserNotFound
	}

	userRecord := entity.UserRecord{}
	if err := result.Decode(&userRecord); err != nil {
		return nil, err
	}

	userEntity, err := entity.ToUser(userRecord)
	if err != nil {
		return nil, err
	}

	return userEntity, nil
}

// FindById returns auth profile by id
func (r *Repository) FindById(ctx context.Context, id uuid.UUID) (*auth.User, error) {
	userModel, err := r.findBy(ctx, "_id", id.String())
	if err != nil {
		return nil, err
	}

	return userModel, nil
}

// FindByLogin returns auth profile by login
func (r *Repository) FindByLogin(ctx context.Context, login string) (*auth.User, error) {
	userModel, err := r.findBy(ctx, "login", login)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}
