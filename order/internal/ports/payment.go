package ports

import (
	"context"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(ctx context.Context, order *domain.Order) error
}
