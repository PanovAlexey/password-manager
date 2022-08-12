//go:build prod

package main

import (
	"client/internal/application"
	"client/internal/config"
	"client/internal/handlers/cli"
	"client/internal/infrastructure/http"
	"fmt"
)

var (
	Version   string
	BuildTime string
)

func main() {
	fmt.Printf("version=%s, time=%s\n", Version, BuildTime)

	config := config.New()

	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())

	userRepository := http.GetUserRepository(client)
	userService := application.GetUserService(userRepository)

	userDataRepository := http.GetUserDataRepository(client)
	userDataService := application.GetUserDataService(userDataRepository)

	consoleCommandHandler := cli.GetConsoleCommandHandler(config, userService, userDataService)
	consoleCommandHandler.RunDialog()
}
