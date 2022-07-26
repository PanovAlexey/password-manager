package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var configSingleton *Config

type Config struct {
	applicationName              string
	serverAddress                string
	userDataManagerGRPCAddress   string
	userAuthorizationGRPCAddress string
	timeoutHttpShutdown          int // seconds
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

func (c Config) GetServerAddress() string {
	return c.serverAddress
}

func (c Config) GetUserDataManagerGRPCServerAddress() string {
	return c.userDataManagerGRPCAddress
}

func (c Config) GetUserAuthorizationGRPCServerAddress() string {
	return c.userAuthorizationGRPCAddress
}

func (c Config) GetTimeoutHttpShutdown() int {
	return c.timeoutHttpShutdown
}

func initConfigByEnv(config Config) Config {
	config.applicationName = getEnv("GW_APPLICATION_NAME")
	config.serverAddress = getEnv("GW_SERVER_ADDRESS")
	config.userDataManagerGRPCAddress = getEnv("GW_SERVICE_USER_DATA_MANAGER_GRPC_ADDRESS")
	config.userAuthorizationGRPCAddress = getEnv("GW_SERVICE_USER_AUTHORIZATION_GRPC_ADDRESS")

	timeoutHttpShutdown := getEnv("GW_TIMEOUT_HTTP_SHUTDOWN")

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
		config.applicationName = "API Gateway"
	}

	if len(config.serverAddress) < 1 {
		config.serverAddress = "localhost:8080"
	}

	if len(config.userDataManagerGRPCAddress) < 1 {
		config.userDataManagerGRPCAddress = "localhost:3200"
	}

	if len(config.userAuthorizationGRPCAddress) < 1 {
		config.userAuthorizationGRPCAddress = "localhost:3202"
	}

	if config.timeoutHttpShutdown == 0 {
		config.timeoutHttpShutdown = 10
	}

	return config
}
