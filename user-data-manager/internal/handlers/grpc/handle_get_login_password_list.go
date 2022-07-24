package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (s *UserDataManagerHandler) GetLoginPasswordList(ctx context.Context, request *pb.GetLoginPasswordListRequest) (*pb.GetLoginPasswordListResponse, error) {
	var response pb.GetLoginPasswordListResponse

	// @ToDo: replace stub data for real data
	var protectedItem pb.ProtectedItem
	protectedItem.Id = "11-22-33"
	protectedItem.Name = "Stub login password for vk.com"
	response.ProtectedItemList = append(response.ProtectedItemList, &protectedItem)
	protectedItem.Id = "77-88-99"
	protectedItem.Name = "Stub 2 login password for google.com"
	response.ProtectedItemList = append(response.ProtectedItemList, &protectedItem)

	s.logger.Info("successful got login-password list. ", request)

	return &response, nil
}
