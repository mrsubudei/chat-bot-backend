package repository

import (
	"context"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
)

type Clients interface {
	StoreDoctor(ctx context.Context, doctor entity.Doctor) error
	DeleteDoctor(ctx context.Context, doctor entity.Doctor) error
	FetchDoctors(ctx context.Context, doctor entity.Doctor) ([]entity.Doctor, error)
	StoreClient(ctx context.Context, client entity.Client) error
	StoreSchedule(ctx context.Context, events []entity.Event) error
	FetchOpenEventsByDoctor(ctx context.Context, doctor entity.Doctor) ([]entity.Event, error)
	FetchReservedEventsByDoctor(ctx context.Context, doctor entity.Doctor) ([]entity.Event, error)
	FetchReservedEventsByClient(ctx context.Context, client entity.Client) ([]entity.Event, error)
	FetchAllEventsByClient(ctx context.Context, client entity.Client) ([]entity.Event, error)
	UpdateEvent(ctx context.Context, event entity.Event) error
}
