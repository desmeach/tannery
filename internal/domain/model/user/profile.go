package user

import (
	errs "tanneryApp/internal/domain/errors"
	"time"

	"github.com/google/uuid"
)

// Profile is entity to store public user Profile
type Profile struct {
	id        uuid.UUID
	userId    uuid.UUID
	firstName string
	lastName  string
	email     *Email
	phone     *Phone
	nickname  Nickname
	createdAt time.Time
	updatedAt time.Time
}

func NewProfile(
	userId uuid.UUID,
	firstName, lastName string,
	nickname Nickname,
	email *Email,
	phone *Phone,
) (*Profile, error) {
	if firstName == "" {
		return nil, errs.ErrProfileFirstNameIsEmpty
	}
	if lastName == "" {
		return nil, errs.ErrProfileLastNameIsEmpty
	}

	return &Profile{
		id:        uuid.New(),
		userId:    userId,
		firstName: firstName,
		lastName:  lastName,
		nickname:  nickname,
		email:     email,
		phone:     phone,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func RestoreProfile(
	id,
	userId uuid.UUID,
	firstName,
	lastName string,
	nickname Nickname,
	email *Email,
	phone *Phone,
	createdAt,
	updatedAt time.Time,
) *Profile {
	return &Profile{
		id:        id,
		userId:    userId,
		firstName: firstName,
		lastName:  lastName,
		nickname:  nickname,
		email:     email,
		phone:     phone,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *Profile) ID() uuid.UUID {
	return u.id
}

func (u *Profile) UserID() uuid.UUID {
	return u.userId
}

func (u *Profile) FirstName() string {
	return u.firstName
}

func (u *Profile) LastName() string {
	return u.lastName
}

func (u *Profile) FullName() string {
	return u.firstName + " " + u.lastName
}

func (u *Profile) Email() *Email {
	return u.email
}

func (u *Profile) Phone() *Phone {
	return u.phone
}

func (u *Profile) Nickname() Nickname {
	return u.nickname
}

func (u *Profile) CreatedAt() time.Time {
	return u.createdAt
}

func (u *Profile) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *Profile) SetPhone(phone Phone) {
	u.phone = &phone
}

func (u *Profile) SetEmail(email Email) {
	u.email = &email
}

func (u *Profile) SetNickname(nickname Nickname) {
	u.nickname = nickname
}
