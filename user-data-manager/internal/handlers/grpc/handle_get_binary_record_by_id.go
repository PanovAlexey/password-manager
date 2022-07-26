package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) GetBinaryRecordById(
	ctx context.Context, request *pb.GetBinaryRecordByIdRequest,
) (*pb.GetBinaryRecordByIdResponse, error) {
	var response pb.GetBinaryRecordByIdResponse
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	binaryRecord, err := h.userDataService.GetBinaryRecordById(request.Id, ctx)

	if err != nil {
		h.logger.Info("getting binary record by id error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.BinaryRecord = &binaryRecord

	h.logger.Info("successful got binary record by id. ", request.Id, ". traceId="+traceId)

	return &response, nil
}
