package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) CreateBinaryRecord(ctx context.Context, request *pb.CreateBinaryRecordRequest) (*pb.CreateBinaryRecordResponse, error) {
	var response pb.CreateBinaryRecordResponse

	binaryRecord, err := h.userDataService.AddBinaryRecord(*request.BinaryRecord, ctx)

	if err != nil {
		h.logger.Info("creating binary record error. "+err.Error(), request)

		return nil, err
	}

	response.BinaryRecord = &binaryRecord

	h.logger.Info("successful created binary record. " + binaryRecord.Id)

	return &response, nil
}
