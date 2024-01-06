package api

import (
<<<<<<< HEAD
	"context"
=======
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
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

<<<<<<< HEAD
func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	err = a.payment.Charge(ctx, &order)
=======
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	err = a.payment.Charge(&order)
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
