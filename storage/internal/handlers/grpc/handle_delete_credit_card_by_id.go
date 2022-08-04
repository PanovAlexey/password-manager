package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) DeleteCreditCardById(ctx context.Context, request *pb.DeleteCreditCardByIdRequest) (*emptypb.Empty, error) {
	userId := h.userMetadataFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	err := h.creditCardService.DeleteCreditCard(request.Id, userId)

	if err != nil {
		return nil, errors.New("credit card deleting by id error: " + err.Error())
	}

	h.logger.Info("successful deleted credit card by id. ", ". traceId="+traceId)

	return &emptypb.Empty{}, nil
}
