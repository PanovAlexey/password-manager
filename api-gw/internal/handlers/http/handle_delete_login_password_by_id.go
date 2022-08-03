package http

import (
	pb "api-gw/pkg/user_data_manager_grpc"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *httpHandler) HandleDeleteLoginPasswordById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))
	id := chi.URLParam(r, "id")

	_, err := (*h.gRPCUserDataManagerClient.GetClient()).DeleteLoginPasswordById(
		r.Context(),
		&pb.DeleteLoginPasswordByIdRequest{
			Id:     id,
			UserId: userId,
		},
	)

	if err != nil {
		info := "error deleting login-password by id: " + err.Error()
		h.logger.Error(info, ". id=", id, ".userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful deleting login-password by id ", id, userId)
	w.WriteHeader(http.StatusNoContent)
}
