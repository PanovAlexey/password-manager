package main

import (
	_ "github.com/lib/pq"
	services "storage/internal/application/service"
	"storage/internal/config"
	grpcHandler "storage/internal/handlers/grpc"
	"storage/internal/infrastructure/logging"
	databases "storage/internal/infrastructure/postgresql"
	"storage/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)

	defer logger.Sync()

	databaseService, err := databases.GetDatabaseService(config)

	if err != nil {
		logger.Error(`Unable to connect to database: %v\n`, err)
	}

	defer databaseService.GetDatabaseConnection()

	databaseUserRepository := databases.GetUserRepository(databaseService.GetDatabaseConnection())
	databaseUserService := services.GetUserService(databaseUserRepository)

	handler := grpcHandler.GetStorageHandler(logger, databaseUserService)

	servers.RunGrpcServer(config, logger, handler)
}
