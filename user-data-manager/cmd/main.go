package main

import (
	"user-data-manager/internal/application/service"
	"user-data-manager/internal/config"
	grpcHandler "user-data-manager/internal/handlers/grpc"
	"user-data-manager/internal/infrastructure/clients/grpc"
	"user-data-manager/internal/infrastructure/logging"
	"user-data-manager/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)
	defer logger.Sync()

	userMetadataFromContextGetterService := service.GetUserMetadataFromContextGetterService()
	storageClient, err := grpc.GetStorageClient(config)

	if err != nil {
		logger.Error("error getting storage client: " + err.Error())
	}

	defer storageClient.GetConnection().Close()

	userDataService := service.GetUserDataService(storageClient)
	userDataManagerHandler := grpcHandler.GetUserDataManagerHandler(logger, userDataService, userMetadataFromContextGetterService)

	servers.RunGrpcServer(config, logger, userDataManagerHandler)
}
