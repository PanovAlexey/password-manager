package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) GetTextRecordList(
	ctx context.Context,
	request *pb.GetTextRecordListRequest,
) (*pb.GetTextRecordListResponse, error) {
	var response pb.GetTextRecordListResponse

	textRecordList, err := h.userDataService.GetTextRecordList(ctx)

	if err != nil {
		h.logger.Info("getting text record list error. "+err.Error(), request)

		return nil, err
	}

	for _, textRecord := range textRecordList {
		protectedItem := pb.ProtectedItem{
			Id:          textRecord.Id,
			Name:        textRecord.Name,
			CreatedDate: textRecord.CreatedDate,
			LastAccess:  textRecord.LastAccess,
		}

		response.ProtectedItemList = append(response.ProtectedItemList, &protectedItem)
	}

	h.logger.Info("successful got text record list. ")

	return &response, nil
}
