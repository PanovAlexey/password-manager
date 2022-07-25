package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (s *StorageHandler) GetTextRecordById(ctx context.Context, request *pb.GetTextRecordByIdRequest) (*pb.GetTextRecordByIdResponse, error) {
	var response pb.GetTextRecordByIdResponse

	// @ToDo: replace stub data for real data
	var textRecord pb.TextRecord
	textRecord.Id = "1234567890"
	textRecord.Note = "Note text etc for example"
	textRecord.Name = "Stub 2 binary record for vk.com"
	textRecord.UserId = "234324-324324-32"
	textRecord.Text = "Temporary text..."
	textRecord.CreatedDate = &timestamp.Timestamp{}
	textRecord.LastAccess = &timestamp.Timestamp{}
	response.TextRecord = &textRecord

	s.logger.Info("successful got text record by id. ", request)

	return &response, nil
}
