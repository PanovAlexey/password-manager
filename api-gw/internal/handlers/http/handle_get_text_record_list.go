package http

import (
	pb "api-gw/pkg/user_data_manager_grpc"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *httpHandler) HandleGetTextRecordList(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetTextRecordList(
		context.Background(),
		&pb.GetTextRecordListRequest{UserId: userId},
	)

	if err != nil {
		h.logger.Error("error getting text record list: "+err.Error(), userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful getting text record list by user id ", userId, response)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling text record list: "+err.Error(), userId, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
