package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) DeleteBinaryRecordById(
	ctx context.Context,
	request *pb.DeleteBinaryRecordByIdRequest,
) (*emptypb.Empty, error) {
	userId := h.userMetadataFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	err := h.binaryRecordService.DeleteBinaryRecord(request.Id, userId)

	if err != nil {
		return nil, errors.New("binary record deleting by id error: " + err.Error())
	}

	h.logger.Info("successful deleted binary record by id. ", ". traceId="+traceId)

	return &emptypb.Empty{}, nil
}
