package db

import (
	"context"
	"fmt"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

type Order struct {
	gorm.Model
	CustomerID int32
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

func NewAdapter(dataSourceUrl string, tp trace.TracerProvider) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	if err := db.Use(otelgorm.NewPlugin(
		otelgorm.WithTracerProvider(tp),
		otelgorm.WithDBName("order"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

	err := db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}
func (a Adapter) Get(ctx context.Context, id int32) (domain.Order, error) {
	var orderEntity Order
	res := a.db.WithContext(ctx).Preload("OrderItems").First(&orderEntity, id)
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	order := domain.Order{
		ID:         int32(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}
	return order, res.Error
}
func (a Adapter) Save(ctx context.Context, order *domain.Order) error {
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {

		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := a.db.WithContext(ctx).Create(&orderModel)
	if res.Error == nil {
		order.ID = int32(orderModel.ID)
	}
	return res.Error
}
