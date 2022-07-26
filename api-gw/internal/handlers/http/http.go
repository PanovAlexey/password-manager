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
	gRPCUserAuthorizationClient grpc.UserAuthorizationClient
	gRPCUserDataManagerClient   grpc.UserDataManagerClient
	logger                      Logger
}

func GetHTTPHandler(
	userAuthorizationClient grpc.UserAuthorizationClient,
	userDataManagerClient grpc.UserDataManagerClient,
	logger Logger,
) *httpHandler {
	return &httpHandler{
		gRPCUserAuthorizationClient: userAuthorizationClient,
		gRPCUserDataManagerClient:   userDataManagerClient,
		logger:                      logger,
	}
}

func (h *httpHandler) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware_custom.Trace(h.logger))
	router.Use(middleware_custom.JSON)
	router.Use(middleware_custom.Authorization)

	router.Get("/api/v1/health/check", h.HandleHealthCheck)

	router.Post("/api/v1/auth", h.HandleAuth)
	router.Put("/api/v1/auth", h.HandleRefreshAuthToken)
	router.Post("/api/v1/signup", h.HandleSignUp)

	router.Get("/api/v1/data/all", h.HandleGetUserAllData)

	router.Get("/api/v1/data/login-password", h.HandleGetLoginPasswordList)
	router.Post("/api/v1/data/login-password", h.HandleCreateLoginPassword)
	router.Get("/api/v1/data/login-password/{id}", h.HandleGetLoginPasswordById)
	router.Patch("/api/v1/data/login-password/{id}", h.HandlePatchLoginPasswordById)
	router.Delete("/api/v1/data/login-password/{id}", h.HandleDeleteLoginPasswordById)

	router.Get("/api/v1/data/credit-card", h.HandleGetCreditCardList)
	router.Post("/api/v1/data/credit-card", h.HandleCreateCreditCard)
	router.Get("/api/v1/data/credit-card/{id}", h.HandleGetCreditCardById)
	router.Patch("/api/v1/data/credit-card/{id}", h.HandlePatchCreditCardById)
	router.Delete("/api/v1/data/credit-card/{id}", h.HandleDeleteCreditCardById)

	router.Get("/api/v1/data/text-record", h.HandleGetTextRecordList)
	router.Post("/api/v1/data/text-record", h.HandleCreateTextRecord)
	router.Get("/api/v1/data/text-record/{id}", h.HandleGetTextRecordById)
	router.Patch("/api/v1/data/text-record/{id}", h.HandlePatchTextRecordById)
	router.Delete("/api/v1/data/text-record/{id}", h.HandleDeleteTextRecordById)

	router.Get("/api/v1/data/binary-record", h.HandleGetBinaryRecordList)
	router.Post("/api/v1/data/binary-record", h.HandleCreateBinaryRecord)
	router.Get("/api/v1/data/binary-record/{id}", h.HandleGetBinaryRecordById)
	router.Patch("/api/v1/data/binary-record/{id}", h.HandlePatchBinaryRecordById)
	router.Delete("/api/v1/data/binary-record/{id}", h.HandleDeleteBinaryRecordById)

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
