package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) CreateTextRecord(ctx context.Context, request *pb.CreateTextRecordRequest) (*pb.CreateTextRecordResponse, error) {
	var response pb.CreateTextRecordResponse

	// @ToDo: replace stub data for real data
	var textRecord pb.TextRecord
	textRecord.Id = "1234567890"
	textRecord.Note = "Note text etc for example"
	textRecord.Name = "Stub 2 binary record for vk.com"
	textRecord.UserId = "234324-324324-32"
	textRecord.Text = "text content"
	textRecord.CreatedDate = &timestamp.Timestamp{}
	textRecord.LastAccess = &timestamp.Timestamp{}
	response.TextRecord = &textRecord

	h.logger.Info("successful created text record. ", request)

	return &response, nil
}
