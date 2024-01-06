package api

import (
	"context"
	"github.com/pwera/grpc-micros-payment/internal/adapters/sse"
	"github.com/pwera/grpc-micros-payment/internal/application/core/domain"
	"github.com/pwera/grpc-micros-payment/internal/ports"
)

type Application struct {
	db ports.DBPort
	p  *sse.Adapter
}

func NewApplication(db ports.DBPort, p *sse.Adapter) *Application {
	return &Application{
		db: db,
		p:  p,
	}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}
	a.p.Send(payment)
	return payment, nil
}
