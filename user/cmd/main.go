package main

import (
	"user-auth/internal/application/service"
	"user-auth/internal/config"
	"user-auth/internal/handlers/grpc"
	clientsGrpc "user-auth/internal/infrastructure/clients/grpc"
	"user-auth/internal/infrastructure/logging"
	"user-auth/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)
	defer logger.Sync()

	jwtAuthorizationService := service.GetJWTAuthorizationService()
	storageClient, err := clientsGrpc.GetStorageClient(config)

	if err != nil {
		logger.Error("error getting storage client: " + err.Error())
	}

	defer storageClient.GetConnection().Close()

	userRegistrationService := service.GetUserRegistrationService(storageClient)
	userMetadataFromContextGetterService := service.GetUserMetadataFromContextGetterService()
	handler := grpc.GetUserAuthorizationHandler(
		logger,
		jwtAuthorizationService,
		userRegistrationService,
		userMetadataFromContextGetterService,
	)

	servers.RunGrpcServer(config, logger, handler)
}
