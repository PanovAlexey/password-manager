package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetLoginPasswordList(ctx context.Context, request *pb.GetLoginPasswordListRequest) (*pb.GetLoginPasswordListResponse, error) {
	var response pb.GetLoginPasswordListResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	loginPasswordEntityList, err := h.loginPasswordService.GetLoginPasswordList(userId)

	if err != nil {
		return nil, errors.New("login password getting list error: " + err.Error())
	}

	for _, loginPasswordEntity := range loginPasswordEntityList {
		loginPassword := pb.LoginPassword{}
		loginPassword.Id = strconv.FormatInt(loginPasswordEntity.Id.Int64, 10)
		loginPassword.Name = loginPasswordEntity.Name
		// loginPassword.Note = loginPasswordEntity.Note
		// loginPassword.Login = loginPasswordEntity.Login
		// loginPassword.Password = loginPasswordEntity.Password

		createDateTime, err := time.Parse(time.RFC3339, loginPasswordEntity.CreatedAt)

		if err != nil {
			return nil, err
		}

		loginPassword.CreatedDate = timestamppb.New(createDateTime)

		if loginPasswordEntity.LastAccessAt.Valid {
			lastAccessAtTime, err := time.Parse(time.RFC3339, loginPasswordEntity.LastAccessAt.String)
			loginPassword.LastAccess = timestamppb.New(lastAccessAtTime)

			if err != nil {
				return nil, err
			}
		}

		response.LoginPasswordList = append(response.LoginPasswordList, &loginPassword)
	}

	h.logger.Info("successful got login-password list. UserId=", userId)

	return &response, nil
}
