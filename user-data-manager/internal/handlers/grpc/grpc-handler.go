package grpc

import (
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type UserDataManagerHandler struct {
	logger Logger
	pb.UnimplementedUserDataManagerServer
}

func GetUserDataManagerHandler(logger Logger) *UserDataManagerHandler {
	return &UserDataManagerHandler{logger: logger}
}
