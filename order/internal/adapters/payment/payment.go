package payment

import (
	"context"
	"github.com/pwera/grpc-micros-order/internal/adapters/grpc"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/pwera/grpc-micros-order/internal/application/core/middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Adapter struct {
	payment grpc.PaymentServiceClient
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

	conn, err := grpc2.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := grpc.NewPaymentServiceClient(conn)
	return &Adapter{payment: client, tracer: tracer}, nil
}

func (a *Adapter) Charge(ctx context.Context, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	ctx, span := a.tracer.Start(ctx, "payments request")
	_, err := a.payment.Create(ctx, &grpc.CreatePaymentRequest{
		Price:      3.0,
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	defer span.End()
	return middleware.HandleError(err, "order creation failed")
}
