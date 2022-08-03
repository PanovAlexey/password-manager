package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) DeleteTextRecordById(ctx context.Context, request *pb.DeleteTextRecordByIdRequest) (*emptypb.Empty, error) {
	h.logger.Info("successful deleted text record by id. ", request)
	// @ToDo handle error
	return &emptypb.Empty{}, errors.New("test error")
}
