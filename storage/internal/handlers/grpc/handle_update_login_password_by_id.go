package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) UpdateLoginPassword(ctx context.Context, request *pb.UpdateLoginPasswordByIdRequest) (*pb.UpdateLoginPasswordByIdResponse, error) {
	var response pb.UpdateLoginPasswordByIdResponse

	// @ToDo: replace stub data for real data
	var loginPassword pb.LoginPassword
	loginPassword.Id = "1234567890"
	loginPassword.Name = "Sber card"
	loginPassword.Login = "test@yandex.ru"
	loginPassword.Password = "23432dsf"
	loginPassword.UserId = "324"
	loginPassword.CreatedDate = &timestamp.Timestamp{}
	loginPassword.LastAccess = &timestamp.Timestamp{}
	response.LoginPassword = &loginPassword

	s.logger.Info("successful updated login password. ", request)

	return &response, nil
}
