package http

import (
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *httpHandler) HandleGetBinaryRecordList(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetBinaryRecordList(
		r.Context(),
		&pb.GetBinaryRecordListRequest{},
	)

	if err != nil {
		h.logger.Error("error getting binary record list: "+err.Error(), userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful getting binary record list by user id ", userId, response)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling binary record list: "+err.Error(), userId, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
