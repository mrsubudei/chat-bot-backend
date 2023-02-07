package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/config"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/api"
	p "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/postgres"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/postgres"
	pb "github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	// config
	cfg, err := config.NewConfig("./config/config.yml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// logger
	l := logger.New(cfg.Logger.Level)

	//Postgres
	pg, err := postgres.New(cfg)
	if err != nil {
		l.Error(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer func() {
		err = pg.Close()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - pg.Close: %w", err))
		}
	}()

	// Repository
	repo := p.NewEventsRepo(pg)

	// grpc server
	lis, err := net.Listen("tcp", cfg.HTTP.Host+cfg.HTTP.Port)
	if err != nil {
		l.Error(fmt.Errorf("app - Run - Listen: %w", err))
	}
	s := grpc.NewServer()

	pb.RegisterAppointmentServer(s, api.NewAppointmentServer(repo, l))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		l.Error(fmt.Errorf("app - Run - Serve: %w", err))
	}
}
