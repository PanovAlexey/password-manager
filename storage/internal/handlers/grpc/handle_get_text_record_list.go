package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetTextRecordList(
	ctx context.Context,
	request *pb.GetTextRecordListRequest,
) (*pb.GetTextRecordListResponse, error) {
	var response pb.GetTextRecordListResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	textRecordEntityList, err := h.textRecordService.GetTextRecordList(userId)

	if err != nil {
		return nil, errors.New("text record getting list error: " + err.Error())
	}

	for _, textRecordEntity := range textRecordEntityList {
		textRecord := pb.TextRecord{}
		textRecord.Id = strconv.FormatInt(textRecordEntity.Id.Int64, 10)
		textRecord.Name = textRecordEntity.Name

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

		response.TextRecordList = append(response.TextRecordList, &textRecord)
	}

	h.logger.Info("successful got text record list. UserId=", userId, ". traceId="+traceId)

	return &response, nil
}
