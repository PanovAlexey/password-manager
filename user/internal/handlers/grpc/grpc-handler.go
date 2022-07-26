package grpc

import (
	pb "user-auth/pkg/user_authorization_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type UserAuthorizationHandler struct {
	logger Logger
	pb.UnimplementedUserAuthorizationServer
}

func GetUserAuthorizationHandler(logger Logger) *UserAuthorizationHandler {
	return &UserAuthorizationHandler{logger: logger}
}
