package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (h *UserAuthorizationHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse

	// @ToDo: replace stub data for real data
	var user pb.User
	user.Id = "123"
	user.RegistrationDate = &timestamp.Timestamp{}
	user.LastLogin = &timestamp.Timestamp{}
	user.Email = "test@ya.ru"
	token, err := h.jwtAuthorizationService.GetJWTToken(user.Id)
	user.Token = token

	if err != nil {
		h.logger.Error("getting JWT token by user id error: "+err.Error(), user.Id)
	}

	response.User = &user

	s.logger.Info("successful registered. ", request)

	return &response, nil
}
