package query

import (
	"tanneryApp/internal/domain/model/user"
	"time"

	"github.com/google/uuid"
)

type FindByIdQuery struct {
	Id uuid.UUID
}

type FindByIdQueryResult struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     *user.Email
	Phone     *user.Phone
	Nickname  user.Nickname
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindByIdQueryResultFromEntity(profile *user.Profile) *FindByIdQueryResult {
	if profile == nil {
		return nil
	}

	return &FindByIdQueryResult{
		Id:        profile.ID(),
		FirstName: profile.FirstName(),
		LastName:  profile.LastName(),
		Email:     profile.Email(),
		Phone:     profile.Phone(),
		Nickname:  profile.Nickname(),
		CreatedAt: profile.CreatedAt(),
		UpdatedAt: profile.UpdatedAt(),
	}
}
