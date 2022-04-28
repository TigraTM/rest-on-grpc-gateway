package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"

	// Driver.
	_ "github.com/lib/pq"
)

// Error names.
const (
	PostgresUniqueViolation     = "unique_violation"
	PostgresForeignKeyViolation = "foreign_key_violation"
)

const postgresDriverName = "postgres"

func NewPostgres(ctx context.Context, dsn, migrationDir, moduleName string) (_ *Repo, err error) {
	db, err := sql.Open(postgresDriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	err = db.PingContext(ctx)
	for err != nil {
		nextErr := db.PingContext(ctx)
		if errors.Is(nextErr, context.DeadlineExceeded) || errors.Is(nextErr, context.Canceled) {
			return nil, fmt.Errorf("db.Ping: %w", err)
		}
		err = nextErr
	}

	if err = migrateUp(db, fmt.Sprintf("file://%s", migrationDir), moduleName); err != nil {
		return nil, fmt.Errorf("migrateUp: %w", err)
	}

	return &Repo{
		DB: sqlx.NewDb(db, postgresDriverName),
	}, nil
}

func migrateUp(dbConn *sql.DB, fileRoot, moduleName string) error {
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{
		MigrationsTable: fmt.Sprintf("%s_migrations", moduleName),
	})
	if err != nil {
		return fmt.Errorf("postgres.WithInstance: %w", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(fileRoot, "postgres", driver)
	if err != nil {
		return fmt.Errorf("migrate.New: %w", err)
	}

	err = mig.Up()
	switch {
	case err == nil:
		return nil
	case errors.Is(err, migrate.ErrNoChange):
		return nil
	default:
		return fmt.Errorf("mig.Up: %w", err)
	}
}
