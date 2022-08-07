package http

import (
	"client/internal/domain"
	"errors"
	"net/http"
	"strconv"
)

type userRepository struct {
	client ApiClient
}

func GetUserRepository(client ApiClient) userRepository {
	return userRepository{
		client: client,
	}
}

func (r userRepository) Register(user domain.User) (string, error) {
	response, err := r.client.client.Do(r.client.getRegistrationRequest(user))

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("registration error. server side status: " + strconv.Itoa(response.StatusCode))
	}

	token, err := r.client.getTokenFromResponse(response)
	tokenValue := *token

	return tokenValue, nil
}

func (r userRepository) Auth(user domain.User) (string, error) {
	authRequest := r.client.getAuthRequest(user)
	response, err := r.client.client.Do(authRequest)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New("login error. server side status: " + strconv.Itoa(response.StatusCode))
	}

	token, err := r.client.getTokenFromResponse(response)
	tokenValue := *token

	return tokenValue, nil
}
