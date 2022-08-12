package http

import (
	customErrors "api-gw/internal/application/errors"
	"api-gw/internal/application/service"
	"api-gw/internal/domain"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func (h *httpHandler) HandleGetTextRecordById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))
	id := chi.URLParam(r, "id")

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetTextRecordById(
		r.Context(),
		&pb.GetTextRecordByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		if errors.As(err, &customErrors.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			h.showError(w, "not found: "+err.Error())
			return
		}

		h.logger.Error("error getting text record by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful getting text record by id ", id, userId, response.TextRecord.Name)

	textRecord := domain.TextRecord{
		Id:           response.TextRecord.Id,
		Name:         response.TextRecord.Name,
		Text:         response.TextRecord.Text,
		Note:         response.TextRecord.Note,
		CreatedAt:    response.TextRecord.CreatedDate.AsTime().Format(time.RFC3339),
		LastAccessAt: response.TextRecord.LastAccess.AsTime().Format(time.RFC3339),
	}

	result, err := json.Marshal(textRecord)

	if err != nil {
		h.logger.Error("error marshalling text record: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
