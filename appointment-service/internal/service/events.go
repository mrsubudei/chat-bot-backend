package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository"
)

type EventsService struct {
	repo repository.Events
}

func NewEventsService(repo repository.Events) *EventsService {
	return &EventsService{repo}
}

func (es *EventsService) CreateDoctor(ctx context.Context, doctor entity.Doctor) error {
	err := es.repo.StoreDoctor(ctx, doctor)
	if err != nil {
		if strings.Contains(err.Error(), DuplicateErrMsg) {
			return entity.ErrEntityAlreadyExists
		}
		return fmt.Errorf("EventsService - CreateDoctor - StoreDoctor: %w", err)
	}
	return nil
}

func (es *EventsService) GetDoctor(ctx context.Context, doctorId int32) (entity.Doctor, error) {
	doctor, err := es.repo.GetDoctor(ctx, doctorId)
	if err != nil {
		if errors.Is(err, entity.ErrEntityDoesNotExist) {
			return doctor, err
		}
		return doctor, fmt.Errorf("EventsService - GetDoctor: %w", err)
	}
	return doctor, nil
}

func (es *EventsService) UpdateDoctor(ctx context.Context,
	doctor entity.Doctor) (entity.Doctor, error) {

	updated, err := es.repo.UpdateDoctor(ctx, doctor)
	if err != nil {
		if errors.Is(err, entity.ErrEntityDoesNotExist) {
			return doctor, err
		}
		return doctor, fmt.Errorf("EventsService - UpdateDoctor - GetDoctor: %w", err)
	}

	return updated, nil
}

func (es *EventsService) DeleteDoctor(ctx context.Context, doctorId int32) error {
	err := es.repo.DeleteDoctor(ctx, doctorId)
	if err != nil {
		if strings.Contains(err.Error(), NoRowsAffected) {
			return entity.ErrEntityDoesNotExist
		}
		return fmt.Errorf("EventsService - DeleteDoctor - DeleteDoctor: %w", err)
	}
	return nil
}

func (es *EventsService) GetAllDoctors(ctx context.Context) ([]entity.Doctor, error) {
	doctors, err := es.repo.FetchDoctors(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventsService - GetAllDoctors - FetchDoctors: %w", err)
	}

	return doctors, nil
}

func (es *EventsService) CreateSchedule(ctx context.Context, schedule entity.Schedule) error {
	dayEvents := []entity.Event{}
	first := schedule.FirstDay
	last := schedule.LastDay.AddDate(0, 0, 1)
	increase := time.Duration(int(schedule.EventDuration) * int(time.Minute))

	startTime := schedule.StartTime.Format(DateFormat)[11:]
	endTime := schedule.EndTime.Format(DateFormat)[11:]
	startBreak := schedule.StartBreak.Format(DateFormat)[11:]
	endBreak := schedule.EndBreak.Format(DateFormat)[11:]

	for first.Before(last) {
		date := first.Format(DateFormat)[:11]
		starts, err := time.Parse(DateFormat, date+startTime)
		if err != nil {
			return fmt.Errorf("EventsService - CreateSchedule - Parse #1: %w", err)
		}

		ends, err := time.Parse(DateFormat, date+endTime)
		if err != nil {
			return fmt.Errorf("EventsService - CreateSchedule - Parse #2: %w", err)
		}

		startsBreak, err := time.Parse(DateFormat, date+startBreak)
		if err != nil {
			return fmt.Errorf("EventsService - CreateSchedule - Parse #3: %w", err)
		}

		endsBreak, err := time.Parse(DateFormat, date+endBreak)
		if err != nil {
			return fmt.Errorf("EventsService - CreateSchedule - Parse #4: %w", err)
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
			for u := 0; u < len(schedule.DoctorIds); u++ {
				event.DoctorId = schedule.DoctorIds[u]
				dayEvents = append(dayEvents, event)
			}
		}

		first = first.AddDate(0, 0, 1)
	}

	err := es.repo.StoreSchedule(ctx, dayEvents)
	if err != nil {
		return fmt.Errorf("EventsService - CreateSchedule - StoreSchedule: %w", err)
	}

	return nil
}

func (es *EventsService) GetOpenEventsByDoctor(ctx context.Context,
	doctorId int32) ([]entity.Event, error) {

	events, err := es.repo.FetchOpenEventsByDoctor(ctx, doctorId)
	if err != nil {
		return nil, fmt.Errorf("EventsService - GetOpenEventsByDoctor: %w", err)
	}

	return events, nil
}
