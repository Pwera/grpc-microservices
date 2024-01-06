package db

import (
	"context"
	"fmt"
	"github.com/pwera/grpc-micros-payment/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int32
	Status     string
	OrderID    int32
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Payment, error) {
	var paymentEntity Payment
	res := a.db.WithContext(ctx).First(&paymentEntity, id)
	payment := domain.Payment{
		ID:         int64(paymentEntity.ID),
		CustomerID: int32(paymentEntity.CustomerID),
		Status:     paymentEntity.Status,
		OrderId:    int32(paymentEntity.OrderID),
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.UnixNano(),
	}
	return payment, res.Error
}

func (a Adapter) Save(ctx context.Context, payment *domain.Payment) error {
	orderModel := Payment{
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		OrderID:    payment.OrderId,
		TotalPrice: payment.TotalPrice,
	}
	res := a.db.WithContext(ctx).Create(&orderModel)
	if res.Error == nil {
		payment.ID = int64(orderModel.ID)
	}
	return res.Error
}

func NewAdapter(dataSourceUrl string, tp trace.TracerProvider) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	if err := db.Use(otelgorm.NewPlugin(
		otelgorm.WithAttributes(semconv.ServiceNameKey.String("payments-v1")),
		otelgorm.WithTracerProvider(tp),
		otelgorm.WithDBName("payments"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}
