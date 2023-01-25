package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/mrsubudei/chat-bot-backend/gateway/proto/appointment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:8081", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAppointmentClient(conn)

	// Contact the server and print out its response.

	layout := "2006-01-02 15:04:05"
	firstDay, err := time.Parse(layout, "2023-01-26 00:00:00")
	if err != nil {
		log.Fatal(err)
	}
	lastDay, _ := time.Parse(layout, "2023-01-26 00:00:00")
	startTime, _ := time.Parse(layout, "2023-01-26 09:00:00")
	endTime, _ := time.Parse(layout, "2023-01-26 15:00:00")
	startBreak, _ := time.Parse(layout, "2023-01-26 12:00:00")
	endBreak, _ := time.Parse(layout, "2023-01-26 13:00:00")

	fd := timestamppb.New(firstDay)
	lt := timestamppb.New(lastDay)
	st := timestamppb.New(startTime)
	et := timestamppb.New(endTime)
	sb := timestamppb.New(startBreak)
	eb := timestamppb.New(endBreak)

	schedule := &pb.Schedule{
		FirstDay:             fd,
		LastDay:              lt,
		StartTime:            st,
		EndTime:              et,
		StartBreak:           sb,
		EndBreak:             eb,
		EventDurationMinutes: 30,
	}
	e, err := c.CreateSchedule(context.Background(), &pb.ScheduleRequest{Value: schedule})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	e.GetValue()
}
