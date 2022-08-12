package domain

import (
	"database/sql"
)

type TextRecord struct {
	Id           sql.NullInt64  `db:"id"`
	Note         string         `db:"note"`
	Name         string         `db:"name"`
	Text         string         `db:"text"`
	UserId       string         `db:"user_id"`
	CreatedAt    string         `db:"created_at"`
	LastAccessAt sql.NullString `db:"last_access_at"`
}
