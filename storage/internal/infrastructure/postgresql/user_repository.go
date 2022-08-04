package postgresql

import (
	"github.com/jmoiron/sqlx"
	"storage/internal/domain"
	"strconv"
	"time"
)

type userRepository struct {
	DB *sqlx.DB
}

func GetUserRepository(DB *sqlx.DB) *userRepository {
	return &userRepository{DB: DB}
}

func (r userRepository) SaveUser(userLogin domain.UserLogin) (*domain.User, error) {
	var user domain.User

	query := "INSERT INTO " +
		TableUsersName +
		" (email, password, created_at, last_access_at) VALUES ($1, $2, $3, $4) RETURNING id, email, created_at, last_access_at"
	err := r.DB.QueryRow(query, userLogin.Email, userLogin.Password, time.Now(), time.Now()).
		Scan(&user.Id, &user.Email, &user.CreatedAt, &user.LastAccessAt)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r userRepository) GetUser(userLogin domain.UserLogin) (*domain.User, error) {
	var user domain.User

	err := r.DB.Get(
		&user,
		"SELECT * FROM "+TableUsersName+" WHERE email = $1 and password = $2 LIMIT 1",
		userLogin.Email,
		userLogin.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r userRepository) UpdateLastAccessAt(entityId int64) error {
	_, err := r.DB.Exec(
		"UPDATE " + TableUsersName +
			" SET last_access_at='" + time.Now().Format(time.RFC3339) + // @ToDo: Replace with prepared queries
			"' WHERE id=" + strconv.FormatInt(entityId, 10),
	)

	return err
}
