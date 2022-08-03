package main

import (
	_ "github.com/lib/pq"
	"storage/internal/application/service"
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
	databaseUserService := service.GetUserService(databaseUserRepository)

	userIdFromContextGetterService := service.GetUserIdFromContextGetterService()

	handler := grpcHandler.GetStorageHandler(logger, databaseUserService, userIdFromContextGetterService)

	servers.RunGrpcServer(config, logger, handler)
}
