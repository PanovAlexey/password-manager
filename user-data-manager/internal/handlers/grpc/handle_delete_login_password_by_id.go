package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) DeleteLoginPasswordById(ctx context.Context, request *pb.DeleteLoginPasswordByIdRequest) (*emptypb.Empty, error) {

	err := h.userDataService.DeleteLoginPasswordById(request.Id, ctx)

	if err != nil {
		h.logger.Info("getting login-password by id error. "+err.Error(), request)

		return nil, err
	}

	h.logger.Info("successful deleted login-password by id. ", request.Id)

	return &emptypb.Empty{}, nil
}
