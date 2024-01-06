package ports

import (
<<<<<<< HEAD
	"context"
=======
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
)

type DBPort interface {
<<<<<<< HEAD
	Get(context.Context, int32) (domain.Order, error)
	Save(context.Context, *domain.Order) error
=======
	Get(id int32) (domain.Order, error)
	Save(*domain.Order) error
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
}
