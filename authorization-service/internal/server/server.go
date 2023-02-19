package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/api"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/config"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/repository"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/auth"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/hasher"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/mailer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"

	pb "github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/proto"
)

type GrpcServer struct {
	repo         repository.Users
	l            logger.Interface
	cfg          *config.Config
	hasher       *hasher.BcryptHasher
	tokenManager *auth.Manager
	mailer       mailer.Interface
}

func NewGrpcServer(repo repository.Users, l logger.Interface, cfg *config.Config,
	hasher *hasher.BcryptHasher, manager *auth.Manager,
	mailer mailer.Interface) *GrpcServer {

	return &GrpcServer{
		repo:         repo,
		l:            l,
		hasher:       hasher,
		cfg:          cfg,
		tokenManager: manager,
		mailer:       mailer,
	}
}

func (gs *GrpcServer) Start(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcAddr := fmt.Sprintf("%s:%s", cfg.Grpc.Host, cfg.Grpc.Port)

	isReady := &atomic.Value{}
	isReady.Store(false)

	// read ca's cert, verify to client's certificate
	caPem, err := ioutil.ReadFile("cert/ca.cert")
	if err != nil {
		log.Fatal(err)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPem) {
		log.Fatal(err)
	}

	// read server cert & key
	serverCert, err := tls.LoadX509KeyPair("cert/service.pem", "cert/service.key")
	if err != nil {
		log.Fatal(err)
	}

	// configuration of the certificate what we want to
	conf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	//create tls certificate
	tlsCredentials := credentials.NewTLS(conf)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Grpc.Timeout) * time.Minute,
		}),
	)

	pb.RegisterAuthorizationServer(grpcServer,
		api.NewAuthorizationServer(gs.repo, gs.l, cfg, gs.hasher,
			gs.tokenManager, gs.mailer))

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
