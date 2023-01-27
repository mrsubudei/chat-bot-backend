package entity

import "errors"

var (
	ErrEventAlreadyReserved = errors.New("event with given date already reserved")
	ErrEventIsNotReserved   = errors.New("event with given date is not reserved")
	ErrEventAlreadyExists   = errors.New("event with given date already exists")
	ErrEventDoesNotExist    = errors.New("event does not exist")
	ErrDoctorAlreadyExists  = errors.New("doctor with these data already exists")
	ErrDoctorDoesNotExist   = errors.New("doctor does not exist")
)
