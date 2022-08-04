package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"storage/internal/domain"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var response pb.CreateUserResponse
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)

	userInput := domain.UserLogin{
		Email:    request.CreateUser.Email,
		Password: request.CreateUser.Password,
	}
	userEntity, err := h.userService.SaveUser(userInput)

	if err != nil {
		h.logger.Error("user dit not save to database: "+err.Error(), ". traceId="+traceId)
		return nil, err
	}

	var user pb.User
	user.Id = strconv.FormatInt(userEntity.Id.Int64, 10)
	user.Email = userEntity.Email

	regDateTime, err := time.Parse(time.RFC3339, userEntity.CreatedAt)
	user.RegistrationDate = timestamppb.New(regDateTime)

	if userEntity.LastAccessAt.Valid {
		lastAccessAtTime, err := time.Parse(time.RFC3339, userEntity.LastAccessAt.String)
		user.LastLogin = timestamppb.New(lastAccessAtTime)

		if err != nil {
			return nil, err
		}
	}

	response.User = &user

	h.logger.Info("successful created user. ", ". traceId="+traceId)

	return &response, nil
}
