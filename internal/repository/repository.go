package repository

import (
	"context"

	"github.com/mrsubudei/chat-bot-backend/internal/entity"
)

type Clients interface {
	CreateSchedule(ctx context.Context, schedule entity.Schedule) error
}
