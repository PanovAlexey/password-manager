package http

import (
	customErrors "api-gw/internal/application/errors"
	"api-gw/internal/application/service"
	"api-gw/internal/domain"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func (h *httpHandler) HandleGetLoginPasswordById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))
	id := chi.URLParam(r, "id")

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetLoginPasswordById(
		r.Context(),
		&pb.GetLoginPasswordByIdRequest{Id: id},
	)

	if err != nil {
		if errors.As(err, &customErrors.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			h.showError(w, "not found: "+err.Error())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		h.showError(w, "error getting login-password by id: "+err.Error())
		return
	}

	loginPassword := domain.LoginPassword{
		Id:           response.LoginPassword.Id,
		Name:         response.LoginPassword.Name,
		Login:        response.LoginPassword.Login,
		Password:     response.LoginPassword.Password,
		Note:         response.LoginPassword.Note,
		CreatedAt:    response.LoginPassword.CreatedDate.AsTime().Format(time.RFC3339),
		LastAccessAt: response.LoginPassword.LastAccess.AsTime().Format(time.RFC3339),
	}

	result, err := json.Marshal(loginPassword)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.showError(w, "error marshalling login-password: "+err.Error())
		return
	}

	h.logger.Info("login password was gotten by id=", id, " and userId=", userId)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
