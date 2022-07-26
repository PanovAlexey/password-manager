package service

import (
	"api-gw/internal/infrastructure/clients/grpc"
	pb "api-gw/pkg/user_authorization_grpc"
	"context"
	"net/http"
)

const userTokenKey = "token"
const userIdKey = "user-id"

type logger interface {
	Error(args ...interface{})
}

type UserAuthorization struct {
	logger                  logger
	userAuthorizationClient grpc.UserAuthorizationClient
}

func GetUserAuthorizationService(logger logger, userAuthorizationClient grpc.UserAuthorizationClient) UserAuthorization {
	return UserAuthorization{
		logger:                  logger,
		userAuthorizationClient: userAuthorizationClient,
	}
}

func (u UserAuthorization) GetTokenFromHeader(r *http.Request) string {
	userToken := ``
	userTokenCookie, err := r.Cookie(userTokenKey)

	if err != nil {
		if err != http.ErrNoCookie {
			u.logger.Error("error with getting token from cookie: " + err.Error())
		}

		return userToken
	}

	userToken = (*userTokenCookie).Value

	return userToken
}

func (u UserAuthorization) IsUserTokenEmpty(userToken string) bool {
	return len(userToken) == 0
}

func (u UserAuthorization) GetUserIdByToken(tokenInput string, ctx context.Context) (string, error) {
	token := pb.Token{Token: tokenInput}
	response, err := (*u.userAuthorizationClient.GetClient()).GetUserIdByJWT(
		ctx,
		&pb.GetUserIdByJWTRequest{
			Token: &token,
		},
	)

	return response.UserId, err
}

func (u UserAuthorization) Auth(ctx context.Context, userEmail, userPassword string) (string, error) {
	authUser := pb.AuthUser{
		Email:    userEmail,
		Password: userPassword,
	}

	response, err := (*u.userAuthorizationClient.GetClient()).Auth(
		ctx,
		&pb.AuthRequest{
			AuthUser: &authUser,
		},
	)

	return response.User.Token.Token, err
}

func (u UserAuthorization) Register(ctx context.Context, userEmail, userPassword, repeatUserPassword string) (string, error) {
	registerUser := pb.RegisterUser{
		Email:          userEmail,
		Password:       userPassword,
		RepeatPassword: repeatUserPassword,
	}

	response, err := (*u.userAuthorizationClient.GetClient()).Register(
		ctx,
		&pb.RegisterRequest{
			RegisterUser: &registerUser,
		},
	)

	return response.User.Token.Token, err
}

func (u UserAuthorization) SetUserIdInContext(userId string, ctx context.Context) context.Context {
	return context.WithValue(ctx, userIdKey, userId)
}
