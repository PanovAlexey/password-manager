package http

import (
	"api-gw/internal/application/service"
	"api-gw/internal/handlers/http/middleware/authorization_by_token"
	"api-gw/internal/handlers/http/middleware/closed_by_authorization"
	"api-gw/internal/handlers/http/middleware/json"
	middleware_custom "api-gw/internal/handlers/http/middleware/trace"
	"api-gw/internal/infrastructure/clients/grpc"
	encodingJson "encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type httpHandler struct {
	gRPCUserDataManagerClient grpc.UserDataManagerClient
	logger                    Logger
	userAuthorizationService  service.UserAuthorization
}

func GetHTTPHandler(
	userDataManagerClient grpc.UserDataManagerClient,
	logger Logger,
	userAuthorizationService service.UserAuthorization,
) *httpHandler {
	return &httpHandler{
		gRPCUserDataManagerClient: userDataManagerClient,
		logger:                    logger,
		userAuthorizationService:  userAuthorizationService,
	}
}

func (h *httpHandler) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware_custom.Trace(h.logger))
	router.Use(json.JSON)
	router.Use(authorization_by_token.AuthorizationByToken(h.userAuthorizationService, h.logger))
	router.Use(middleware.Timeout(60 * time.Second)) // @ToDo move 60 to conf

	router.Route("/api/v1/data", func(r chi.Router) {
		r.Use(closed_by_authorization.ClosedByAuthorization(h.userAuthorizationService))

		r.Get("/all", h.HandleGetUserAllData)

		r.Get("/login-password", h.HandleGetLoginPasswordList)
		r.Post("/login-password", h.HandleCreateLoginPassword)
		r.Get("/login-password/{id}", h.HandleGetLoginPasswordById)
		r.Patch("/login-password/{id}", h.HandlePatchLoginPasswordById)
		r.Delete("/login-password/{id}", h.HandleDeleteLoginPasswordById)

		r.Get("/credit-card", h.HandleGetCreditCardList)
		r.Post("/credit-card", h.HandleCreateCreditCard)
		r.Get("/credit-card/{id}", h.HandleGetCreditCardById)
		r.Patch("/credit-card/{id}", h.HandlePatchCreditCardById)
		r.Delete("/credit-card/{id}", h.HandleDeleteCreditCardById)

		r.Get("/text-record", h.HandleGetTextRecordList)
		r.Post("/text-record", h.HandleCreateTextRecord)
		r.Get("/text-record/{id}", h.HandleGetTextRecordById)
		r.Patch("/text-record/{id}", h.HandlePatchTextRecordById)
		r.Delete("/text-record/{id}", h.HandleDeleteTextRecordById)

		r.Get("/binary-record", h.HandleGetBinaryRecordList)
		r.Post("/binary-record", h.HandleCreateBinaryRecord)
		r.Get("/binary-record/{id}", h.HandleGetBinaryRecordById)
		r.Patch("/binary-record/{id}", h.HandlePatchBinaryRecordById)
		r.Delete("/binary-record/{id}", h.HandleDeleteBinaryRecordById)
	})

	router.Get("/api/v1/health/check", h.HandleHealthCheck)

	router.Post("/api/v1/auth", h.HandleAuth)
	router.Put("/api/v1/auth", h.HandleRefreshAuthToken)
	router.Post("/api/v1/signup", h.HandleSignUp)

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

func (h *httpHandler) showError(w http.ResponseWriter, errorText string, args ...interface{}) {
	h.logger.Debug(errorText, args)

	err := encodingJson.NewEncoder(w).Encode(errorText)

	if err != nil {
		h.logger.Error("unknown error while preparing response " + err.Error())
	}
}
