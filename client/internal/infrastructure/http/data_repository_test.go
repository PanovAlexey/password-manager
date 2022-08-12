package http

import (
	"client/internal/config"
	"client/internal/domain"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestGetUserDataRepository(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())

	type args struct {
		client ApiClient
	}
	tests := []struct {
		name string
		args args
		want userDataRepository
	}{
		{
			name: "successful getting user data repository test",
			args: args{client: client},
			want: GetUserDataRepository(client),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetUserDataRepository(tt.args.client), "GetUserDataRepository(%v)", tt.args.client)
		})
	}
}

func Test_userDataRepository_CreateBinaryRecord(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()
	timestampString := time.Now().Format("20060102150405")

	type fields struct {
		client ApiClient
	}
	type args struct {
		token        string
		binaryRecord domain.BinaryRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.BinaryRecord
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "successful creating binary record",
			fields: fields{client: client},
			args: args{
				token: token,
				binaryRecord: domain.BinaryRecord{
					Name:   "Test binary record  №" + timestampString,
					Note:   "created by auto tests",
					Binary: "01010101010101010101010101010101",
				},
			},
			want: domain.BinaryRecord{
				Name:   "Test binary record  №" + timestampString,
				Note:   "created by auto tests",
				Binary: "01010101010101010101010101010101",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		log.Println(tt)
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}

			got, err := r.CreateBinaryRecord(tt.args.token, tt.args.binaryRecord)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
			assert.Equal(t, tt.want.Name, got.Name)
		})
	}
}

func Test_userDataRepository_CreateCreditCard(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()
	timestampString := time.Now().Format("20060102150405")

	type fields struct {
		client ApiClient
	}
	type args struct {
		token      string
		creditCard domain.CreditCard
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.CreditCard
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "successful creating binary record",
			fields: fields{client: client},
			args: args{
				token: token,
				creditCard: domain.CreditCard{
					Name:       "Test credit card  №" + timestampString,
					Note:       "created by auto tests",
					Number:     "9123 3131 13131 14412",
					Expiration: "01/29",
					Owner:      "Mikluho Maclay",
					Cvv:        "029",
				},
			},
			want: domain.CreditCard{
				Name:       "Test credit card  №" + timestampString,
				Note:       "created by auto tests",
				Number:     "9123 3131 13131 14412",
				Expiration: "01/29",
				Owner:      "Mikluho Maclay",
				Cvv:        "029",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.CreateCreditCard(tt.args.token, tt.args.creditCard)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
			assert.Equal(t, tt.want.Name, got.Name)
		})
	}
}

func Test_userDataRepository_CreateLoginPassword(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()
	timestampString := time.Now().Format("20060102150405")

	type fields struct {
		client ApiClient
	}
	type args struct {
		token         string
		loginPassword domain.LoginPassword
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.LoginPassword
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "successful creating login password",
			fields: fields{client: client},
			args: args{
				token: token,
				loginPassword: domain.LoginPassword{
					Name:     "Test login password № " + timestampString,
					Note:     "created by auto tests",
					Login:    "root",
					Password: "qwerty",
				},
			},
			want: domain.LoginPassword{
				Name:     "Test login password № " + timestampString,
				Note:     "created by auto tests",
				Login:    "root",
				Password: "qwerty",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.CreateLoginPassword(tt.args.token, tt.args.loginPassword)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
			assert.Equal(t, tt.want.Name, got.Name)
		})
	}
}

func Test_userDataRepository_CreateTextRecord(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()
	timestampString := time.Now().Format("20060102150405")

	type fields struct {
		client ApiClient
	}
	type args struct {
		token      string
		textRecord domain.TextRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.TextRecord
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "successful creating text record",
			fields: fields{client: client},
			args: args{
				token: token,
				textRecord: domain.TextRecord{
					Name: "Test text record № " + timestampString,
					Note: "created by auto tests",
					Text: "abcdefghijklmnopqrstuvwxyz",
				},
			},
			want: &domain.TextRecord{
				Name: "Test text record № " + timestampString,
				Note: "created by auto tests",
				Text: "abcdefghijklmnopqrstuvwxyz",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.CreateTextRecord(tt.args.token, tt.args.textRecord)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
			assert.Equal(t, tt.want.Name, got.Name)
		})
	}
}

func Test_userDataRepository_GetAllData(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.UserData
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "successful getting all data",
			fields:  fields{client: client},
			args:    args{token: token},
			want:    &domain.UserData{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetAllData(tt.args.token)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func Test_userDataRepository_GetBinaryRecordById(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
		id    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.BinaryRecord
	}{
		{
			name:   "negative getting binary record by id",
			fields: fields{client: client},
			args: args{
				id:    "1",
				token: token,
			},
			want: &domain.BinaryRecord{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetBinaryRecordById(tt.args.token, tt.args.id)

			assert.NotEmpty(t, err)
			assert.Empty(t, got)
		})
	}
}

func Test_userDataRepository_GetBinaryRecordCollection(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "successful getting binary record collection",
			fields:  fields{client: client},
			args:    args{token: token},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetBinaryRecordCollection(tt.args.token)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func Test_userDataRepository_GetCreditCardById(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
		id    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.CreditCard
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "negative getting credit card by id",
			fields: fields{client: client},
			args: args{
				id:    "1",
				token: token,
			},
			want: &domain.CreditCard{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetCreditCardById(tt.args.token, tt.args.id)

			assert.NotEmpty(t, err)
			assert.Empty(t, got)
		})
	}
}

func Test_userDataRepository_GetCreditCardCollection(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ProtectedItem
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "successful getting credit card collection",
			fields:  fields{client: client},
			args:    args{token: token},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetCreditCardCollection(tt.args.token)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func Test_userDataRepository_GetLoginPasswordById(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
		id    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.LoginPassword
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "negative getting login password by id",
			fields: fields{client: client},
			args: args{
				id:    "1",
				token: token,
			},
			want: &domain.LoginPassword{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetLoginPasswordById(tt.args.token, tt.args.id)

			assert.NotEmpty(t, err)
			assert.Empty(t, got)
		})
	}
}

func Test_userDataRepository_GetLoginPasswordCollection(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ProtectedItem
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "successful getting login password collection",
			fields:  fields{client: client},
			args:    args{token: token},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetLoginPasswordCollection(tt.args.token)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func Test_userDataRepository_GetTextRecordById(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
		id    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.TextRecord
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "negative getting text record by id",
			fields: fields{client: client},
			args: args{
				id:    "1",
				token: token,
			},
			want: &domain.TextRecord{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetTextRecordById(tt.args.token, tt.args.id)

			assert.NotEmpty(t, err)
			assert.Empty(t, got)
		})
	}
}

func Test_userDataRepository_GetTextRecordCollection(t *testing.T) {
	config := config.New()
	client := GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserWithContentOrLogin()

	type fields struct {
		client ApiClient
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ProtectedItem
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "successful getting text record collection",
			fields:  fields{client: client},
			args:    args{token: token},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userDataRepository{
				client: tt.fields.client,
			}
			got, err := r.GetTextRecordCollection(tt.args.token)

			assert.Empty(t, err)
			assert.NotEmpty(t, got)
		})
	}
}
