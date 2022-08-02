package domain

import (
	"database/sql"
)

type User struct {
	Id           sql.NullInt64  `db:"id"`
	Email        string         `db:"email"`
	Password     string         `db:"password"`
	CreatedAt    string         `db:"created_at"`
	LastAccessAt sql.NullString `db:"last_access_at"`
}
