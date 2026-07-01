package entity

import (
	model2 "tanneryApp/internal/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToAuthUserRecord(user model2.AuthUser) (AuthUserRecord, error) {
	id, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return AuthUserRecord{}, err
	}
	return AuthUserRecord{
		Id:       id,
		Login:    user.Login,
		Password: user.Password.String(),
	}, nil
}

func ToAuthUser(record AuthUserRecord) (model2.AuthUser, error) {
	passHash := model2.NewUserPasswordHashed(record.Password)

	return model2.AuthUser{
		Id:       record.Id.Hex(),
		Login:    record.Login,
		Password: passHash,
	}, nil
}
