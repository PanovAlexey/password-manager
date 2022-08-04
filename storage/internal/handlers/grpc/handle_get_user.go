package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"storage/internal/domain"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var response pb.GetUserResponse
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)

	userLogin := domain.UserLogin{
		Email:    request.GetUser.Email,
		Password: request.GetUser.Password,
	}
	user, err := h.userService.GetUser(userLogin)

	if err != nil {
		h.logger.Error("user dit not save to database: "+err.Error(), ". traceId="+traceId)
		return nil, err
	}

	var userOutput pb.User
	userOutput.Id = strconv.FormatInt(user.Id.Int64, 10)
	userOutput.Email = user.Email

	regDateTime, err := time.Parse(time.RFC3339, user.CreatedAt)
	userOutput.RegistrationDate = timestamppb.New(regDateTime)

	if user.LastAccessAt.Valid {
		lastAccessAtTime, err := time.Parse(time.RFC3339, user.LastAccessAt.String)
		userOutput.LastLogin = timestamppb.New(lastAccessAtTime)

		if err != nil {
			return nil, err
		}
	}

	response.User = &userOutput

	h.logger.Info("successful created user. ", ". traceId="+traceId)

	return &response, nil
}
