package application

import (
	"client/internal/config"
	"client/internal/domain"
	"client/internal/infrastructure/http"
	"reflect"
	"testing"
)

func getTestUserWithCorrectFieldsValue() domain.User {
	return domain.User{
		Email:          "test@test.test",
		Password:       "qwerty",
		RepeatPassword: "qwerty",
	}
}

// create user and getting correct user token for tests
func createUserOrLogin() string {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	userRepository := http.GetUserRepository(client)
	userService := GetUserService(userRepository)

	token, err := userService.Register(getTestUserWithCorrectFieldsValue())

	if err != nil || token == "" {
		token, _ = userService.Auth(getTestUserWithCorrectFieldsValue())
	}

	return token
}

func TestGetUserDataService(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	userDataRepository := http.GetUserDataRepository(client)

	type args struct {
		repository dataRepository
	}
	tests := []struct {
		name string
		args args
		want UserDataService
	}{
		{
			name: "success getting user data service",
			args: args{repository: userDataRepository},
			want: GetUserDataService(userDataRepository),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserDataService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserDataService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserService(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	userRepositoryStruct := http.GetUserRepository(client)

	type args struct {
		repository userRepository
	}
	tests := []struct {
		name string
		args args
		want UserService
	}{
		{
			name: "success getting user service",
			args: args{repository: userRepositoryStruct},
			want: GetUserService(userRepositoryStruct),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_CreateBinaryRecord(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	userDataRepository := http.GetUserDataRepository(client)

	type fields struct {
		repository dataRepository
	}
	type args struct {
		token        string
		binaryRecord domain.BinaryRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.BinaryRecord
		wantErr bool
	}{
		{
			name:    "negative getting binary record for wrong user token",
			fields:  fields{repository: userDataRepository},
			args:    args{token: "token"},
			want:    &domain.BinaryRecord{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.CreateBinaryRecord(tt.args.token, tt.args.binaryRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBinaryRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBinaryRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_CreateCreditCard(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	userDataRepository := http.GetUserDataRepository(client)

	type fields struct {
		repository dataRepository
	}
	type args struct {
		token      string
		creditCard domain.CreditCard
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.CreditCard
		wantErr bool
	}{
		{
			name:   "successful creating credit card",
			fields: fields{repository: userDataRepository},
			args: args{
				token:      "token",
				creditCard: domain.CreditCard{},
			},
			want:    &domain.CreditCard{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.CreateCreditCard(tt.args.token, tt.args.creditCard)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCreditCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCreditCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_CreateLoginPassword(t *testing.T) {
	type fields struct {
		repository dataRepository
	}
	type args struct {
		token         string
		loginPassword domain.LoginPassword
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.LoginPassword
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.CreateLoginPassword(tt.args.token, tt.args.loginPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateLoginPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLoginPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_CreateTextRecord(t *testing.T) {
	type fields struct {
		repository dataRepository
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.CreateTextRecord(tt.args.token, tt.args.textRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTextRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTextRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetAllData(t *testing.T) {
	type fields struct {
		repository dataRepository
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.UserData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetAllData(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetBinaryRecordById(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository dataRepository
	}
	type args struct {
		token string
		id    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.BinaryRecord
		wantErr bool
	}{
		{
			name:   "negative getting binary record by id",
			fields: fields{repository: http.GetUserDataRepository(client)},
			args: args{
				id:    "1",
				token: token,
			},
			want:    &domain.BinaryRecord{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetBinaryRecordById(tt.args.token, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBinaryRecordById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBinaryRecordById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetBinaryRecordCollection(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository dataRepository
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ProtectedItem
		wantErr bool
	}{
		{
			name:    "successful getting binary record collection",
			fields:  fields{repository: http.GetUserDataRepository(client)},
			args:    args{token: token},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetBinaryRecordCollection(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBinaryRecordCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBinaryRecordCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetCreditCardById(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository dataRepository
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
		wantErr bool
	}{
		{
			name:   "negative getting credit card by id",
			fields: fields{repository: http.GetUserDataRepository(client)},
			args: args{
				id:    "1",
				token: token,
			},
			want:    &domain.CreditCard{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetCreditCardById(tt.args.token, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCreditCardById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCreditCardById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetCreditCardCollection(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository dataRepository
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ProtectedItem
		wantErr bool
	}{
		{
			name:    "successful getting credit card collection",
			fields:  fields{repository: http.GetUserDataRepository(client)},
			args:    args{token: token},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetCreditCardCollection(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCreditCardCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCreditCardCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetLoginPasswordById(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository dataRepository
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
		wantErr bool
	}{
		{
			name:   "negative getting login password by id",
			fields: fields{repository: http.GetUserDataRepository(client)},
			args: args{
				id:    "1",
				token: token,
			},
			want:    &domain.LoginPassword{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetLoginPasswordById(tt.args.token, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoginPasswordById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLoginPasswordById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetLoginPasswordCollection(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository dataRepository
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ProtectedItem
		wantErr bool
	}{
		{
			name:    "successful getting login password collection",
			fields:  fields{repository: http.GetUserDataRepository(client)},
			args:    args{token: token},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetLoginPasswordCollection(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoginPasswordCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLoginPasswordCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetTextRecordById(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository dataRepository
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
		wantErr bool
	}{
		{
			name:   "negative getting text record by id",
			fields: fields{repository: http.GetUserDataRepository(client)},
			args: args{
				id:    "1",
				token: token,
			},
			want:    &domain.TextRecord{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetTextRecordById(tt.args.token, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTextRecordById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTextRecordById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDataService_GetTextRecordCollection(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository dataRepository
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.ProtectedItem
		wantErr bool
	}{
		{
			name:    "successful getting text record collection",
			fields:  fields{repository: http.GetUserDataRepository(client)},
			args:    args{token: token},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserDataService{
				repository: tt.fields.repository,
			}
			got, err := s.GetTextRecordCollection(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTextRecordCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTextRecordCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Auth(t *testing.T) {
	config := config.New()
	client := http.GetApiClient(config.GetServerAddress(), config.GetMaxIdleConnections(), config.GetHttpTimeout())
	token := createUserOrLogin()

	type fields struct {
		repository userRepository
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
			name:    "successful user authorization",
			fields:  fields{repository: http.GetUserRepository(client)},
			args:    args{user: getTestUserWithCorrectFieldsValue()},
			want:    token,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				repository: tt.fields.repository,
			}
			got, err := s.Auth(tt.args.user)
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
