package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/config"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
	p "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/postgres"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/service"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/postgres"
	pb "github.com/mrsubudei/chat-bot-backend/gateway/proto/appointment"
	"google.golang.org/grpc"
)

type AppointmentServer struct {
	pb.UnimplementedAppointmentServer
	service *service.ClientsService
}

func NewAppointmentServer(srv *service.ClientsService) *AppointmentServer {
	return &AppointmentServer{service: srv}
}

func (as *AppointmentServer) CreateSchedule(ctx context.Context,
	in *pb.ScheduleRequest) (*pb.ScheduleResponse, error) {
	fd := in.Value.FirstDay.AsTime()
	ld := in.Value.LastDay.AsTime()
	st := in.Value.StartTime.AsTime()
	et := in.Value.EndTime.AsTime()
	sb := in.Value.StartBreak.AsTime()
	eb := in.Value.EndBreak.AsTime()
	fmt.Println(in)
	schedule := entity.Schedule{
		FirstDay:      fd,
		LastDay:       ld,
		StartTime:     st,
		EndTime:       et,
		StartBreak:    sb,
		EndBreak:      eb,
		EventDuration: int(in.Value.EventDurationMinutes),
	}

	err := as.service.CreateSchedule(ctx, schedule)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.ScheduleResponse{}, nil
}

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
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

	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterAppointmentServer(s, NewAppointmentServer(service))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}