package grpc

import (
	"user-data-manager/internal/application/service"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type UserDataManagerHandler struct {
	logger          Logger
	userDataService service.UserData
	pb.UnimplementedUserDataManagerServer
}

func GetUserDataManagerHandler(
	logger Logger,
	userDataService service.UserData,
) *UserDataManagerHandler {
	return &UserDataManagerHandler{
		logger:          logger,
		userDataService: userDataService,
	}
}
