package repository

import (
	"context"

	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/entity"
)

type Users interface {
	Store(ctx context.Context, user entity.User) (int32, error)
	Delete(ctx context.Context, userId int32) error
	GetByPhone(ctx context.Context, phone string) (entity.User, error)
	GetById(ctx context.Context, id int32) (entity.User, error)
	GetByToken(ctx context.Context, token string) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateSession(ctx context.Context, user entity.User) error
	UpdateVerification(ctx context.Context, user entity.User) error
}
