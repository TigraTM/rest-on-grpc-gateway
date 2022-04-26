package repo

import (
	"context"
	"fmt"

	"rest-on-grpc-gateway/modules/payment/internal/config"
	"rest-on-grpc-gateway/pkg/repo"
)

const (
	dbMaxOpenConns = 20
	dbMaxIdleConns = 5
)

type Repo struct {
	*repo.Repo
}

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
