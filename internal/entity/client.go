package entity

import "time"

type Client struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Event struct {
	Id       int64     `json:"id"`
	Type     string    `json:"event_type"`
	ClientId int64     `json:"client_id"`
	Title    string    `json:"title"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}
