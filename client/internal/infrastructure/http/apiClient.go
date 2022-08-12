package http

import (
	"bytes"
	"client/internal/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ApiClient struct {
	serverAddress string
	client        http.Client
}

func GetApiClient(serverAddress string, maxIdleConnections, httpTimeout int) ApiClient {
	return ApiClient{
		serverAddress: serverAddress,
		client:        getHttpClient(maxIdleConnections, httpTimeout),
	}
}

func getHttpClient(maxIdleConnections, httpTimeout int) http.Client {
	client := http.Client{}
	transport := &http.Transport{}
	transport.MaxIdleConns = maxIdleConnections
	client.Transport = transport
	client.Timeout = time.Second * time.Duration(httpTimeout)

	return client
}

func (c ApiClient) getRegistrationRequest(user domain.User) *http.Request {
	body, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}

	request, err := http.NewRequest(http.MethodPost, c.serverAddress+"/api/v1/signup", bytes.NewBuffer(body))

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(body)))

	return request
}

func (c ApiClient) getAuthRequest(user domain.User) *http.Request {
	body, err := json.Marshal(user)

	if err != nil {
		log.Println(err)
	}

	request, err := http.NewRequest(http.MethodPost, c.serverAddress+"/api/v1/auth", bytes.NewBuffer(body))

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(body)))

	return request
}

func (c ApiClient) getAllDataRequest(token string) *http.Request {
	request, err := http.NewRequest(
		http.MethodGet,
		c.serverAddress+"/api/v1/data/all",
		nil,
	)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getLoginPasswordRequest(token string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, c.serverAddress+"/api/v1/data/login-password", nil)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getTextRecordRequest(token string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, c.serverAddress+"/api/v1/data/text-record", nil)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getBinaryRecordRequest(token string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, c.serverAddress+"/api/v1/data/binary-record", nil)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getCreditCardRequest(token string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, c.serverAddress+"/api/v1/data/credit-card", nil)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getCreateLoginPasswordRequest(token string, loginPassword domain.LoginPassword) *http.Request {
	body, err := json.Marshal(loginPassword)

	if err != nil {
		log.Println(err)
	}

	request, err := http.NewRequest(
		http.MethodPost,
		c.serverAddress+"/api/v1/data/login-password",
		bytes.NewBuffer(body),
	)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(body)))
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getCreateTextRecordRequest(token string, textRecord domain.TextRecord) *http.Request {
	body, err := json.Marshal(textRecord)

	if err != nil {
		log.Println(err)
	}

	request, err := http.NewRequest(
		http.MethodPost,
		c.serverAddress+"/api/v1/data/text-record",
		bytes.NewBuffer(body),
	)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(body)))
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getCreateBinaryRecordRequest(token string, binaryRecord domain.BinaryRecord) *http.Request {
	body, err := json.Marshal(binaryRecord)

	if err != nil {
		log.Println(err)
	}

	request, err := http.NewRequest(
		http.MethodPost,
		c.serverAddress+"/api/v1/data/binary-record",
		bytes.NewBuffer(body),
	)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(body)))
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getCreateCreditCardRequest(token string, creditCard domain.CreditCard) *http.Request {
	body, err := json.Marshal(creditCard)

	if err != nil {
		log.Println(err)
	}

	request, err := http.NewRequest(http.MethodPost, c.serverAddress+"/api/v1/data/credit-card", bytes.NewBuffer(body))

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(body)))
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getLoginPasswordByIdRequest(token, id string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, c.serverAddress+"/api/v1/data/login-password/"+id, nil)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getTextRecordByIdRequest(token, id string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, c.serverAddress+"/api/v1/data/text-record/"+id, nil)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getBinaryRecordByIdRequest(token, id string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, c.serverAddress+"/api/v1/data/binary-record/"+id, nil)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getCreditCardByIdRequest(token, id string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, c.serverAddress+"/api/v1/data/credit-card/"+id, nil)

	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Token", token)

	return request
}

func (c ApiClient) getContentFromResponse(response *http.Response) (*string, error) {
	content := ""
	err := json.NewDecoder(response.Body).Decode(&content)

	if err != nil {
		return nil, err
	}

	return &content, nil
}
