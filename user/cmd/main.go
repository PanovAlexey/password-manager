package main

import (
	"user-auth/internal/application/service"
	"user-auth/internal/config"
	"user-auth/internal/handlers/grpc"
	"user-auth/internal/infrastructure/logging"
	"user-auth/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)
	jwtAuthorizationService := service.GetJWTAuthorizationService()
	handler := grpc.GetUserAuthorizationHandler(logger, jwtAuthorizationService)

	defer logger.Sync()

	servers.RunGrpcServer(config, logger, handler)
}
