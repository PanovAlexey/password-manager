package http

import (
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *httpHandler) HandleGetTextRecordById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))
	id := chi.URLParam(r, "id")

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetTextRecordById(
		r.Context(),
		&pb.GetTextRecordByIdRequest{
			Id:     id,
			UserId: userId,
		},
	)

	if err != nil {
		h.logger.Error("error getting text record by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful getting text record by id ", id, userId, response.TextRecord.Name)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling text record: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
