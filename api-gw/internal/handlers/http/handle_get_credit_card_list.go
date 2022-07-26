package http

import (
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *httpHandler) HandleGetCreditCardList(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetCreditCardList(
		r.Context(),
		&pb.GetCreditCardListRequest{UserId: userId},
	)

	if err != nil {
		h.logger.Error("error getting credit card list: "+err.Error(), userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful getting credit card list by user id ", userId, response)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling credit card list: "+err.Error(), userId, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
