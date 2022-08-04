package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) CreateTextRecord(ctx context.Context, request *pb.CreateTextRecordRequest) (*pb.CreateTextRecordResponse, error) {
	var response pb.CreateTextRecordResponse
	userid := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	textRecord, err := h.userDataService.AddTextRecord(*request.TextRecord, userid, ctx)

	if err != nil {
		h.logger.Info("creating text record error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.TextRecord = &textRecord

	h.logger.Info("successful created text record. "+textRecord.Id, ". traceId="+traceId)

	return &response, nil
}
