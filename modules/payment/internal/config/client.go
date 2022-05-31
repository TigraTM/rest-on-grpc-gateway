package config

// Clients contains all data for integration with third-parti clients.
type Clients struct {
	APILayerAPIKey   string `envconfig:"API_LAYER_API_KEY" required:"true"`
	APILayerBasePath string `envconfig:"API_LAYER_BASE_PATH" required:"true" default:"5432"`

	UserClientHost string `envconfig:"USER_CLIENT_HOST" required:"true" default:"0.0.0.0"`
	UserClientPort int    `envconfig:"USER_CLIENT_PORT" required:"true" default:"8081"`
}
