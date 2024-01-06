package ports

import (
	"context"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
)

type DBPort interface {
	Get(context.Context, int32) (domain.Order, error)
	Save(context.Context, *domain.Order) error
}
