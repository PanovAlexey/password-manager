package main

import (
	"user-data-manager/internal/application/service"
	"user-data-manager/internal/config"
	grpcHandler "user-data-manager/internal/handlers/grpc"
	"user-data-manager/internal/infrastructure/logging"
	"user-data-manager/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)

	defer logger.Sync()

	userAuthorizationService := service.GetUserAuthorizationService()
	userDataService := service.GetUserDataService(userAuthorizationService)

	userDataManagerHandler := grpcHandler.GetUserDataManagerHandler(logger, userDataService)

	servers.RunGrpcServer(config, logger, userDataManagerHandler)
}
