package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	if err != nil {
		fmt.Errorf("Error running migrations %w", err)
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("Could not run migrations %w", err)
		}
	}

	return nil
}
