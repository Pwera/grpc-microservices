package payment

import (
	"context"
	"github.com/pwera/grpc-micros-order/internal/adapters/grpc"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/pwera/grpc-micros-order/internal/application/core/middleware"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Adapter struct {
	payment grpc.PaymentServiceClient
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
	conn, err := grpc2.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
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
	_, err := a.payment.Create(ctx, &grpc.CreatePaymentRequest{
		Price:      3.0,
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return middleware.HandleError(err, "order creation failed")
}
