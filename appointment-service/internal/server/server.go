package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/api"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/config"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/proto"
)

type GrpcServer struct {
	repo repository.Events
	l    logger.Interface
}

func NewGrpcServer(repo repository.Events, l logger.Interface) *GrpcServer {
	return &GrpcServer{
		repo: repo,
		l:    l,
	}
}

func (gs *GrpcServer) Start(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcAddr := fmt.Sprintf("%s:%s", cfg.Grpc.Host, cfg.Grpc.Port)

	isReady := &atomic.Value{}
	isReady.Store(false)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Grpc.Timeout) * time.Minute,
		}),
	)

	pb.RegisterAppointmentServer(grpcServer, api.NewAppointmentServer(gs.repo, gs.l))

	go func() {
		gs.l.Info("GRPC Server is listening on: %s", grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			gs.l.Fatal("Failed running gRPC server", err)
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		gs.l.Info("The service is ready to accept requests")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		gs.l.Info("signal.Notify: %v", v)
	case done := <-ctx.Done():
		gs.l.Info("ctx.Done: %v", done)
	}

	isReady.Store(false)

	grpcServer.GracefulStop()
	gs.l.Info("grpcServer shut down correctly")

	return nil
}
