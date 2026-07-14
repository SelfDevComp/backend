package errors

import "errors"

var (
	ErrInvalidEmail          = errors.New("invalid email")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username is unavailable")

	ErrUserNotFound = errors.New("user not found")

	ErrEmailNotVerified = errors.New("email not verified")
	ErrUnauthorized     = errors.New("incorrect password")
	ErrTokenWasExpired  = errors.New("token was expired")
)
