package main

import (
	"api-gw/internal/application/service"
	"api-gw/internal/config"
	"api-gw/internal/handlers/http"
	"api-gw/internal/infrastructure/clients/grpc"
	"api-gw/internal/infrastructure/logging"
	"api-gw/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)
	defer logger.Sync()

	userAuthorizationClient, err := grpc.GetUserAuthorizationClient(config)
	defer userAuthorizationClient.GetConnection().Close()

	if err != nil {
		logger.Error("error getting user authorization client: " + err.Error())
	}

	userDataManagerClient, err := grpc.GetUserDataManagerClient(config)
	defer userDataManagerClient.GetConnection().Close()

	if err != nil {
		logger.Error("error getting user data manager client: " + err.Error())
	}

	userAuthorizationService := service.GetUserAuthorizationService(logger, userAuthorizationClient)

	httpHandler := http.GetHTTPHandler(userDataManagerClient, logger, userAuthorizationService)

	servers.RunHttpServer(httpHandler, config, logger)
}
