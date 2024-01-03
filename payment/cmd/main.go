package main

import (
	"github.com/pwera/grpc-micros-payment/config"
	"github.com/pwera/grpc-micros-payment/internal/adapters/db"
	"github.com/pwera/grpc-micros-payment/internal/adapters/grpc"
	"github.com/pwera/grpc-micros-payment/internal/adapters/sse"
	"github.com/pwera/grpc-micros-payment/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	sseAdapter := sse.NewAdapter()

	application := api.NewApplication(dbAdapter, sseAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	go sseAdapter.Run()
	grpcAdapter.Run()
}
