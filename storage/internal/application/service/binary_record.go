package service

import (
	"errors"
	"fmt"
	customErrors "storage/internal/application/errors"
	"storage/internal/domain"
	"strconv"
)

type BinaryRecordRepository interface {
	Add(binaryRecord domain.BinaryRecord) (*domain.BinaryRecord, error)
	GetById(id, userId int) (*domain.BinaryRecord, error)
	GetList(userId int) ([]domain.ProtectedItem, error)
	UpdateLastAccessAt(entityId int64) error
	Delete(id, userId int) error
	Update(binaryRecord domain.BinaryRecord) (*domain.BinaryRecord, error)
}

type BinaryRecord struct {
	binaryRecordRepository BinaryRecordRepository
}

func GetBinaryRecordService(binaryRecordRepository BinaryRecordRepository) BinaryRecord {
	return BinaryRecord{binaryRecordRepository: binaryRecordRepository}
}

func (s BinaryRecord) GetBinaryRecordById(idString, userIdString string) (*domain.BinaryRecord, error) {
	id, err := strconv.Atoi(idString)

	if err != nil {
		return nil, errors.New("parsing binary record id error: " + err.Error())
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return nil, errors.New("parsing user id error: " + err.Error())
	}

	binaryRecord, err := s.binaryRecordRepository.GetById(id, userId)

	if binaryRecord == nil || !binaryRecord.Id.Valid {
		return nil, fmt.Errorf("%v: %w", id, customErrors.ErrNotFound)
	}

	err = s.binaryRecordRepository.UpdateLastAccessAt(binaryRecord.Id.Int64)

	if err != nil {
		return nil, err
	}

	return binaryRecord, err
}

func (s BinaryRecord) GetBinaryRecordList(userIdString string) ([]domain.ProtectedItem, error) {
	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return nil, errors.New("parsing user id error: " + err.Error())
	}

	return s.binaryRecordRepository.GetList(userId)
}

func (s BinaryRecord) AddBinaryRecord(binaryRecord domain.BinaryRecord, userId string) (*domain.BinaryRecord, error) {
	binaryRecord.UserId = userId

	return s.binaryRecordRepository.Add(binaryRecord)
}

func (s BinaryRecord) UpdateBinaryRecord(binaryRecord domain.BinaryRecord, userId string) (*domain.BinaryRecord, error) {
	binaryRecord.UserId = userId
	binaryRecordResult, err := s.binaryRecordRepository.Update(binaryRecord)
	err = s.binaryRecordRepository.UpdateLastAccessAt(binaryRecordResult.Id.Int64)

	if err != nil {
		return nil, err
	}

	return binaryRecordResult, nil
}

func (s BinaryRecord) DeleteBinaryRecord(idString, userIdString string) error {
	id, err := strconv.Atoi(idString)

	if err != nil {
		return errors.New("parsing id error: " + err.Error())
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return errors.New("parsing user id error: " + err.Error())
	}

	return s.binaryRecordRepository.Delete(id, userId)
}
