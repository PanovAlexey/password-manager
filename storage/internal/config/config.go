package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var configSingleton *Config

type Config struct {
	applicationName  string
	grpcServerPort   string
	databaseUser     string
	databasePassword string
	databasePort     string
	databaseAddress  string
	databaseName     string
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

func (c Config) GetGrpcServerPort() string {
	return c.grpcServerPort
}

func (c Config) GetDatabaseUser() string {
	return c.databaseUser
}

func (c Config) GetDatabasePassword() string {
	return c.databasePassword
}

func (c Config) GetDatabasePort() string {
	return c.databasePort
}

func (c Config) GetDatabaseAddress() string {
	return c.databaseAddress
}

func (c Config) GetDatabaseName() string {
	return c.databaseName
}

func initConfigByEnv(config Config) Config {
	config.applicationName = getEnv("STORAGE_APPLICATION_NAME")
	config.grpcServerPort = getEnv("STORAGE_GRPC_SERVER_PORT")
	config.databasePort = getEnv("STORAGE_DB_MASTER_PORT")
	config.databaseAddress = getEnv("STORAGE_DB_MASTER_ADDRESS")
	config.databaseName = getEnv("STORAGE_DB_MASTER_DATABASE")

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

	if len(config.grpcServerPort) < 1 {
		config.grpcServerPort = "3205"
	}

	if len(config.databaseUser) < 1 {
		config.databaseUser = "pw_user"
	}

	if len(config.databasePassword) < 1 {
		config.databasePassword = "secret"
	}

	if len(config.databasePort) < 1 {
		config.databasePort = "5432"
	}

	if len(config.databaseAddress) < 1 {
		config.databaseAddress = "password-manager-storage-master-postgres"
	}

	if len(config.databaseName) < 1 {
		config.databaseName = "password-manager"
	}

	return config
}
