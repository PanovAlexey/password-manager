package grpc

import (
	"context"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (s *UserAuthorizationHandler) GetUserIdByJWT(ctx context.Context, request *pb.GetUserIdByJWTRequest) (*pb.GetUserIdByJWTResponse, error) {
	var response pb.GetUserIdByJWTResponse

	// @ToDo: replace stub data for real data
	response.UserId = "23213"

	s.logger.Info("successful got user id by token. ", request, request.Token)

	return &response, nil
}
