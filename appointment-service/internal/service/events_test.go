package service_test

import (
	"context"
	"errors"
	"testing"

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

	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().StoreDoctor(context.Background(), gomock.Any()).Return(nil)
			},
			res: nil,
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()
			err := eventsService.CreateDoctor(context.Background(), entity.Doctor{})
			require.ErrorIs(t, err, tc.err)
		})
	}
}
