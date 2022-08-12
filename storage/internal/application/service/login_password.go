package service

import (
	"errors"
	"fmt"
	customErrors "storage/internal/application/errors"
	"storage/internal/domain"
	"strconv"
)

type LoginPasswordRepository interface {
	Add(loginPassword domain.LoginPassword) (*domain.LoginPassword, error)
	GetById(id, userId int) (*domain.LoginPassword, error)
	GetList(userId int) ([]domain.ProtectedItem, error)
	UpdateLastAccessAt(entityId int64) error
	Delete(id, userId int) error
	Update(loginPassword domain.LoginPassword) (*domain.LoginPassword, error)
}

type LoginPassword struct {
	loginPasswordRepository LoginPasswordRepository
}

func GetLoginPasswordService(loginPasswordRepository LoginPasswordRepository) LoginPassword {
	return LoginPassword{loginPasswordRepository: loginPasswordRepository}
}

func (s LoginPassword) GetLoginPasswordById(idString, userIdString string) (*domain.LoginPassword, error) {
	id, err := strconv.Atoi(idString)

	if err != nil {
		return nil, errors.New("parsing login password id error: " + err.Error())
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return nil, errors.New("parsing user id error: " + err.Error())
	}

	loginPassword, err := s.loginPasswordRepository.GetById(id, userId)

	if loginPassword == nil || !loginPassword.Id.Valid {
		return nil, fmt.Errorf("%v: %w", id, customErrors.ErrNotFound)
	}

	err = s.loginPasswordRepository.UpdateLastAccessAt(loginPassword.Id.Int64)

	if err != nil {
		return nil, err
	}

	return loginPassword, err

}

func (s LoginPassword) GetLoginPasswordList(userIdString string) ([]domain.ProtectedItem, error) {
	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return nil, errors.New("parsing user id error: " + err.Error())
	}

	return s.loginPasswordRepository.GetList(userId)
}

func (s LoginPassword) AddLoginPassword(loginPassword domain.LoginPassword, userId string) (*domain.LoginPassword, error) {
	loginPassword.UserId = userId

	return s.loginPasswordRepository.Add(loginPassword)
}

func (s LoginPassword) UpdateLoginPassword(loginPassword domain.LoginPassword, userId string) (*domain.LoginPassword, error) {
	loginPassword.UserId = userId
	loginPasswordResult, err := s.loginPasswordRepository.Update(loginPassword)
	err = s.loginPasswordRepository.UpdateLastAccessAt(loginPasswordResult.Id.Int64)

	if err != nil {
		return nil, err
	}

	return loginPasswordResult, nil
}

func (s LoginPassword) DeleteLoginPassword(idString, userIdString string) error {
	id, err := strconv.Atoi(idString)

	if err != nil {
		return errors.New("parsing id error: " + err.Error())
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return errors.New("parsing user id error: " + err.Error())
	}

	return s.loginPasswordRepository.Delete(id, userId)
}
