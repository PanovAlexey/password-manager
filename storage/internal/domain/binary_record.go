package domain

import (
	"database/sql"
)

type BinaryRecord struct {
	Id           sql.NullInt64  `db:"id"`
	Note         string         `db:"note"`
	Name         string         `db:"name"`
	Binary       string         `db:"binary_data"`
	UserId       string         `db:"user_id"`
	CreatedAt    string         `db:"created_at"`
	LastAccessAt sql.NullString `db:"last_access_at"`
}
