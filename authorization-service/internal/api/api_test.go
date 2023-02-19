package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/api"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/config"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/entity"
	mock_repository "github.com/mrsubudei/chat-bot-backend/authorization-service/internal/repository/mock"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/auth"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/hasher"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/logger"
	mock_mailer "github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/mailer/mock"
	pb "github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type test struct {
	name string
	mock func()
	user *pb.UserSingle
	res  interface{}
	err  error
}

func mockUserApi(t *testing.T) (*api.AuthorizationServer, *mock_repository.MockUsersRepo) {
	t.Helper()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := mock_repository.NewMockUsersRepo(mockCtl)
	cfg, err := config.NewConfig("../../config.yml", "../../env")
	if err != nil {
		t.Fatalf("Config error: %s", err)
	}

	l := logger.New(cfg.Logger.Level)
	hasher := hasher.NewBcryptHasher()
	tokenManager := auth.NewManager(cfg)
	mockMailer := mock_mailer.NewMockMailer()

	eventsService := api.NewAuthorizationServer(repo, l, cfg, hasher, tokenManager, mockMailer)

	return eventsService, repo
}

func TestCreateUser(t *testing.T) {

	apiService, repo := mockUserApi(t)
	ctx := context.Background()
	var id int32
	tests := []test{
		{
			name: "OK",
			mock: func() {
				repo.EXPECT().Store(ctx, gomock.Any()).Return(int32(0), nil)
			},
			user: &pb.UserSingle{
				Value: &pb.User{
					Name:     "Kilka",
					Phone:    "+7-701-798-87-77",
					Password: "pass",
				},
			},
			res: &pb.IdRequest{},
			err: nil,
		},
		{
			name: "Error user already exists",
			mock: func() {
				repo.EXPECT().Store(ctx, gomock.Any()).Return(id,
					errors.New(api.DuplicateErrMsg))
			},
			user: &pb.UserSingle{
				Value: &pb.User{
					Name:     "Kilka",
					Phone:    "+7-701-798-87-77",
					Password: "pass",
				},
			},
			res: &pb.IdRequest{},
			err: status.Error(codes.AlreadyExists,
				entity.ErrUserAlreadyExists.Error()),
		},
		{
			name: "Error internal",
			mock: func() {
				repo.EXPECT().Store(ctx, gomock.Any()).Return(id,
					errors.New("Internal error"))
			},
			user: &pb.UserSingle{
				Value: &pb.User{
					Name:     "Kilka",
					Phone:    "+7-701-798-87-77",
					Password: "pass",
				},
			},
			res: &pb.IdRequest{},
			err: status.Error(codes.Internal,
				api.InternalErr),
		},
		{
			name: "Error empty name field",
			mock: func() {
				repo.EXPECT().Store(ctx, gomock.Any()).Return(int32(0), nil)
			},
			user: &pb.UserSingle{
				Value: &pb.User{
					Phone:    "+7-701-798-87-77",
					Password: "pass",
				},
			},
			res: &pb.IdRequest{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": Name"),
		},
		{
			name: "Error empty phone field",
			mock: func() {
				repo.EXPECT().Store(ctx, gomock.Any()).Return(int32(0), nil)
			},
			user: &pb.UserSingle{
				Value: &pb.User{
					Name:     "Kilka",
					Password: "pass",
				},
			},
			res: &pb.IdRequest{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": Phone"),
		},
		{
			name: "Error wrong phone format",
			mock: func() {
				repo.EXPECT().Store(ctx, gomock.Any()).Return(int32(0), nil)
			},
			user: &pb.UserSingle{
				Value: &pb.User{
					Name:     "Kilka",
					Phone:    "87019875421",
					Password: "pass",
				},
			},
			res: &pb.IdRequest{},
			err: status.Error(codes.InvalidArgument,
				api.WrongPhoneFormat),
		},
		{
			name: "Error empty field password",
			mock: func() {
				repo.EXPECT().Store(ctx, gomock.Any()).Return(id, nil)
			},
			user: &pb.UserSingle{
				Value: &pb.User{
					Name:  "Kilka",
					Phone: "+7-701-798-87-77",
				},
			},
			res: &pb.IdRequest{},
			err: status.Error(codes.InvalidArgument,
				api.RequestZeroValue+": Password"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := apiService.SignUp(ctx, tc.user)
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}
