package http

import (
	middleware_custom "api-gw/internal/handlers/http/middleware"
	"api-gw/internal/infrastructure/clients/grpc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type httpHandler struct {
	gRPCUserDataManagerClient grpc.UserDataManagerClient
	logger                    Logger
}

func GetHTTPHandler(userDataManagerClient grpc.UserDataManagerClient, logger Logger) *httpHandler {
	return &httpHandler{gRPCUserDataManagerClient: userDataManagerClient, logger: logger}
}

func (h *httpHandler) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware_custom.JSON)
	router.Use(middleware_custom.Authorization)

	router.Get("/api/v1/data/login-password", h.HandleGetLoginPasswordList)
	router.Get("/api/v1/data/login-password/{id}", h.HandleGetLoginPasswordById)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	return router
}
