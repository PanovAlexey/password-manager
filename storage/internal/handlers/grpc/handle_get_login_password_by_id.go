package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/metadata"
	"log"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) GetLoginPasswordById(ctx context.Context, request *pb.GetLoginPasswordByIdRequest) (*pb.GetLoginPasswordByIdResponse, error) {
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

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	log.Println(userId)

	h.logger.Info("successful got login-password by id. ", request)

	return &response, nil
}
