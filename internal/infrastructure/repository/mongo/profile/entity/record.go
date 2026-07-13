package entity

import (
	"time"
)

type ProfileRecord struct {
	Id        string    `bson:"_id,omitempty"`
	UserId    string    `bson:"userId"`
	FirstName string    `bson:"firstname"`
	LastName  string    `bson:"lastname"`
	Nickname  string    `bson:"nickname"`
	Email     *string   `bson:"email"`
	Phone     *string   `bson:"phone"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}
