package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetBinaryRecordList(
	ctx context.Context,
	request *pb.GetBinaryRecordListRequest,
) (*pb.GetBinaryRecordListResponse, error) {
	var response pb.GetBinaryRecordListResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	binaryRecordEntityList, err := h.binaryRecordService.GetBinaryRecordList(userId)

	if err != nil {
		return nil, errors.New("binary record getting list error: " + err.Error())
	}

	for _, binaryRecordEntity := range binaryRecordEntityList {
		binaryRecord := pb.BinaryRecord{}
		binaryRecord.Id = strconv.FormatInt(binaryRecordEntity.Id.Int64, 10)
		binaryRecord.Name = binaryRecordEntity.Name

		createDateTime, err := time.Parse(time.RFC3339, binaryRecordEntity.CreatedAt)

		if err != nil {
			return nil, err
		}

		binaryRecord.CreatedDate = timestamppb.New(createDateTime)

		if binaryRecordEntity.LastAccessAt.Valid {
			lastAccessAtTime, err := time.Parse(time.RFC3339, binaryRecordEntity.LastAccessAt.String)
			binaryRecord.LastAccess = timestamppb.New(lastAccessAtTime)

			if err != nil {
				return nil, err
			}
		}

		response.BinaryRecordList = append(response.BinaryRecordList, &binaryRecord)
	}

	h.logger.Info("successful got binary record list. UserId="+userId, ". traceId="+traceId)

	return &response, nil
}
