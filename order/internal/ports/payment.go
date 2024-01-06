package ports

<<<<<<< HEAD
import (
	"context"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(ctx context.Context, order *domain.Order) error
=======
import "github.com/pwera/grpc-micros-order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
}
