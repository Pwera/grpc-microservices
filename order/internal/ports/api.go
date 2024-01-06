package ports

import (
	"context"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}
