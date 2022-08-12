package grpc

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "storage/pkg/storage_grpc"
	"strconv"
	"time"
)

func (h *StorageHandler) GetCreditCardList(ctx context.Context, request *pb.GetCreditCardListRequest) (*pb.GetCreditCardListResponse, error) {
	var response pb.GetCreditCardListResponse

	userId := h.userMetadataFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	creditCardEntityList, err := h.creditCardService.GetCreditCardList(userId)

	if err != nil {
		return nil, errors.New("credit card getting list error: " + err.Error())
	}

	for _, creditCardEntity := range creditCardEntityList {
		creditCard := pb.CreditCard{}
		creditCard.Id = strconv.FormatInt(creditCardEntity.Id.Int64, 10)
		creditCard.Name = creditCardEntity.Name
		// creditCard.Note = creditCardEntity.Note
		// creditCard.Login = creditCardEntity.Login
		// creditCard.Password = creditCardEntity.Password

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

		response.CreditCardList = append(response.CreditCardList, &creditCard)
	}

	h.logger.Info("successful got credit card list. UserId="+userId, ". traceId="+traceId)

	return &response, nil
}
