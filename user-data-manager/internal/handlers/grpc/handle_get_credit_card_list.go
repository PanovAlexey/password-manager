package grpc

import (
	"context"
	pb "user-data-manager/pkg/user_data_manager_grpc"
)

func (h *UserDataManagerHandler) GetCreditCardList(ctx context.Context, request *pb.GetCreditCardListRequest) (*pb.GetCreditCardListResponse, error) {
	var response pb.GetCreditCardListResponse

	creditCardList, err := h.userDataService.GetCreditCardList(ctx)

	if err != nil {
		h.logger.Info("getting credit card list error. "+err.Error(), request)

		return nil, err
	}

	for _, creditCard := range creditCardList {
		protectedItem := pb.ProtectedItem{
			Id:          creditCard.Id,
			Name:        creditCard.Name,
			CreatedDate: creditCard.CreatedDate,
			LastAccess:  creditCard.LastAccess,
		}

		response.ProtectedItemList = append(response.ProtectedItemList, &protectedItem)
	}

	h.logger.Info("successful got credit card list. ")

	return &response, nil
}
