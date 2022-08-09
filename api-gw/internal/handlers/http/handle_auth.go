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
		w.WriteHeader(http.StatusForbidden)
		h.showError(w, "error authorization: already logged in. ")
		return
	}

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.showError(w, "error auth: "+err.Error())
		return
	}

	authUserDto := dto.AuthUser{}
	err = json.Unmarshal(bodyJSON, &authUserDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.showError(w, "error user auth by wrong request: "+err.Error())
		return
	}

	// @ToDo: move validation to user service
	if len(authUserDto.Email) == 0 ||
		len(authUserDto.Password) == 0 {

		w.WriteHeader(http.StatusBadRequest)
		h.showError(w, "error user auth by wrong request: fields values are incorrect. ")
		return
	}

	userToken, err := h.userAuthorizationService.Auth(
		r.Context(),
		authUserDto.Email,
		authUserDto.Password,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.showError(w, "error user auth: "+err.Error())
		return
	}

	if userToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		h.showError(w, "user not found")
		return
	}

	h.logger.Info("successful user auth ", userToken)
	result, err := json.Marshal(userToken)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.showError(w, "error marshalling user: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
