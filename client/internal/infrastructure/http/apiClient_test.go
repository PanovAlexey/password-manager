package http

import (
	"client/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestApiClient_getTextRecordByIdRequest(t *testing.T) {
	config := config.New()

	type fields struct {
		serverAddress string
		client        http.Client
	}
	type args struct {
		token string
		id    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "successfully getting text record by id request test",
			fields: fields{
				serverAddress: config.GetServerAddress(),
			},
			args: args{
				token: "token",
				id:    "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ApiClient{
				serverAddress: tt.fields.serverAddress,
				client:        tt.fields.client,
			}
			request := c.getTextRecordByIdRequest(tt.args.token, tt.args.id)

			assert.NotEmpty(t, request.Header.Get("Content-Type"))
			assert.NotEmpty(t, request.Header.Get("Token"))
			assert.Equal(t, request.Header.Get("Token"), "token")
		})
	}
}

func TestGetApiClient(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())

	type args struct {
		serverAddress      string
		maxIdleConnections int
		httpTimeout        int
	}
	tests := []struct {
		name string
		args args
		want ApiClient
	}{
		{
			name: "successfully getting api client test",
			args: args{
				serverAddress:      config.GetServerAddress(),
				maxIdleConnections: config.GetMaxIdleConnections(),
				httpTimeout:        config.GetHttpTimeout(),
			},
			want: client,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetApiClient(tt.args.serverAddress, tt.args.maxIdleConnections, tt.args.httpTimeout), "GetApiClient(%v, %v, %v)", tt.args.serverAddress, tt.args.maxIdleConnections, tt.args.httpTimeout)
		})
	}
}
