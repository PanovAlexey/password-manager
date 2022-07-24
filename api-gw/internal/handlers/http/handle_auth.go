package http

import "net/http"

func (h *httpHandler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
