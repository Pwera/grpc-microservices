package ports

import (
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
)

type DBPort interface {
	Get(id int32) (domain.Order, error)
	Save(*domain.Order) error
}
