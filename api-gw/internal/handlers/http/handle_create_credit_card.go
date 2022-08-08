package http

import (
	"api-gw/internal/application/service"
	"api-gw/internal/domain"
	"api-gw/internal/handlers/http/dto"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"time"
)

func (h *httpHandler) HandleCreateCreditCard(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value(service.UserIdKey))
	id := chi.URLParam(r, "id")

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		info := "error creating credit cart by id: " + err.Error()
		h.logger.Error(info, id, userId)
		w.Write([]byte(info))
		return
	}

	createCreditCardDto := dto.CreateCreditCard{}
	err = json.Unmarshal(bodyJSON, &createCreditCardDto)

	if err != nil ||
		len(createCreditCardDto.Name) == 0 ||
		len(createCreditCardDto.Number) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		info := "error creating credit card by wrong request: " + err.Error()
		h.logger.Error(info, bodyJSON)
		w.Write([]byte(info))
		return
	}

	createCreditCard := pb.CreateCreditCard{
		Name:       createCreditCardDto.Name,
		Number:     createCreditCardDto.Number,
		Expiration: createCreditCardDto.Expiration,
		Cvv:        createCreditCardDto.Cvv,
		Owner:      createCreditCardDto.Owner,
		Note:       createCreditCardDto.Note,
	}

	response, err := (*h.gRPCUserDataManagerClient.GetClient()).CreateCreditCard(
		r.Context(),
		&pb.CreateCreditCardRequest{
			CreateCreditCard: &createCreditCard,
		},
	)

	if err != nil {
		info := "error creating credit card by id: " + err.Error()
		h.logger.Error(info, id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful creating credit card by id ", id, userId, response.CreditCard.Name)

	creditCard := domain.CreditCard{
		Id:           response.CreditCard.Id,
		Name:         response.CreditCard.Name,
		Number:       response.CreditCard.Number,
		Cvv:          response.CreditCard.Cvv,
		Expiration:   response.CreditCard.Expiration,
		Owner:        response.CreditCard.Owner,
		Note:         response.CreditCard.Note,
		CreatedAt:    response.CreditCard.CreatedDate.AsTime().Format(time.RFC3339),
		LastAccessAt: response.CreditCard.LastAccess.AsTime().Format(time.RFC3339),
	}

	result, err := json.Marshal(creditCard)

	if err != nil {
		info := "error marshalling credit card: " + err.Error()
		h.logger.Error(info, id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
