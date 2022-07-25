package main

import (
	"storage/internal/config"
	"storage/internal/infrastructure/logging"
	"storage/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)

	defer logger.Sync()

	servers.RunGrpcServer(config, logger)
}
