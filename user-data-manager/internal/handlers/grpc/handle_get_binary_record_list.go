package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) GetBinaryRecordList(
	ctx context.Context,
	request *pb.GetBinaryRecordListRequest,
) (*pb.GetBinaryRecordListResponse, error) {
	var response pb.GetBinaryRecordListResponse
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	binaryRecordList, err := h.userDataService.GetBinaryRecordList(ctx)

	if err != nil {
		h.logger.Info("getting binary record list error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	for _, binaryRecord := range binaryRecordList {
		protectedItem := pb.ProtectedItem{
			Id:          binaryRecord.Id,
			Name:        binaryRecord.Name,
			CreatedDate: binaryRecord.CreatedDate,
			LastAccess:  binaryRecord.LastAccess,
		}

		response.ProtectedItemList = append(response.ProtectedItemList, &protectedItem)
	}

	h.logger.Info("successful got binary record list. ", ". traceId="+traceId)

	return &response, nil
}
