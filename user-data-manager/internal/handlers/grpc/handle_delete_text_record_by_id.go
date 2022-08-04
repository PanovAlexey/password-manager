package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) DeleteTextRecordById(
	ctx context.Context,
	request *pb.DeleteTextRecordByIdRequest,
) (*emptypb.Empty, error) {
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	err := h.userDataService.DeleteTextRecordById(request.Id, ctx)

	if err != nil {
		h.logger.Info("getting text record by id error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	h.logger.Info("successful deleted text record by id. ", request.Id, ". traceId="+traceId)

	return &emptypb.Empty{}, nil
}
