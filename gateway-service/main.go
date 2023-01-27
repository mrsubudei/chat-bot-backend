package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/mrsubudei/chat-bot-backend/appointment-service/proto"
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
	firstDay, err := time.Parse(layout, "2023-01-28 00:00:00")
	if err != nil {
		log.Fatal(err)
	}
	lastDay, _ := time.Parse(layout, "2023-01-28 00:00:00")
	startTime, _ := time.Parse(layout, "2023-01-28 09:00:00")
	endTime, _ := time.Parse(layout, "2023-01-28 15:00:00")
	startBreak, _ := time.Parse(layout, "2023-01-28 12:00:00")
	endBreak, _ := time.Parse(layout, "2023-01-28 13:00:00")

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
		DoctorId:             []int32{9, 10},
	}
	e, err := c.CreateSchedule(context.Background(), &pb.ScheduleRequest{Value: schedule})
	if err != nil {
		fmt.Println(err)
	}
	e.GetValue()
}
