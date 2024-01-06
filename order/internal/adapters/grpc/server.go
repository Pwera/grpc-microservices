package grpc

import (
	"context"
	"fmt"
<<<<<<< HEAD
	"github.com/pwera/grpc-micros-commons/config"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/pwera/grpc-micros-order/internal/application/core/middleware"
	"github.com/pwera/grpc-micros-order/internal/ports"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
=======
	"github.com/pwera/grpc-micros-order/config"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/pwera/grpc-micros-order/internal/application/core/middleware"
	"github.com/pwera/grpc-micros-order/internal/ports"
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Adapter struct {
	api  ports.APIPort
	port int
	UnimplementedOrderServer
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

	RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}

}

func (a Adapter) Create(ctx context.Context, request *CreateOrderRequest) (*CreateOrderResponse, error) {
	err := middleware.HandleError(request.Validate(), "Validation is not passing")
	if err != nil {
		return nil, err
	}
	var orderItems []domain.OrderItem
	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
<<<<<<< HEAD
	result, err := a.api.PlaceOrder(ctx, newOrder)
=======
	result, err := a.api.PlaceOrder(newOrder)
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	if err != nil {
		return nil, err
	}
	return &CreateOrderResponse{OrderId: result.ID}, nil
}
