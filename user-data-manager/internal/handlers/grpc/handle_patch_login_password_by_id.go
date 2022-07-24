package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (s *UserDataManagerHandler) PatchLoginPasswordById(ctx context.Context, request *pb.PatchLoginPasswordByIdRequest) (*pb.PatchLoginPasswordByIdResponse, error) {
	var response pb.PatchLoginPasswordByIdResponse

	// @ToDo: replace stub data for real data
	var loginPassword pb.LoginPassword
	loginPassword.Id = "444444"
	loginPassword.Note = "Note text etc for example"
	loginPassword.Name = "Stub 4 login password for vk.com"
	loginPassword.Login = "login"
	loginPassword.Password = "pass"
	loginPassword.CreatedDate = &timestamp.Timestamp{}
	loginPassword.LastAccess = &timestamp.Timestamp{}
	response.LoginPassword = &loginPassword

	s.logger.Info("successful patched login-password by id ", request)

	return &response, nil
}
