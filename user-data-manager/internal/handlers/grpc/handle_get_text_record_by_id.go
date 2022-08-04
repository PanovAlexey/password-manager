package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) GetTextRecordById(
	ctx context.Context, request *pb.GetTextRecordByIdRequest,
) (*pb.GetTextRecordByIdResponse, error) {
	var response pb.GetTextRecordByIdResponse
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	textRecord, err := h.userDataService.GetTextRecordById(request.Id, ctx)

	if err != nil {
		h.logger.Info("getting text record by id error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.TextRecord = &textRecord

	h.logger.Info("successful got text record by id. ", request.Id, ". traceId="+traceId)

	return &response, nil
}
