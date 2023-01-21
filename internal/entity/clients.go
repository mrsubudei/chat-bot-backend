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

type Day struct {
	DayStarts      time.Time `json:"day_starts"`
	DayEnds        time.Time `json:"day_ends"`
	EventsQuantity int       `json:"events_quantity"`
	EventDuration  int       `json:"event_duration"`
	BreakStarts    time.Time `json:"break_starts"`
	BreakEnds      time.Time `json:"break_ends"`
}
