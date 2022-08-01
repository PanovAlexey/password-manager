package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var response pb.UpdateUserResponse

	// @ToDo: replace stub data for real data
	var user pb.User
	user.Id = "1234567890"
	user.Email = "test@gmail.com"
	user.RegistrationDate = &timestamp.Timestamp{}
	user.LastLogin = &timestamp.Timestamp{}
	response.User = &user

	s.logger.Info("successful updated user. ", request)

	return &response, nil
}
