package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) UpdateBinaryRecord(ctx context.Context, request *pb.UpdateBinaryRecordByIdRequest) (*pb.UpdateBinaryRecordByIdResponse, error) {
	var response pb.UpdateBinaryRecordByIdResponse

	// @ToDo: replace stub data for real data
	var binaryRecord pb.BinaryRecord
	binaryRecord.Id = "1234567890"
	binaryRecord.Binary = "0101010101010111"
	binaryRecord.UserId = "324"
	binaryRecord.Name = "File photo"
	binaryRecord.CreatedDate = &timestamp.Timestamp{}
	binaryRecord.LastAccess = &timestamp.Timestamp{}
	response.BinaryRecord = &binaryRecord

	s.logger.Info("successful updated binary record. ", request)

	return &response, nil
}
