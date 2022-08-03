package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) GetCreditCardById(
	ctx context.Context, request *pb.GetCreditCardByIdRequest,
) (*pb.GetCreditCardByIdResponse, error) {
	var response pb.GetCreditCardByIdResponse

	creditCard, err := h.userDataService.GetCreditCardById(request.Id, ctx)

	if err != nil {
		h.logger.Info("getting credit card by id error. "+err.Error(), request)

		return nil, err
	}

	response.CreditCard = &creditCard

	h.logger.Info("successful got credit card by id. ", request.Id)

	return &response, nil
}
