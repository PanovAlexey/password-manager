package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetCreditCardById(ctx context.Context, request *pb.GetCreditCardByIdRequest) (*pb.GetCreditCardByIdResponse, error) {
	var response pb.GetCreditCardByIdResponse
	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userIdFromContextGetter.GetTraceIdFromContext(ctx)
	creditCardEntity, err := h.creditCardService.GetCreditCardById(request.Id, userId)

	if err != nil {
		return nil, errors.New("credit card getting by id error: " + err.Error())
	}

	var creditCard pb.CreditCard
	creditCard.Id = strconv.FormatInt(creditCardEntity.Id.Int64, 10)
	creditCard.Note = creditCardEntity.Note
	creditCard.Name = creditCardEntity.Name
	creditCard.Number = creditCardEntity.Number
	creditCard.Expiration = creditCardEntity.Expiration
	creditCard.Cvv = creditCardEntity.Cvv
	creditCard.Owner = creditCardEntity.Owner
	creditCard.UserId = creditCardEntity.UserId

	createDateTime, err := time.Parse(time.RFC3339, creditCardEntity.CreatedAt)

	if err != nil {
		return nil, err
	}

	creditCard.CreatedDate = timestamppb.New(createDateTime)

	if creditCardEntity.LastAccessAt.Valid {
		lastAccessAtTime, err := time.Parse(time.RFC3339, creditCardEntity.LastAccessAt.String)
		creditCard.LastAccess = timestamppb.New(lastAccessAtTime)

		if err != nil {
			return nil, err
		}
	}

	response.CreditCard = &creditCard

	h.logger.Info("successful got credit card by id. ", ". traceId="+traceId)

	return &response, nil
}
