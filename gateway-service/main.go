package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	addr = flag.String("addr", "127.0.0.1:8087", "the address to connect to")
)

func main() {
	flag.Parse()
	// read ca's cert
	caCert, err := ioutil.ReadFile("cert/ca.cert")
	if err != nil {
		log.Fatal(caCert)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(err)
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair("cert/service.pem", "cert/service.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(tlsCredential))
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
	startTime, _ := time.Parse(layout, "0001-01-01 01:00:00")
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
		DoctorIds:            []int32{1},
	}
	_, err = c.CreateSchedule(context.Background(), &pb.ScheduleSingle{Value: schedule})
	if err != nil {
		fmt.Println(err)
	}
}
