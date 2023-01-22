package service

import (
	"context"

	"github.com/mrsubudei/chat-bot-backend/internal/entity"
)

type Service interface {
	CreateSchedule(ctx context.Context, schedule entity.Schedule) error
}
