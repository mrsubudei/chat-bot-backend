package api

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/config"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/entity"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/repository"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/auth"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/hasher"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/mailer"
	pb "github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthorizationServer struct {
	pb.UnimplementedAuthorizationServer
	repo         repository.Users
	l            logger.Interface
	cfg          *config.Config
	hasher       *hasher.BcryptHasher
	tokenManager *auth.Manager
	mailer       mailer.Interface
}

func NewAuthorizationServer(repo repository.Users, l logger.Interface,
	cfg *config.Config, hasher *hasher.BcryptHasher, manager *auth.Manager,
	mailer mailer.Interface) *AuthorizationServer {

	return &AuthorizationServer{
		repo:         repo,
		l:            l,
		cfg:          cfg,
		hasher:       hasher,
		tokenManager: manager,
		mailer:       mailer,
	}
}

func (as *AuthorizationServer) SignUp(ctx context.Context,
	in *pb.UserSingle) (*pb.IdRequest, error) {

	//validate data
	if err := as.checkUserData(in); err != nil {
		return &pb.IdRequest{}, err
	}

	// hash password
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

	// generate verification token
	generatedBytes := make([]byte, 32)
	_, err = rand.Read(generatedBytes)
	if err != nil {
		as.l.Error(fmt.Errorf("api - CreateUser - rand.Read: %w", err))
		return &pb.IdRequest{}, status.Error(codes.Internal, InternalErr)
	}
	verificationToken := fmt.Sprintf("%x", generatedBytes)
	user.VerificationToken = verificationToken
	user.VerificationTTL = time.Now().Add(time.Hour *
		time.Duration(as.cfg.TokenManager.VerificationExpiringTime))

	// store in database
	id, err := as.repo.Store(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), DuplicateErrMsg) {
			return &pb.IdRequest{}, status.Error(codes.AlreadyExists,
				entity.ErrUserAlreadyExists.Error())
		}
		as.l.Error(fmt.Errorf("api - CreateUser - Store: %w", err))
		return &pb.IdRequest{}, status.Error(codes.Internal, InternalErr)
	}

	// send verification email
	err = as.mailer.DialAndSend(user.Email, id, string(verificationToken))
	if err != nil {
		// if error, delete created user
		err = as.repo.Delete(ctx, id)
		if err != nil {
			as.l.Error(fmt.Errorf("api - CreateUser - Delete: %w", err))
			return &pb.IdRequest{}, status.Error(codes.Internal, InternalErr)
		}
		as.l.Error(fmt.Errorf("api - CreateUser - DialAndSend: %w", err))
		return &pb.IdRequest{}, status.Error(codes.Internal, InternalErr)
	}

	idRequest := &pb.IdRequest{
		Id: id,
	}
	return idRequest, nil
}

func (as *AuthorizationServer) SignIn(ctx context.Context,
	in *pb.UserSingle) (*pb.Empty, error) {

	user := entity.User{
		Email:    in.Value.Email,
		Password: in.Value.Password,
	}

	found, err := as.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		as.l.Error(fmt.Errorf("api - UpdateSession - GetByEmail: %w", err))
		return &pb.Empty{}, status.Error(codes.Internal, InternalErr)
	}

	if !found.Verified {
		return &pb.Empty{}, status.Error(codes.PermissionDenied,
			entity.ErrUserNotVerified.Error())
	}

	err = as.repo.UpdateSession(ctx, user)
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

func (as *AuthorizationServer) VerifyRegistration(ctx context.Context,
	in *pb.UserSingle) (*pb.Empty, error) {

	user := entity.User{
		Id:                in.Value.Id,
		VerificationToken: in.Value.VerificationToken,
	}

	found, err := as.repo.GetById(ctx, user.Id)
	if err != nil {
		as.l.Error(fmt.Errorf("api - VerifyRegistration - GetById: %w", err))
		return &pb.Empty{}, status.Error(codes.Internal, InternalErr)
	}

	if found.VerificationToken != user.VerificationToken {
		return &pb.Empty{}, status.Error(codes.InvalidArgument,
			entity.ErrTokenNotValid.Error())
	}

	expired, err := as.tokenManager.CheckTTLExpired(found.VerificationTTL)
	if err != nil {
		as.l.Error(fmt.Errorf("api - VerifyRegistration - CheckTTLExpired: %w", err))
		return &pb.Empty{}, status.Error(codes.Internal, InternalErr)
	}

	if expired {
		return &pb.Empty{}, status.Error(codes.DeadlineExceeded,
			entity.ErrTokenTTLExpired.Error())
	}

	user.Verified = true

	err = as.repo.UpdateVerification(ctx, user)
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
			Id:                found.Id,
			Name:              found.Name,
			Phone:             found.Phone,
			Email:             found.Email,
			Password:          found.Password,
			Role:              found.Role,
			Verified:          found.Verified,
			VerificationToken: found.VerificationToken,
			SessionToken:      found.SessionToken,
			SessionTtl:        timestamppb.New(found.SessionTTL),
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
			Id:                found.Id,
			Name:              found.Name,
			Phone:             found.Phone,
			Email:             found.Email,
			Password:          found.Password,
			Role:              found.Role,
			Verified:          found.Verified,
			VerificationToken: found.VerificationToken,
			SessionToken:      found.SessionToken,
			SessionTtl:        timestamppb.New(found.SessionTTL),
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
			Id:                found.Id,
			Name:              found.Name,
			Phone:             found.Phone,
			Email:             found.Email,
			Password:          found.Password,
			Role:              found.Role,
			Verified:          found.Verified,
			VerificationToken: found.VerificationToken,
			SessionToken:      found.SessionToken,
			SessionTtl:        timestamppb.New(found.SessionTTL),
		},
	}

	return user, nil
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
