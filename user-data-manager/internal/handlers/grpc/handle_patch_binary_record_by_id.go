package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) PatchBinaryRecordById(
	ctx context.Context,
	request *pb.PatchBinaryRecordByIdRequest,
) (*pb.PatchBinaryRecordByIdResponse, error) {
	var response pb.PatchBinaryRecordByIdResponse
	binaryRecord, err := h.userDataService.UpdateBinaryRecord(*request.BinaryRecord, ctx)

	if err != nil {
		h.logger.Info("updating binary record error. "+err.Error(), request)

		return nil, err
	}

	response.BinaryRecord = &binaryRecord
	h.logger.Info("successful updated binary record. " + request.BinaryRecord.Id)

	return &response, nil
}
