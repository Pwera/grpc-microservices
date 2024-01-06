package ports

import (
<<<<<<< HEAD
	"context"
=======
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
)

type APIPort interface {
<<<<<<< HEAD
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
=======
	PlaceOrder(order domain.Order) (domain.Order, error)
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
}
