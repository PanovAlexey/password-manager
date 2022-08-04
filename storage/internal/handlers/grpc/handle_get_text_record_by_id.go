package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetTextRecordById(
	ctx context.Context,
	request *pb.GetTextRecordByIdRequest,
) (*pb.GetTextRecordByIdResponse, error) {
	var response pb.GetTextRecordByIdResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	textRecordEntity, err := h.textRecordService.GetTextRecordById(request.Id, userId)

	if err != nil {
		return nil, errors.New("text record getting by id error: " + err.Error())
	}

	var textRecord pb.TextRecord
	textRecord.Id = strconv.FormatInt(textRecordEntity.Id.Int64, 10)
	textRecord.Note = textRecordEntity.Note
	textRecord.Name = textRecordEntity.Name
	textRecord.Text = textRecordEntity.Text
	textRecord.UserId = textRecordEntity.UserId

	createDateTime, err := time.Parse(time.RFC3339, textRecordEntity.CreatedAt)

	if err != nil {
		return nil, err
	}

	textRecord.CreatedDate = timestamppb.New(createDateTime)

	if textRecordEntity.LastAccessAt.Valid {
		lastAccessAtTime, err := time.Parse(time.RFC3339, textRecordEntity.LastAccessAt.String)
		textRecord.LastAccess = timestamppb.New(lastAccessAtTime)

		if err != nil {
			return nil, err
		}
	}

	response.TextRecord = &textRecord

	h.logger.Info("successful got text record by id. ", ". traceId="+traceId)

	return &response, nil
}
