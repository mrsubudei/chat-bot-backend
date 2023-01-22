package repository

import (
	"context"

	"github.com/mrsubudei/chat-bot-backend/internal/entity"
)

type Clients interface {
	StoreSchedule(ctx context.Context, events []entity.Event) error
}
