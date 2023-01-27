package entity

import "time"

type Client struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

type Event struct {
	Id         int       `json:"id,omitempty"`
	DoctorName string    `json:"doctor_name,omitempty"`
	ClientId   int       `json:"client_id,omitempty"`
	StartsAt   time.Time `json:"starts_at,omitempty"`
	EndsAt     time.Time `json:"ends_at,omitempty"`
}

type Schedule struct {
	FirstDay      time.Time `json:"first_day,omitempty"`
	LastDay       time.Time `json:"last_day,omitempty"`
	StartTime     time.Time `json:"start_time,omitempty"`
	EndTime       time.Time `json:"end_time,omitempty"`
	StartBreak    time.Time `json:"start_break,omitempty"`
	EndBreak      time.Time `json:"end_break,omitempty"`
	EventDuration int       `json:"event_duration,omitempty"`
	DoctorName    []string  `json:"doctor_name,omitempty"`
}
