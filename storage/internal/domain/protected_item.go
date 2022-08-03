package domain

import "database/sql"

type ProtectedItem struct {
	Id           sql.NullInt64  `db:"id"`
	Name         string         `db:"name"`
	CreatedAt    string         `db:"created_at"`
	LastAccessAt sql.NullString `db:"last_access_at"`
}
