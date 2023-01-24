package entity

import "time"

type Client struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Event struct {
	Id       int       `json:"id"`
	Category string    `json:"category"`
	ClientId int       `json:"client_id"`
	Title    string    `json:"title"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}

type Schedule struct {
	FirstDay      time.Time `json:"first_day"`
	LastDay       time.Time `json:"last_day"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	StartBreak    time.Time `json:"start_break"`
	EndBreak      time.Time `json:"end_break"`
	EventDuration int       `json:"event_duration"`
}
