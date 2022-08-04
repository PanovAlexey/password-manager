package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) PatchLoginPasswordById(
	ctx context.Context,
	request *pb.PatchLoginPasswordByIdRequest,
) (*pb.PatchLoginPasswordByIdResponse, error) {
	var response pb.PatchLoginPasswordByIdResponse
	userid := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	loginPassword, err := h.userDataService.UpdateLoginPassword(*request.LoginPassword, userid, ctx)

	if err != nil {
		h.logger.Info("updating login-password error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.LoginPassword = &loginPassword
	h.logger.Info("successful updated login-password. "+request.LoginPassword.Id, ". traceId="+traceId)

	return &response, nil
}
