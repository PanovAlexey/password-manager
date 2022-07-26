package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (s *UserAuthorizationHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse

	// @ToDo: replace stub data for real data
	var user pb.User
	var token pb.Token
	user.Id = "123"
	token.Token = "bearer token 234324324"
	user.Token = &token
	user.RegistrationDate = &timestamp.Timestamp{}
	user.LastLogin = &timestamp.Timestamp{}
	user.Email = "test@ya.ru"
	response.User = &user

	s.logger.Info("successful registered. ", request)

	return &response, nil
}
