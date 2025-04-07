package database

import (
	"fmt"
	"github.com/ArtyomYatsenko/currency/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg config.DataBaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
