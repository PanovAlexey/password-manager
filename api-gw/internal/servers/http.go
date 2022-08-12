package servers

import (
	"api-gw/internal/config"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type HandlerInterface interface {
	NewRouter() chi.Router
}

type mainHttpServer struct {
	httpServer *http.Server
}

func RunHttpServer(handler HandlerInterface, config config.Config, logger Logger) {
	logger.Info(config.GetApplicationName() + " http server starting...")

	srv := new(mainHttpServer)
	srv.httpServer = &http.Server{
		Addr:    ":" + config.GetServerPort(),
		Handler: handler.NewRouter(),
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info(srv.httpServer.ListenAndServe())
	}()

	logger.Info(config.GetApplicationName() + " http server started")

	logger.Info("Signal detected: ", <-sigs)

	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(config.GetTimeoutHttpShutdown())*time.Second,
	)

	defer cancel()

	err := srv.httpServer.Shutdown(ctx)
	logger.Info(config.GetApplicationName() + " http server is shutdowning...")

	if err != nil {
		logger.Info(config.GetApplicationName() + " http server shutdowning error: " + err.Error())
	}

	logger.Info(config.GetApplicationName() + " http server has been stopped.")
}
