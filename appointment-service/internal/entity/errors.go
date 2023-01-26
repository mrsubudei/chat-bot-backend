package entity

import "errors"

var (
	ErrNoData              = errors.New("there are no data")
	ErrUniqueDateViolation = errors.New("events with date already exist")
)
