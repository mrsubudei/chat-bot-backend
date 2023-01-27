package entity

import "errors"

var (
	ErrNoData           = errors.New("there are no data")
	ErrDateAlreadyExist = errors.New("events with date already exist")
)
