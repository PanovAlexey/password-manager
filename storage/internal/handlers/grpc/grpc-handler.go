package grpc

import (
	service "storage/internal/application/service"
	pb "storage/pkg/storage_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type StorageHandler struct {
	logger                  Logger
	userService             service.UserService
	userIdFromContextGetter service.UserIdFromContextGetter
	loginPasswordService    service.LoginPassword
	pb.UnimplementedStorageServer
}

func GetStorageHandler(
	logger Logger,
	userService service.UserService,
	userIdFromContextGetter service.UserIdFromContextGetter,
	loginPasswordService service.LoginPassword,
) *StorageHandler {
	return &StorageHandler{
		logger:                  logger,
		userService:             userService,
		userIdFromContextGetter: userIdFromContextGetter,
		loginPasswordService:    loginPasswordService,
	}
}
