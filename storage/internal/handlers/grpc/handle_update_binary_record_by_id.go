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

func (h *StorageHandler) UpdateBinaryRecord(
	ctx context.Context,
	request *pb.UpdateBinaryRecordByIdRequest,
) (*pb.UpdateBinaryRecordByIdResponse, error) {
	var response pb.UpdateBinaryRecordByIdResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	binaryRecordEntity, err := h.binaryRecordService.UpdateBinaryRecord(
		domain.BinaryRecord{
			Name:   request.CreateBinaryRecord.Name,
			Binary: request.CreateBinaryRecord.Binary,
			UserId: request.CreateBinaryRecord.UserId,
			Note:   request.CreateBinaryRecord.Note,
		},
		userId,
	)

	if err != nil {
		return nil, errors.New("binary record updating error: " + err.Error())
	}

	var binaryRecord pb.BinaryRecord
	binaryRecord.Id = strconv.FormatInt(binaryRecordEntity.Id.Int64, 10)
	binaryRecord.Note = binaryRecordEntity.Note
	binaryRecord.Name = binaryRecordEntity.Name
	binaryRecord.Binary = binaryRecordEntity.Binary

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

	h.logger.Info("successful updated binary record by userId=" + userId)

	return &response, nil
}
