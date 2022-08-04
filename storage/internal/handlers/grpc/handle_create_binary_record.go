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

func (h *StorageHandler) CreateBinaryRecord(
	ctx context.Context,
	request *pb.CreateBinaryRecordRequest,
) (*pb.CreateBinaryRecordResponse, error) {
	var response pb.CreateBinaryRecordResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	binaryRecordEntity, err := h.binaryRecordService.AddBinaryRecord(
		domain.BinaryRecord{
			Name:   request.CreateBinaryRecord.Name,
			Binary: request.CreateBinaryRecord.Binary,
			UserId: request.CreateBinaryRecord.UserId,
			Note:   request.CreateBinaryRecord.Note,
		},
		userId,
	)

	if err != nil {
		return nil, errors.New("binary record creating error: " + err.Error())
	}

	var binaryRecord pb.BinaryRecord
	binaryRecord.Id = strconv.FormatInt(binaryRecordEntity.Id.Int64, 10)
	binaryRecord.Note = binaryRecordEntity.Note
	binaryRecord.Name = binaryRecordEntity.Name
	binaryRecord.Binary = binaryRecordEntity.Binary
	binaryRecord.UserId = binaryRecordEntity.UserId

	createDateTime, err := time.Parse(time.RFC3339, binaryRecordEntity.CreatedAt)
	binaryRecord.CreatedDate = timestamppb.New(createDateTime)

	if binaryRecordEntity.LastAccessAt.Valid {
		lastAccessAtTime, err := time.Parse(time.RFC3339, binaryRecordEntity.LastAccessAt.String)
		binaryRecord.LastAccess = timestamppb.New(lastAccessAtTime)

		if err != nil {
			return nil, err
		}
	}

	response.BinaryRecord = &binaryRecord

	h.logger.Info("successful created binary record by userId="+userId, ". traceId="+traceId)
	return &response, nil
}
