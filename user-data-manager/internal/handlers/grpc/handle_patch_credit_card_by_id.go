package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) PatchCreditCardById(
	ctx context.Context,
	request *pb.PatchCreditCardByIdRequest,
) (*pb.PatchCreditCardByIdResponse, error) {
	var response pb.PatchCreditCardByIdResponse
	creditCard, err := h.userDataService.UpdateCreditCard(*request.CreditCard, ctx)

	if err != nil {
		h.logger.Info("updating credit card error. "+err.Error(), request)

		return nil, err
	}

	response.CreditCard = &creditCard
	h.logger.Info("successful updated credit card. " + request.CreditCard.Id)

	return &response, nil
}
