package entity

import (
	"fmt"
	"tanneryApp/internal/domain/model/user"

	"github.com/google/uuid"
)

func ToProfileRecord(user user.Profile) (*ProfileRecord, error) {
	fmt.Println(user)

	var email *string
	if user.Email() != nil {
		emailStr := user.Email().String()
		email = &emailStr
	}

	var phone *string
	if user.Phone() != nil {
		phoneStr := user.Phone().String()
		phone = &phoneStr
	}

	return &ProfileRecord{
		Id:        user.ID().String(),
		UserId:    user.UserID().String(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     email,
		Phone:     phone,
		Nickname:  user.Nickname().String(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}

func ToProfile(record ProfileRecord) (*user.Profile, error) {
	var email *user.Email
	if record.Email != nil {
		emailVO, err := user.NewEmail(*record.Email)
		if err != nil {
			return nil, err
		}
		email = emailVO
	}

	var phone *user.Phone
	if record.Phone != nil {
		phoneVO, err := user.NewPhone(*record.Phone)
		if err != nil {
			return nil, err
		}
		phone = phoneVO
	}

	nickname, err := user.NewNickname(record.Nickname)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(record.Id)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(record.UserId)
	if err != nil {
		return nil, err
	}

	profile := user.RestoreProfile(
		id,
		userId,
		record.FirstName,
		record.LastName,
		*nickname,
		email,
		phone,
		record.CreatedAt,
		record.UpdatedAt,
	)

	return profile, nil
}
