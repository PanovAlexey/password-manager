package http

import (
	"api-gw/internal/application/service"
	"api-gw/internal/handlers/http/dto"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *httpHandler) HandleGetUserAllData(w http.ResponseWriter, r *http.Request) {
	UserData := dto.UserData{}
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))

	binaryRecordListResponse, err := (*h.gRPCUserDataManagerClient.GetClient()).GetBinaryRecordList(
		r.Context(),
		&pb.GetBinaryRecordListRequest{},
	)

	if err != nil {
		info := "error getting binary record list: " + err.Error()
		h.logger.Error(info, userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	UserData.BinaryRecordCollection = binaryRecordListResponse.ProtectedItemList

	textRecordListResponse, err := (*h.gRPCUserDataManagerClient.GetClient()).GetTextRecordList(
		r.Context(),
		&pb.GetTextRecordListRequest{},
	)

	if err != nil {
		info := "error getting text record list: " + err.Error()
		h.logger.Error(info, userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	UserData.TextRecordCollection = textRecordListResponse.ProtectedItemList

	creditCardListResponse, err := (*h.gRPCUserDataManagerClient.GetClient()).GetCreditCardList(
		r.Context(),
		&pb.GetCreditCardListRequest{},
	)

	if err != nil {
		info := "error getting credit card list: " + err.Error()
		h.logger.Error(info, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	UserData.CreditCardCollection = creditCardListResponse.ProtectedItemList

	loginPasswordListResponse, err := (*h.gRPCUserDataManagerClient.GetClient()).GetLoginPasswordList(
		r.Context(),
		&pb.GetLoginPasswordListRequest{},
	)

	if err != nil {
		info := "error getting login-password list: " + err.Error()
		h.logger.Error(info, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	UserData.LoginPasswordCollection = loginPasswordListResponse.ProtectedItemList

	result, err := json.Marshal(UserData)

	if err != nil {
		info := "error marshalling user data: " + err.Error()
		h.logger.Error(info, ". userId=", userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful getting user data by userId=", userId)

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
