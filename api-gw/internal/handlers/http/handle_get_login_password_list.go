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

func (h *httpHandler) HandleGetLoginPasswordList(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetLoginPasswordList(
		r.Context(),
		&pb.GetLoginPasswordListRequest{},
	)

	if err != nil {
		info := "error getting login-password list: " + err.Error()
		h.logger.Error(info, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

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
		info := "error marshalling login-password list: " + err.Error()
		h.logger.Error(info, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful getting login-password list by userId=", userId)

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
