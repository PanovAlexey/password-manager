package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) CreateBinaryRecord(ctx context.Context, request *pb.CreateBinaryRecordRequest) (*pb.CreateBinaryRecordResponse, error) {
	var response pb.CreateBinaryRecordResponse

	// @ToDo: replace stub data for real data
	var binaryRecord pb.BinaryRecord
	binaryRecord.Id = "1234567890"
	binaryRecord.Note = "Note text etc for example"
	binaryRecord.Name = "Stub 2 binary record for vk.com"
	binaryRecord.UserId = "234324-324324-32"
	binaryRecord.Binary = "01010101010101"
	binaryRecord.CreatedDate = &timestamp.Timestamp{}
	binaryRecord.LastAccess = &timestamp.Timestamp{}
	response.BinaryRecord = &binaryRecord

	h.logger.Info("successful created binary record. ", request)

	return &response, nil
}
