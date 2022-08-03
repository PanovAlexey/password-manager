package postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"storage/internal/domain"
	"strconv"
	"time"
)

type creditCardRepository struct {
	DB *sqlx.DB
}

func GetCreditCardRepository(DB *sqlx.DB) *creditCardRepository {
	return &creditCardRepository{DB: DB}
}

func (r creditCardRepository) Add(creditCard domain.CreditCard) (*domain.CreditCard, error) {
	query := "INSERT INTO " +
		TableCreditCardName +
		" (name, number, expiration, cvv, owner, note, user_id, created_at, last_access_at) VALUES" +
		" ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, number, expiration, cvv, owner, note, user_id, created_at, last_access_at"
	err := r.DB.QueryRow(
		query,
		creditCard.Name,
		creditCard.Number,
		creditCard.Expiration,
		creditCard.Cvv,
		creditCard.Owner,
		creditCard.Note,
		creditCard.UserId,
		time.Now(),
		time.Now(),
	).
		Scan(
			&creditCard.Id,
			&creditCard.Name,
			&creditCard.Number,
			&creditCard.Expiration,
			&creditCard.Cvv,
			&creditCard.Owner,
			&creditCard.Note,
			&creditCard.UserId,
			&creditCard.CreatedAt,
			&creditCard.LastAccessAt,
		)

	if err != nil {
		return nil, err
	}

	return &creditCard, err
}

func (r creditCardRepository) UpdateLastAccessAt(entityId int64) error {
	_, err := r.DB.Exec(
		"UPDATE " + TableCreditCardName +
			" SET last_access_at='" + time.Now().Format(time.RFC3339) +
			"' WHERE id=" + strconv.FormatInt(entityId, 10),
	)

	return err
}

func (r creditCardRepository) GetList(userId int) ([]domain.ProtectedItem, error) {
	creditCard := domain.ProtectedItem{}
	creditCardCollection := []domain.ProtectedItem{}

	query := "SELECT id, name, created_at, last_access_at FROM " + TableCreditCardName + " WHERE user_id = $1"
	rows, err := r.DB.Query(query, userId)

	if err != nil || rows.Err() != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&creditCard.Id,
			&creditCard.Name,
			&creditCard.CreatedAt,
			&creditCard.LastAccessAt,
		); err != nil {
			return creditCardCollection, err
		}

		creditCardCollection = append(creditCardCollection, creditCard)
	}

	return creditCardCollection, nil
}

func (r creditCardRepository) GetById(id, userId int) (*domain.CreditCard, error) {
	var creditCard domain.CreditCard

	err := r.DB.Get(
		&creditCard,
		"SELECT * FROM "+TableCreditCardName+" WHERE id = $1 and user_id = $2 LIMIT 1",
		id,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return &creditCard, err
}

func (r creditCardRepository) Delete(id, userId int) error {
	result := ""
	err := r.DB.Get(
		&result,
		"DELETE FROM "+TableCreditCardName+" WHERE id = $1 AND user_id = $2",
		id,
		userId,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r creditCardRepository) Update(creditCard domain.CreditCard) (*domain.CreditCard, error) {
	query := "UPDATE " +
		TableCreditCardName +
		" SET name=$1, number=$2, expiration=$3, cvv=$4, owner=$5, note=$6, user_id=$7, last_access_at=$8" +
		" WHERE id = $9 AND user_id = $10" +
		" RETURNING id, name, number, expiration, cvv, owner, note, user_id, created_at, last_access_at"
	err := r.DB.QueryRow(
		query,
		creditCard.Name,
		creditCard.Number,
		creditCard.Expiration,
		creditCard.Cvv,
		creditCard.Owner,
		creditCard.Note,
		creditCard.UserId,
		time.Now().Format(time.RFC3339),
		creditCard.Id,
		creditCard.UserId,
	).
		Scan(
			&creditCard.Id,
			&creditCard.Name,
			&creditCard.Number,
			&creditCard.Expiration,
			&creditCard.Cvv,
			&creditCard.Owner,
			&creditCard.Note,
			&creditCard.UserId,
			&creditCard.CreatedAt,
			&creditCard.LastAccessAt,
		)

	if err != nil {
		return nil, err
	}

	return &creditCard, err
}
