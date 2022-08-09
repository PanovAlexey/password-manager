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

func (h *httpHandler) HandleGetBinaryRecordById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))
	id := chi.URLParam(r, "id")

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).GetBinaryRecordById(
		r.Context(),
		&pb.GetBinaryRecordByIdRequest{
			Id: id,
		},
	)

	if err != nil {
		if errors.As(err, &customErrors.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			h.showError(w, "not found: "+err.Error())
			return
		}

		h.logger.Error("error getting binary record by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful getting binary record by id ", id, userId, response.BinaryRecord.Name)

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
