package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetBinaryRecordById(
	ctx context.Context,
	request *pb.GetBinaryRecordByIdRequest,
) (*pb.GetBinaryRecordByIdResponse, error) {
	var response pb.GetBinaryRecordByIdResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	binaryRecordEntity, err := h.binaryRecordService.GetBinaryRecordById(request.Id, userId)

	if err != nil {
		return nil, errors.New("binary record getting by id error: " + err.Error())
	}

	var binaryRecord pb.BinaryRecord
	binaryRecord.Id = strconv.FormatInt(binaryRecordEntity.Id.Int64, 10)
	binaryRecord.Note = binaryRecordEntity.Note
	binaryRecord.Name = binaryRecordEntity.Name
	binaryRecord.Binary = binaryRecordEntity.Binary
	binaryRecord.UserId = binaryRecordEntity.UserId

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

	response.BinaryRecord = &binaryRecord

	h.logger.Info("successful got binary record by id. ", ". traceId="+traceId)

	return &response, nil
}
