package migrations

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"path/filepath"
)

type Migrator struct {
	migrationsPath string
	logger         *zap.Logger
}

func NewMigrator(migrationsPath string, logger *zap.Logger) (*Migrator, error) {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return nil, err
	}

	return &Migrator{
		migrationsPath: absPath,
		logger:         logger,
	}, nil
}

func (m *Migrator) ApplyMigrations(db *sqlx.DB) error {
	sqlDB := db.DB // Так как использую *sqlx.DB получаю *sql.DB

	// Создаём драйвер для файловой системы
	fileSource, err := (&file.File{}).Open(m.migrationsPath)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", fileSource, "psql_db", driver)
	if err != nil {
		return err
	}

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
