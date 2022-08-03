package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) CreateLoginPassword(ctx context.Context, request *pb.CreateLoginPasswordRequest) (*pb.CreateLoginPasswordResponse, error) {
	var response pb.CreateLoginPasswordResponse

	loginPassword, err := h.userDataService.AddLoginPassword(*request.CreateLoginPassword, ctx)

	if err != nil {
		h.logger.Info("creating login-password error. "+err.Error(), request)

		return nil, err
	}

	response.LoginPassword = &loginPassword

	h.logger.Info("successful created login-password. " + loginPassword.Id)

	return &response, nil
}
