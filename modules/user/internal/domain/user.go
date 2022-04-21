// Package domain contains domain models.
package domain

import (
	"time"
)

// User domain model user.
type User struct {
	ID           int
	CreateAt     time.Time
	UpdateAt     time.Time
	Name         string
	Email        string
	Password     string
	PasswordHash []byte
}
