package grpc

import (
	services "storage/internal/application/service"
	pb "storage/pkg/storage_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type StorageHandler struct {
	logger      Logger
	userService services.UserService
	pb.UnimplementedStorageServer
}

func GetStorageHandler(logger Logger, userService services.UserService) *StorageHandler {
	return &StorageHandler{
		logger:      logger,
		userService: userService,
	}
}
