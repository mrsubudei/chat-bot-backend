package entity

import "time"

type User struct {
	Id                int32
	Name              string
	Phone             string
	Email             string
	Password          string
	Role              string
	Verified          bool
	VerificationToken string
	VerificationTTL   time.Time
	SessionToken      string
	SessionTTL        time.Time
}
