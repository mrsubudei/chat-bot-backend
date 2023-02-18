package entity

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with these data already exists")
	ErrUserDoesNotExist  = errors.New("user does not exist")
)
