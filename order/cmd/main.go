package main

import (
	"github.com/pwera/grpc-micros-order/config"
	"github.com/pwera/grpc-micros-order/internal/adapters/db"
	"github.com/pwera/grpc-micros-order/internal/adapters/grpc"
	"github.com/pwera/grpc-micros-order/internal/adapters/payment"
	"github.com/pwera/grpc-micros-order/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v\n", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v\n", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
