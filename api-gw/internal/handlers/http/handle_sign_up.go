package http

import (
	"api-gw/internal/handlers/http/dto"
	"encoding/json"
	"io"
	"net/http"
)

func (h *httpHandler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	userId := h.userAuthorizationService.GetUserIdFromContext(r.Context())

	if !h.userAuthorizationService.IsUserIdEmpty(userId) {
		w.WriteHeader(http.StatusForbidden)
		h.showError(w, "error registration: already logged in. ")
		return
	}

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.showError(w, "error registration: "+err.Error())
		return
	}

	registerUserDto := dto.RegisterUser{}
	err = json.Unmarshal(bodyJSON, &registerUserDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.showError(w, "error user registration by wrong request: "+err.Error())
		return
	}

	// @ToDo: move validation from handler
	if len(registerUserDto.Email) == 0 ||
		len(registerUserDto.Password) == 0 ||
		len(registerUserDto.RepeatPassword) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		h.showError(w, "error user registration by wrong request: incorrect fields value")
		return
	}

	userToken, err := h.userAuthorizationService.Register(
		r.Context(),
		registerUserDto.Email,
		registerUserDto.Password,
		registerUserDto.RepeatPassword,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.showError(w, "error user registration: "+err.Error())
		return
	}

	h.logger.Info("successful user registration ", userToken)
	result, err := json.Marshal(userToken)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.showError(w, "error marshalling user: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
