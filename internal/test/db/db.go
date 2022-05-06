package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"go-scheduler-api/internal/container"
)

type DatabaseTestFunc = func(tx *sqlx.Tx)

func WithTestDatabase(handler DatabaseTestFunc) {
	tx, err := testDatabaseInit()
	if err != nil {
		panic(fmt.Sprintf("failed to create a test database: %v", err))
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			panic(fmt.Sprintf("failed to rollback a transaction: %v", err))
		}
	}()

	handler(tx)
}

func testDatabaseInit() (*sqlx.Tx, error) {
	migration, err := container.NewMigrate()
	if err != nil {
		return nil, err
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	db, err := container.NewDatabase()
	if err != nil {
		return nil, err
	}

	return db.Beginx()
}
