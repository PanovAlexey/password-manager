package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "storage/pkg/storage_grpc"
)

func (h *StorageHandler) UpdateCreditCard(ctx context.Context, request *pb.UpdateCreditCardByIdRequest) (*pb.UpdateCreditCardByIdResponse, error) {
	var response pb.UpdateCreditCardByIdResponse

	// @ToDo: replace stub data for real data
	var creditCard pb.CreditCard
	creditCard.Id = "1234567890"
	creditCard.Cvv = "123"
	creditCard.Expiration = "02/23"
	creditCard.Name = "Sber card"
	creditCard.Number = "22223333444422222"
	creditCard.Owner = "Ivanov Ivan"
	creditCard.UserId = "324"
	creditCard.CreatedDate = &timestamp.Timestamp{}
	creditCard.LastAccess = &timestamp.Timestamp{}
	response.CreditCard = &creditCard

	h.logger.Info("successful updated credit card. ", request)

	return &response, nil
}
