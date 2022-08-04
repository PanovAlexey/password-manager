package http

import (
	"api-gw/internal/application/service"
	"api-gw/internal/handlers/http/dto"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func (h *httpHandler) HandleCreateLoginPassword(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))
	id := chi.URLParam(r, "id")

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		info := "error creating login-password by id: " + err.Error()
		h.logger.Error(info, id, "userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	createLoginPasswordDto := dto.CreateLoginPassword{}
	err = json.Unmarshal(bodyJSON, &createLoginPasswordDto)

	if err != nil ||
		len(createLoginPasswordDto.Name) == 0 ||
		len(createLoginPasswordDto.Login) == 0 ||
		len(createLoginPasswordDto.Note) == 0 {

		info := "error creating login-password by wrong request: " + err.Error()
		h.logger.Error(info, bodyJSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(info))
		return
	}

	createLoginPassword := pb.CreateLoginPassword{
		Name:     createLoginPasswordDto.Name,
		Login:    createLoginPasswordDto.Login,
		Password: createLoginPasswordDto.Password,
		Note:     createLoginPasswordDto.Note,
	}

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).CreateLoginPassword(
		r.Context(),
		&pb.CreateLoginPasswordRequest{
			CreateLoginPassword: &createLoginPassword,
		},
	)

	if err != nil {
		info := "error creating login-password by id: " + err.Error()
		h.logger.Error(info, id, "userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful creating login-password by id ", id, userId, response.LoginPassword.Name)
	result, err := json.Marshal(response)

	if err != nil {
		info := "error marshalling login-password: " + err.Error()
		h.logger.Error(info, id, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
