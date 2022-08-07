package http

import (
	"api-gw/internal/application/service"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *httpHandler) HandleGetBinaryRecordList(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetBinaryRecordList(
		r.Context(),
		&pb.GetBinaryRecordListRequest{},
	)

	if err != nil {
		info := "error getting binary record list: " + err.Error()
		h.logger.Error(info, userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful getting binary record list by user id ", userId, response)
	result, err := json.Marshal(response.ProtectedItemList)

	if err != nil {
		info := "error marshalling binary record list: " + err.Error()
		h.logger.Error(info, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
