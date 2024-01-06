package db

import (
<<<<<<< HEAD
	"context"
	"fmt"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
=======
	"fmt"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
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

<<<<<<< HEAD
func NewAdapter(dataSourceUrl string, tp trace.TracerProvider, serviceName string) (*Adapter, error) {
=======
func NewAdapter(dataSourceUrl string) (*Adapter, error) {
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

<<<<<<< HEAD
	if err := db.Use(otelgorm.NewPlugin(
		otelgorm.WithTracerProvider(tp),
		otelgorm.WithAttributes(semconv.ServiceNameKey.String(serviceName)),
		otelgorm.WithDBName("order"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

=======
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	err := db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}
<<<<<<< HEAD
func (a Adapter) Get(ctx context.Context, id int32) (domain.Order, error) {
	var orderEntity Order
	res := a.db.WithContext(ctx).Preload("OrderItems").First(&orderEntity, id)
=======
func (a Adapter) Get(id int32) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
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
<<<<<<< HEAD
func (a Adapter) Save(ctx context.Context, order *domain.Order) error {
=======
func (a Adapter) Save(order *domain.Order) error {
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
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
<<<<<<< HEAD
	res := a.db.WithContext(ctx).Create(&orderModel)
=======
	res := a.db.Create(&orderModel)
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	if res.Error == nil {
		order.ID = int32(orderModel.ID)
	}
	return res.Error
}
