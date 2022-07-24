package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (s *UserDataManagerHandler) CreateLoginPassword(ctx context.Context, request *pb.CreateLoginPasswordRequest) (*pb.CreateLoginPasswordResponse, error) {
	var response pb.CreateLoginPasswordResponse

	// @ToDo: replace stub data for real data
	var loginPassword pb.LoginPassword
	loginPassword.Id = "1234567890"
	loginPassword.Note = "Note text etc for example"
	loginPassword.Name = "Stub 2 login password for vk.com"
	loginPassword.Login = "login"
	loginPassword.Password = "pass"
	loginPassword.CreatedDate = &timestamp.Timestamp{}
	loginPassword.LastAccess = &timestamp.Timestamp{}
	response.LoginPassword = &loginPassword

	s.logger.Info("successful created login-password. ", request)

	return &response, nil
}
