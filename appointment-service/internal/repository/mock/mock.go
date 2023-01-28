package mock_repository

import (
	"context"
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
)

type MockEventsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockEventsRepoRecorder
}

type MockEventsRepoRecorder struct {
	mock *MockEventsRepo
}

func NewMockEventsRepo(ctrl *gomock.Controller) *MockEventsRepo {
	mock := &MockEventsRepo{ctrl: ctrl}
	mock.recorder = &MockEventsRepoRecorder{mock}
	return mock
}

func (m *MockEventsRepo) EXPECT() *MockEventsRepoRecorder {
	return m.recorder
}

func (m *MockEventsRepo) StoreDoctor(ctx context.Context, doctor entity.Doctor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreDoctor", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEventsRepoRecorder) StoreDoctor(ctx, doctor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreDoctor", reflect.TypeOf((*MockEventsRepo)(nil).StoreDoctor), ctx)
}

func (m *MockEventsRepo) GetDoctor(ctx context.Context, doctorId int32) (entity.Doctor, error) {
	panic(ctx)
}

func (m *MockEventsRepo) UpdateDoctor(ctx context.Context, doctor entity.Doctor) (entity.Doctor, error) {
	panic(ctx)
}

func (m *MockEventsRepo) DeleteDoctor(ctx context.Context, id int32) error {
	panic(ctx)
}

func (m *MockEventsRepo) FetchDoctors(ctx context.Context) ([]entity.Doctor, error) {
	panic(ctx)
}

func (m *MockEventsRepo) StoreSchedule(ctx context.Context, events []entity.Event) (time.Time, error) {
	panic(ctx)
}

func (m *MockEventsRepo) FetchOpenEventsByDoctor(ctx context.Context, doctorId int32) ([]entity.Event, error) {
	panic(ctx)
}

func (m *MockEventsRepo) FetchReservedEventsByDoctor(ctx context.Context, doctorId int32) ([]entity.Event, error) {
	panic(ctx)
}

func (m *MockEventsRepo) FetchReservedEventsByClient(ctx context.Context, clientId int32) ([]entity.Event, error) {
	panic(ctx)
}

func (m *MockEventsRepo) FetchAllEventsByClient(ctx context.Context, clientId int32) ([]entity.Event, error) {
	panic(ctx)
}

func (m *MockEventsRepo) UpdateEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
	panic(ctx)
}

func (m *MockEventsRepo) ClearEvent(ctx context.Context, event entity.Event) error {
	panic(ctx)
}
