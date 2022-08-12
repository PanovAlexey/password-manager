package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) CreateCreditCard(ctx context.Context, request *pb.CreateCreditCardRequest) (*pb.CreateCreditCardResponse, error) {
	var response pb.CreateCreditCardResponse

	userid := h.userMetadataFromContextGetter.GetUserIdFromContext(ctx)
	traceId := h.userMetadataFromContextGetter.GetTraceIdFromContext(ctx)
	creditCard, err := h.userDataService.AddCreditCard(*request.CreateCreditCard, userid, ctx)

	if err != nil {
		h.logger.Info("creating credit card error. "+err.Error(), ". traceId="+traceId)

		return nil, err
	}

	response.CreditCard = &creditCard

	h.logger.Info("successful created credit card. "+creditCard.Id, ". traceId="+traceId)

	return &response, nil
}
