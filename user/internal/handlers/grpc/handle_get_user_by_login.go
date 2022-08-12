package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (h *UserAuthorizationHandler) GetUserByLogin(
	ctx context.Context, request *pb.GetUserByLoginRequest,
) (*pb.GetUserByLoginResponse, error) {
	var response pb.GetUserByLoginResponse
	traceId := h.userMetadataFromContextGetterService.GetTraceIdFromContext(ctx)

	// @ToDo: replace stub data for real data
	var user pb.User
	user.Id = "123"
	user.Token = "bearer token 234324324"
	user.RegistrationDate = &timestamp.Timestamp{}
	user.LastLogin = &timestamp.Timestamp{}
	user.Email = "test@ya.ru"
	response.User = &user

	h.logger.Info("successful got user by login. ", request.Login, ". traceId="+traceId)

	return &response, nil
}
