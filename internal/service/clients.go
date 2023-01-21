package service

import (
	"context"
	"time"

	"github.com/mrsubudei/chat-bot-backend/internal/entity"
	"github.com/mrsubudei/chat-bot-backend/internal/repository"
)

type ClientsService struct {
	repo repository.Clients
}

func NewClientsService(repo repository.Clients) *ClientsService {
	return &ClientsService{repo}
}

func (cs *ClientsService) CreateSchedule(ctx context.Context, day entity.Day) error {
	dayEvents := []entity.Event{}

	starts := day.DayStarts
	increase := time.Duration(day.EventDuration * int(time.Second))
	ends := starts.Before(day.DayEnds.Add(increase))
	i := starts

	for ; ends; i = i.Add(increase) {
		if i.Equal(day.BreakStarts) {
			break
		}
		event := entity.Event{}
		event.StartsAt = i
		event.EndsAt = i.Add(increase)
		dayEvents = append(dayEvents, event)
	}

	if !day.BreakEnds.IsZero() {
		i = day.BreakEnds
		for ; ends; i = i.Add(increase) {
			if i.Equal(day.BreakStarts) {
				break
			}
			event := entity.Event{}
			event.StartsAt = i
			event.EndsAt = i.Add(increase)
			dayEvents = append(dayEvents, event)
		}
	}

	return nil
}
