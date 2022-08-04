package service

import (
	"errors"
	"storage/internal/domain"
	"strconv"
)

type TextRecordRepository interface {
	Add(textRecord domain.TextRecord) (*domain.TextRecord, error)
	GetById(id, userId int) (*domain.TextRecord, error)
	GetList(userId int) ([]domain.ProtectedItem, error)
	UpdateLastAccessAt(entityId int64) error
	Delete(id, userId int) error
	Update(textRecord domain.TextRecord) (*domain.TextRecord, error)
}

type TextRecord struct {
	textRecordRepository TextRecordRepository
}

func GetTextRecordService(textRecordRepository TextRecordRepository) TextRecord {
	return TextRecord{textRecordRepository: textRecordRepository}
}

func (s TextRecord) GetTextRecordById(idString, userIdString string) (*domain.TextRecord, error) {
	id, err := strconv.Atoi(idString)

	if err != nil {
		return nil, errors.New("parsing text record id error: " + err.Error())
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return nil, errors.New("parsing user id error: " + err.Error())
	}

	textRecord, err := s.textRecordRepository.GetById(id, userId)

	if textRecord == nil || !textRecord.Id.Valid {
		return nil, errors.New("not found") //@ToDo: replace with custom error
	}

	err = s.textRecordRepository.UpdateLastAccessAt(textRecord.Id.Int64)

	if err != nil {
		return nil, err
	}

	return textRecord, err
}

func (s TextRecord) GetTextRecordList(userIdString string) ([]domain.ProtectedItem, error) {
	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return nil, errors.New("parsing user id error: " + err.Error())
	}

	return s.textRecordRepository.GetList(userId)
}

func (s TextRecord) AddTextRecord(textRecord domain.TextRecord, userId string) (*domain.TextRecord, error) {
	textRecord.UserId = userId

	return s.textRecordRepository.Add(textRecord)
}

func (s TextRecord) UpdateTextRecord(textRecord domain.TextRecord, userId string) (*domain.TextRecord, error) {
	textRecord.UserId = userId
	textRecordResult, err := s.textRecordRepository.Update(textRecord)
	err = s.textRecordRepository.UpdateLastAccessAt(textRecordResult.Id.Int64)

	if err != nil {
		return nil, err
	}

	return textRecordResult, nil
}

func (s TextRecord) DeleteTextRecord(idString, userIdString string) error {
	id, err := strconv.Atoi(idString)

	if err != nil {
		return errors.New("parsing id error: " + err.Error())
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return errors.New("parsing user id error: " + err.Error())
	}

	return s.textRecordRepository.Delete(id, userId)
}
