package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var configSingleton *Config

type Config struct {
	applicationName    string
	grpcServerAddress  string
	storageGrpcAddress string
}

func New() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
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

func (c Config) GetStorageGRPCServerAddress() string {
	return c.storageGrpcAddress
}

func initConfigByEnv(config Config) Config {
	config.applicationName = getEnv("UDM_APPLICATION_NAME")
	config.grpcServerAddress = getEnv("UDM_GRPC_SERVER_ADDRESS")
	config.storageGrpcAddress = getEnv("UDM_SERVICE_STORAGE_GRPC_SERVER_ADDRESS")

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
		config.applicationName = "User data manager"
	}

	if len(config.grpcServerAddress) < 1 {
		config.grpcServerAddress = "localhost:3200"
	}

	if len(config.storageGrpcAddress) < 1 {
		config.storageGrpcAddress = "localhost:3205"
	}

	return config
}
