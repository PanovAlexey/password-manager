package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (s *UserAuthorizationHandler) GetUserByLogin(
	ctx context.Context, request *pb.GetUserByLoginRequest,
) (*pb.GetUserByLoginResponse, error) {
	var response pb.GetUserByLoginResponse

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

	s.logger.Info("successful got user by login. ", request)

	return &response, nil
}
