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

func (h *StorageHandler) CreateCreditCard(ctx context.Context, request *pb.CreateCreditCardRequest) (*pb.CreateCreditCardResponse, error) {
	var response pb.CreateCreditCardResponse

	userId := h.userIdFromContextGetter.GetUserIdFromContext(ctx)
	creditCardEntity, err := h.creditCardService.AddCreditCard(
		domain.CreditCard{
			Name:       request.CreateCreditCard.Name,
			Number:     request.CreateCreditCard.Number,
			Expiration: request.CreateCreditCard.Expiration,
			Cvv:        request.CreateCreditCard.Cvv,
			Owner:      request.CreateCreditCard.Owner,
			UserId:     request.CreateCreditCard.UserId,
			Note:       request.CreateCreditCard.Note,
		},
		userId,
	)

	if err != nil {
		return nil, errors.New("credit card creating error: " + err.Error())
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
	creditCard.CreatedDate = timestamppb.New(createDateTime)

	if creditCardEntity.LastAccessAt.Valid {
		lastAccessAtTime, err := time.Parse(time.RFC3339, creditCardEntity.LastAccessAt.String)
		creditCard.LastAccess = timestamppb.New(lastAccessAtTime)

		if err != nil {
			return nil, err
		}
	}

	response.CreditCard = &creditCard

	h.logger.Info("successful created credit card by userId=" + userId)

	return &response, nil
}
