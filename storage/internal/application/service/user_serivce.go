package services

import (
	"storage/internal/domain"
)

type UserRepository interface {
	SaveUser(user domain.User) (int, error)
}

type UserService struct {
	userRepository UserRepository
}

func GetUserService(
	userRepository UserRepository,
) UserService {
	return UserService{userRepository: userRepository}
}

func (s UserService) SaveUser(user domain.User) (int, error) {
	return s.userRepository.SaveUser(user)
}
