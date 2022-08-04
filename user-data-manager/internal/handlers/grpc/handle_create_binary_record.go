package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) CreateBinaryRecord(ctx context.Context, request *pb.CreateBinaryRecordRequest) (*pb.CreateBinaryRecordResponse, error) {
	var response pb.CreateBinaryRecordResponse

	userid := h.userMetadataFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	binaryRecord, err := h.userDataService.AddBinaryRecord(*request.BinaryRecord, userid, ctx)

	if err != nil {
		h.logger.Info("creating binary record error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.BinaryRecord = &binaryRecord

	h.logger.Info("successful created binary record. "+binaryRecord.Id, ". traceId="+traceId)

	return &response, nil
}
