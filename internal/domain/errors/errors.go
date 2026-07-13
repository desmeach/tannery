package errors

import "errors"

var (
	ErrEmailTooLong              = errors.New("email too long")
	ErrNicknameTooLong           = errors.New("nickname too long")
	ErrPasswordInvalidCharacters = errors.New("password contains invalid characters")
	ErrPasswordTooShort          = errors.New("password is too short")
	ErrPasswordTooLong           = errors.New("password is too long")
	ErrLoginTooShort             = errors.New("login is too short")
	ErrLoginTooLong              = errors.New("login is too long")
	ErrLoginInvalidCharacters    = errors.New("login contains invalid characters")
	ErrGroupMemberCountInvalid   = errors.New("member count is invalid")
	ErrSignInInvalidData         = errors.New("credentials is invalid")
	ErrUserNotFound              = errors.New("user not found")
	ErrProfileNotFound           = errors.New("profile not found")
	ErrProfileFirstNameIsEmpty   = errors.New("firstname is empty")
	ErrProfileLastNameIsEmpty    = errors.New("lastname is empty")
	ErrInvalidEmailFormat        = errors.New("invalid email format: must be *@*.*")
	ErrInvalidNicknameFormat     = errors.New("invalid nickname format: must contain only letters, numbers, '_' and '-'")
	ErrInvalidPhoneFormat        = errors.New("invalid phone format: must start with '+'")
	ErrInvalidPhoneLength        = errors.New("invalid phone length: must be 12 digits including country code")
	ErrPhoneContainsLetters      = errors.New("phone number cannot contain letters")
	ErrPhoneEmpty                = errors.New("phone number cannot be empty")
)
