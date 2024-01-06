package ports

import (
	"context"
	"github.com/pwera/grpc-micros-payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
