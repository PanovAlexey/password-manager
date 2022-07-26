package main

import (
	"user-auth/internal/config"
	"user-auth/internal/infrastructure/logging"
	"user-auth/internal/servers"
)

func main() {
	config := config.New()
	logger := logging.GetLogger(config)

	defer logger.Sync()

	servers.RunGrpcServer(config, logger)
}
