package migrations

import (
	"embed"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

//go:embed *.sql
var MigrationsFS embed.FS

type Migrator struct {
	srcDriver source.Driver
}

func NewMigrator(dirName string) (*Migrator, error) {
	d, err := iofs.New(MigrationsFS, ".")
	if err != nil {
		return nil, err
	}

	return &Migrator{
		srcDriver: d,
	}, nil
}

func (m *Migrator) ApplyMigrations(db *sqlx.DB) error {
	sqlDB := db.DB // Так как использую *sqlx.DB получаю *sql.DB

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", m.srcDriver, "psql_db", driver)
	if err != nil {
		return err
	}

	defer func() {
		migrator.Close()
	}()

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
