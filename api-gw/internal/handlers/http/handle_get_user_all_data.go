package http

import (
	"api-gw/internal/application/service"
	"api-gw/internal/domain"
	"api-gw/internal/handlers/http/dto"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (h *httpHandler) HandleGetUserAllData(w http.ResponseWriter, r *http.Request) {
	userData := dto.UserData{}
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

	for _, item := range binaryRecordListResponse.ProtectedItemList {
		userData.BinaryRecordCollection = append(userData.BinaryRecordCollection, domain.ProtectedItem{
			Id:           item.Id,
			Name:         item.Name,
			CreatedAt:    item.CreatedDate.AsTime().Format(time.RFC3339),
			LastAccessAt: item.LastAccess.AsTime().Format(time.RFC3339),
		})
	}

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

	for _, item := range textRecordListResponse.ProtectedItemList {
		userData.TextRecordCollection = append(userData.TextRecordCollection, domain.ProtectedItem{
			Id:           item.Id,
			Name:         item.Name,
			CreatedAt:    item.CreatedDate.AsTime().Format(time.RFC3339),
			LastAccessAt: item.LastAccess.AsTime().Format(time.RFC3339),
		})
	}

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

	for _, item := range creditCardListResponse.ProtectedItemList {
		userData.TextRecordCollection = append(userData.CreditCardCollection, domain.ProtectedItem{
			Id:           item.Id,
			Name:         item.Name,
			CreatedAt:    item.CreatedDate.AsTime().Format(time.RFC3339),
			LastAccessAt: item.LastAccess.AsTime().Format(time.RFC3339),
		})
	}

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

	for _, item := range loginPasswordListResponse.ProtectedItemList {
		userData.LoginPasswordCollection = append(userData.LoginPasswordCollection, domain.ProtectedItem{
			Id:           item.Id,
			Name:         item.Name,
			CreatedAt:    item.CreatedDate.AsTime().Format(time.RFC3339),
			LastAccessAt: item.LastAccess.AsTime().Format(time.RFC3339),
		})
	}

	result, err := json.Marshal(userData)

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
