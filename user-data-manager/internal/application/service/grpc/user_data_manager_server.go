package grpc

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

type Logger interface {
	Error(args ...interface{})
	Info(args ...interface{})
}

type UserDataManagerServer struct {
	logger Logger
	pb.UnimplementedUserDataManagerServer
}

func GetUserDataManagerServer(logger Logger) *UserDataManagerServer {
	return &UserDataManagerServer{logger: logger}
}

func (s *UserDataManagerServer) GetLoginPasswordList(ctx context.Context, request *pb.GetLoginPasswordListRequest) (*pb.GetLoginPasswordListResponse, error) {
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

func (s *UserDataManagerServer) CreateLoginPassword(ctx context.Context, request *pb.CreateLoginPasswordRequest) (*pb.CreateLoginPasswordResponse, error) {
	var response pb.CreateLoginPasswordResponse

	// @ToDo: replace stub data for real data
	var loginPassword pb.LoginPassword
	loginPassword.Id = "1234567890"
	loginPassword.Note = "Note text etc for example"
	loginPassword.Name = "Stub 2 login password for vk.com"
	loginPassword.Login = "login"
	loginPassword.Password = "pass"
	loginPassword.CreatedDate = &timestamp.Timestamp{}
	loginPassword.LastAccess = &timestamp.Timestamp{}
	response.LoginPassword = &loginPassword

	s.logger.Info("successful created login-password. ", request)

	return &response, nil
}

func (s *UserDataManagerServer) GetLoginPasswordById(ctx context.Context, request *pb.GetLoginPasswordByIdRequest) (*pb.GetLoginPasswordByIdResponse, error) {
	var response pb.GetLoginPasswordByIdResponse

	// @ToDo: replace stub data for real data
	var loginPassword pb.LoginPassword
	loginPassword.Id = "33333"
	loginPassword.Note = "Note text etc for example"
	loginPassword.Name = "Stub 2 login password for vk.com"
	loginPassword.Login = "login"
	loginPassword.Password = "pass"
	loginPassword.CreatedDate = &timestamp.Timestamp{}
	loginPassword.LastAccess = &timestamp.Timestamp{}
	response.LoginPassword = &loginPassword

	s.logger.Info("successful got login-password by id. ", request)

	return &response, nil
}

func (s *UserDataManagerServer) PatchLoginPasswordById(ctx context.Context, request *pb.PatchLoginPasswordByIdRequest) (*pb.PatchLoginPasswordByIdResponse, error) {
	var response pb.PatchLoginPasswordByIdResponse

	// @ToDo: replace stub data for real data
	var loginPassword pb.LoginPassword
	loginPassword.Id = "444444"
	loginPassword.Note = "Note text etc for example"
	loginPassword.Name = "Stub 4 login password for vk.com"
	loginPassword.Login = "login"
	loginPassword.Password = "pass"
	loginPassword.CreatedDate = &timestamp.Timestamp{}
	loginPassword.LastAccess = &timestamp.Timestamp{}
	response.LoginPassword = &loginPassword

	s.logger.Info("successful patched login-password by id ", request)

	return &response, nil
}

func (s *UserDataManagerServer) DeleteLoginPasswordById(ctx context.Context, request *pb.DeleteLoginPasswordByIdRequest) (*emptypb.Empty, error) {
	s.logger.Info("successful deleted login-password by id. ", request)
	// @ToDo handle error
	return &emptypb.Empty{}, errors.New("test error")
}
