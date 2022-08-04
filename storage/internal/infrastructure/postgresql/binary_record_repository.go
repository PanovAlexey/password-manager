package postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"storage/internal/domain"
	"strconv"
	"time"
)

type binaryRecordRepository struct {
	DB *sqlx.DB
}

func GetBinaryRecordRepository(DB *sqlx.DB) *binaryRecordRepository {
	return &binaryRecordRepository{DB: DB}
}

func (r binaryRecordRepository) Add(binaryRecord domain.BinaryRecord) (*domain.BinaryRecord, error) {
	query := "INSERT INTO " +
		TableBinaryRecordName +
		" (name, binary_data, note, user_id, created_at, last_access_at) VALUES" +
		" ($1, $2, $3, $4, $5, $6) RETURNING id, name, binary_data, note, user_id, created_at, last_access_at"
	err := r.DB.QueryRow(
		query,
		binaryRecord.Name,
		binaryRecord.Binary,
		binaryRecord.Note,
		binaryRecord.UserId,
		time.Now(),
		time.Now(),
	).
		Scan(
			&binaryRecord.Id,
			&binaryRecord.Name,
			&binaryRecord.Binary,
			&binaryRecord.Note,
			&binaryRecord.UserId,
			&binaryRecord.CreatedAt,
			&binaryRecord.LastAccessAt,
		)

	if err != nil {
		return nil, err
	}

	return &binaryRecord, err
}

func (r binaryRecordRepository) UpdateLastAccessAt(entityId int64) error {
	_, err := r.DB.Exec(
		"UPDATE " + TableBinaryRecordName +
			" SET last_access_at='" + time.Now().Format(time.RFC3339) + // @ToDo: Replace with prepared queries
			"' WHERE id=" + strconv.FormatInt(entityId, 10),
	)

	return err
}

func (r binaryRecordRepository) GetList(userId int) ([]domain.ProtectedItem, error) {
	binaryRecord := domain.ProtectedItem{}
	binaryRecordCollection := []domain.ProtectedItem{}

	query := "SELECT id, name, created_at, last_access_at FROM " + TableBinaryRecordName + " WHERE user_id = $1"
	rows, err := r.DB.Query(query, userId)

	if err != nil || rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&binaryRecord.Id,
			&binaryRecord.Name,
			&binaryRecord.CreatedAt,
			&binaryRecord.LastAccessAt,
		); err != nil {
			return binaryRecordCollection, err
		}

		binaryRecordCollection = append(binaryRecordCollection, binaryRecord)
	}

	return binaryRecordCollection, nil
}

func (r binaryRecordRepository) GetById(id, userId int) (*domain.BinaryRecord, error) {
	var binaryRecord domain.BinaryRecord

	err := r.DB.Get(
		&binaryRecord,
		"SELECT id, name, binary_data, note, user_id, created_at, last_access_at FROM "+
			TableBinaryRecordName+" WHERE id = $1 and user_id = $2 LIMIT 1",
		id,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return &binaryRecord, err
}

func (r binaryRecordRepository) Delete(id, userId int) error {
	result := ""
	err := r.DB.Get(
		&result,
		"DELETE FROM "+TableBinaryRecordName+" WHERE id = $1 AND user_id = $2",
		id,
		userId,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r binaryRecordRepository) Update(binaryRecord domain.BinaryRecord) (*domain.BinaryRecord, error) {
	query := "UPDATE " +
		TableBinaryRecordName +
		" SET name=$1, binary_data=$2, note=$3, last_access_at=$4" +
		" WHERE id=$5 AND user_id=$6" +
		" RETURNING id, name, binary_data, note, user_id, created_at, last_access_at"
	err := r.DB.QueryRow(
		query,
		binaryRecord.Name,
		binaryRecord.Binary,
		binaryRecord.Note,
		time.Now().Format(time.RFC3339),
		binaryRecord.Id,
		binaryRecord.UserId,
	).
		Scan(
			&binaryRecord.Id,
			&binaryRecord.Name,
			&binaryRecord.Binary,
			&binaryRecord.Note,
			&binaryRecord.UserId,
			&binaryRecord.CreatedAt,
			&binaryRecord.LastAccessAt,
		)

	if err != nil {
		return nil, err
	}

	return &binaryRecord, err
}
