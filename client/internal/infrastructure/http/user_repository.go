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
		content, err := r.client.getContentFromResponse(response)

		if err != nil {
			return "", errors.New(
				"registration error. can not parse answer. " +
					err.Error() +
					". server side status: " +
					strconv.Itoa(response.StatusCode),
			)
		} else {
			return "", errors.New(
				"registration error." + *content + ". server side status: " + strconv.Itoa(response.StatusCode),
			)
		}
	}

	content, err := r.client.getContentFromResponse(response)
	tokenValue := *content

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
		content, err := r.client.getContentFromResponse(response)

		if err != nil {
			return "", errors.New(
				"login error. can not parse answer. " + err.Error() +
					". server side status: " +
					strconv.Itoa(response.StatusCode),
			)
		} else {
			return "", errors.New(
				"login error." + *content + ". server side status: " + strconv.Itoa(response.StatusCode),
			)
		}
	}

	token, err := r.client.getContentFromResponse(response)
	tokenValue := *token

	return tokenValue, nil
}
