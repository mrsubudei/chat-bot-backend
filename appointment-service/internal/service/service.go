package service

import (
	"context"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
)

type Service interface {
	CreateDoctor(ctx context.Context, doctor entity.Doctor) error
	GetDoctor(ctx context.Context, doctorId int32) (entity.Doctor, error)
	UpdateDoctor(ctx context.Context, doctor entity.Doctor) (entity.Doctor, error)
	DeleteDoctor(ctx context.Context, doctorId int32) error
	GetAllDoctors(ctx context.Context) ([]entity.Doctor, error)
	CreateSchedule(ctx context.Context, schedule entity.Schedule) error
}
