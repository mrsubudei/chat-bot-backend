package api_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/api"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/config"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
	mock_repository "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/mock"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
	pb "github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type test struct {
	name     string
	mock     func()
	res      interface{}
	err      error
	schedule *pb.ScheduleSingle
}

func mockEventApi(t *testing.T) (*api.AppointmentServer, *mock_repository.MockEventsRepo) {
	t.Helper()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := mock_repository.NewMockEventsRepo(mockCtl)

	cfg, err := config.NewConfig("../../config.yml")
	if err != nil {
		t.Fatalf("Config error: %s", err)
	}
	l := logger.New(cfg.Logger.Level)
	eventsService := api.NewAppointmentServer(repo, l)

	return eventsService, repo
}

func TestCreateDoctor(t *testing.T) {
	t.Parallel()
	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().StoreDoctor(ctx, gomock.Any()).Return(nil)
			},
			res: &pb.Empty{},
			err: nil,
		},
		{
			name: "Error already exists",
			mock: func() {
				repo.EXPECT().StoreDoctor(ctx, gomock.Any()).
					Return(errors.New(api.DuplicateErrMsg))
			},
			res: &pb.Empty{},
			err: status.Error(codes.AlreadyExists,
				entity.ErrDoctorAlreadyExists.Error()),
		},
		{
			name: "Error internal",
			mock: func() {
				repo.EXPECT().StoreDoctor(ctx, gomock.Any()).
					Return(errors.New("internal error"))
			},
			res: &pb.Empty{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.CreateDoctor(ctx, &pb.DoctorSingle{
				Value: &pb.Doctor{}})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestGetDoctor(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().GetDoctor(ctx, gomock.Any()).
					Return(entity.Doctor{}, nil)
			},
			res: &pb.DoctorSingle{
				Value: &pb.Doctor{},
			},
			err: nil,
		},
		{
			name: "Error does not exist",
			mock: func() {
				repo.EXPECT().GetDoctor(ctx, gomock.Any()).
					Return(entity.Doctor{}, entity.ErrDoctorDoesNotExist)
			},
			res: &pb.DoctorSingle{},
			err: status.Error(codes.NotFound,
				entity.ErrDoctorDoesNotExist.Error()),
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().GetDoctor(ctx, gomock.Any()).
					Return(entity.Doctor{}, errors.New("internal error"))
			},
			res: &pb.DoctorSingle{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.GetDoctor(ctx, &pb.IdRequest{Id: 1})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestUpdateDoctor(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().UpdateDoctor(ctx, gomock.Any()).
					Return(entity.Doctor{}, nil)
			},
			res: &pb.DoctorSingle{
				Value: &pb.Doctor{},
			},
			err: nil,
		},
		{
			name: "Error does not exist",
			mock: func() {
				repo.EXPECT().UpdateDoctor(ctx, gomock.Any()).
					Return(entity.Doctor{}, entity.ErrDoctorDoesNotExist)
			},
			res: &pb.DoctorSingle{},
			err: status.Error(codes.NotFound,
				entity.ErrDoctorDoesNotExist.Error()),
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().UpdateDoctor(ctx, gomock.Any()).
					Return(entity.Doctor{}, errors.New("internal error"))
			},
			res: &pb.DoctorSingle{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.UpdateDoctor(ctx, &pb.DoctorSingle{
				Value: &pb.Doctor{},
			})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestDeleteDoctor(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().DeleteDoctor(ctx, gomock.Any()).Return(nil)
			},
			res: &pb.Empty{},
			err: nil,
		},
		{
			name: "Error does not exist",
			mock: func() {
				repo.EXPECT().DeleteDoctor(ctx, gomock.Any()).
					Return(errors.New(api.NoRowsAffected))
			},
			res: &pb.Empty{},
			err: status.Error(codes.NotFound,
				entity.ErrDoctorDoesNotExist.Error()),
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().DeleteDoctor(ctx, gomock.Any()).
					Return(errors.New("internal error"))
			},
			res: &pb.Empty{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.DeleteDoctor(ctx, &pb.IdRequest{})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestGetAllDoctors(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().FetchDoctors(ctx).Return([]entity.Doctor{
					{Id: 1},
					{Id: 2},
				}, nil)
			},
			res: &pb.DoctorMultiple{
				Value: []*pb.Doctor{
					{Id: 1},
					{Id: 2},
				},
			},
			err: nil,
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().FetchDoctors(ctx).
					Return([]entity.Doctor{}, errors.New("internal error"))
			},
			res: &pb.DoctorMultiple{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.GetAllDoctors(ctx, &pb.Empty{})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestCreateSchedule(t *testing.T) {

	apiService, repo := mockEventApi(t)
	ctx := context.Background()
	now := time.Now()
	zeroTime, err := time.Parse(api.DateAndTimeFormat, "2023-01-28 00:00:00")
	if err != nil {
		t.Fatal(err)
	}

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{}, nil)
			},
			schedule: &pb.ScheduleSingle{
				Value: &pb.Schedule{
					FirstDay:             timestamppb.New(now),
					LastDay:              timestamppb.New(now.Add(time.Hour * 24)),
					StartTime:            timestamppb.New(now),
					EndTime:              timestamppb.New(now.Add(time.Hour * 8)),
					StartBreak:           timestamppb.New(now.Add(time.Hour * 4)),
					EndBreak:             timestamppb.New(now.Add(time.Hour * 5)),
					DoctorIds:            []int32{4, 5},
					EventDurationMinutes: 30,
				},
			},
			res: &pb.Empty{},
			err: nil,
		},
		{
			name: "Error empty FirstDay",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{}, nil)
			},
			schedule: &pb.ScheduleSingle{
				Value: &pb.Schedule{
					FirstDay:             timestamppb.New(time.Time{}),
					LastDay:              timestamppb.New(now.Add(time.Hour * 24)),
					StartTime:            timestamppb.New(now),
					EndTime:              timestamppb.New(now.Add(time.Hour * 8)),
					StartBreak:           timestamppb.New(now.Add(time.Hour * 4)),
					EndBreak:             timestamppb.New(now.Add(time.Hour * 5)),
					DoctorIds:            []int32{4, 5},
					EventDurationMinutes: 30,
				},
			},
			res: &pb.Empty{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": FirstDay"),
		},
		{
			name: "Error empty LastDay",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{}, nil)
			},
			schedule: &pb.ScheduleSingle{
				Value: &pb.Schedule{
					FirstDay:             timestamppb.New(now),
					LastDay:              timestamppb.New(time.Time{}),
					StartTime:            timestamppb.New(now),
					EndTime:              timestamppb.New(now.Add(time.Hour * 8)),
					StartBreak:           timestamppb.New(now.Add(time.Hour * 4)),
					EndBreak:             timestamppb.New(now.Add(time.Hour * 5)),
					DoctorIds:            []int32{4, 5},
					EventDurationMinutes: 30,
				},
			},
			res: &pb.Empty{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": LastDay"),
		},
		{
			name: "Error StartTime zero value",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{}, nil)
			},
			schedule: &pb.ScheduleSingle{
				Value: &pb.Schedule{
					FirstDay:             timestamppb.New(now),
					LastDay:              timestamppb.New(now.Add(time.Hour * 24)),
					StartTime:            timestamppb.New(zeroTime),
					EndTime:              timestamppb.New(now.Add(time.Hour * 8)),
					StartBreak:           timestamppb.New(now.Add(time.Hour * 4)),
					EndBreak:             timestamppb.New(now.Add(time.Hour * 5)),
					DoctorIds:            []int32{4, 5},
					EventDurationMinutes: 30,
				},
			},
			res: &pb.Empty{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": StartTime"),
		},
		{
			name: "Error EndTime zero value",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{}, nil)
			},
			schedule: &pb.ScheduleSingle{
				Value: &pb.Schedule{
					FirstDay:             timestamppb.New(now),
					LastDay:              timestamppb.New(now.Add(time.Hour * 24)),
					StartTime:            timestamppb.New(now),
					EndTime:              timestamppb.New(zeroTime),
					StartBreak:           timestamppb.New(now.Add(time.Hour * 4)),
					EndBreak:             timestamppb.New(now.Add(time.Hour * 5)),
					DoctorIds:            []int32{4, 5},
					EventDurationMinutes: 30,
				},
			},
			res: &pb.Empty{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": EndTime"),
		},
		{
			name: "Error StartBreak zero value",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{}, nil)
			},
			schedule: &pb.ScheduleSingle{
				Value: &pb.Schedule{
					FirstDay:             timestamppb.New(now),
					LastDay:              timestamppb.New(now.Add(time.Hour * 24)),
					StartTime:            timestamppb.New(now),
					EndTime:              timestamppb.New(now.Add(time.Hour * 8)),
					StartBreak:           timestamppb.New(zeroTime),
					EndBreak:             timestamppb.New(now.Add(time.Hour * 5)),
					DoctorIds:            []int32{4, 5},
					EventDurationMinutes: 30,
				},
			},
			res: &pb.Empty{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": StartBreak"),
		},
		{
			name: "Error EndBreak zero value",
			mock: func() {
				repo.EXPECT().StoreSchedule(ctx, gomock.Any()).Return(time.Time{}, nil)
			},
			schedule: &pb.ScheduleSingle{
				Value: &pb.Schedule{
					FirstDay:             timestamppb.New(now),
					LastDay:              timestamppb.New(now.Add(time.Hour * 24)),
					StartTime:            timestamppb.New(now),
					EndTime:              timestamppb.New(now.Add(time.Hour * 8)),
					StartBreak:           timestamppb.New(now.Add(time.Hour * 4)),
					EndBreak:             timestamppb.New(zeroTime),
					DoctorIds:            []int32{4, 5},
					EventDurationMinutes: 30,
				},
			},
			res: &pb.Empty{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": EndBreak"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.CreateSchedule(ctx, tc.schedule)
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestGetOpenEventsByDoctor(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().FetchOpenEventsByDoctor(ctx, gomock.Any()).
					Return([]entity.Event(nil), nil)
			},
			res: &pb.EventMultiple{
				Value: []*pb.Event{},
			},
			err: nil,
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().FetchOpenEventsByDoctor(ctx, gomock.Any()).
					Return([]entity.Event(nil), errors.New("Internal error"))
			},
			res: &pb.EventMultiple{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.GetOpenEventsByDoctor(ctx, &pb.IdRequest{Id: 2})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestGetReservedEventsByDoctor(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().FetchReservedEventsByDoctor(ctx, gomock.Any()).
					Return([]entity.Event(nil), nil)
			},
			res: &pb.EventMultiple{
				Value: []*pb.Event{},
			},
			err: nil,
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().FetchReservedEventsByDoctor(ctx, gomock.Any()).
					Return([]entity.Event(nil), errors.New("Internal error"))
			},
			res: &pb.EventMultiple{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.GetReservedEventsByDoctor(ctx, &pb.IdRequest{Id: 2})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestGetReservedEventsByClient(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().FetchReservedEventsByClient(ctx, gomock.Any()).
					Return([]entity.Event(nil), nil)
			},
			res: &pb.EventMultiple{
				Value: []*pb.Event{},
			},
			err: nil,
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().FetchReservedEventsByClient(ctx, gomock.Any()).
					Return([]entity.Event(nil), errors.New("Internal error"))
			},
			res: &pb.EventMultiple{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.GetReservedEventsByClient(ctx, &pb.IdRequest{Id: 2})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestGetAllEventsByClient(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().FetchAllEventsByClient(ctx, gomock.Any()).
					Return([]entity.Event(nil), nil)
			},
			res: &pb.EventMultiple{
				Value: []*pb.Event{},
			},
			err: nil,
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().FetchAllEventsByClient(ctx, gomock.Any()).
					Return([]entity.Event(nil), errors.New("Internal error"))
			},
			res: &pb.EventMultiple{},
			err: status.Error(codes.Internal, api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.GetAllEventsByClient(ctx, &pb.IdRequest{Id: 2})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}

func TestRegisterToEvent(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().UpdateEvent(ctx, gomock.Any()).
					Return(entity.Event{}, nil)
			},
			res: &pb.EventSingle{
				Value: &pb.Event{
					StartsAt: timestamppb.New(time.Time{}),
					EndsAt:   timestamppb.New(time.Time{}),
				},
			},
			err: nil,
		},
		{
			name: "Error does not exist",
			mock: func() {
				repo.EXPECT().UpdateEvent(ctx, gomock.Any()).
					Return(entity.Event{}, entity.ErrEventDoesNotExist)
			},
			res: &pb.EventSingle{},
			err: status.Error(codes.NotFound,
				entity.ErrEventDoesNotExist.Error()),
		},
		{
			name: "Error already reserved",
			mock: func() {
				repo.EXPECT().UpdateEvent(ctx, gomock.Any()).
					Return(entity.Event{}, entity.ErrEventAlreadyReserved)
			},
			res: &pb.EventSingle{},
			err: status.Error(codes.Canceled,
				entity.ErrEventAlreadyReserved.Error()),
		},
		{
			name: "Internal error",
			mock: func() {
				repo.EXPECT().UpdateEvent(ctx, gomock.Any()).
					Return(entity.Event{}, errors.New("Internal error"))
			},
			res: &pb.EventSingle{},
			err: status.Error(codes.Internal,
				api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.RegisterToEvent(ctx, &pb.EventSingle{
				Value: &pb.Event{},
			})
			require.ErrorIs(t, err, tc.err)
			require.True(t, proto.Equal(res, tc.res.(*pb.EventSingle)),
				"expected: %v\nactual: %v\n", tc.res, res)
		})
	}
}

func TestUnregisterEvent(t *testing.T) {
	t.Parallel()

	apiService, repo := mockEventApi(t)
	ctx := context.Background()

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().ClearEvent(ctx, gomock.Any()).
					Return(nil)
			},
			res: &pb.Empty{},
			err: nil,
		},
		{
			name: "Error does not exist",
			mock: func() {
				repo.EXPECT().ClearEvent(ctx, gomock.Any()).
					Return(entity.ErrEventDoesNotExist)
			},
			res: &pb.Empty{},
			err: status.Error(codes.NotFound,
				entity.ErrEventDoesNotExist.Error()),
		},
		{
			name: "Error is not reserved",
			mock: func() {
				repo.EXPECT().ClearEvent(ctx, gomock.Any()).
					Return(entity.ErrEventIsNotReserved)
			},
			res: &pb.Empty{},
			err: status.Error(codes.Canceled,
				entity.ErrEventIsNotReserved.Error()),
		},
		{
			name: "Error is not reserved",
			mock: func() {
				repo.EXPECT().ClearEvent(ctx, gomock.Any()).
					Return(errors.New("Internal error"))
			},
			res: &pb.Empty{},
			err: status.Error(codes.Internal,
				api.InternalErr),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.UnregisterEvent(ctx, &pb.EventSingle{
				Value: &pb.Event{},
			})
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}
