package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) DeleteTextRecordById(ctx context.Context, request *pb.DeleteTextRecordByIdRequest) (*emptypb.Empty, error) {
	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	err := h.textRecordService.DeleteTextRecord(request.Id, userId)

	if err != nil {
		return nil, errors.New("text record deleting by id error: " + err.Error())
	}

	h.logger.Info("successful deleted text record by id. ", ". traceId="+traceId)

	return &emptypb.Empty{}, nil
}
