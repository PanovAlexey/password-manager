package http

import (
	"client/internal/domain"
	"encoding/json"
	"io"
)

type userDataRepository struct {
	client ApiClient
}

func GetUserDataRepository(client ApiClient) userDataRepository {
	return userDataRepository{
		client: client,
	}
}

func (r userDataRepository) CreateLoginPassword(
	token string,
	loginPassword domain.LoginPassword,
) (*domain.LoginPassword, error) {
	response, err := r.client.client.Do(r.client.getCreateLoginPasswordRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var loginPasswordOut domain.LoginPassword

	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &loginPasswordOut)

	return &loginPasswordOut, err
}

func (r userDataRepository) CreateCreditCard(token string, creditCard domain.CreditCard) (*domain.CreditCard, error) {
	response, err := r.client.client.Do(r.client.getCreateCreditCardRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var creditCardOut domain.CreditCard
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &creditCardOut)

	return &creditCardOut, err
}

func (r userDataRepository) CreateTextRecord(token string, textRecord domain.TextRecord) (*domain.TextRecord, error) {
	response, err := r.client.client.Do(r.client.getCreateTextRecordRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var textRecordOut domain.TextRecord
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &textRecordOut)

	return &textRecordOut, err
}

func (r userDataRepository) CreateBinaryRecord(
	token string,
	binaryRecord domain.BinaryRecord,
) (*domain.BinaryRecord, error) {
	response, err := r.client.client.Do(r.client.getCreateBinaryRecordRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var binaryRecordOut domain.BinaryRecord
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &binaryRecordOut)

	return &binaryRecordOut, err
}

func (r userDataRepository) GetAllData(token string) (*domain.UserData, error) {
	response, err := r.client.client.Do(r.client.getAllDataRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var userData domain.UserData
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &userData)

	return &userData, err
}

func (r userDataRepository) GetLoginPasswordCollection(token string) ([]domain.ProtectedItem, error) {
	response, err := r.client.client.Do(r.client.getLoginPasswordRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var collection []domain.ProtectedItem
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &collection)

	return collection, err
}

func (r userDataRepository) GetCreditCardCollection(token string) ([]domain.ProtectedItem, error) {
	response, err := r.client.client.Do(r.client.getCreditCardRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var collection []domain.ProtectedItem
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &collection)

	return collection, err
}

func (r userDataRepository) GetTextRecordCollection(token string) ([]domain.ProtectedItem, error) {
	response, err := r.client.client.Do(r.client.getTextRecordRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var collection []domain.ProtectedItem
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &collection)

	return collection, err
}

func (r userDataRepository) GetBinaryRecordCollection(token string) ([]domain.ProtectedItem, error) {
	response, err := r.client.client.Do(r.client.getBinaryRecordRequest(token))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var collection []domain.ProtectedItem
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &collection)

	return collection, err
}

func (r userDataRepository) GetLoginPasswordById(token, id string) (*domain.LoginPassword, error) {
	response, err := r.client.client.Do(r.client.getLoginPasswordByIdRequest(token, id))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var item domain.LoginPassword
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &item)

	return &item, err
}

func (r userDataRepository) GetCreditCardById(token, id string) (*domain.CreditCard, error) {
	response, err := r.client.client.Do(r.client.getCreditCardByIdRequest(token, id))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var item domain.CreditCard
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &item)

	return &item, err
}

func (r userDataRepository) GetTextRecordById(token, id string) (*domain.TextRecord, error) {
	response, err := r.client.client.Do(r.client.getCreditCardByIdRequest(token, id))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var item domain.TextRecord
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &item)

	return &item, err
}

func (r userDataRepository) GetBinaryRecordById(token, id string) (*domain.BinaryRecord, error) {
	response, err := r.client.client.Do(r.client.getCreditCardByIdRequest(token, id))
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var item domain.BinaryRecord
	bodyJSON, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyJSON, &item)

	return &item, err
}
