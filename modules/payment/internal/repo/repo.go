package repo

import (
	"context"
	"fmt"
	"rest-on-grpc-gateway/modules/payment/internal/config"
	"rest-on-grpc-gateway/pkg/repo"

	"github.com/jmoiron/sqlx"
)

const (
	dbMaxOpenConns = 20
	dbMaxIdleConns = 5
)

// Repo structure for work with database.
type Repo struct {
	*repo.Repo
}

// WrapperTx structure for work with transaction database.
type WrapperTx struct {
	tx *sqlx.Tx
}

// New build and return new Repo.
func New(ctx context.Context, dbCfg *config.Database, moduleName string) (_ *Repo, err error) {
	r := &Repo{}
	r.Repo, err = repo.NewPostgres(ctx, dbCfg.DSN(), dbCfg.MigrationsDir, moduleName)
	if err != nil {
		return nil, fmt.Errorf("repo.NewPostgres: %w", err)
	}

	r.DB.SetMaxOpenConns(dbMaxOpenConns)
	r.DB.SetMaxIdleConns(dbMaxIdleConns)

	return r, nil
}
