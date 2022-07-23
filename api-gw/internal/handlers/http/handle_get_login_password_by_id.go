package http

import (
	pb "api-gw/pkg/user_data_manager_grpc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *httpHandler) HandleGetLoginPasswordById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))
	id := chi.URLParam(r, "id")

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetLoginPasswordById(
		context.Background(), &pb.GetLoginPasswordByIdRequest{
			Id:     id,
			UserId: userId,
		},
	)

	if err != nil {
		h.logger.Error("error getting login-password by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful getting login-password by id successful ", id, userId, response.LoginPassword.Name)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling login-password: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
