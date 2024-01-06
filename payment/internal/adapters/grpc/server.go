package grpc

import (
	"fmt"
<<<<<<< HEAD
	"github.com/pwera/grpc-micros-commons/config"
	"github.com/pwera/grpc-micros-payment/internal/ports"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
=======
	"github.com/pwera/grpc-micros-payment/config"
	"github.com/pwera/grpc-micros-payment/internal/ports"
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	UnimplementedPaymentServiceServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

<<<<<<< HEAD
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
=======
	grpcServer := grpc.NewServer()
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	a.server = grpcServer
	RegisterPaymentServiceServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("starting payment service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
