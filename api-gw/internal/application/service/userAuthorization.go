package service

import (
	"api-gw/internal/infrastructure/clients/grpc"
	pb "api-gw/pkg/user_authorization_grpc"
	"context"
	"net/http"
)

const userTokenKey = "token"
const UserIdKey = "user-id"
const TraceIdKey = "trace-id"

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
	userTokenHeader := r.Header.Get(userTokenKey)

	return userTokenHeader
}

func (u UserAuthorization) IsUserTokenEmpty(userToken string) bool {
	return len(userToken) == 0
}

func (u UserAuthorization) IsUserIdEmpty(userId string) bool {
	return len(userId) == 0
}

func (u UserAuthorization) GetUserIdByToken(tokenInput string, ctx context.Context) (string, error) {
	userId := ""

	response, err := (*u.userAuthorizationClient.GetClient()).GetUserIdByJWT(
		ctx,
		&pb.GetUserIdByJWTRequest{
			Token: tokenInput,
		},
	)

	if err != nil {
		return userId, err
	}

	userId = response.UserId

	return userId, err
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

	if err != nil {
		return "", err
	}

	return response.User.Token, err
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

	if err != nil {
		return "", err
	}

	return response.User.Token, nil
}

func (u UserAuthorization) GetUserIdFromContext(ctx context.Context) string {
	userToken := ""
	if ctx.Value(UserIdKey) != nil {
		userToken = ctx.Value(UserIdKey).(string)
	}

	return userToken
}

func (u UserAuthorization) SetUserIdInContext(userId string, ctx context.Context) context.Context {
	return context.WithValue(ctx, UserIdKey, userId)
}
