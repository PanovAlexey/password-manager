package http

import (
	"api-gw/internal/handlers/http/dto"
	pb "api-gw/pkg/user_authorization_grpc"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *httpHandler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))

	// @ToDO: disable re-auth if userId exists and correct.
	if len(userId) > 0 {

	}

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("error auth: " + err.Error())
		return
	}

	authUserDto := dto.AuthUser{}
	err = json.Unmarshal(bodyJSON, &authUserDto)

	if err != nil || // @ToDo: move validation from handler
		len(authUserDto.Email) == 0 ||
		len(authUserDto.Password) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("error user auth by wrong request: "+err.Error(), bodyJSON)
		return
	}

	authUser := pb.AuthUser{
		Email:    authUserDto.Email,
		Password: authUserDto.Password,
	}

	response, err := (*h.gRPCUserAuthorizationClient.GetClient()).Auth(
		context.Background(), &pb.AuthRequest{
			AuthUser: &authUser,
		},
	)

	if err != nil {
		h.logger.Error("error user auth: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful user auth ", response.User.Id, response.User.Email)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling user: "+err.Error(), response.User.Id, response.User.Email)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
