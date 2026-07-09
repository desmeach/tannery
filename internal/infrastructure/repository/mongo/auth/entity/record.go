package entity

import "time"

type UserRecord struct {
	Id        string    `bson:"_id,omitempty"`
	Login     string    `bson:"login"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}
