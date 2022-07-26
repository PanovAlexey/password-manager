package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) GetLoginPasswordById(
	ctx context.Context, request *pb.GetLoginPasswordByIdRequest,
) (*pb.GetLoginPasswordByIdResponse, error) {
	var response pb.GetLoginPasswordByIdResponse
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	loginPassword, err := h.userDataService.GetLoginPasswordById(request.Id, ctx)

	if err != nil {
		h.logger.Info("getting login-password by id error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.LoginPassword = &loginPassword

	h.logger.Info("successful got login-password by id. ", request.Id, ". traceId="+traceId)

	return &response, nil
}
