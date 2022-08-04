package service

import (
	"errors"
	"storage/internal/domain"
	"strconv"
)

type CreditCardRepository interface {
	Add(creditCard domain.CreditCard) (*domain.CreditCard, error)
	GetById(id, userId int) (*domain.CreditCard, error)
	GetList(userId int) ([]domain.ProtectedItem, error)
	UpdateLastAccessAt(entityId int64) error
	Delete(id, userId int) error
	Update(creditCard domain.CreditCard) (*domain.CreditCard, error)
}

type CreditCard struct {
	creditCardRepository CreditCardRepository
}

func GetCreditCardService(creditCardRepository CreditCardRepository) CreditCard {
	return CreditCard{creditCardRepository: creditCardRepository}
}

func (s CreditCard) GetCreditCardById(idString, userIdString string) (*domain.CreditCard, error) {
	id, err := strconv.Atoi(idString)

	if err != nil {
		return nil, errors.New("parsing credit card id error: " + err.Error())
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return nil, errors.New("parsing user id error: " + err.Error())
	}

	creditCard, err := s.creditCardRepository.GetById(id, userId)

	if creditCard == nil || !creditCard.Id.Valid {
		return nil, errors.New("not found") //@ToDo: replace with custom error
	}

	err = s.creditCardRepository.UpdateLastAccessAt(creditCard.Id.Int64)

	if err != nil {
		return nil, err
	}

	return creditCard, err

}

func (s CreditCard) GetCreditCardList(userIdString string) ([]domain.ProtectedItem, error) {
	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return nil, errors.New("parsing user id error: " + err.Error())
	}

	return s.creditCardRepository.GetList(userId)
}

func (s CreditCard) AddCreditCard(creditCard domain.CreditCard, userId string) (*domain.CreditCard, error) {
	creditCard.UserId = userId

	return s.creditCardRepository.Add(creditCard)
}

func (s CreditCard) UpdateCreditCard(creditCard domain.CreditCard, userId string) (*domain.CreditCard, error) {
	creditCard.UserId = userId
	creditCardResult, err := s.creditCardRepository.Update(creditCard)
	err = s.creditCardRepository.UpdateLastAccessAt(creditCardResult.Id.Int64)

	if err != nil {
		return nil, err
	}

	return creditCardResult, nil
}

func (s CreditCard) DeleteCreditCard(idString, userIdString string) error {
	id, err := strconv.Atoi(idString)

	if err != nil {
		return errors.New("parsing id error: " + err.Error())
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return errors.New("parsing user id error: " + err.Error())
	}

	return s.creditCardRepository.Delete(id, userId)
}
