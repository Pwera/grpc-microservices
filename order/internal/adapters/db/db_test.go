package db

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"testing"
	"time"
)

type OrderDatabaseTestSuite struct {
	suite.Suite
	DataSourceUrl string
}

func (o *OrderDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()
	port := "3306/tcp"
	dbURL := func(host string, port nat.Port) string {
		return fmt.Sprintf("root:s3cr3t@tcp(localhost:%s)/orders?charset=utf8mb4&parseTime=True&loc=Local", port.Port())
	}
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/mysql:8.0.30",
		ExposedPorts: []string{port},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "s3cr3t",
			"MYSQL_DATABASE":      "orders",
		},
		WaitingFor: wait.ForSQL(nat.Port(port), "mysql", dbURL).WithStartupTimeout(30 * time.Second),
	}

	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start MYSQL %v\n", err)
	}
	endpoint, _ := mysqlContainer.Endpoint(ctx, "")
	o.DataSourceUrl = fmt.Sprintf("root:s3cr3t@tcp(%s)/orders?charset=utf8mb4&parseTime=True&loc=Local", endpoint)
}

func (o *OrderDatabaseTestSuite) Test_Should_Save_Order() {
<<<<<<< HEAD
	adapter, err := NewAdapter(o.DataSourceUrl, nil, "orders-v1")
	o.Nil(err)
	saveErr := adapter.Save(context.Background(), &domain.Order{})
=======
	adapter, err := NewAdapter(o.DataSourceUrl)
	o.Nil(err)
	saveErr := adapter.Save(&domain.Order{})
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	o.Nil(saveErr)
}

func (o *OrderDatabaseTestSuite) Test_Should_Get_Order() {
<<<<<<< HEAD
	adapter, _ := NewAdapter(o.DataSourceUrl, nil, "orders-v1")
=======
	adapter, _ := NewAdapter(o.DataSourceUrl)
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	order := domain.NewOrder(2, []domain.OrderItem{
		{
			ProductCode: "CAM",
			Quantity:    5,
			UnitPrice:   1.01,
		},
	})
<<<<<<< HEAD
	ctx := context.Background()
	err := adapter.Save(ctx, &order)
	o.Nil(err)
	ord, err := adapter.Get(ctx, order.ID)
=======
	err := adapter.Save(&order)
	o.Nil(err)
	ord, err := adapter.Get(order.ID)
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	o.Nil(err)
	o.Equal(int32(2), ord.CustomerID)
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDatabaseTestSuite))
}
