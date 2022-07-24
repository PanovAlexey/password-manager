package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var configSingleton *Config

type Config struct {
	applicationName     string
	grpcServerAddress   string
	timeoutHttpShutdown int // seconds
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

func (c Config) GetTimeoutHttpShutdown() int {
	return c.timeoutHttpShutdown
}

func initConfigByEnv(config Config) Config {
	config.applicationName = getEnv("UDM_APPLICATION_NAME")
	config.grpcServerAddress = getEnv("UDM_GRPC_SERVER_ADDRESS_ADDRESS")

	timeoutHttpShutdown := getEnv("UDM_TIMEOUT_HTTP_SHUTDOWN")

	if len(timeoutHttpShutdown) > 0 {
		config.timeoutHttpShutdown, _ = strconv.Atoi(timeoutHttpShutdown)
	}

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

	if config.timeoutHttpShutdown == 0 {
		config.timeoutHttpShutdown = 10
	}

	return config
}
