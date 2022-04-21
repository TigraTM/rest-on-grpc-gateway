package config

// Transport contains data for start and connected backend.
type Transport struct {
	Host       string `envconfig:"USER_HOST" default:"0.0.0.0"`
	GRPCPort   int    `envconfig:"USER_GRPC_PORT" default:"8080"`
	GRPCGWPort int    `envconfig:"USER_GRPC_GW_PORT" default:"8081"`
}
