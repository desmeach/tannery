package response

import (
	"tanneryApp/internal/application/user/query"
	"time"

	"github.com/google/uuid"
)

type GetCurrentProfileResponse struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToGetCurrentProfileResponse(q query.FindByUserIdQueryResult) *GetCurrentProfileResponse {
	var email string
	if q.Email != nil {
		email = q.Email.String()
	}

	var phone string
	if q.Phone != nil {
		email = q.Phone.String()
	}

	return &GetCurrentProfileResponse{
		Id:        q.Id,
		FirstName: q.FirstName,
		LastName:  q.LastName,
		Email:     email,
		Phone:     phone,
		Nickname:  q.Nickname.String(),
		CreatedAt: q.CreatedAt,
		UpdatedAt: q.UpdatedAt,
	}
}
