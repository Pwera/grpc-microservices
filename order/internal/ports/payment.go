package ports

import "github.com/pwera/grpc-micros-order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
