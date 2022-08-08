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

func (h *httpHandler) HandleCreateBinaryRecord(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))
	id := chi.URLParam(r, "id")

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("error creating binary record by id: "+err.Error(), id, userId)
		return
	}

	createBinaryRecordDto := dto.CreateBinaryRecord{}
	err = json.Unmarshal(bodyJSON, &createBinaryRecordDto)

	if err != nil ||
		len(createBinaryRecordDto.Name) == 0 ||
		len(createBinaryRecordDto.Binary) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("error creating binary record by wrong request: "+err.Error(), bodyJSON)
		return
	}

	createBinaryRecord := pb.BinaryRecord{
		Name:   createBinaryRecordDto.Name,
		Binary: createBinaryRecordDto.Binary,
		Note:   createBinaryRecordDto.Note,
	}

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).CreateBinaryRecord(
		r.Context(),
		&pb.CreateBinaryRecordRequest{
			BinaryRecord: &createBinaryRecord,
		},
	)

	if err != nil {
		h.logger.Error("error creating binary record by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful creating binary record by id ", id, userId, response.BinaryRecord.Name)

	binaryRecord := domain.BinaryRecord{
		Id:           response.BinaryRecord.Id,
		Name:         response.BinaryRecord.Name,
		Binary:       response.BinaryRecord.Binary,
		Note:         response.BinaryRecord.Note,
		CreatedAt:    response.BinaryRecord.CreatedDate.AsTime().Format(time.RFC3339),
		LastAccessAt: response.BinaryRecord.LastAccess.AsTime().Format(time.RFC3339),
	}

	result, err := json.Marshal(binaryRecord)

	if err != nil {
		h.logger.Error("error marshalling binary record: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
