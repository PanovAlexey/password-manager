package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) DeleteBinaryRecordById(
	ctx context.Context,
	request *pb.DeleteBinaryRecordByIdRequest,
) (*emptypb.Empty, error) {
	err := h.userDataService.DeleteBinaryRecordById(request.Id, ctx)
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)

	if err != nil {
		h.logger.Info("getting binary record by id error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	h.logger.Info("successful deleted binary record by id. ", request.Id, ". traceId="+traceId)

	return &emptypb.Empty{}, nil
}
