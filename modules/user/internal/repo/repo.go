// Package repo contains methods for work with database.
package repo

import (
	"context"
	"fmt"

	"rest-on-grpc-gateway/modules/user/internal/config"
	"rest-on-grpc-gateway/pkg/repo"
)

const (
	dbMaxOpenConns = 20
	dbMaxIdleConns = 5
)

// Repo wrapper on repo.Repo.
type Repo struct {
	*repo.Repo
}

// New build and return new Repo.
func New(ctx context.Context, dbCfg *config.Database) (_ *Repo, err error) {
	r := &Repo{}
	r.Repo, err = repo.NewPostgres(ctx, dbCfg.DSN(), dbCfg.MigrationsDir)
	if err != nil {
		return nil, fmt.Errorf("repo.NewPostgres: %w", err)
	}

	r.DB.SetMaxOpenConns(dbMaxOpenConns)
	r.DB.SetMaxIdleConns(dbMaxIdleConns)

	return r, nil
}
