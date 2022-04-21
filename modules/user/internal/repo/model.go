package repo

import (
	"time"
)

// User model for work with database.
type User struct {
	ID           int       `db:"id"`
	CreateAt     time.Time `db:"create_at"`
	UpdateAt     time.Time `db:"update_at"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash []byte    `db:"password_hash"`
}
