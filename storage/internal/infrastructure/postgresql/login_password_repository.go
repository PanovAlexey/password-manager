package postgresql

import (
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

func (r loginPasswordRepository) Save(loginPassword domain.LoginPassword) (*domain.LoginPassword, error) {
	query := "INSERT INTO " +
		TableLoginPasswordName +
		" (email, password, created_at, last_access_at) VALUES ($1, $2, $3, $4) RETURNING id, email, created_at, last_access_at"
	err := r.DB.QueryRow(query, loginPassword.Password, time.Now(), time.Now()).
		Scan(&loginPassword.Id, &loginPassword.CreatedAt, &loginPassword.LastAccessAt)

	if err != nil {
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
			" SET last_access_at='" + time.Now().Format(time.RFC3339) +
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
