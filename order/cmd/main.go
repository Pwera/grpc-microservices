package main

import (
<<<<<<< HEAD
	"context"
	"github.com/pwera/grpc-micros-commons/config"
	"github.com/pwera/grpc-micros-commons/telemetry"
=======
	"github.com/pwera/grpc-micros-order/config"
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	"github.com/pwera/grpc-micros-order/internal/adapters/db"
	"github.com/pwera/grpc-micros-order/internal/adapters/grpc"
	"github.com/pwera/grpc-micros-order/internal/adapters/payment"
	"github.com/pwera/grpc-micros-order/internal/application/core/api"
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

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL(), tp, serviceName)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v\n", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl(), tp.Tracer("order-client-tracer"))
=======
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v\n", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v\n", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
