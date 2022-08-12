package http

import "net/http"

func (h *httpHandler) HandleRefreshAuthToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
