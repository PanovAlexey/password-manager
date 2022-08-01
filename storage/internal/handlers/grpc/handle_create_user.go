package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var response pb.CreateUserResponse

	// @ToDo: replace stub data for real data
	var user pb.User
	user.Id = "1234567890"
	user.Email = "test@gmail.com"
	user.RegistrationDate = &timestamp.Timestamp{}
	user.LastLogin = &timestamp.Timestamp{}
	response.User = &user

	s.logger.Info("successful created user. ", request)

	return &response, nil
}
