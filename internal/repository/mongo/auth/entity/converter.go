package entity

import (
	"tanneryApp/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToAuthUserRecord(user model.AuthUser) (AuthUserRecord, error) {
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

//func ToAuthUserRecordCreate(user model.AuthUser) AuthUserRecordCreate {
//	return AuthUserRecordCreate{
//		Login:    user.Login,
//		Password: user.Password.String(),
//	}
//}

func ToAuthUser(record AuthUserRecord) (model.AuthUser, error) {
	passHash := model.NewUserPasswordHashed(record.Password)

	return model.AuthUser{
		Id:       record.Id.Hex(),
		Login:    record.Login,
		Password: passHash,
	}, nil
}
