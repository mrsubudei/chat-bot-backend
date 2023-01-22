package service

import (
	"context"
	"fmt"
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

var layout = "2006-01-02 15:04:05"

func (cs *ClientsService) CreateSchedule(ctx context.Context, schedule entity.Schedule) error {
	dayEvents := []entity.Event{}

	first := schedule.FirstDay
	last := schedule.LastDay.AddDate(0, 0, 1)
	increase := time.Duration(schedule.EventDuration * int(time.Minute))

	startTime := schedule.StartTime.Format(layout)[11:]
	endTime := schedule.EndTime.Format(layout)[11:]
	startBreak := schedule.StartBreak.Format(layout)[11:]
	endBreak := schedule.EndBreak.Format(layout)[11:]

	for first.Before(last) {
		date := first.Format(layout)[:11]
		starts, err := time.Parse(layout, date+startTime)
		if err != nil {
			return fmt.Errorf("ClentService - CreateSchedule - Parse #1: %w", err)
		}

		ends, err := time.Parse(layout, date+endTime)
		if err != nil {
			return fmt.Errorf("ClentService - CreateSchedule - Parse #2: %w", err)
		}

		startsBreak, err := time.Parse(layout, date+startBreak)
		if err != nil {
			return fmt.Errorf("ClentService - CreateSchedule - Parse #3: %w", err)
		}

		endsBreak, err := time.Parse(layout, date+endBreak)
		if err != nil {
			return fmt.Errorf("ClentService - CreateSchedule - Parse #4: %w", err)
		}

		was := true
		for i := starts; i.Before(ends); i = i.Add(increase) {
			if was {
				if i.Equal(startsBreak) || i.After(startsBreak) {
					i = endsBreak
					was = false
				}
			}
			event := entity.Event{}
			event.StartsAt = i
			event.EndsAt = i.Add(increase)
			dayEvents = append(dayEvents, event)
		}

		first = first.AddDate(0, 0, 1)
	}

	err := cs.repo.StoreSchedule(ctx, dayEvents)
	if err != nil {
		return fmt.Errorf("ClentService - CreateSchedule: %w", err)
	}

	return nil
}
