package repository

// import (
// 	"context"
// 	"time"

// 	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
// )

// type Events interface {
// 	StoreDoctor(ctx context.Context, doctor entity.Doctor) error
// 	GetDoctor(ctx context.Context, doctorId int32) (entity.Doctor, error)
// 	UpdateDoctor(ctx context.Context, doctor entity.Doctor) (entity.Doctor, error)
// 	DeleteDoctor(ctx context.Context, id int32) error
// 	FetchDoctors(ctx context.Context) ([]entity.Doctor, error)
// 	StoreSchedule(ctx context.Context, events []entity.Event) (time.Time, error)
// 	FetchOpenEventsByDoctor(ctx context.Context, doctorId int32) ([]entity.Event, error)
// 	FetchReservedEventsByDoctor(ctx context.Context, doctorId int32) ([]entity.Event, error)
// 	FetchReservedEventsByClient(ctx context.Context, clientId int32) ([]entity.Event, error)
// 	FetchAllEventsByClient(ctx context.Context, clientId int32) ([]entity.Event, error)
// 	UpdateEvent(ctx context.Context, event entity.Event) (entity.Event, error)
// 	ClearEvent(ctx context.Context, event entity.Event) error
// }
