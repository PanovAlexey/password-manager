package postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"storage/internal/domain"
	"strconv"
	"time"
)

type textRecordRepository struct {
	DB *sqlx.DB
}

func GetTextRecordRepository(DB *sqlx.DB) *textRecordRepository {
	return &textRecordRepository{DB: DB}
}

func (r textRecordRepository) Add(textRecord domain.TextRecord) (*domain.TextRecord, error) {
	query := "INSERT INTO " +
		TableTextRecordName +
		" (name, text, note, user_id, created_at, last_access_at) VALUES" +
		" ($1, $2, $3, $4, $5, $6) RETURNING id, name, text, note, user_id, created_at, last_access_at"
	err := r.DB.QueryRow(
		query,
		textRecord.Name,
		textRecord.Text,
		textRecord.Note,
		textRecord.UserId,
		time.Now(),
		time.Now(),
	).
		Scan(
			&textRecord.Id,
			&textRecord.Name,
			&textRecord.Text,
			&textRecord.Note,
			&textRecord.UserId,
			&textRecord.CreatedAt,
			&textRecord.LastAccessAt,
		)

	if err != nil {
		return nil, err
	}

	return &textRecord, err
}

func (r textRecordRepository) UpdateLastAccessAt(entityId int64) error {
	_, err := r.DB.Exec(
		"UPDATE " + TableTextRecordName +
			" SET last_access_at='" + time.Now().Format(time.RFC3339) + // @ToDo: Replace with prepared queries
			"' WHERE id=" + strconv.FormatInt(entityId, 10),
	)

	return err
}

func (r textRecordRepository) GetList(userId int) ([]domain.ProtectedItem, error) {
	textRecord := domain.ProtectedItem{}
	textRecordCollection := []domain.ProtectedItem{}

	query := "SELECT id, name, created_at, last_access_at FROM " + TableTextRecordName + " WHERE user_id = $1"
	rows, err := r.DB.Query(query, userId)

	if err != nil || rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&textRecord.Id,
			&textRecord.Name,
			&textRecord.CreatedAt,
			&textRecord.LastAccessAt,
		); err != nil {
			return textRecordCollection, err
		}

		textRecordCollection = append(textRecordCollection, textRecord)
	}

	return textRecordCollection, nil
}

func (r textRecordRepository) GetById(id, userId int) (*domain.TextRecord, error) {
	var textRecord domain.TextRecord

	err := r.DB.Get(
		&textRecord,
		"SELECT id, name, text, note, user_id, created_at, last_access_at FROM "+
			TableTextRecordName+" WHERE id = $1 and user_id = $2 LIMIT 1",
		id,
		userId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &textRecord, err
}

func (r textRecordRepository) Delete(id, userId int) error {
	result := ""
	err := r.DB.Get(
		&result,
		"DELETE FROM "+TableTextRecordName+" WHERE id = $1 AND user_id = $2",
		id,
		userId,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r textRecordRepository) Update(textRecord domain.TextRecord) (*domain.TextRecord, error) {
	query := "UPDATE " +
		TableTextRecordName +
		" SET name=$1, text=$2, note=$3, last_access_at=$4" +
		" WHERE id=$5 AND user_id=$6" +
		" RETURNING id, name, text, note, user_id, created_at, last_access_at"
	err := r.DB.QueryRow(
		query,
		textRecord.Name,
		textRecord.Text,
		textRecord.Note,
		time.Now().Format(time.RFC3339),
		textRecord.Id,
		textRecord.UserId,
	).
		Scan(
			&textRecord.Id,
			&textRecord.Name,
			&textRecord.Text,
			&textRecord.Note,
			&textRecord.UserId,
			&textRecord.CreatedAt,
			&textRecord.LastAccessAt,
		)

	if err != nil {
		return nil, err
	}

	return &textRecord, err
}
