package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetLoginPasswordById(ctx context.Context, request *pb.GetLoginPasswordByIdRequest) (*pb.GetLoginPasswordByIdResponse, error) {
	var response pb.GetLoginPasswordByIdResponse
	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	loginPasswordEntity, err := h.loginPasswordService.GetLoginPasswordById(request.Id, userId)

	if err != nil {
		return nil, errors.New("login password getting by id error: " + err.Error())
	}

	var loginPassword pb.LoginPassword
	loginPassword.Id = strconv.FormatInt(loginPasswordEntity.Id.Int64, 10)
	loginPassword.Note = loginPasswordEntity.Note
	loginPassword.Name = loginPasswordEntity.Name
	loginPassword.Login = loginPasswordEntity.Login
	loginPassword.Password = loginPasswordEntity.Password

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

	response.LoginPassword = &loginPassword

	h.logger.Info("successful got login-password by id. ", ". traceId="+traceId)

	return &response, nil
}
