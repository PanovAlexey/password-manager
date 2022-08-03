package postgresql

import (
	"github.com/jmoiron/sqlx"
	"storage/internal/domain"
	"time"
)

type loginPasswordRepository struct {
	DB *sqlx.DB
}

func GetLoginPasswordRepository(DB *sqlx.DB) *loginPasswordRepository {
	return &loginPasswordRepository{DB: DB}
}

func (r loginPasswordRepository) SaveLoginPassword(loginPassword domain.LoginPassword) (*domain.User, error) {
	var user domain.User

	query := "INSERT INTO " +
		TableLoginPasswordName +
		" (email, password, created_at, last_access_at) VALUES ($1, $2, $3, $4) RETURNING id, email, created_at, last_access_at"
	err := r.DB.QueryRow(query, loginPassword.Password, time.Now(), time.Now()).
		Scan(&user.Id, &user.Email, &user.CreatedAt, &user.LastAccessAt)

	if err != nil {
		return nil, err
	}

	return &user, err
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

/* @ToDo: move to common repository
func (r loginPasswordRepository) UpdateLastAccessAt(entityId int64) error {
	_, err := r.DB.Exec(
		"UPDATE " + TableLoginPasswordName +
			" SET last_access_at='" + time.Now().Format(time.RFC3339) +
			"' WHERE id=" + strconv.FormatInt(entityId, 10),
	)

	return err
}*/

func (r loginPasswordRepository) GetList() ([]domain.LoginPassword, error) {
	loginPasswordCollection := []domain.LoginPassword{}

	query := "SELECT id, guid FROM " + TableLoginPasswordName
	rows, err := r.DB.Query(query)

	if err != nil || rows.Err() != nil {
		return loginPasswordCollection, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id   int64
			guid string
		)

		if err := rows.Scan(&id, &guid); err != nil {
			return loginPasswordCollection, err
		}

		loginPasswordCollection = append(loginPasswordCollection, domain.LoginPassword{})
	}

	return loginPasswordCollection, nil
}

func (r loginPasswordRepository) GetByID(id int) (domain.LoginPassword, error) {
	query := "SELECT id FROM " + TableLoginPasswordName + " WHERE id=($1)"
	row := r.DB.QueryRow(query, id)

	var userGUID string
	err := row.Scan(&userGUID)

	if err != nil {
		return domain.LoginPassword{}, err
	}

	return domain.LoginPassword{}, nil
}
