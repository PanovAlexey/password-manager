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

func (h *StorageHandler) UpdateTextRecord(
	ctx context.Context,
	request *pb.UpdateTextRecordByIdRequest,
) (*pb.UpdateTextRecordByIdResponse, error) {
	var response pb.UpdateTextRecordByIdResponse

	userId := h.userMetadataFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	textRecordEntity, err := h.textRecordService.UpdateTextRecord(
		domain.TextRecord{
			Name:   request.CreateTextRecord.Name,
			Text:   request.CreateTextRecord.Text,
			UserId: request.CreateTextRecord.UserId,
			Note:   request.CreateTextRecord.Note,
		},
		userId,
	)

	if err != nil {
		return nil, errors.New("text record updating error: " + err.Error())
	}

	var textRecord pb.TextRecord
	textRecord.Id = strconv.FormatInt(textRecordEntity.Id.Int64, 10)
	textRecord.Note = textRecordEntity.Note
	textRecord.Name = textRecordEntity.Name
	textRecord.Text = textRecordEntity.Text

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

	h.logger.Info("successful updated text record by userId="+userId, ". traceId="+traceId)

	return &response, nil
}
