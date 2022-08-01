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
	jwtAuthorizationService := service.GetJWTAuthorizationService()
	storageClient, err := clientsGrpc.GetStorageClient(config)
	defer storageClient.GetConnection().Close()

	if err != nil {
		logger.Error("error getting storage client: " + err.Error())
	}

	handler := grpc.GetUserAuthorizationHandler(logger, jwtAuthorizationService, storageClient)

	defer logger.Sync()

	servers.RunGrpcServer(config, logger, handler)
}
