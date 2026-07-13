package profile

import (
	"context"
	"errors"
	"sync"
	errs "tanneryApp/internal/domain/errors"
	"tanneryApp/internal/domain/model/user"
	"tanneryApp/internal/infrastructure/repository/mongo/profile/entity"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	mu sync.Mutex
	db *mongo.Collection
}

func NewRepository(db mongo.Database) *Repository {
	return &Repository{
		db: db.Collection("profile"),
	}
}

func (r *Repository) Create(ctx context.Context, profile user.Profile) (*uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	record, err := entity.ToProfileRecord(profile)
	if err != nil {
		return nil, err
	}

	res, err := r.db.InsertOne(ctx, record)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(res.InsertedID.(string))
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *Repository) Update(ctx context.Context, profile user.Profile) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	update := bson.M{}

	if firstName := profile.FirstName(); firstName != "" {
		update["firstname"] = firstName
	}
	if lastName := profile.LastName(); lastName != "" {
		update["lastname"] = lastName
	}
	if nickname := profile.Nickname(); nickname.String() != "" {
		update["nickname"] = nickname.String()
	}
	if email := profile.Email(); email != nil {
		update["email"] = email.String()
	}
	if phone := profile.Phone(); phone != nil {
		update["phone"] = phone.String()
	}

	update["updated_at"] = time.Now()

	if len(update) == 0 {
		return nil
	}

	res, err := r.db.UpdateByID(ctx, profile.ID(), bson.M{"$set": update})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errs.ErrProfileNotFound
	}

	return nil
}

func (r *Repository) findBy(ctx context.Context, column string, value interface{}) (*user.Profile, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.FindOne(ctx, bson.M{column: value})
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return nil, errs.ErrProfileNotFound
	}

	profileRecord := entity.ProfileRecord{}
	if err := result.Decode(&profileRecord); err != nil {
		return nil, err
	}

	profileEntity, err := entity.ToProfile(profileRecord)
	if err != nil {
		return nil, err
	}

	return profileEntity, nil
}

func (r *Repository) FindByNickname(ctx context.Context, nickname string) (*user.Profile, error) {
	res, err := r.findBy(ctx, "nickname", nickname)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) FindByPhone(ctx context.Context, phone string) (*user.Profile, error) {
	res, err := r.findBy(ctx, "phone", phone)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) FindById(ctx context.Context, id uuid.UUID) (*user.Profile, error) {
	res, err := r.findBy(ctx, "_id", id.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) FindByUserId(ctx context.Context, userId uuid.UUID) (*user.Profile, error) {
	res, err := r.findBy(ctx, "userId", userId.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}
