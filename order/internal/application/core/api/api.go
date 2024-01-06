package api

import (
	"context"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/pwera/grpc-micros-order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	err = a.payment.Charge(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
