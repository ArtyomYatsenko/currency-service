package migrations

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"path/filepath"
)

type Migrator struct {
	migrationsPath string
}

func NewMigrator(migrationsPath string) (*Migrator, error) {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return nil, err
	}

	return &Migrator{
		migrationsPath: absPath,
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

	//defer func() {  // Закомментил, не знаю как правильно его закрыть и нужно ли, так как если закрываю здесь, то закрывается подключение к БД
	//	migrator.Close()
	//}()

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
