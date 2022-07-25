package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) DeleteCreditCardById(ctx context.Context, request *pb.DeleteCreditCardByIdRequest) (*emptypb.Empty, error) {
	s.logger.Info("successful deleted credit card by id. ", request)
	// @ToDo handle error
	return &emptypb.Empty{}, errors.New("test error")
}
