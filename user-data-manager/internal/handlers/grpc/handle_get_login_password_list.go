package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) GetLoginPasswordList(ctx context.Context, request *pb.GetLoginPasswordListRequest) (*pb.GetLoginPasswordListResponse, error) {
	var response pb.GetLoginPasswordListResponse

	loginPasswordList, err := h.userDataService.GetLoginPasswordList(ctx)

	if err != nil {
		h.logger.Info("getting login-password list error. "+err.Error(), request)

		return nil, err
	}

	for _, loginPassword := range loginPasswordList {
		protectedItem := pb.ProtectedItem{
			Id:          loginPassword.Id,
			Name:        loginPassword.Name,
			CreatedDate: loginPassword.CreatedDate,
			LastAccess:  loginPassword.LastAccess,
		}

		response.ProtectedItemList = append(response.ProtectedItemList, &protectedItem)
	}

	h.logger.Info("successful got login-password list. ")

	return &response, nil
}
