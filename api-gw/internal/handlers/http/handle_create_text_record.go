package http

import (
	"api-gw/internal/application/service"
	"api-gw/internal/domain"
	"api-gw/internal/handlers/http/dto"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"time"
)

func (h *httpHandler) HandleCreateTextRecord(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))
	id := chi.URLParam(r, "id")

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("error creating text record by id: "+err.Error(), id, userId)
		return
	}

	createTextRecordDto := dto.CreateTextRecord{}
	err = json.Unmarshal(bodyJSON, &createTextRecordDto)

	if err != nil ||
		len(createTextRecordDto.Name) == 0 ||
		len(createTextRecordDto.Text) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("error creating text record by wrong request: "+err.Error(), bodyJSON)
		return
	}

	createTextRecord := pb.TextRecord{
		Name: createTextRecordDto.Name,
		Text: createTextRecordDto.Text,
		Note: createTextRecordDto.Note,
	}

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).CreateTextRecord(
		r.Context(),
		&pb.CreateTextRecordRequest{
			TextRecord: &createTextRecord,
		},
	)

	if err != nil {
		h.logger.Error("error creating text record by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful creating text record by id ", id, userId, response.TextRecord.Name)

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
