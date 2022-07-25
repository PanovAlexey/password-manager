package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) UpdateTextRecord(ctx context.Context, request *pb.UpdateTextRecordByIdRequest) (*pb.UpdateTextRecordByIdResponse, error) {
	var response pb.UpdateTextRecordByIdResponse

	// @ToDo: replace stub data for real data
	var textRecord pb.TextRecord
	textRecord.Id = "1234567890"
	textRecord.Name = "War and peace. Tolstoy."
	textRecord.Text = "Temporary text etc..."
	textRecord.UserId = "324"
	textRecord.CreatedDate = &timestamp.Timestamp{}
	textRecord.LastAccess = &timestamp.Timestamp{}
	response.TextRecord = &textRecord

	s.logger.Info("successful updated text record. ", request)

	return &response, nil
}
