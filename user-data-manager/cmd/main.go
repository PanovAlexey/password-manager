package main

import (
	"user-data-manager/internal/config"
	"user-data-manager/internal/infrastructure/logging"
	"user-data-manager/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)

	defer logger.Sync()

	servers.RunGrpcServer(config, logger)
}
