package postgresql

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"storage/internal/config"
)

const TableUsersName = `users`
const TableLoginPasswordName = `login_password`
const TableCreditCardName = `credit_card`

type databaseService struct {
	db *sqlx.DB
}

func GetDatabaseService(config config.Config) (*databaseService, error) {
	databaseService := databaseService{}
	db, err := databaseService.initDatabaseConnection(config)

	if err != nil {
		return nil, err
	}

	databaseService.db = db

	return &databaseService, nil
}

func (s databaseService) GetDatabaseConnection() *sqlx.DB {
	return s.db
}

func (s databaseService) initDatabaseConnection(c config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		"postgresql://"+c.GetDatabaseUser()+":"+c.GetDatabasePassword()+"@0.0.0.0:"+c.GetDatabasePort()+"/"+c.GetDatabaseName()+"?sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	s.db = db
	err = s.db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
