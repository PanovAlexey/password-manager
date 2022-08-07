package main

import (
	"client/internal/application"
	"client/internal/config"
	"client/internal/handlers/cli"
	"client/internal/infrastructure/http"
)

func main() {
	config := config.New()

	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())

	userRepository := http.GetUserRepository(client)
	userService := application.GetUserService(userRepository)

	userDataRepository := http.GetUserDataRepository(client)
	userDataService := application.GetUserDataService(userDataRepository)

	consoleCommandHandler := cli.GetConsoleCommandHandler(config, userService, userDataService)
	consoleCommandHandler.RunDialog()
}
