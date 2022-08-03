package http

import (
	"api-gw/internal/domain"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func (h *httpHandler) HandlePatchLoginPasswordById(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))
	id := chi.URLParam(r, "id")

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		info := "error patching login-password by id: " + err.Error()
		h.logger.Error(info, ". id=", id, ". userId=", userId)
		w.Write([]byte(info))
		return
	}

	loginPassword := domain.LoginPassword{}
	err = json.Unmarshal(bodyJSON, &loginPassword)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		info := "error patching login-password by wrong request: " + err.Error()
		h.logger.Error(info, bodyJSON)
		w.Write([]byte(info))
		return
	}

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).PatchLoginPasswordById(
		r.Context(),
		&pb.PatchLoginPasswordByIdRequest{
			LoginPassword: &pb.LoginPassword{
				Name:     loginPassword.Name,
				Login:    loginPassword.Login,
				Password: loginPassword.Password,
				Note:     loginPassword.Note,
			},
		},
	)

	if err != nil {
		info := "error patching login-password by id: " + err.Error()
		h.logger.Error(info, "id=", id, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful patching login-password by id ", id, userId, response.LoginPassword.Name)
	result, err := json.Marshal(response)

	if err != nil {
		info := "error marshalling login-password: " + err.Error()
		h.logger.Error(info, ". id=", id, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
