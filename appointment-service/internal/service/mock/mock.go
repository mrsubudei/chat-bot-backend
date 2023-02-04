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
	ret := m.ctrl.Call(m, "CreateDoctor", ctx, doctor)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEventsRecorder) CreateDoctor(ctx, doctor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDoctor",
		reflect.TypeOf((*MockEvents)(nil).CreateDoctor), ctx, doctor)
}

func (m *MockEvents) GetDoctor(ctx context.Context, doctorId int32) (entity.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDoctor", ctx, doctorId)
	ret0, _ := ret[0].(entity.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventsRecorder) GetDoctor(ctx, doctorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoctor",
		reflect.TypeOf((*MockEvents)(nil).GetDoctor), ctx, doctorId)
}

func (m *MockEvents) UpdateDoctor(ctx context.Context, doctor entity.Doctor) (entity.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDoctor", ctx, doctor)
	ret0, _ := ret[0].(entity.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventsRecorder) UpdateDoctor(ctx, doctor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDoctor",
		reflect.TypeOf((*MockEvents)(nil).UpdateDoctor), ctx, doctor)
}

func (m *MockEvents) DeleteDoctor(ctx context.Context, doctorId int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDoctor", ctx, doctorId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEventsRecorder) DeleteDoctor(ctx, doctorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDoctor",
		reflect.TypeOf((*MockEvents)(nil).DeleteDoctor), ctx, doctorId)
}

func (m *MockEvents) GetAllDoctors(ctx context.Context) ([]entity.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDoctors", ctx)
	ret0, _ := ret[0].([]entity.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventsRecorder) GetAllDoctors(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDoctors",
		reflect.TypeOf((*MockEvents)(nil).GetAllDoctors), ctx)
}

func (m *MockEvents) CreateSchedule(ctx context.Context, schedule entity.Schedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchedule", ctx, schedule)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEventsRecorder) CreateSchedule(ctx, schedule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchedule",
		reflect.TypeOf((*MockEvents)(nil).CreateSchedule), ctx, schedule)
}

func (m *MockEvents) GetOpenEventsByDoctor(ctx context.Context, doctorId int32) ([]entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenEventsByDoctor", ctx, doctorId)
	ret0, _ := ret[0].([]entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventsRecorder) GetOpenEventsByDoctor(ctx, doctorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenEventsByDoctor",
		reflect.TypeOf((*MockEvents)(nil).GetOpenEventsByDoctor), ctx, doctorId)
}

func (m *MockEvents) GetReservedEventsByDoctor(ctx context.Context, doctorId int32) ([]entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservedEventsByDoctor", ctx, doctorId)
	ret0, _ := ret[0].([]entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventsRecorder) GetReservedEventsByDoctor(ctx, doctorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservedEventsByDoctor",
		reflect.TypeOf((*MockEvents)(nil).GetReservedEventsByDoctor), ctx, doctorId)
}

func (m *MockEvents) GetReservedEventsByClient(ctx context.Context, clientId int32) ([]entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservedEventsByClient", ctx, clientId)
	ret0, _ := ret[0].([]entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventsRecorder) GetReservedEventsByClient(ctx, clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservedEventsByClient",
		reflect.TypeOf((*MockEvents)(nil).GetReservedEventsByClient), ctx, clientId)
}

func (m *MockEvents) GetAllEventsByClient(ctx context.Context, clientId int32) ([]entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEventsByClient", ctx, clientId)
	ret0, _ := ret[0].([]entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventsRecorder) GetAllEventsByClient(ctx, clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEventsByClient",
		reflect.TypeOf((*MockEvents)(nil).GetAllEventsByClient), ctx, clientId)
}

func (m *MockEvents) RegisterToEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterToEvent", ctx, event)
	ret0, _ := ret[0].(entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventsRecorder) RegisterToEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterToEvent",
		reflect.TypeOf((*MockEvents)(nil).RegisterToEvent), ctx, event)
}

func (m *MockEvents) UnregisterEvent(ctx context.Context, event entity.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterEvent", ctx, event)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockEventsRecorder) UnregisterEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterEvent",
		reflect.TypeOf((*MockEvents)(nil).UnregisterEvent), ctx, event)
}
