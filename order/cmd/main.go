package main

import (
	"context"
	"github.com/pwera/grpc-micros-commons/config"
	"github.com/pwera/grpc-micros-commons/telemetry"
	"github.com/pwera/grpc-micros-order/internal/adapters/db"
	"github.com/pwera/grpc-micros-order/internal/adapters/grpc"
	"github.com/pwera/grpc-micros-order/internal/adapters/payment"
	"github.com/pwera/grpc-micros-order/internal/application/core/api"
	"log"
)

func main() {
	jaegerUrl := config.GetEnvironmentValue("JAEGER_SERVICE_NAME") // add defaults
	serviceName := config.GetEnvironmentValue("SERVICE_NAME")
	env := config.GetEnvironmentValue("ENVIRONMENT")
	tp, err := telemetry.TracerProvider(context.Background(), jaegerUrl, serviceName, env)
	if err != nil {
		log.Fatal(err)
	}
	_ = tp

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL(), tp)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v\n", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl(), tp.Tracer("order-client-tracer"))
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v\n", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
