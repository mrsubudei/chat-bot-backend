package mock_repository

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/entity"
)

type MockUsersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepoRecorder
}

type MockUsersRepoRecorder struct {
	mock *MockUsersRepo
}

func NewMockUsersRepo(ctrl *gomock.Controller) *MockUsersRepo {
	mock := &MockUsersRepo{ctrl: ctrl}
	mock.recorder = &MockUsersRepoRecorder{mock}
	return mock
}

func (m *MockUsersRepo) EXPECT() *MockUsersRepoRecorder {
	return m.recorder
}

func (m *MockUsersRepo) Store(ctx context.Context, user entity.User) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, user)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersRepoRecorder) Store(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store",
		reflect.TypeOf((*MockUsersRepo)(nil).Store), ctx, user)
}

func (m *MockUsersRepo) GetByPhone(ctx context.Context,
	phone string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPhone", ctx, phone)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersRepoRecorder) GetByPhone(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPhone",
		reflect.TypeOf((*MockUsersRepo)(nil).GetByPhone), ctx, phone)
}

func (m *MockUsersRepo) GetById(ctx context.Context, id int32) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersRepoRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById",
		reflect.TypeOf((*MockUsersRepo)(nil).GetById), ctx, id)
}

func (m *MockUsersRepo) GetByToken(ctx context.Context,
	token string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByToken", ctx, token)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersRepoRecorder) GetByToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByToken",
		reflect.TypeOf((*MockUsersRepo)(nil).GetByToken), ctx, token)
}

func (m *MockUsersRepo) UpdateSession(ctx context.Context, user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUsersRepoRecorder) UpdateSession(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession",
		reflect.TypeOf((*MockUsersRepo)(nil).UpdateSession), ctx, user)
}

func (m *MockUsersRepo) UpdateVerification(ctx context.Context, user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVerification", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUsersRepoRecorder) UpdateVerification(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerification",
		reflect.TypeOf((*MockUsersRepo)(nil).UpdateVerification), ctx, user)
}
