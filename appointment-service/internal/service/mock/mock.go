package mock_service

import (
	"context"
	"reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
)

type MockEvents struct {
	ctrl     *gomock.Controller
	recorder *MockEventsRecorder
}

type MockEventsRecorder struct {
	mock *MockEvents
}

func NewMockEvents(ctrl *gomock.Controller) *MockEvents {
	mock := &MockEvents{ctrl: ctrl}
	mock.recorder = &MockEventsRecorder{mock}
	return mock
}

func (m *MockEvents) EXPECT() *MockEventsRecorder {
	return m.recorder
}

func (m *MockEvents) CreateDoctor(ctx context.Context, doctor entity.Doctor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDoctor", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEventsRecorder) CreateDoctor(ctx, doctor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDoctor",
		reflect.TypeOf((*MockEvents)(nil).CreateDoctor), ctx, doctor)
}
