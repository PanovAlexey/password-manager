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
		info := "error registration: already logged in. "
		h.logger.Debug(info, userId)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(info))
		return
	}

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		info := "error registration: " + err.Error()
		h.logger.Error(info)
		w.Write([]byte(info))
		return
	}

	registerUserDto := dto.RegisterUser{}
	err = json.Unmarshal(bodyJSON, &registerUserDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		info := "error user registration by wrong request: "
		h.logger.Error(info, err, string(bodyJSON))
		w.Write([]byte(info))
		return
	}

	// @ToDo: move validation from handler
	if len(registerUserDto.Email) == 0 ||
		len(registerUserDto.Password) == 0 ||
		len(registerUserDto.RepeatPassword) == 0 {

		info := "error user registration by wrong request: incorrect fields value"
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error(info, string(bodyJSON))
		w.Write([]byte(info))
		return
	}

	userToken, err := h.userAuthorizationService.Register(
		r.Context(),
		registerUserDto.Email,
		registerUserDto.Password,
		registerUserDto.RepeatPassword,
	)

	if err != nil {
		info := "error user registration: " + err.Error()
		h.logger.Error(info)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful user registration ", userToken)
	result, err := json.Marshal(userToken)

	if err != nil {
		info := "error marshalling user: " + err.Error()
		h.logger.Error(info, userToken)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
