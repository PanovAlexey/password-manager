package http

import (
	"net/http"
)

func (h *httpHandler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
