package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) DeleteBinaryRecordById(ctx context.Context, request *pb.DeleteBinaryRecordByIdRequest) (*emptypb.Empty, error) {
	h.logger.Info("successful deleted binary record by id. ", request)
	// @ToDo handle error
	return &emptypb.Empty{}, errors.New("test error")
}
