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
		info := "error authorization: already logged in. "
		h.logger.Debug(info, userId)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(info))
		return
	}

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		info := "error auth: " + err.Error()
		h.logger.Error(info)
		w.Write([]byte(info))
		return
	}

	authUserDto := dto.AuthUser{}
	err = json.Unmarshal(bodyJSON, &authUserDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		info := "error user auth by wrong request: " + err.Error()
		h.logger.Error(info, string(bodyJSON))
		w.Write([]byte(info))
		return
	}

	// @ToDo: move validation to user service
	if len(authUserDto.Email) == 0 ||
		len(authUserDto.Password) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		info := "error user auth by wrong request: fields values are incorrect. "
		h.logger.Error(info, string(bodyJSON))
		w.Write([]byte(info))
		return
	}

	userToken, err := h.userAuthorizationService.Auth(
		r.Context(),
		authUserDto.Email,
		authUserDto.Password,
	)

	if err != nil {
		info := "error user auth: " + err.Error()
		h.logger.Error(info)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(info))
		return
	}

	h.logger.Info("successful user auth ", userToken)
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
