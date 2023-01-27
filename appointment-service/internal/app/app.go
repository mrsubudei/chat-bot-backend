package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	config "github.com/mrsubudei/chat-bot-backend/appointment-service/config"
	v1 "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/controller/http/v1"
	p "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/postgres"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/service"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/httpserver"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/postgres"
)

func Run(cfg *config.Config) {
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

	// Service
	service := service.NewEventsService(repo)

	handler := gin.New()
	v1.NewRouter(handler, service, l)

	httpServer := httpserver.New(handler)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	//Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
