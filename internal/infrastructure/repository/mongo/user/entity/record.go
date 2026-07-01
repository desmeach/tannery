package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRecord struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstname"`
	LastName  string             `bson:"lastname"`
	Email     string             `bson:"email,omitempty"`
	Phone     string             `bson:"phone,omitempty"`
	Nickname  string             `bson:"nickname"`
}

func (u *UserRecord) GetID() string {
	return u.Id.Hex()
}
