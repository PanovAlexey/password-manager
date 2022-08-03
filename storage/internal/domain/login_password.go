package domain

import (
	"database/sql"
)

type LoginPassword struct {
	Id           sql.NullInt64  `db:"id"`
	Note         string         `db:"note"`
	Name         string         `db:"name"`
	Login        string         `db:"login"`
	Password     string         `db:"password"`
	UserId       string         `db:"user_id"`
	CreatedAt    string         `db:"created_at"`
	LastAccessAt sql.NullString `db:"last_access_at"`
}
