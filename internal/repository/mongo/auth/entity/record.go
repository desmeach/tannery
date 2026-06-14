package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

//type AuthUserRecord struct {
//	Id       string
//	Login    string
//	Password string
//}
//
//type AuthUserRecordCreate struct {
//	Login    string
//	Password string
//}

type AuthUserRecord struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Login    string             `bson:"login"`
	Password string             `bson:"password"`
}

func (u *AuthUserRecord) GetID() string {
	return u.Id.Hex()
}
