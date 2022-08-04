package grpc

import (
	"context"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (h *UserAuthorizationHandler) GetUserIdByJWT(
	ctx context.Context,
	r *pb.GetUserIdByJWTRequest,
) (*pb.GetUserIdByJWTResponse, error) {
	var response pb.GetUserIdByJWTResponse
	traceId := h.userMetadataFromContextGetterService.GetTraceIdFromContext(ctx)
	payload, err := h.jwtAuthorizationService.CheckGetJWTToken(r.Token)

	if err != nil {
		h.logger.Error(err, ". traceId="+traceId)

		return nil, err
	}

	response.UserId = payload.UserId
	h.logger.Info("successful got user id by token. ", r.Token, ". traceId="+traceId)

	return &response, nil
}
