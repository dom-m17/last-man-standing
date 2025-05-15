package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/peterldowns/pgtestdb"
)

// NewDB is a helper that returns an open connection to a unique and isolated
// test database, fully migrated and ready for you to query.
func NewDB(t *testing.T) *sql.DB {
	t.Helper()
	conf := pgtestdb.Config{
		DriverName: "pgx",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		Port:       "5433",
		Options:    "sslmode=disable",
	}
	migrator := MigrateWrapper{
		MigrationsDir: "../migration",
	}
	return pgtestdb.New(t, conf, migrator)
}

type MigrateWrapper struct {
	MigrationsDir string
}

func (m MigrateWrapper) Hash() (string, error) {
	// For simplicity, just return a fixed string or hash the contents of your migration directory
	return "always-run", nil
}

func (m MigrateWrapper) Migrate(ctx context.Context, db *sql.DB, conf pgtestdb.Config) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create migrate db driver: %w", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://"+m.MigrationsDir,
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("could not create migrator: %w", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
