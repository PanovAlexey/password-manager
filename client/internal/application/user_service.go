package application

import "client/internal/domain"

type userRepository interface {
	Register(user domain.User) (string, error)
	Auth(user domain.User) (string, error)
}

type UserService struct {
	repository userRepository
}

func GetUserService(repository userRepository) UserService {
	return UserService{repository: repository}
}

func (s UserService) Register(user domain.User) (string, error) {
	return s.repository.Register(user)
}

func (s UserService) Auth(user domain.User) (string, error) {
	return s.repository.Auth(user)
}
