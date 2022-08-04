package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) CreateTextRecord(ctx context.Context, request *pb.CreateTextRecordRequest) (*pb.CreateTextRecordResponse, error) {
	var response pb.CreateTextRecordResponse

	textRecord, err := h.userDataService.AddTextRecord(*request.TextRecord, ctx)

	if err != nil {
		h.logger.Info("creating text record error. "+err.Error(), request)

		return nil, err
	}

	response.TextRecord = &textRecord

	h.logger.Info("successful created text record. " + textRecord.Id)

	return &response, nil
}
