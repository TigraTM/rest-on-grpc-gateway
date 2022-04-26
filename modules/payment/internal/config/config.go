// Package config contains load config for user module from env.
package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	devEnvPath = "environments/dev"
	configName = "configs.yaml"
)

// Config is a common struct contains other configs.
type Config struct {
	Env
	Transport
	Database
}

// IsDev check dev mode.
func (c *Config) IsDev() bool {
	return c.ProgramEnv == devEnv
}

// LoadConfig loads config.
func LoadConfig(envPath string) (*Config, error) {
	if err := godotenv.Load(
		fmt.Sprintf("%s/%s", envPath, configName),
	); err != nil {
		log.Printf("Can't load env: %+v \n", err)
	}

	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("unable to decode env into struct: %w", err)
	}

	return &cfg, nil
}

// LoadDevConfig loads config for dev env.
func LoadDevConfig() (*Config, error) {
	return LoadConfig(devEnvPath)
}
