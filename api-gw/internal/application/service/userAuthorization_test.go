package service

import (
	"api-gw/internal/config"
	"api-gw/internal/infrastructure/clients/grpc"
	"api-gw/internal/infrastructure/logging"
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestUserAuthorization_Auth(t *testing.T) {
	config := config.New()
	log := logging.GetLogger(config)
	defer log.Sync()

	userAuthorizationClient, _ := grpc.GetUserAuthorizationClient(config)
	defer userAuthorizationClient.GetConnection().Close()

	type fields struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}
	type args struct {
		ctx          context.Context
		userEmail    string
		userPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "failed authorization",
			fields:  fields{log, userAuthorizationClient},
			args:    args{ctx: context.Background(), userEmail: "email", userPassword: "password"},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserAuthorization{
				logger:                  tt.fields.logger,
				userAuthorizationClient: tt.fields.userAuthorizationClient,
			}
			got, err := u.Auth(tt.args.ctx, tt.args.userEmail, tt.args.userPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Auth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAuthorization_GetTokenFromHeader(t *testing.T) {
	type fields struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserAuthorization{
				logger:                  tt.fields.logger,
				userAuthorizationClient: tt.fields.userAuthorizationClient,
			}
			if got := u.GetTokenFromHeader(tt.args.r); got != tt.want {
				t.Errorf("GetTokenFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAuthorization_GetUserIdByToken(t *testing.T) {
	type fields struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}
	type args struct {
		tokenInput string
		ctx        context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserAuthorization{
				logger:                  tt.fields.logger,
				userAuthorizationClient: tt.fields.userAuthorizationClient,
			}
			got, err := u.GetUserIdByToken(tt.args.tokenInput, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserIdByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserIdByToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAuthorization_GetUserIdFromContext(t *testing.T) {
	type fields struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserAuthorization{
				logger:                  tt.fields.logger,
				userAuthorizationClient: tt.fields.userAuthorizationClient,
			}
			if got := u.GetUserIdFromContext(tt.args.ctx); got != tt.want {
				t.Errorf("GetUserIdFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAuthorization_IsUserIdEmpty(t *testing.T) {
	type fields struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserAuthorization{
				logger:                  tt.fields.logger,
				userAuthorizationClient: tt.fields.userAuthorizationClient,
			}
			if got := u.IsUserIdEmpty(tt.args.userId); got != tt.want {
				t.Errorf("IsUserIdEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAuthorization_IsUserTokenEmpty(t *testing.T) {
	type fields struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}
	type args struct {
		userToken string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserAuthorization{
				logger:                  tt.fields.logger,
				userAuthorizationClient: tt.fields.userAuthorizationClient,
			}
			if got := u.IsUserTokenEmpty(tt.args.userToken); got != tt.want {
				t.Errorf("IsUserTokenEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAuthorization_Register(t *testing.T) {
	type fields struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}
	type args struct {
		ctx                context.Context
		userEmail          string
		userPassword       string
		repeatUserPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserAuthorization{
				logger:                  tt.fields.logger,
				userAuthorizationClient: tt.fields.userAuthorizationClient,
			}
			got, err := u.Register(tt.args.ctx, tt.args.userEmail, tt.args.userPassword, tt.args.repeatUserPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAuthorization_SetUserIdInContext(t *testing.T) {
	type fields struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}
	type args struct {
		userId string
		ctx    context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserAuthorization{
				logger:                  tt.fields.logger,
				userAuthorizationClient: tt.fields.userAuthorizationClient,
			}
			if got := u.SetUserIdInContext(tt.args.userId, tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserIdInContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserAuthorizationService(t *testing.T) {
	config := config.New()
	log := logging.GetLogger(config)
	defer log.Sync()

	userAuthorizationClient, _ := grpc.GetUserAuthorizationClient(config)
	defer userAuthorizationClient.GetConnection().Close()
	userAuthorizationService := GetUserAuthorizationService(log, userAuthorizationClient)
	//

	type args struct {
		logger                  logger
		userAuthorizationClient grpc.UserAuthorizationClient
	}

	tests := []struct {
		name string
		args args
		want UserAuthorization
	}{
		{
			name: "successful creation",
			args: args{log, userAuthorizationClient},
			want: userAuthorizationService,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserAuthorizationService(tt.args.logger, tt.args.userAuthorizationClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserAuthorizationService() = %v, want %v", got, tt.want)
			}
		})
	}
}
