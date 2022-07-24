package http

import (
	pb "api-gw/pkg/user_data_manager_grpc"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *httpHandler) HandleDeleteCreditCardById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))
	id := chi.URLParam(r, "id")

	_, err := (*h.gRPCUserDataManagerClient.GetClient()).DeleteCreditCardById(
		context.Background(), &pb.DeleteCreditCardByIdRequest{
			Id:     id,
			UserId: userId,
		},
	)

	if err != nil {
		h.logger.Error("error deleting credit card by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful deleting credit card by id ", id, userId)
	w.WriteHeader(http.StatusNoContent)
}
