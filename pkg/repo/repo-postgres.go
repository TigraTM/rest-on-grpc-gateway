package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func NewPostgres(ctx context.Context, dsn string) (_ *Repo, err error) {
	db, err := sql.Open(postgresDriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	// TODO: add auto migration.

	err = db.PingContext(ctx)
	for err != nil {
		nextErr := db.PingContext(ctx)
		if errors.Is(nextErr, context.DeadlineExceeded) || errors.Is(nextErr, context.Canceled) {
			return nil, fmt.Errorf("db.Ping: %w", err)
		}
		err = nextErr
	}

	return &Repo{
		DB: sqlx.NewDb(db, postgresDriverName),
	}, nil
}
