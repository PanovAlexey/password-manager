package grpc

import (
	"context"
	"storage/internal/domain"
	pb "storage/pkg/storage_grpc"
	"strconv"
)

func (h *StorageHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var response pb.CreateUserResponse

	userInput := domain.User{
		Email:    request.CreateUser.Email,
		Password: request.CreateUser.Password,
	}
	id, err := h.userService.SaveUser(userInput)

	if err != nil {
		h.logger.Error("user dit not save to database: " + err.Error())
		return nil, err
	}

	var user pb.User
	user.Id = strconv.Itoa(id)
	user.Email = userInput.Email

	/* @ToDo: add
	user.RegistrationDate = &timestamp.Timestamp{}
	user.LastLogin = &timestamp.Timestamp{}
	*/

	response.User = &user

	h.logger.Info("successful created user. ", request)

	return &response, nil
}
