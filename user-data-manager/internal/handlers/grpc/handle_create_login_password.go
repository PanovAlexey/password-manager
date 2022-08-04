package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) CreateLoginPassword(ctx context.Context, request *pb.CreateLoginPasswordRequest) (*pb.CreateLoginPasswordResponse, error) {
	var response pb.CreateLoginPasswordResponse

	userid := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	loginPassword, err := h.userDataService.AddLoginPassword(*request.CreateLoginPassword, userid, ctx)

	if err != nil {
		h.logger.Info("creating login-password error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.LoginPassword = &loginPassword

	h.logger.Info("successful created login-password. "+loginPassword.Id, ". traceId="+traceId)

	return &response, nil
}
