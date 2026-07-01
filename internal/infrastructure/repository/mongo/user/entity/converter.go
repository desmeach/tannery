package entity

import (
	model2 "tanneryApp/internal/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToUserRecord(user model2.User) (UserRecord, error) {
	id, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return UserRecord{}, err
	}

	return UserRecord{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     string(user.Email),
		Phone:     user.Phone.String(),
		Nickname:  user.Nickname.String(),
	}, nil
}

func ToUser(record UserRecord) (model2.User, error) {
	email, err := model2.NewUserEmail(record.Email)
	if err != nil {
		return model2.User{}, err
	}

	phone, err := model2.NewUserPhone(record.Phone)
	if err != nil {
		return model2.User{}, err
	}

	nickname, err := model2.NewUserNickname(record.Nickname)
	if err != nil {
		return model2.User{}, err
	}

	return model2.User{
		Id:        record.Id.Hex(),
		FirstName: record.FirstName,
		LastName:  record.LastName,
		Email:     email,
		Phone:     phone,
		Nickname:  nickname,
	}, nil
}
