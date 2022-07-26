package http

import (
	"api-gw/internal/handlers/http/dto"
	"encoding/json"
	"io"
	"net/http"
)

func (h *httpHandler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	userId := h.userAuthorizationService.GetUserIdFromContext(r.Context())

	if !h.userAuthorizationService.IsUserIdEmpty(userId) {
		h.logger.Debug("error registration: already logged in. ", userId)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("error auth: " + err.Error())
		return
	}

	authUserDto := dto.AuthUser{}
	err = json.Unmarshal(bodyJSON, &authUserDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("error user auth by wrong request: "+err.Error(), string(bodyJSON))
		return
	}

	// @ToDo: move validation from handler
	if len(authUserDto.Email) == 0 ||
		len(authUserDto.Password) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		h.logger.Error("error user auth by wrong request: fields values are incorrect.", string(bodyJSON))
		return
	}

	userToken, err := h.userAuthorizationService.Auth(
		r.Context(),
		authUserDto.Email,
		authUserDto.Password,
	)

	if err != nil {
		h.logger.Error("error user auth: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("successful user auth ", userToken)
	result, err := json.Marshal(userToken)

	if err != nil {
		h.logger.Error("error marshalling user: "+err.Error(), userToken)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
