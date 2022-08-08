package http

import (
	"api-gw/internal/application/service"
	"api-gw/internal/domain"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (h *httpHandler) HandleGetCreditCardList(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetCreditCardList(
		r.Context(),
		&pb.GetCreditCardListRequest{},
	)

	if err != nil {
		info := "error getting credit card list: " + err.Error()
		h.logger.Error(info, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful getting credit card list by user id ", userId)

	var collection []domain.ProtectedItem

	for _, item := range response.ProtectedItemList {
		collection = append(
			collection,
			domain.ProtectedItem{
				Id:           item.Id,
				Name:         item.Name,
				LastAccessAt: item.LastAccess.AsTime().Format(time.RFC3339),
				CreatedAt:    item.CreatedDate.AsTime().Format(time.RFC3339),
			},
		)
	}

	result, err := json.Marshal(collection)

	if err != nil {
		info := "error marshalling credit card list: " + err.Error()
		h.logger.Error(info, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
