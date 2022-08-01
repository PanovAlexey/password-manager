package grpc

import (
	"user-auth/internal/application/service"
	"user-auth/internal/infrastructure/clients/grpc"
	pb "user-auth/pkg/user_authorization_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type UserAuthorizationHandler struct {
	logger                  Logger
	jwtAuthorizationService service.JWTAuthorization
	storageClient           grpc.StorageClient
	pb.UnimplementedUserAuthorizationServer
}

func GetUserAuthorizationHandler(
	logger Logger,
	jwtAuthorizationService service.JWTAuthorization,
	storageClient grpc.StorageClient,
) *UserAuthorizationHandler {
	return &UserAuthorizationHandler{
		logger:                  logger,
		jwtAuthorizationService: jwtAuthorizationService,
		storageClient:           storageClient,
	}
}
