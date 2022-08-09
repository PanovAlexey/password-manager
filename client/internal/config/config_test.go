package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_SetToken(t *testing.T) {
	type fields struct {
		serverAddress      string
		userToken          string
		httpTimeout        int
		maxIdleConnections int
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success test with empty token",
			fields: fields{
				serverAddress:      "localhost",
				userToken:          "",
				httpTimeout:        30,
				maxIdleConnections: 30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				serverAddress:      tt.fields.serverAddress,
				userToken:          tt.fields.userToken,
				httpTimeout:        tt.fields.httpTimeout,
				maxIdleConnections: tt.fields.maxIdleConnections,
			}
			token := c.GetToken()
			c.SetToken(tt.args.token)
			assert.Equal(t, token, c.GetToken())
		})
	}
}

func TestConfig_GetHttpTimeout(t *testing.T) {
	type fields struct {
		serverAddress      string
		userToken          string
		httpTimeout        int
		maxIdleConnections int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Success test with getting http timeout",
			fields: fields{
				serverAddress:      "localhost",
				userToken:          "token",
				httpTimeout:        30,
				maxIdleConnections: 30,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				serverAddress:      tt.fields.serverAddress,
				userToken:          tt.fields.userToken,
				httpTimeout:        tt.fields.httpTimeout,
				maxIdleConnections: tt.fields.maxIdleConnections,
			}
			assert.Equalf(t, tt.want, c.GetHttpTimeout(), "GetHttpTimeout()")
		})
	}
}

func TestConfig_GetMaxIdleConnections(t *testing.T) {
	type fields struct {
		serverAddress      string
		userToken          string
		httpTimeout        int
		maxIdleConnections int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Success test with getting max idle connections",
			fields: fields{
				serverAddress:      "localhost",
				userToken:          "token",
				httpTimeout:        30,
				maxIdleConnections: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				serverAddress:      tt.fields.serverAddress,
				userToken:          tt.fields.userToken,
				httpTimeout:        tt.fields.httpTimeout,
				maxIdleConnections: tt.fields.maxIdleConnections,
			}
			assert.Equalf(t, tt.want, c.GetMaxIdleConnections(), "GetMaxIdleConnections()")
		})
	}
}

func TestConfig_GetServerAddress(t *testing.T) {
	type fields struct {
		serverAddress      string
		userToken          string
		httpTimeout        int
		maxIdleConnections int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Success test with getting server address",
			fields: fields{
				serverAddress:      "localhost",
				userToken:          "token",
				httpTimeout:        30,
				maxIdleConnections: 10,
			},
			want: "localhost",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				serverAddress:      tt.fields.serverAddress,
				userToken:          tt.fields.userToken,
				httpTimeout:        tt.fields.httpTimeout,
				maxIdleConnections: tt.fields.maxIdleConnections,
			}
			assert.Equalf(t, tt.want, c.GetServerAddress(), "GetServerAddress()")
		})
	}
}

func TestConfig_GetToken(t *testing.T) {
	type fields struct {
		serverAddress      string
		userToken          string
		httpTimeout        int
		maxIdleConnections int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Success test with getting token",
			fields: fields{
				serverAddress:      "localhost",
				userToken:          "token",
				httpTimeout:        30,
				maxIdleConnections: 10,
			},
			want: "token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				serverAddress:      tt.fields.serverAddress,
				userToken:          tt.fields.userToken,
				httpTimeout:        tt.fields.httpTimeout,
				maxIdleConnections: tt.fields.maxIdleConnections,
			}
			assert.Equalf(t, tt.want, c.GetToken(), "GetToken()")
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Config
	}{
		{
			name: "Success test with getting new Config",
			want: Config{
				serverAddress:      "http://0.0.0.0:8081",
				userToken:          "",
				httpTimeout:        3,
				maxIdleConnections: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, New(), "New()")
		})
	}
}
