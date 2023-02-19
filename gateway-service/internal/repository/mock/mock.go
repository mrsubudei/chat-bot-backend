package mock_repository

// import (
// 	"context"
// 	"reflect"
// 	"time"

// 	"github.com/golang/mock/gomock"
// 	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
// )

// type MockEventsRepo struct {
// 	ctrl     *gomock.Controller
// 	recorder *MockEventsRepoRecorder
// }

// type MockEventsRepoRecorder struct {
// 	mock *MockEventsRepo
// }

// func NewMockEventsRepo(ctrl *gomock.Controller) *MockEventsRepo {
// 	mock := &MockEventsRepo{ctrl: ctrl}
// 	mock.recorder = &MockEventsRepoRecorder{mock}
// 	return mock
// }

// func (m *MockEventsRepo) EXPECT() *MockEventsRepoRecorder {
// 	return m.recorder
// }

// func (m *MockEventsRepo) StoreDoctor(ctx context.Context, doctor entity.Doctor) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "StoreDoctor", ctx, doctor)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// func (mr *MockEventsRepoRecorder) StoreDoctor(ctx, doctor interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreDoctor",
// 		reflect.TypeOf((*MockEventsRepo)(nil).StoreDoctor), ctx, doctor)
// }

// func (m *MockEventsRepo) GetDoctor(ctx context.Context, doctorId int32) (entity.Doctor, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "GetDoctor", ctx, doctorId)
// 	ret0, _ := ret[0].(entity.Doctor)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) GetDoctor(ctx, doctorId interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoctor",
// 		reflect.TypeOf((*MockEventsRepo)(nil).GetDoctor), ctx, doctorId)
// }

// func (m *MockEventsRepo) UpdateDoctor(ctx context.Context, doctor entity.Doctor) (entity.Doctor, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateDoctor", ctx, doctor)
// 	ret0, _ := ret[0].(entity.Doctor)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) UpdateDoctor(ctx, doctor interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDoctor",
// 		reflect.TypeOf((*MockEventsRepo)(nil).UpdateDoctor), ctx, doctor)
// }

// func (m *MockEventsRepo) DeleteDoctor(ctx context.Context, id int32) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "DeleteDoctor", ctx, id)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// func (mr *MockEventsRepoRecorder) DeleteDoctor(ctx, id interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDoctor",
// 		reflect.TypeOf((*MockEventsRepo)(nil).DeleteDoctor), ctx, id)
// }

// func (m *MockEventsRepo) FetchDoctors(ctx context.Context) ([]entity.Doctor, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FetchDoctors", ctx)
// 	ret0, _ := ret[0].([]entity.Doctor)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) FetchDoctors(ctx interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDoctors",
// 		reflect.TypeOf((*MockEventsRepo)(nil).FetchDoctors), ctx)
// }

// func (m *MockEventsRepo) StoreSchedule(ctx context.Context, events []entity.Event) (time.Time, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "StoreSchedule", ctx, events)
// 	ret0, _ := ret[0].(time.Time)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) StoreSchedule(ctx, events interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSchedule",
// 		reflect.TypeOf((*MockEventsRepo)(nil).StoreSchedule), ctx, events)
// }

// func (m *MockEventsRepo) FetchOpenEventsByDoctor(ctx context.Context, doctorId int32) ([]entity.Event, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FetchOpenEventsByDoctor", ctx, doctorId)
// 	ret0, _ := ret[0].([]entity.Event)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) FetchOpenEventsByDoctor(ctx, doctorId interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOpenEventsByDoctor",
// 		reflect.TypeOf((*MockEventsRepo)(nil).FetchOpenEventsByDoctor), ctx, doctorId)
// }

// func (m *MockEventsRepo) FetchReservedEventsByDoctor(ctx context.Context, doctorId int32) ([]entity.Event, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FetchReservedEventsByDoctor", ctx, doctorId)
// 	ret0, _ := ret[0].([]entity.Event)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) FetchReservedEventsByDoctor(ctx, doctorId interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchReservedEventsByDoctor",
// 		reflect.TypeOf((*MockEventsRepo)(nil).FetchReservedEventsByDoctor), ctx, doctorId)
// }

// func (m *MockEventsRepo) FetchReservedEventsByClient(ctx context.Context, clientId int32) ([]entity.Event, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FetchReservedEventsByClient", ctx, clientId)
// 	ret0, _ := ret[0].([]entity.Event)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) FetchReservedEventsByClient(ctx, clientId interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchReservedEventsByClient",
// 		reflect.TypeOf((*MockEventsRepo)(nil).FetchReservedEventsByClient), ctx, clientId)
// }

// func (m *MockEventsRepo) FetchAllEventsByClient(ctx context.Context, clientId int32) ([]entity.Event, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "FetchAllEventsByClient", ctx, clientId)
// 	ret0, _ := ret[0].([]entity.Event)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) FetchAllEventsByClient(ctx, clientId interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllEventsByClient",
// 		reflect.TypeOf((*MockEventsRepo)(nil).FetchAllEventsByClient), ctx, clientId)
// }

// func (m *MockEventsRepo) UpdateEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "UpdateEvent", ctx, event)
// 	ret0, _ := ret[0].(entity.Event)
// 	ret1, _ := ret[1].(error)
// 	return ret0, ret1
// }

// func (mr *MockEventsRepoRecorder) UpdateEvent(ctx, event interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent",
// 		reflect.TypeOf((*MockEventsRepo)(nil).UpdateEvent), ctx, event)
// }

// func (m *MockEventsRepo) ClearEvent(ctx context.Context, event entity.Event) error {
// 	m.ctrl.T.Helper()
// 	ret := m.ctrl.Call(m, "ClearEvent", ctx, event)
// 	ret0, _ := ret[0].(error)
// 	return ret0
// }

// func (mr *MockEventsRepoRecorder) ClearEvent(ctx, event interface{}) *gomock.Call {
// 	mr.mock.ctrl.T.Helper()
// 	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearEvent",
// 		reflect.TypeOf((*MockEventsRepo)(nil).ClearEvent), ctx, event)
// }
