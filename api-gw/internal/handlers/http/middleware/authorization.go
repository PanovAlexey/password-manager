package middleware

import (
	"context"
	"net/http"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//@ToDo: add JWT authorization
		userTokenName := "token"
		userToken := "temporary stub token value"
		ctx := context.WithValue(r.Context(), userTokenName, userToken)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
