package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) DeleteCreditCardById(ctx context.Context, request *pb.DeleteCreditCardByIdRequest) (*emptypb.Empty, error) {
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	err := h.userDataService.DeleteCreditCardById(request.Id, ctx)

	if err != nil {
		h.logger.Info("getting credit card by id error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	h.logger.Info("successful deleted credit card by id. ", request.Id, ". traceId="+traceId)

	return &emptypb.Empty{}, nil
}
