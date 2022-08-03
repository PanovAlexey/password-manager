package service

import (
	"errors"
	"storage/internal/domain"
	"strconv"
)

type LoginPasswordRepository interface {
	Save(loginPassword domain.LoginPassword) (*domain.LoginPassword, error)
	GetById(id, userId int) (*domain.LoginPassword, error)
	GetList() ([]domain.LoginPassword, error)
	UpdateLastAccessAt(entityId int64) error
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

	return s.loginPasswordRepository.GetById(id, userId)
}
