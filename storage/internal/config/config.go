package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var configSingleton *Config

type Config struct {
	applicationName   string
	grpcServerAddress string
}

func New() Config {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("error loading global env variables: %s", err.Error())
	}

	if err := godotenv.Load("./.env"); err != nil {
		log.Printf("error loading personal env variables: %s", err.Error())
	}

	if configSingleton == nil {
		config := Config{}
		config = initConfigByEnv(config)
		config = initConfigByDefault(config)

		configSingleton = &config
	}

	return *configSingleton
}

func (c Config) GetApplicationName() string {
	return c.applicationName
}

func (c Config) GetGrpcServerAddress() string {
	return c.grpcServerAddress
}

func initConfigByEnv(config Config) Config {
	config.applicationName = getEnv("STORAGE_APPLICATION_NAME")
	config.grpcServerAddress = getEnv("STORAGE_GRPC_SERVER_ADDRESS")

	return config
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}

func initConfigByDefault(config Config) Config {
	if len(config.applicationName) < 1 {
		config.applicationName = "Storage"
	}

	if len(config.grpcServerAddress) < 1 {
		config.grpcServerAddress = "localhost:3201"
	}

	return config
}
