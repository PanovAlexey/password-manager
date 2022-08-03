package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) CreateCreditCard(ctx context.Context, request *pb.CreateCreditCardRequest) (*pb.CreateCreditCardResponse, error) {
	var response pb.CreateCreditCardResponse

	// @ToDo: replace stub data for real data
	var creditCard pb.CreditCard
	creditCard.Id = "1234567890"
	creditCard.Note = "Note text etc for example"
	creditCard.Name = "Stub 2 binary record for vk.com"
	creditCard.UserId = "234324-324324-32"
	creditCard.Cvv = "4242"
	creditCard.Owner = "Mark Cukenberg"
	creditCard.Expiration = "11/24"
	creditCard.Number = "44443343434333292"
	creditCard.CreatedDate = &timestamp.Timestamp{}
	creditCard.LastAccess = &timestamp.Timestamp{}
	response.CreditCard = &creditCard

	h.logger.Info("successful created credit card. ", request)

	return &response, nil
}
