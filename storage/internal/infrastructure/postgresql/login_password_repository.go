package postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"storage/internal/domain"
	"strconv"
	"time"
)

type loginPasswordRepository struct {
	DB *sqlx.DB
}

func GetLoginPasswordRepository(DB *sqlx.DB) *loginPasswordRepository {
	return &loginPasswordRepository{DB: DB}
}

func (r loginPasswordRepository) Add(loginPassword domain.LoginPassword) (*domain.LoginPassword, error) {
	query := "INSERT INTO " +
		TableLoginPasswordName +
		" (name, login, password, note, user_id, created_at, last_access_at) VALUES" +
		" ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, login, password, note, user_id, created_at, last_access_at"
	err := r.DB.QueryRow(
		query,
		loginPassword.Name,
		loginPassword.Login,
		loginPassword.Password,
		loginPassword.Note,
		loginPassword.UserId,
		time.Now(),
		time.Now(),
	).
		Scan(
			&loginPassword.Id,
			&loginPassword.Name,
			&loginPassword.Login,
			&loginPassword.Password,
			&loginPassword.Note,
			&loginPassword.UserId,
			&loginPassword.CreatedAt,
			&loginPassword.LastAccessAt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &loginPassword, err
}

/*
func (r loginPasswordRepository) GetLoginPassword(loginPassword domain.LoginPassword) (*domain.LoginPassword, error) {
	var user domain.User

	err := r.DB.Get(
		&user,
		"SELECT * FROM "+TableLoginPasswordName+" WHERE email = $1 and password = $2 LIMIT 1",
		loginPassword.Email,
		loginPassword.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, err
}
*/

func (r loginPasswordRepository) UpdateLastAccessAt(entityId int64) error {
	_, err := r.DB.Exec(
		"UPDATE " + TableLoginPasswordName +
			" SET last_access_at='" + time.Now().Format(time.RFC3339) + // @ToDo: Replace with prepared queries
			"' WHERE id=" + strconv.FormatInt(entityId, 10),
	)

	return err
}

func (r loginPasswordRepository) GetList(userId int) ([]domain.ProtectedItem, error) {
	loginPassword := domain.ProtectedItem{}
	loginPasswordCollection := []domain.ProtectedItem{}

	query := "SELECT id, name, created_at, last_access_at FROM " + TableLoginPasswordName + " WHERE user_id = $1"
	rows, err := r.DB.Query(query, userId)

	if err != nil || rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&loginPassword.Id,
			&loginPassword.Name,
			&loginPassword.CreatedAt,
			&loginPassword.LastAccessAt,
		); err != nil {
			return loginPasswordCollection, err
		}

		loginPasswordCollection = append(loginPasswordCollection, loginPassword)
	}

	return loginPasswordCollection, nil
}

func (r loginPasswordRepository) GetById(id, userId int) (*domain.LoginPassword, error) {
	var loginPassword domain.LoginPassword

	err := r.DB.Get(
		&loginPassword,
		"SELECT * FROM "+TableLoginPasswordName+" WHERE id = $1 and user_id = $2 LIMIT 1",
		id,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return &loginPassword, err
}

func (r loginPasswordRepository) Delete(id, userId int) error {
	result := ""
	err := r.DB.Get(
		&result,
		"DELETE FROM "+TableLoginPasswordName+" WHERE id = $1 AND user_id = $2",
		id,
		userId,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r loginPasswordRepository) Update(loginPassword domain.LoginPassword) (*domain.LoginPassword, error) {
	query := "UPDATE " +
		TableLoginPasswordName +
		" SET name=$1, login=$2, password=$3, note=$4, last_access_at=$5" +
		" WHERE id = $6 AND user_id = $7" +
		" RETURNING id, name, login, password, note, user_id, created_at, last_access_at"
	err := r.DB.QueryRow(
		query,
		loginPassword.Name,
		loginPassword.Login,
		loginPassword.Password,
		loginPassword.Note,
		time.Now().Format(time.RFC3339),
		loginPassword.Id,
		loginPassword.UserId,
	).
		Scan(
			&loginPassword.Id,
			&loginPassword.Name,
			&loginPassword.Login,
			&loginPassword.Password,
			&loginPassword.Note,
			&loginPassword.UserId,
			&loginPassword.CreatedAt,
			&loginPassword.LastAccessAt,
		)

	if err != nil {
		return nil, err
	}

	return &loginPassword, err
}
