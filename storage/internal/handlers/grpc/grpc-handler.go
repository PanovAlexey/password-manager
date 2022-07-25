package grpc

import (
	pb "storage/pkg/storage_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type StorageHandler struct {
	logger Logger
	pb.UnimplementedStorageServer
}

func GetStorageHandler(logger Logger) *StorageHandler {
	return &StorageHandler{logger: logger}
}
