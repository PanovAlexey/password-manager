package http

import (
	"api-gw/internal/handlers/http/dto"
	pb "api-gw/pkg/user_data_manager_grpc"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func (h *httpHandler) HandleCreateCreditCard(w http.ResponseWriter, r *http.Request) {
	userId := fmt.Sprintf("%v", r.Context().Value("token"))
	id := chi.URLParam(r, "id")

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("error creating credit cart by id: "+err.Error(), id, userId)
		return
	}

	createCreditCardDto := dto.CreateCreditCard{}
	err = json.Unmarshal(bodyJSON, &createCreditCardDto)

	if err != nil ||
		len(createCreditCardDto.Name) == 0 ||
		len(createCreditCardDto.Number) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("error creating credit card by wrong request: "+err.Error(), bodyJSON)
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
			UserId:           userId,
		},
	)

	if err != nil {
		h.logger.Error("error creating credit card by id: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful creating credit card by id ", id, userId, response.CreditCard.Name)
	result, err := json.Marshal(response)

	if err != nil {
		h.logger.Error("error marshalling credit card: "+err.Error(), id, userId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
