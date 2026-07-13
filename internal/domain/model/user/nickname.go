package user

import (
	"regexp"
	"strings"
	"tanneryApp/internal/domain/errors"
)

// Nickname is value object to store and validate profile's nickname
type Nickname string

const MaxNicknameLength = 30

// NewNickname create profile nickname value object
func NewNickname(value string) (*Nickname, error) {
	value = strings.TrimSpace(value)

	if len(value) > MaxNicknameLength {
		return nil, errors.ErrNicknameTooLong
	}

	if !validateNicknameFormat(value) {
		return nil, errors.ErrInvalidNicknameFormat
	}

	nickname := Nickname(value)
	return &nickname, nil
}

// validateNicknameFormat validate nickname on allowed chars
func validateNicknameFormat(nickname string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return re.MatchString(nickname)
}

// String type cast from Nickname to string
func (u Nickname) String() string {
	return string(u)
}

func (u Nickname) Equals(value Nickname) bool {
	return u.String() == value.String()
}
