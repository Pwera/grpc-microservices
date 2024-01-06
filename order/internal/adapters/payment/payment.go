package payment

import (
	"context"
	"github.com/pwera/grpc-micros-order/internal/adapters/grpc"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/pwera/grpc-micros-order/internal/application/core/middleware"
<<<<<<< HEAD
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
=======
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Adapter struct {
	payment grpc.PaymentServiceClient
<<<<<<< HEAD
	tracer  trace.Tracer
}

func NewAdapter(paymentServiceUrl string, tracer trace.Tracer) (*Adapter, error) {
	var opts []grpc2.DialOption
	opts = append(opts,
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
		grpc2.WithChainUnaryInterceptor(
			otelgrpc.UnaryClientInterceptor(),
			middleware.CircuitBreakerClientInterceptor("cbreaker")),
	)
	if false {
		opts = append(opts, grpc2.WithUnaryInterceptor(middleware.RetryInterceptor()))
	}

=======
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc2.DialOption
	opts = append(opts, grpc2.WithTransportCredentials(insecure.NewCredentials()))
	if true {
		opts = append(opts, grpc2.WithUnaryInterceptor(middleware.CircuitBreakerClientInterceptor("cbreaker")))
	}
	if false {
		opts = append(opts, grpc2.WithUnaryInterceptor(middleware.RetryInterceptor()))
	}
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	conn, err := grpc2.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	client := grpc.NewPaymentServiceClient(conn)
	return &Adapter{payment: client, tracer: tracer}, nil
}

func (a *Adapter) Charge(ctx context.Context, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	ctx, span := a.tracer.Start(ctx, "payments request")
=======
	// TODO: These should be closed
	//if conn != nil {
	//	err = conn.Close()
	//}
	//defer conn.Close()
	client := grpc.NewPaymentServiceClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	_, err := a.payment.Create(ctx, &grpc.CreatePaymentRequest{
		Price:      3.0,
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
<<<<<<< HEAD
	defer span.End()
=======
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	return middleware.HandleError(err, "order creation failed")
}
