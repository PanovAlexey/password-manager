package closed_by_authorization

import (
	"api-gw/internal/application/service"
	"net/http"
)

func ClosedByAuthorization(userAuthorizationService service.UserAuthorization) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := userAuthorizationService.GetUserIdFromContext(r.Context())

			if userAuthorizationService.IsUserIdEmpty(userId) {
				http.Error(w, "it is forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
