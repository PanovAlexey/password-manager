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

func (h *StorageHandler) CreateTextRecord(
	ctx context.Context,
	request *pb.CreateTextRecordRequest,
) (*pb.CreateTextRecordResponse, error) {
	var response pb.CreateTextRecordResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	textRecordEntity, err := h.textRecordService.AddTextRecord(
		domain.TextRecord{
			Name:   request.CreateTextRecord.Name,
			Text:   request.CreateTextRecord.Text,
			UserId: request.CreateTextRecord.UserId,
			Note:   request.CreateTextRecord.Note,
		},
		userId,
	)

	if err != nil {
		return nil, errors.New("text record creating error: " + err.Error())
	}

	var textRecord pb.TextRecord
	textRecord.Id = strconv.FormatInt(textRecordEntity.Id.Int64, 10)
	textRecord.Note = textRecordEntity.Note
	textRecord.Name = textRecordEntity.Name
	textRecord.Text = textRecordEntity.Text
	textRecord.UserId = textRecordEntity.UserId

	createDateTime, err := time.Parse(time.RFC3339, textRecordEntity.CreatedAt)
	textRecord.CreatedDate = timestamppb.New(createDateTime)

	if textRecordEntity.LastAccessAt.Valid {
		lastAccessAtTime, err := time.Parse(time.RFC3339, textRecordEntity.LastAccessAt.String)
		textRecord.LastAccess = timestamppb.New(lastAccessAtTime)

		if err != nil {
			return nil, err
		}
	}

	response.TextRecord = &textRecord

	h.logger.Info("successful created text record by userId=" + userId)
	return &response, nil
}
