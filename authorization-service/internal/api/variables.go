package api

import "time"

const (
	DateAndTimeFormat = "2006-01-02 15:04:05"
	DateFormat        = "2006-01-02"
	DefaultAttempts   = 20
	DefaultTimeout    = time.Second
)

const (
	DuplicateErrMsg  = "duplicate key value violates unique constraint"
	NoRowsAffected   = "RowsAffected: %!w(<nil>)"
	CanNotParseTime  = "Can not parse time"
	RequestZeroValue = "Request has zero value"
	WrongEmailFormat = "Email has wrong format"
	WrongPhoneFormat = "Phone has wrong format"
	InternalErr      = "Internal Error"
)
