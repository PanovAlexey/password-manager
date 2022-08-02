package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"user-auth/internal/domain"
	"user-auth/internal/infrastructure/clients/grpc"
	grpcClient "user-auth/pkg/storage_grpc"
	pb "user-auth/pkg/user_authorization_grpc"
)

const (
	salt = "wertyuiopasdfghjkl"
)

type UserRegistration struct {
	storage grpc.StorageClient
}

func GetUserRegistrationService(storage grpc.StorageClient) UserRegistration {
	return UserRegistration{storage: storage}
}

func (s UserRegistration) Validate(user pb.RegisterUser) error {
	if user.Password != user.RepeatPassword {
		return errors.New("password and password confirmation must be identical")
	}

	return nil
}

func (s UserRegistration) Register(user pb.RegisterUser, ctx context.Context) (*domain.User, error) {
	passwordHash, err := s.generatePasswordHash(user.Password)

	if err != nil {
		return nil, err
	}

	createUser := grpcClient.CreateUser{
		Email:    user.Email,
		Password: passwordHash,
	}

	createUserResponse, err := (*s.storage.GetClient()).CreateUser(
		ctx,
		&grpcClient.CreateUserRequest{
			CreateUser: &createUser,
		},
	)

	if err != nil {
		return nil, err
	}

	outputUser := domain.User{
		Id:               createUserResponse.User.Id,
		Email:            createUserResponse.User.Email,
		RegistrationDate: createUserResponse.User.RegistrationDate,
		LastLogin:        createUserResponse.User.LastLogin,
	}

	return &outputUser, nil
}

func (s UserRegistration) generatePasswordHash(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), err
}
