package application

import "client/internal/domain"

type dataRepository interface {
	CreateLoginPassword(token string, loginPassword domain.LoginPassword) (*domain.LoginPassword, error)
	CreateCreditCard(token string, creditCard domain.CreditCard) (*domain.CreditCard, error)
	CreateTextRecord(token string, textRecord domain.TextRecord) (*domain.TextRecord, error)
	CreateBinaryRecord(token string, binaryRecord domain.BinaryRecord) (*domain.BinaryRecord, error)
	GetAllData(token string) (*domain.UserData, error)
	GetLoginPasswordCollection(token string) ([]domain.ProtectedItem, error)
	GetCreditCardCollection(token string) ([]domain.ProtectedItem, error)
	GetTextRecordCollection(token string) ([]domain.ProtectedItem, error)
	GetBinaryRecordCollection(token string) ([]domain.ProtectedItem, error)
	GetLoginPasswordById(token, id string) (*domain.LoginPassword, error)
	GetCreditCardById(token, id string) (*domain.CreditCard, error)
	GetTextRecordById(token, id string) (*domain.TextRecord, error)
	GetBinaryRecordById(token, id string) (*domain.BinaryRecord, error)
}

type UserDataService struct {
	repository dataRepository
}

func GetUserDataService(repository dataRepository) UserDataService {
	return UserDataService{repository: repository}
}

func (s UserDataService) CreateLoginPassword(
	token string,
	loginPassword domain.LoginPassword,
) (*domain.LoginPassword, error) {
	return s.repository.CreateLoginPassword(token, loginPassword)
}

func (s UserDataService) CreateCreditCard(token string, creditCard domain.CreditCard) (*domain.CreditCard, error) {
	return s.repository.CreateCreditCard(token, creditCard)
}

func (s UserDataService) CreateTextRecord(token string, textRecord domain.TextRecord) (*domain.TextRecord, error) {
	return s.repository.CreateTextRecord(token, textRecord)
}

func (s UserDataService) CreateBinaryRecord(
	token string,
	binaryRecord domain.BinaryRecord,
) (*domain.BinaryRecord, error) {
	return s.repository.CreateBinaryRecord(token, binaryRecord)
}

func (s UserDataService) GetAllData(token string) (*domain.UserData, error) {
	return s.repository.GetAllData(token)
}

func (s UserDataService) GetLoginPasswordCollection(token string) ([]domain.ProtectedItem, error) {
	return s.repository.GetLoginPasswordCollection(token)
}

func (s UserDataService) GetCreditCardCollection(token string) ([]domain.ProtectedItem, error) {
	return s.repository.GetCreditCardCollection(token)
}

func (s UserDataService) GetTextRecordCollection(token string) ([]domain.ProtectedItem, error) {
	return s.repository.GetTextRecordCollection(token)
}

func (s UserDataService) GetBinaryRecordCollection(token string) ([]domain.ProtectedItem, error) {
	return s.repository.GetBinaryRecordCollection(token)
}

func (s UserDataService) GetLoginPasswordById(token, id string) (*domain.LoginPassword, error) {
	return s.repository.GetLoginPasswordById(token, id)
}

func (s UserDataService) GetCreditCardById(token, id string) (*domain.CreditCard, error) {
	return s.repository.GetCreditCardById(token, id)
}

func (s UserDataService) GetTextRecordById(token, id string) (*domain.TextRecord, error) {
	return s.repository.GetTextRecordById(token, id)
}

func (s UserDataService) GetBinaryRecordById(token, id string) (*domain.BinaryRecord, error) {
	return s.repository.GetBinaryRecordById(token, id)
}
