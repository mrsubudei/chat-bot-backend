package entity

import "time"

type Doctor struct {
	Id      int32
	Name    string
	Surname string
	Phone   string
}

type Event struct {
	Id       int32
	ClientId int32
	DoctorId int32
	StartsAt time.Time
	EndsAt   time.Time
}

type Schedule struct {
	FirstDay      time.Time
	LastDay       time.Time
	StartTime     time.Time
	EndTime       time.Time
	StartBreak    time.Time
	EndBreak      time.Time
	EventDuration int32
	DoctorIds     []int32
}
