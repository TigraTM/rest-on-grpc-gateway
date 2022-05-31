// Package repo helper for database connect.
package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// Repo structure for working with the database.
type Repo struct {
	DB *sqlx.DB
}

// Turn sqlx errors like `missing destination …` into panics
// https://github.com/jmoiron/sqlx/issues/529. As we can't distinguish
// between sqlx and other errors except driver ones, let's hope filtering
// driver errors is enough and there are no other non-driver regular errors.
func (*Repo) strict(err error) error {
	switch {
	case err == nil:
	case errors.As(err, new(*pq.Error)):
	case errors.Is(err, sql.ErrNoRows):
	case errors.Is(err, context.Canceled):
	case errors.Is(err, context.DeadlineExceeded):
	default:
		return err
	}

	return err
}

// Tx provides DAL method wrapper with:
// - wrapping errors with DAL method name,
// - transaction.
func (r *Repo) Tx(ctx context.Context, opts *sql.TxOptions, f func(*sqlx.Tx) error) (err error) {
	return r.strict(func() error {
		tx, err := r.DB.BeginTxx(ctx, opts)
		if err == nil { //nolint:nestif // No idea how to simplify.
			defer func() {
				if err := recover(); err != nil {
					if errRollback := tx.Rollback(); errRollback != nil {
						err = fmt.Errorf("%v: %w", err, errRollback)
					}
					panic(err)
				}
			}()
			err = f(tx)
			if err == nil {
				err = tx.Commit()
			} else if errRollback := tx.Rollback(); errRollback != nil {
				err = fmt.Errorf("%w: %s", err, errRollback)
			}
		}
		if err != nil {
			err = fmt.Errorf("failed tx: %w", err)
		}

		return err
	}())
}
