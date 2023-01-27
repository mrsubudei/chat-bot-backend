package app

import (
	"context"
	"fmt"

	config "github.com/mrsubudei/chat-bot-backend/appointment-service/config"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
	p "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/postgres"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/service"
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
	// docktor := entity.Doctor{
	// 	// Id:    int32(456),
	// 	Name:    "ab3c",
	// 	Surname: "se4f",
	// 	Phone:   "222",
	// }
	// schedule := entity.Schedule{

	// }
	event := entity.Event{
		Id:       834,
		ClientId: 77,
	}
	service := service.NewEventsService(repo)
	err = service.RegisterToEvent(context.Background(), event)
	if err != nil {
		fmt.Println(err)
	}

	// for i := 0; i < len(d); i++ {
	// fmt.Println(d)
	// }

	/*
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
	*/
}
