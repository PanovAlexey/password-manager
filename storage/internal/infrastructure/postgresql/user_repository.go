package postgresql

import (
	"github.com/jmoiron/sqlx"
	"storage/internal/domain"
	"time"
)

type userRepository struct {
	DB *sqlx.DB
}

func GetUserRepository(DB *sqlx.DB) *userRepository {
	return &userRepository{DB: DB}
}

func (r userRepository) SaveUser(user domain.User) (int, error) {
	var insertedID int

	query := "INSERT INTO " + TableUsersName + " (email, password, created_at) VALUES ($1, $2, $3) RETURNING id"
	err := r.DB.QueryRow(query, user.Email, user.Password, time.Now()).Scan(&insertedID)

	if err != nil {
		return 0, err // ToDo: 0 - is a crutch
	}

	return insertedID, err
}
