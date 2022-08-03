package http

import (
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *httpHandler) HandleGetLoginPasswordById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("user-id"))
	id := chi.URLParam(r, "id")

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetLoginPasswordById(
		r.Context(),
		&pb.GetLoginPasswordByIdRequest{
			Id:     id,
			UserId: userId,
		},
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		info := "error getting login-password by id: " + err.Error()
		h.logger.Error(info, id, userId)
		w.Write([]byte(info))
		return
	}

	result, err := json.Marshal(response)

	if err != nil {
		info := "error marshalling login-password: " + err.Error()
		h.logger.Error(info, id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("login password was gotten by id=", id, " and userId=", userId)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
