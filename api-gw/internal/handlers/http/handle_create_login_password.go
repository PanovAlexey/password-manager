package http

import (
	"api-gw/internal/handlers/http/dto"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func (h *httpHandler) HandleCreateLoginPassword(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))
	id := chi.URLParam(r, "id")

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("error creating login-password by id: "+err.Error(), id, userId)
		return
	}

	createLoginPasswordDto := dto.CreateLoginPassword{}
	err = json.Unmarshal(bodyJSON, &createLoginPasswordDto)

	if err != nil ||
		len(createLoginPasswordDto.Name) == 0 ||
		len(createLoginPasswordDto.Login) == 0 ||
		len(createLoginPasswordDto.Note) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("error creating login-password by wrong request: "+err.Error(), bodyJSON)
		return
	}

	createLoginPassword := pb.CreateLoginPassword{
		Name:     createLoginPasswordDto.Name,
		Login:    createLoginPasswordDto.Name,
		Password: createLoginPasswordDto.Password,
		Note:     createLoginPasswordDto.Note,
	}

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).CreateLoginPassword(
		r.Context(),
		&pb.CreateLoginPasswordRequest{
			CreateLoginPassword: &createLoginPassword,
			UserId:              userId,
		},
	)

	if err != nil {
		h.logger.Error("error creating login-password by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful creating login-password by id ", id, userId, response.LoginPassword.Name)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling login-password: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
