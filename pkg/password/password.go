// Package password will hash and comparable hash-pass.
package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Manager contains method for hashing and comparable value.
type Manager struct {
	cost int
}

// New creates and returns new Hasher.
func New() *Manager {
	return &Manager{cost: bcrypt.DefaultCost}
}

// Hashing value and returns bytes.
func (m *Manager) Hashing(val string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(val), m.cost)
}

// Compare comparable two hash.
func (*Manager) Compare(val1 []byte, val2 []byte) bool {
	return bcrypt.CompareHashAndPassword(val1, val2) == nil
}
