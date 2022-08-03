package authorization_by_token

import (
	"api-gw/internal/application/service"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type Logger interface {
	Error(args ...interface{})
}

func AuthorizationByToken(userAuthorizationService service.UserAuthorization, logger Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := userAuthorizationService.GetTokenFromHeader(r)

			if userAuthorizationService.IsUserTokenEmpty(token) {
				next.ServeHTTP(w, r)
				return
			}

			userId, err := userAuthorizationService.GetUserIdByToken(token, r.Context())

			if err != nil {
				logger.Error("getting user id by token error: " + err.Error())
			} else {
				ctx := userAuthorizationService.SetUserIdInContext(userId, r.Context())

				// @ToDo: move to service
				// Transferring userId to other microservices by grpc
				ctx = metadata.NewOutgoingContext(
					ctx,
					metadata.New(map[string]string{
						service.UserIdKey: userId,
					}),
				)

				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}
