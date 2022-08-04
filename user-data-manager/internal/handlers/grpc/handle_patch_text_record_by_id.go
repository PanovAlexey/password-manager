package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) PatchTextRecordById(
	ctx context.Context,
	request *pb.PatchTextRecordByIdRequest,
) (*pb.PatchTextRecordByIdResponse, error) {
	var response pb.PatchTextRecordByIdResponse
	userid := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	textRecord, err := h.userDataService.UpdateTextRecord(*request.TextRecord, userid, ctx)

	if err != nil {
		h.logger.Info("updating text record error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.TextRecord = &textRecord
	h.logger.Info("successful updated text record. "+request.TextRecord.Id, ". traceId="+traceId)

	return &response, nil
}
