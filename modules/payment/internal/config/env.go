package config

const devEnv = "dev"

// Env is a config for environments for application, e.g. log file path, application mode.
type Env struct {
	ProgramEnv string `envconfig:"ENV" default:"dev"`
}
