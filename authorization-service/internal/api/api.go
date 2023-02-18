package api

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/entity"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/repository"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/auth"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/hasher"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/logger"
	pb "github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthorizationServer struct {
	pb.UnimplementedAuthorizationServer
	repo         repository.Users
	l            logger.Interface
	hasher       *hasher.BcryptHasher
	tokenManager *auth.Manager
}

func NewAuthorizationServer(repo repository.Users, l logger.Interface,
	hasher *hasher.BcryptHasher, manager *auth.Manager) *AuthorizationServer {
	return &AuthorizationServer{
		repo:         repo,
		l:            l,
		hasher:       hasher,
		tokenManager: manager,
	}
}

func (as *AuthorizationServer) CreateUser(ctx context.Context,
	in *pb.UserSingle) (*pb.IdRequest, error) {
	if err := as.checkUserData(in); err != nil {
		return &pb.IdRequest{}, err
	}

	hashed, err := as.hasher.Hash(in.Value.Password)
	if err != nil {
		return &pb.IdRequest{}, status.Error(codes.Internal, InternalErr)
	}

	user := entity.User{
		Name:     in.Value.Name,
		Phone:    in.Value.Phone,
		Email:    in.Value.Email,
		Password: hashed,
	}
	id, err := as.repo.Store(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), DuplicateErrMsg) {
			return &pb.IdRequest{}, status.Error(codes.AlreadyExists,
				entity.ErrUserAlreadyExists.Error())
		}
		as.l.Error(fmt.Errorf("api - CreateUser - Store: %w", err))
		return &pb.IdRequest{}, status.Error(codes.Internal, InternalErr)
	}

	idRequest := &pb.IdRequest{
		Id: id,
	}
	return idRequest, nil
}

func (as *AuthorizationServer) GetByPhone(ctx context.Context,
	in *pb.StringRequest) (*pb.UserSingle, error) {

	if err := as.checkPhoneFormat(in.Str); err != nil {
		return &pb.UserSingle{}, err
	}

	found, err := as.repo.GetByPhone(ctx, in.Str)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) {
			return &pb.UserSingle{}, status.Error(codes.NotFound,
				entity.ErrUserDoesNotExist.Error())
		}
		as.l.Error(fmt.Errorf("api - GetByPhone: %w", err))
		return &pb.UserSingle{}, status.Error(codes.Internal, InternalErr)
	}

	user := &pb.UserSingle{
		Value: &pb.User{
			Id:           found.Id,
			Name:         found.Name,
			Phone:        found.Phone,
			Email:        found.Email,
			Password:     found.Password,
			Role:         found.Role,
			Verified:     found.Verified,
			SmsCode:      found.SmsCode,
			SessionToken: found.SessionToken,
			SessionTtl:   timestamppb.New(found.SessionTTL),
		},
	}

	return user, nil
}

func (as *AuthorizationServer) GetById(ctx context.Context,
	in *pb.IdRequest) (*pb.UserSingle, error) {

	found, err := as.repo.GetById(ctx, in.Id)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) {
			return &pb.UserSingle{}, status.Error(codes.NotFound,
				entity.ErrUserDoesNotExist.Error())
		}
		as.l.Error(fmt.Errorf("api - GetById: %w", err))
		return &pb.UserSingle{}, status.Error(codes.Internal, InternalErr)
	}

	user := &pb.UserSingle{
		Value: &pb.User{
			Id:           found.Id,
			Name:         found.Name,
			Phone:        found.Phone,
			Email:        found.Email,
			Password:     found.Password,
			Role:         found.Role,
			Verified:     found.Verified,
			SmsCode:      found.SmsCode,
			SessionToken: found.SessionToken,
			SessionTtl:   timestamppb.New(found.SessionTTL),
		},
	}

	return user, nil
}

func (as *AuthorizationServer) GetByToken(ctx context.Context,
	in *pb.StringRequest) (*pb.UserSingle, error) {

	found, err := as.repo.GetByToken(ctx, in.Str)
	if err != nil {
		if errors.Is(err, entity.ErrUserDoesNotExist) {
			return &pb.UserSingle{}, status.Error(codes.NotFound,
				entity.ErrUserDoesNotExist.Error())
		}
		as.l.Error(fmt.Errorf("api - GetByToken: %w", err))
		return &pb.UserSingle{}, status.Error(codes.Internal, InternalErr)
	}

	user := &pb.UserSingle{
		Value: &pb.User{
			Id:           found.Id,
			Name:         found.Name,
			Phone:        found.Phone,
			Email:        found.Email,
			Password:     found.Password,
			Role:         found.Role,
			Verified:     found.Verified,
			SmsCode:      found.SmsCode,
			SessionToken: found.SessionToken,
			SessionTtl:   timestamppb.New(found.SessionTTL),
		},
	}

	return user, nil
}

func (as *AuthorizationServer) UpdateSession(ctx context.Context,
	in *pb.UserSingle) (*pb.Empty, error) {

	user := entity.User{
		Id:           in.Value.Id,
		SessionToken: in.Value.SessionToken,
		SessionTTL:   in.Value.SessionTtl.AsTime(),
	}

	err := as.repo.UpdateSession(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), NoRowsAffected) {
			return &pb.Empty{}, status.Error(codes.NotFound,
				entity.ErrUserDoesNotExist.Error())
		}
		as.l.Error(fmt.Errorf("api - UpdateSession: %w", err))
		return &pb.Empty{}, status.Error(codes.Internal, InternalErr)
	}

	return &pb.Empty{}, nil
}

func (as *AuthorizationServer) UpdateVerification(ctx context.Context,
	in *pb.UserSingle) (*pb.Empty, error) {

	user := entity.User{
		Id:       in.Value.Id,
		SmsCode:  in.Value.SmsCode,
		Verified: in.Value.Verified,
	}

	err := as.repo.UpdateVerification(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), NoRowsAffected) {
			return &pb.Empty{}, status.Error(codes.NotFound,
				entity.ErrUserDoesNotExist.Error())
		}
		as.l.Error(fmt.Errorf("api - UpdateVerification: %w", err))
		return &pb.Empty{}, status.Error(codes.Internal, InternalErr)
	}

	return &pb.Empty{}, nil
}

func (as *AuthorizationServer) checkUserData(in *pb.UserSingle) error {
	switch {
	case in.Value.Name == "":
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": Name")
	case in.Value.Password == "":
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": Password")
	}

	if _, err := mail.ParseAddress(in.Value.Email); in.Value.Email != "" &&
		err != nil {
		return status.Error(codes.InvalidArgument,
			WrongEmailFormat)
	}

	if err := as.checkPhoneFormat(in.Value.Phone); err != nil {
		return err
	}

	return nil
}

func (as *AuthorizationServer) checkPhoneFormat(phone string) error {
	if phone == "" {
		return status.Error(codes.InvalidArgument,
			RequestZeroValue+": Phone")
	}

	re := regexp.MustCompile(`^\+7-\d{3}-\d{3}-\d{2}-\d{2}$`)
	if match := re.MatchString(phone); !match {
		return status.Error(codes.InvalidArgument,
			WrongPhoneFormat)
	}

	return nil
}
