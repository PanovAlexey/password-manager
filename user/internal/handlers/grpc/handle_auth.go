package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (s *UserAuthorizationHandler) Auth(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	var response pb.AuthResponse

	// @ToDo: replace stub data for real data
	var user pb.User
	user.Id = "123"
	user.Token = "bearer token 234324324"
	user.RegistrationDate = &timestamp.Timestamp{}
	user.LastLogin = &timestamp.Timestamp{}
	user.Email = "test@ya.ru"
	response.User = &user

	s.logger.Info("successful auth. ", request)

	return &response, nil
}
