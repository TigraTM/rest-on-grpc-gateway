package config

import (
	"fmt"
	"net/url"
)

const (
	sslModeURLParameter = "sslmode"
)

// Database config for connect to databse.
type Database struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     int    `envconfig:"DB_PORT" default:"5432"`
	DBName   string `envconfig:"DB_NAME" default:"postgres"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`

	Parameters *DBParameters
}

type DBParameters struct {
	SSLMode string `envconfig:"DB_SSL_MODE" default:"disable"`
}

// DSN convert struct to DSN and returns connection string.
func (d *Database) DSN() string {
	vlues := make(url.Values)
	vlues.Set(sslModeURLParameter, d.Parameters.SSLMode)

	uri := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(d.User, d.Password),
		Host:     fmt.Sprintf("%s:%d", d.Host, d.Port),
		Path:     d.DBName,
		RawQuery: vlues.Encode(),
	}

	return uri.String()
}
