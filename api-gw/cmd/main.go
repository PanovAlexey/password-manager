package main

import (
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

	userDataManagerClient, err := grpc.GetUserDataManagerClient(config)
	defer userDataManagerClient.GetConnection().Close()

	if err != nil {
		logger.Error("error getting user data manager client: " + err.Error())
	}

	httpHandler := http.GetHTTPHandler(userDataManagerClient, logger)

	servers.RunHttpServer(httpHandler, config, logger)
}
