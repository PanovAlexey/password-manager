package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"storage/internal/domain"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) UpdateLoginPassword(ctx context.Context, request *pb.UpdateLoginPasswordByIdRequest) (*pb.UpdateLoginPasswordByIdResponse, error) {
	var response pb.UpdateLoginPasswordByIdResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	loginPasswordEntity, err := h.loginPasswordService.UpdateLoginPassword(
		domain.LoginPassword{
			Name:     request.CreateLoginPassword.Name,
			Login:    request.CreateLoginPassword.Login,
			Password: request.CreateLoginPassword.Password,
			UserId:   request.CreateLoginPassword.UserId,
			Note:     request.CreateLoginPassword.Note,
		},
		userId,
	)

	if err != nil {
		return nil, errors.New("login password updating error: " + err.Error())
	}

	var loginPassword pb.LoginPassword
	loginPassword.Id = strconv.FormatInt(loginPasswordEntity.Id.Int64, 10)
	loginPassword.Note = loginPasswordEntity.Note
	loginPassword.Name = loginPasswordEntity.Name
	loginPassword.Login = loginPasswordEntity.Login
	loginPassword.Password = loginPasswordEntity.Password

	createDateTime, err := time.Parse(time.RFC3339, loginPasswordEntity.CreatedAt)
	loginPassword.CreatedDate = timestamppb.New(createDateTime)

	if loginPasswordEntity.LastAccessAt.Valid {
		lastAccessAtTime, err := time.Parse(time.RFC3339, loginPasswordEntity.LastAccessAt.String)
		loginPassword.LastAccess = timestamppb.New(lastAccessAtTime)

		if err != nil {
			return nil, err
		}
	}

	response.LoginPassword = &loginPassword

	h.logger.Info("successful updated login-password by userId=" + userId)

	return &response, nil
}
