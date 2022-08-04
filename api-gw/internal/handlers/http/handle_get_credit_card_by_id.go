package http

import (
	"api-gw/internal/application/service"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *httpHandler) HandleGetCreditCardById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))
	id := chi.URLParam(r, "id")

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetCreditCardById(
		r.Context(),
		&pb.GetCreditCardByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		h.logger.Error("error getting credit card by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful getting credit card by id ", id, userId, response.CreditCard.Name)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling credit card: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
