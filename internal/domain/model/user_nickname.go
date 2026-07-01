package model

import (
	"regexp"
	"strings"
	"tanneryApp/internal/domain/errors"
)

// UserNickname is value object to store and validate user's nickname
type UserNickname string

const MaxNicknameLength = 30

// NewUserNickname create user nickname value object
func NewUserNickname(value string) (UserNickname, error) {
	value = strings.TrimSpace(value)

	if len(value) > MaxNicknameLength {
		return "", errors.ErrNicknameTooLong
	}

	if !validateNicknameFormat(value) {
		return "", errors.ErrInvalidNicknameFormat
	}

	return UserNickname(value), nil
}

// validateNicknameFormat validate nickname on allowed chars
func validateNicknameFormat(nickname string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return re.MatchString(nickname)
}

// String type cast from UserNickname to string
func (u UserNickname) String() string {
	return string(u)
}
