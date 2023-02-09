package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/api"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/config"
	p "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/postgres"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/server"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/postgres"
)

func main() {
	// config
	cfg, err := config.NewConfig("config.yml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// logger
	l := logger.New(cfg.Logger.Level)

	time.Sleep(time.Second * 10)

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

	// Migrate
	err = api.Migrate(cfg, l)
	if err != nil {
		l.Error(fmt.Errorf("app - Run - Migrate: %w", err))
	}

	// Repository
	repo := p.NewEventsRepo(pg)

	// Grpc server
	if err := server.NewGrpcServer(repo, l).Start(cfg); err != nil {
		l.Error("Failed creating gRPC server", err)
		return
	}
}
