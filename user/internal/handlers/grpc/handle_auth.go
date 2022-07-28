package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (h *UserAuthorizationHandler) Auth(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	var response pb.AuthResponse

	// @ToDo: replace stub data for real data from storage
	var user pb.User
	user.Id = "1231488"
	user.RegistrationDate = &timestamp.Timestamp{}
	user.LastLogin = &timestamp.Timestamp{}
	user.Email = "test@ya.ru"
	response.User = &user

	token, err := h.jwtAuthorizationService.GetJWTToken(user.Id)

	if err != nil {
		h.logger.Error("getting JWT token by user id error: "+err.Error(), user.Id)
	}

	user.Token = token

	h.logger.Info("successful auth. ", request)

	return &response, nil
}
