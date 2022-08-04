package grpc

import (
	"context"
	"errors"
	pb "user-auth/pkg/user_authorization_grpc"
)

func (h *UserAuthorizationHandler) Auth(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	traceId := h.userMetadataFromContextGetterService.GetTraceIdFromContext(ctx)
	authUser, err := h.userRegistrationService.Auth(request.AuthUser.Email, request.AuthUser.Password, ctx)

	if err != nil {
		h.logger.Error("auth error: "+err.Error(), ". traceId="+traceId)
		return nil, err
	}

	token, err := h.jwtAuthorizationService.GetJWTToken(authUser.Id)

	if err != nil {
		h.logger.Error("getting JWT token by user id error: "+err.Error(), authUser.Id, ". traceId="+traceId)
		return nil, errors.New("getting JWT token by user id error: " + err.Error())
	}

	user := pb.User{
		Id:               authUser.Id,
		Email:            authUser.Email,
		Token:            token,
		RegistrationDate: authUser.RegistrationDate,
		LastLogin:        authUser.LastLogin,
	}

	h.logger.Info("successful auth. ", user, ". traceId="+traceId)

	return &pb.AuthResponse{User: &user}, nil
}
