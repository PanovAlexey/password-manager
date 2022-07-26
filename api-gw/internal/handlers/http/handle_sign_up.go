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

func (h *httpHandler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))

	// @ToDO: disable re-registration if userId exists and correct.
	if len(userId) > 0 {

	}

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("error registration: " + err.Error())
		return
	}

	registerUserDto := dto.RegisterUser{}
	err = json.Unmarshal(bodyJSON, &registerUserDto)

	if err != nil || // @ToDo: move validation from handler
		len(registerUserDto.Email) == 0 ||
		len(registerUserDto.Password) == 0 ||
		len(registerUserDto.RepeatPassword) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("error user registration by wrong request: "+err.Error(), bodyJSON)
		return
	}

	registerUser := pb.RegisterUser{
		Email:          registerUserDto.Email,
		Password:       registerUserDto.Password,
		RepeatPassword: registerUserDto.RepeatPassword,
	}

	response, err := (*h.gRPCUserAuthorizationClient.GetClient()).Register(
		context.Background(), &pb.RegisterRequest{
			RegisterUser: &registerUser,
		},
	)

	if err != nil {
		h.logger.Error("error user registration: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful user registration ", response.User.Id, response.User.Email)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling user: "+err.Error(), response.User.Id, response.User.Email)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
