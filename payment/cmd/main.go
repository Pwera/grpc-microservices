package main

import (
<<<<<<< HEAD
	"context"
	"github.com/pwera/grpc-micros-commons/config"
	"github.com/pwera/grpc-micros-commons/telemetry"
=======
	"github.com/pwera/grpc-micros-payment/config"
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	"github.com/pwera/grpc-micros-payment/internal/adapters/db"
	"github.com/pwera/grpc-micros-payment/internal/adapters/grpc"
	"github.com/pwera/grpc-micros-payment/internal/adapters/sse"
	"github.com/pwera/grpc-micros-payment/internal/application/core/api"
	"log"
)

func main() {
<<<<<<< HEAD
	jaegerUrl := config.GetEnvironmentValue("JAEGER_SERVICE_NAME") // add defaults
	serviceName := config.GetEnvironmentValue("SERVICE_NAME")
	env := config.GetEnvironmentValue("ENVIRONMENT")
	tp, err := telemetry.TracerProvider(context.Background(), jaegerUrl, serviceName, env)
	if err != nil {
		log.Fatal(err)
	}
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL(), tp)
=======
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	sseAdapter := sse.NewAdapter()

	application := api.NewApplication(dbAdapter, sseAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	go sseAdapter.Run()
	grpcAdapter.Run()
}
