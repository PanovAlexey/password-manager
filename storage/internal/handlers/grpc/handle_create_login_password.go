package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) CreateLoginPassword(ctx context.Context, request *pb.CreateLoginPasswordRequest) (*pb.CreateLoginPasswordResponse, error) {
	var response pb.CreateLoginPasswordResponse

	// @ToDo: replace stub data for real data
	var loginPassword pb.LoginPassword
	loginPassword.Id = "1234567890"
	loginPassword.Note = "Note text etc for example"
	loginPassword.Name = "Stub 2 binary record for vk.com"
	loginPassword.UserId = "234324-324324-32"
	loginPassword.Password = "4242232323232323"
	loginPassword.Login = "alex@gmail.com"
	loginPassword.CreatedDate = &timestamp.Timestamp{}
	loginPassword.LastAccess = &timestamp.Timestamp{}
	response.LoginPassword = &loginPassword

	s.logger.Info("successful created login password. ", request)

	return &response, nil
}
