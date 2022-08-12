package domain

import (
	"database/sql"
)

type CreditCard struct {
	Id           sql.NullInt64  `db:"id"`
	Note         string         `db:"note"`
	Name         string         `db:"name"`
	Number       string         `db:"number"`
	Expiration   string         `db:"expiration"`
	Cvv          string         `db:"cvv"`
	Owner        string         `db:"owner"`
	UserId       string         `db:"user_id"`
	CreatedAt    string         `db:"created_at"`
	LastAccessAt sql.NullString `db:"last_access_at"`
}
