package domain

import "database/sql"

type User struct {
	ID        sql.NullInt64
	Email     string
	Password  string
	Token     string
	CreatedAt string
}
