package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/mrsubudei/chat-bot-backend/config"
	v1 "github.com/mrsubudei/chat-bot-backend/internal/controller/http/v1"
	p "github.com/mrsubudei/chat-bot-backend/internal/repository/postgres"
	"github.com/mrsubudei/chat-bot-backend/internal/service"
	"github.com/mrsubudei/chat-bot-backend/pkg/httpserver"
	"github.com/mrsubudei/chat-bot-backend/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/pkg/postgres"
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
	repo := p.NewClientsRepo(pg)

	// Service
	service := service.NewClientsService(repo)

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
