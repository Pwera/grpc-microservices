package ports

import (
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
