package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (s *UserDataManagerHandler) DeleteLoginPasswordById(ctx context.Context, request *pb.DeleteLoginPasswordByIdRequest) (*emptypb.Empty, error) {
	s.logger.Info("successful deleted login-password by id. ", request)
	// @ToDo handle error
	return &emptypb.Empty{}, errors.New("test error")
}
