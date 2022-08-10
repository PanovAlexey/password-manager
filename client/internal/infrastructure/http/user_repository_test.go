package http

import (
	"client/internal/application"
	"client/internal/config"
	"client/internal/domain"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

// create user and getting correct user token for tests
// this user has content
func createUserWithContentOrLogin() string {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	userRepository := GetUserRepository(client)
	userService := application.GetUserService(userRepository)

	token, err := userService.Register(getTestUserWithCorrectFieldsValue())

	if err != nil || token == "" {
		token, _ = userService.Auth(getTestUserWithCorrectFieldsValue())
	}

	return token
}

func getTestUserWithCorrectFieldsValue() domain.User {
	return domain.User{
		Email:          "test@test-r.test",
		Password:       "qwerty",
		RepeatPassword: "qwerty",
	}
}

func TestGetUserRepository(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())

	type args struct {
		client ApiClient
	}
	tests := []struct {
		name string
		args args
		want userRepository
	}{
		{
			name: "successful getting user repository test",
			args: args{client: client},
			want: GetUserRepository(client),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserRepository(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Auth(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "successful auth by repository test",
			fields: fields{client: client},
			args: args{
				user: getTestUserWithCorrectFieldsValue(),
			},
			want:    token,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				client: tt.fields.client,
			}
			got, err := r.Auth(tt.args.user)
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

func Test_userRepository_Register(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	timestampString := time.Now().String()

	user := getTestUserWithCorrectFieldsValue()
	user.Email = user.Email + timestampString

	type fields struct {
		client ApiClient
	}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "successful registration in repository test",
			fields: fields{client: client},
			args: args{
				user: user,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				client: tt.fields.client,
			}
			got, err := r.Register(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEmpty(t, got)
		})
	}
}

func Test_userRepository_Register_by_different_password_and_confirmation(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())

	type fields struct {
		client ApiClient
	}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "negative registration by different password and confirmation in repository test",
			fields: fields{client: client},
			args: args{
				user: domain.User{
					Email:          "test-wrong@test.test",
					Password:       "qwerty",
					RepeatPassword: "qwerty2",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				client: tt.fields.client,
			}
			got, err := r.Register(tt.args.user)
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
