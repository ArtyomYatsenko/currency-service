package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type CurrencyRepository struct {
	db *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{
		db: db,
	}
}

func (c *CurrencyRepository) AddCurrency(date map[string]interface{}) error {

	dateNow, ok := date["date"].(string)
	if !ok {
		return nil
	}

	currencies, ok := date["rub"].(map[string]interface{})
	if !ok {
		return nil
	}

	query := "INSERT INTO currencies (other_currency, basic_currency, meaning, created_date) VALUES "

	placeholders := make([]string, 0, len(currencies))
	args := make([]interface{}, 0, len(currencies))

	i := 1

	for currency, value := range currencies {

		args = append(args, currency)
		args = append(args, "rub")
		args = append(args, value)
		args = append(args, dateNow)

		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", i, i+1, i+2, i+3))
		i += 4
	}

	query += strings.Join(placeholders, ", ")

	_, err := c.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
