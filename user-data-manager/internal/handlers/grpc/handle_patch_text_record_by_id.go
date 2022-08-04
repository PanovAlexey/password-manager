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
	textRecord, err := h.userDataService.UpdateTextRecord(*request.TextRecord, ctx)

	if err != nil {
		h.logger.Info("updating text record error. "+err.Error(), request)

		return nil, err
	}

	response.TextRecord = &textRecord
	h.logger.Info("successful updated text record. " + request.TextRecord.Id)

	return &response, nil
}
