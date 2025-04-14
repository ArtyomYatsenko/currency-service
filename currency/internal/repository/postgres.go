package repository

import (
	"fmt"
	"github.com/ArtyomYatsenko/currency/internal/dto"
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

func (c *CurrencyRepository) AddCurrencies(data dto.Currency) error {

	currencies := data.Rub

	query := "INSERT INTO currencies (other_currency, basic_currency, meaning, created_date) VALUES "

	placeholders := make([]string, 0, len(currencies))
	args := make([]interface{}, 0, len(currencies))

	i := 1

	for currency, value := range currencies {
		args = append(args, currency, "rub", value, data.Date)
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
