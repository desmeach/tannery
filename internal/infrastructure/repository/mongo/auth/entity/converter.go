package entity

import (
	model "tanneryApp/internal/domain/model"
)

func ToAuthUserRecord(user model.AuthUser) (AuthUserRecord, error) {
	return AuthUserRecord{
		Id:       user.Id(),
		Login:    user.Login(),
		Password: user.Password().String(),
	}, nil
}

func ToAuthUser(record AuthUserRecord) model.AuthUser {
	authUser := model.RestoreAuthUser(record.Id, record.Login, record.Password)
	return authUser
}
