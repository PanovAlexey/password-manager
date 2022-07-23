package config

import "os"

var configSingleton *Config

type Config struct {
	applicationName            string
	serverAddress              string
	userDataManagerGRPCAddress string
}

func New() Config {
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

func initConfigByEnv(config Config) Config {
	config.applicationName = getEnv("APPLICATION_NAME")
	config.serverAddress = getEnv("SERVER_ADDRESS")
	config.userDataManagerGRPCAddress = getEnv("SERVICE_USER_DATA_MANAGER_GRPC_ADDRESS")

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

	return config
}
