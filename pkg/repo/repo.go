// Package repo helper for database connect.
package repo

import (
	"github.com/jmoiron/sqlx"
)

// Repo structure for working with the database.
type Repo struct {
	DB *sqlx.DB
}
