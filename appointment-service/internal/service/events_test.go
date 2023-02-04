package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
	mock_repository "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/mock"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/service"
	"github.com/stretchr/testify/require"
)

var errInternalServErr = errors.New("test: internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func mockEventService(t *testing.T) (*service.EventsService, *mock_repository.MockEventsRepo) {
	t.Helper()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := mock_repository.NewMockEventsRepo(mockCtl)
	eventsService := service.NewEventsService(repo)

	return eventsService, repo
}

func TestCreateDoctor(t *testing.T) {
	t.Parallel()

	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().StoreDoctor(ctx, gomock.Any()).Return(nil)
			},
			err: nil,
		},
		{
			name: "Error doctor exists",
			mock: func() {
				repo.EXPECT().StoreDoctor(ctx, gomock.Any()).Return(errors.New(service.DuplicateErrMsg))
			},
			err: entity.ErrDoctorAlreadyExists,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := eventsService.CreateDoctor(ctx, entity.Doctor{})
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestGetDoctor(t *testing.T) {
	t.Parallel()

	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().GetDoctor(ctx, gomock.Any()).Return(entity.Doctor{}, nil)
			},
			res: entity.Doctor{},
			err: nil,
		},
		{
			name: "Error doctor does not exist",
			mock: func() {
				repo.EXPECT().GetDoctor(ctx, gomock.Any()).Return(entity.Doctor{},
					entity.ErrDoctorDoesNotExist)
			},
			res: entity.Doctor{},
			err: entity.ErrDoctorDoesNotExist,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := eventsService.GetDoctor(ctx, int32(1))
			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestUpdateDoctor(t *testing.T) {
	t.Parallel()

	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().UpdateDoctor(ctx, gomock.Any()).Return(entity.Doctor{}, nil)
			},
			res: entity.Doctor{},
			err: nil,
		},
		{
			name: "Error doctor does not exist",
			mock: func() {
				repo.EXPECT().UpdateDoctor(ctx, gomock.Any()).Return(entity.Doctor{},
					entity.ErrDoctorDoesNotExist)
			},
			res: entity.Doctor{},
			err: entity.ErrDoctorDoesNotExist,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := eventsService.UpdateDoctor(ctx, entity.Doctor{})
			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestDeleteDoctor(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().DeleteDoctor(ctx, gomock.Any()).Return(nil)
			},
			err: nil,
		},
		{
			name: "Error event does not exist",
			mock: func() {
				repo.EXPECT().DeleteDoctor(ctx, gomock.Any()).
					Return(errors.New(service.NoRowsAffected))
			},
			err: entity.ErrEventDoesNotExist,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := eventsService.DeleteDoctor(ctx, int32(1))
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestGetAllDoctors(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	repo.EXPECT().FetchDoctors(ctx).Return([]entity.Doctor(nil), nil)
	res, err := eventsService.GetAllDoctors(ctx)
	require.Equal(t, res, []entity.Doctor(nil))
	require.ErrorIs(t, err, nil)
}

func TestCreateSchedule(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{}, nil)
			},
			res: time.Time{},
			err: nil,
		},
		{
			name: "Error event already exists",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{},
					entity.ErrEventAlreadyExists)
			},
			res: time.Time{},
			err: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := eventsService.CreateSchedule(ctx, entity.Schedule{})
			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestGetOpenEventsByDoctor(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	repo.EXPECT().FetchOpenEventsByDoctor(ctx, gomock.Any()).
		Return([]entity.Event(nil), nil)
	res, err := eventsService.GetOpenEventsByDoctor(ctx, int32(2))

	require.Equal(t, res, []entity.Event(nil))
	require.ErrorIs(t, err, nil)
}

func TestGetReservedEventsByDoctor(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	repo.EXPECT().FetchReservedEventsByDoctor(ctx, gomock.Any()).
		Return([]entity.Event(nil), nil)
	res, err := eventsService.GetReservedEventsByDoctor(ctx, int32(2))

	require.Equal(t, res, []entity.Event(nil))
	require.ErrorIs(t, err, nil)
}

func TestGetReservedEventsByClient(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	repo.EXPECT().FetchReservedEventsByClient(ctx, gomock.Any()).
		Return([]entity.Event(nil), nil)
	res, err := eventsService.GetReservedEventsByClient(ctx, int32(2))

	require.Equal(t, res, []entity.Event(nil))
	require.ErrorIs(t, err, nil)
}

func TestGetAllEventsByClient(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	repo.EXPECT().FetchAllEventsByClient(ctx, gomock.Any()).
		Return([]entity.Event(nil), nil)
	res, err := eventsService.GetAllEventsByClient(ctx, int32(2))

	require.Equal(t, res, []entity.Event(nil))
	require.ErrorIs(t, err, nil)
}

func TestRegisterToEvent(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().UpdateEvent(ctx, gomock.Any()).Return(entity.Event{}, nil)
			},
			res: entity.Event{},
			err: nil,
		},
		{
			name: "Error entity does not exist",
			mock: func() {
				repo.EXPECT().UpdateEvent(ctx, gomock.Any()).Return(entity.Event{}, entity.ErrEventDoesNotExist)
			},
			res: entity.Event{},
			err: entity.ErrEventDoesNotExist,
		},
		{
			name: "Error entity already reserved",
			mock: func() {
				repo.EXPECT().UpdateEvent(ctx, gomock.Any()).Return(entity.Event{}, entity.ErrEventAlreadyReserved)
			},
			res: entity.Event{},
			err: entity.ErrEventAlreadyReserved,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := eventsService.RegisterToEvent(ctx, entity.Event{})
			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestUnregisterEvent(t *testing.T) {
	eventsService, repo := mockEventService(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().ClearEvent(ctx, gomock.Any()).Return(nil)
			},
			err: nil,
		},
		{
			name: "Error entity does not exist",
			mock: func() {
				repo.EXPECT().ClearEvent(ctx, gomock.Any()).Return(entity.ErrEventDoesNotExist)
			},
			err: entity.ErrEventDoesNotExist,
		},
		{
			name: "Error entity is not reserved",
			mock: func() {
				repo.EXPECT().ClearEvent(ctx, gomock.Any()).Return(entity.ErrEventIsNotReserved)
			},
			err: entity.ErrEventIsNotReserved,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := eventsService.UnregisterEvent(ctx, entity.Event{})
			require.ErrorIs(t, err, tc.err)
		})
	}
}
