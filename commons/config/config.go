package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return GetEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return GetEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := GetEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

func GetPaymentServiceUrl() string {
	return GetEnvironmentValue("PAYMENT_SERVICE_URL")
}

func GetEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}
