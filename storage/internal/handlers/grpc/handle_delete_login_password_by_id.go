package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) DeleteLoginPasswordById(ctx context.Context, request *pb.DeleteLoginPasswordByIdRequest) (*emptypb.Empty, error) {
	userId := h.userMetadataFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	err := h.loginPasswordService.DeleteLoginPassword(request.Id, userId)

	if err != nil {
		return nil, errors.New("login password deleting by id error: " + err.Error())
	}

	h.logger.Info("successful deleted login-password by id. ", ". traceId="+traceId)

	return &emptypb.Empty{}, nil
}
