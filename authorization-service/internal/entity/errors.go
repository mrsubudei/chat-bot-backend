package entity

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with these data already exists")
	ErrUserDoesNotExist  = errors.New("user does not exist")
	ErrUserNotVerified   = errors.New("user has not verified his email")
	ErrTokenNotValid     = errors.New("token is not valid")
	ErrTokenTTLExpired   = errors.New("token live time is expired")
)
