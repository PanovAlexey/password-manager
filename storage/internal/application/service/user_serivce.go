package service

import (
	"storage/internal/domain"
)

type UserRepository interface {
	SaveUser(user domain.UserLogin) (*domain.User, error)
	GetUser(user domain.UserLogin) (*domain.User, error)
	UpdateLastAccessAt(entityId int64) error
}

type UserService struct {
	userRepository UserRepository
}

func GetUserService(
	userRepository UserRepository,
) UserService {
	return UserService{userRepository: userRepository}
}

func (s UserService) SaveUser(user domain.UserLogin) (*domain.User, error) {
	return s.userRepository.SaveUser(user)
}

func (s UserService) GetUser(userLogin domain.UserLogin) (*domain.User, error) {
	user, err := s.userRepository.GetUser(userLogin)

	if err != nil {
		return nil, err
	}

	err = s.userRepository.UpdateLastAccessAt(user.Id.Int64)

	if err != nil {
		return nil, err
	}

	return user, err
}
