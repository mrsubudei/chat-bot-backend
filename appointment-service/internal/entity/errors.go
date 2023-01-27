package entity

import "errors"

var (
	ErrDateAlreadyExists   = errors.New("events with given dates already exist")
	ErrEntityAlreadyExists = errors.New("entity with these data already exists")
	ErrEntityDoesNotExist  = errors.New("entity with these data does not exist")
)
