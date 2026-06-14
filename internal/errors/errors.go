package errors

import "errors"

var (
	ErrEmailTooLong              = errors.New("email too long")
	ErrPasswordInvalidCharacters = errors.New("password contains invalid characters")
	ErrPasswordTooShort          = errors.New("password is too short")
	ErrGroupMemberCountInvalid   = errors.New("member count is invalid")
	ErrSignInInvalidData         = errors.New("credentials is invalid")
	ErrUserNotFound              = errors.New("user not found")
)
