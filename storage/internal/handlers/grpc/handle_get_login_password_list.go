package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) GetLoginPasswordList(ctx context.Context, request *pb.GetLoginPasswordListRequest) (*pb.GetLoginPasswordListResponse, error) {
	var response pb.GetLoginPasswordListResponse

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
	response.LoginPasswordList = append(response.LoginPasswordList, &loginPassword)

	s.logger.Info("successful got login-password list. ", request)

	return &response, nil
}
