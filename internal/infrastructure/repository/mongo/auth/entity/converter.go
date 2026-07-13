package entity

import (
	"tanneryApp/internal/domain/model/auth"

	"github.com/google/uuid"
)

func ToUserRecord(user auth.User) (*UserRecord, error) {
	return &UserRecord{
		Id:        user.Id().String(),
		Login:     user.Login().String(),
		Password:  user.Password().String(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}

func ToUser(record UserRecord) (*auth.User, error) {
	login, err := auth.NewLogin(record.Login)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(record.Id)
	if err != nil {
		return nil, err
	}

	user := auth.RestoreUser(id, *login, record.Password, record.CreatedAt, record.UpdatedAt)

	return user, nil
}
