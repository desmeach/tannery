package mongo

import (
	"context"
	"tanneryApp/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db.Collection("group"),
	}
}

func (r *Repository) Create(ctx context.Context, group *model.Group) error {
	_, err := r.db.InsertOne(ctx, group)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddUser(ctx context.Context, groupId string, userId string) error {
	_, err := r.db.Find(ctx, bson.D{{"id", groupId}})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindById(ctx context.Context, id string) (*model.Group, error) {
	result := r.db.FindOne(ctx, bson.M{"_id": id})
	if result.Err() != nil {
		return nil, result.Err()
	}

	group := &model.Group{}
	if err := result.Decode(group); err != nil {
		return nil, err
	}

	return group, nil
}
